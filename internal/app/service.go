package app

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/wangle201210/keylogger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// KeyRecord 数据库记录模型
type KeyRecord struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	KeyCode       int       `json:"key_code"`
	KeyName       string    `json:"key_name"`
	IsDown        bool      `json:"is_down"`
	ModifierFlags int       `json:"modifier_flags"`
}

// TableName 指定表名
func (KeyRecord) TableName() string {
	return "key_records"
}

// AppService 主应用服务
type AppService struct {
	mu              sync.RWMutex
	db              *gorm.DB
	storage         *keylogger.SQLiteStorage
	isRecording     bool
	cancelRecording context.CancelFunc
}

// NewAppService 创建新的应用服务
func NewAppService() *AppService {
	return &AppService{
		isRecording: false,
	}
}

// Init 初始化数据库并自动启动键盘监听
func (s *AppService) Init() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 获取用户数据目录
	homeDir, err := getUserDataDir()
	if err != nil {
		return fmt.Errorf("failed to get user data dir: %w", err)
	}

	dbPath := filepath.Join(homeDir, "keyview.db")

	// 打开数据库连接
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(sqlite.Open(dbPath), config)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	s.db = db

	// 自动迁移表结构
	if err := db.AutoMigrate(&KeyRecord{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// 创建 keylogger 存储
	s.storage, err = keylogger.NewSQLiteStorage(dbPath)
	if err != nil {
		return fmt.Errorf("failed to create storage: %w", err)
	}

	// 自动启动键盘监听
	go s.startRecordingInBackground()

	return nil
}

// startRecordingInBackground 在后台启动键盘监听
func (s *AppService) startRecordingInBackground() {
	keylogger.StartWithStorage(func(event keylogger.KeyEvent) {
		fmt.Println(event.KeyName)
		// 事件已通过存储自动保存
	}, s.storage)
	s.isRecording = true
}

// StartRecording 开始记录键盘事件
func (s *AppService) StartRecording() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRecording {
		return fmt.Errorf("recording is already in progress")
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		keylogger.StartWithStorage(func(event keylogger.KeyEvent) {
			// 事件已通过存储自动保存
		}, s.storage)

		<-ctx.Done()
	}()

	s.isRecording = true
	s.cancelRecording = cancel

	return nil
}

// StopRecording 停止记录键盘事件
func (s *AppService) StopRecording() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRecording {
		return fmt.Errorf("no recording in progress")
	}

	if s.cancelRecording != nil {
		s.cancelRecording()
		s.cancelRecording = nil
	}

	s.isRecording = false

	return nil
}

// IsRecording 返回是否正在记录
func (s *AppService) IsRecording() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isRecording
}

// GetRecords 获取键盘记录
func (s *AppService) GetRecords(offset, limit int) ([]KeyRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var records []KeyRecord
	err := s.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, err
}

// GetRecordsByFilter 根据条件筛选获取键盘记录
func (s *AppService) GetRecordsByFilter(keyName string, date string, offset, limit int) ([]KeyRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	query := s.db.Model(&KeyRecord{})

	if keyName != "" {
		query = query.Where("key_name = ?", keyName)
	}

	if date != "" {
		query = query.Where("DATE(created_at) = ?", date)
	}

	var records []KeyRecord
	err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&records).Error

	return records, err
}

// GetTotalCount 获取总记录数
func (s *AppService) GetTotalCount() (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.db == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	var count int64
	err := s.db.Model(&KeyRecord{}).Count(&count).Error
	return count, err
}

// GetTodayKeystrokes 获取今日按键次数
func (s *AppService) GetTodayKeystrokes() (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.db == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	var count int64
	today := time.Now().Format("2006-01-02")
	err := s.db.Model(&KeyRecord{}).
		Where("DATE(created_at) = ?", today).
		Count(&count).Error

	return count, err
}

// GetUniqueKeyNames 获取所有唯一的按键名称
func (s *AppService) GetUniqueKeyNames() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var keyNames []string
	err := s.db.Model(&KeyRecord{}).
		Distinct("key_name").
		Order("key_name ASC").
		Pluck("key_name", &keyNames).Error

	return keyNames, err
}

// DeleteRecordsBefore 删除指定日期之前的记录
func (s *AppService) DeleteRecordsBefore(date string) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.db == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	result := s.db.Where("created_at < ?", date).Delete(&KeyRecord{})
	return result.RowsAffected, result.Error
}

// KeyStats 按键统计
type KeyStats struct {
	KeyName string `json:"key_name"`
	Count   int64  `json:"count"`
}

// GetKeyStats 获取所有按键的统计次数（支持日期范围筛选）
func (s *AppService) GetKeyStats(startDate, endDate string) ([]KeyStats, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	query := s.db.Model(&KeyRecord{})

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

	if err != nil {
		return nil, err
	}

	// 添加修饰键统计
	modifierStats, err := s.getModifierStats(startDate, endDate)
	if err != nil {
		return nil, err
	}

	stats = append(stats, modifierStats...)
	return stats, err
}

// getModifierStats 获取修饰键的统计次数（支持日期范围筛选）
func (s *AppService) getModifierStats(startDate, endDate string) ([]KeyStats, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	// 定义修饰键标志位
	// 0x10000 (65536) = Caps Lock
	// 0x20000 (131072) = Shift
	// 0x40000 (262144) = Control
	// 0x80000 (524288) = Option
	// 0x100000 (1048576) = Command

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
		query := s.db.Model(&KeyRecord{}).Where("modifier_flags & ? > 0", mod.flag)

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

// Close 关闭服务
func (s *AppService) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRecording && s.cancelRecording != nil {
		s.cancelRecording()
	}

	if s.storage != nil {
		_ = s.storage.Close()
	}

	return nil
}

// getUserDataDir 获取用户数据目录
func getUserDataDir() (string, error) {
	// 简单实现，返回当前目录
	// 实际应该根据平台返回合适的目录
	return ".", nil
}
