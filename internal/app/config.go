package app

import (
	"os"
	"path/filepath"
)

// GetUserDataDir 获取用户数据目录
func GetUserDataDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".keyview"), nil
}

// GetDatabasePath 获取数据库文件路径
func GetDatabasePath() (string, error) {
	homeDir, err := GetUserDataDir()
	if err != nil {
		return "", err
	}

	// 确保目录存在
	if err := os.MkdirAll(homeDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "keyview.db"), nil
}
