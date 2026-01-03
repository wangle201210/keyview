package app

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Repository 数据库访问层
type Repository struct {
	db *gorm.DB
}

// NewRepository 创建新的数据库仓库
func NewRepository(dbPath string) (*Repository, error) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(sqlite.Open(dbPath), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 自动迁移表结构
	if err := db.AutoMigrate(&KeyRecord{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Repository{db: db}, nil
}

// Close 关闭数据库连接
func (r *Repository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Create 创建记录
func (r *Repository) Create(record *KeyRecord) error {
	return r.db.Create(record).Error
}

// FindByPage 分页查找记录
func (r *Repository) FindByPage(offset, limit int) ([]KeyRecord, error) {
	var records []KeyRecord
	err := r.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&records).Error
	return records, err
}

// FindByFilter 根据条件筛选记录
func (r *Repository) FindByFilter(keyName string, startDate, endDate string, offset, limit int) ([]KeyRecord, error) {
	query := r.db.Model(&KeyRecord{})

	if keyName != "" {
		query = query.Where("key_name = ?", keyName)
	}

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	var records []KeyRecord
	err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&records).Error
	return records, err
}

// Count 统计总记录数
func (r *Repository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&KeyRecord{}).Count(&count).Error
	return count, err
}

// CountByDate 统计指定日期的记录数
func (r *Repository) CountByDate(date string) (int64, error) {
	var count int64
	err := r.db.Model(&KeyRecord{}).
		Where("DATE(created_at) = ?", date).
		Count(&count).Error
	return count, err
}

// GetUniqueKeyNames 获取所有唯一的按键名称
func (r *Repository) GetUniqueKeyNames() ([]string, error) {
	var keyNames []string
	err := r.db.Model(&KeyRecord{}).
		Distinct("key_name").
		Order("key_name ASC").
		Pluck("key_name", &keyNames).Error
	return keyNames, err
}

// DeleteBefore 删除指定日期之前的记录
func (r *Repository) DeleteBefore(date string) (int64, error) {
	result := r.db.Where("created_at < ?", date).Delete(&KeyRecord{})
	return result.RowsAffected, result.Error
}

// GetKeyStats 获取按键统计（支持日期范围）
func (r *Repository) GetKeyStats(startDate, endDate string) ([]KeyStats, error) {
	query := r.db.Model(&KeyRecord{})

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	var stats []KeyStats
	err := query.
		Select("key_name, count(*) as count").
		Group("key_name").
		Order("count DESC").
		Find(&stats).Error
	return stats, err
}

// GetModifierStats 获取修饰键统计（支持日期范围）
func (r *Repository) GetModifierStats(startDate, endDate string) ([]KeyStats, error) {
	modifiers := []struct {
		name string
		flag int
	}{
		{"Shift", 0x20000},
		{"Control", 0x40000},
		{"Option", 0x80000},
		{"Command", 0x100000},
		{"Caps Lock", 0x10000},
	}

	var stats []KeyStats

	for _, mod := range modifiers {
		var count int64
		query := r.db.Model(&KeyRecord{}).Where("modifier_flags & ? > 0", mod.flag)

		if startDate != "" {
			query = query.Where("created_at >= ?", startDate)
		}
		if endDate != "" {
			query = query.Where("created_at <= ?", endDate)
		}

		err := query.Count(&count).Error
		if err != nil {
			return nil, err
		}

		if count > 0 {
			stats = append(stats, KeyStats{
				KeyName: mod.name,
				Count:   count,
			})
		}
	}

	return stats, nil
}
