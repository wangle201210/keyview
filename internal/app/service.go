package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/wangle201210/keylogger"
)

// AppService 主应用服务
type AppService struct {
	mu              sync.RWMutex
	repository      *Repository
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

	// 获取数据库路径
	dbPath, err := GetDatabasePath()
	if err != nil {
		return fmt.Errorf("failed to get database path: %w", err)
	}

	// 创建数据库仓库
	repo, err := NewRepository(dbPath)
	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}
	s.repository = repo

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

	if s.repository == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	return s.repository.FindByPage(offset, limit)
}

// GetRecordsByFilter 根据条件筛选获取键盘记录
func (s *AppService) GetRecordsByFilter(keyName, startDate, endDate string, offset, limit int) ([]KeyRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.repository == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	return s.repository.FindByFilter(keyName, startDate, endDate, offset, limit)
}

// GetTotalCount 获取总记录数
func (s *AppService) GetTotalCount() (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.repository == nil {
		return 0, fmt.Errorf("repository not initialized")
	}

	return s.repository.Count()
}

// GetTodayKeystrokes 获取今日按键次数
func (s *AppService) GetTodayKeystrokes() (int64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.repository == nil {
		return 0, fmt.Errorf("repository not initialized")
	}

	today := getCurrentDate()
	return s.repository.CountByDate(today)
}

// GetUniqueKeyNames 获取所有唯一的按键名称
func (s *AppService) GetUniqueKeyNames() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.repository == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	return s.repository.GetUniqueKeyNames()
}

// GetKeyStats 获取所有按键的统计次数（支持日期范围筛选）
func (s *AppService) GetKeyStats(startDate, endDate string) ([]KeyStats, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.repository == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	// 获取普通按键统计
	stats, err := s.repository.GetKeyStats(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// 添加修饰键统计
	modifierStats, err := s.repository.GetModifierStats(startDate, endDate)
	if err != nil {
		return nil, err
	}

	stats = append(stats, modifierStats...)
	return stats, nil
}

// DeleteRecordsBefore 删除指定日期之前的记录
func (s *AppService) DeleteRecordsBefore(date string) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.repository == nil {
		return 0, fmt.Errorf("repository not initialized")
	}

	return s.repository.DeleteBefore(date)
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

	if s.repository != nil {
		return s.repository.Close()
	}

	return nil
}

// getCurrentDate 获取当前日期字符串
func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}
