package app

import (
	"path/filepath"
)

// GetUserDataDir 获取用户数据目录
func GetUserDataDir() (string, error) {
	// TODO: 实现跨平台的用户数据目录获取
	// macOS: ~/Library/Application Support/KeyView
	// Windows: %APPDATA%/KeyView
	// Linux: ~/.config/keyview
	// 目前简单实现，返回当前目录
	return ".", nil
}

// GetDatabasePath 获取数据库文件路径
func GetDatabasePath() (string, error) {
	homeDir, err := GetUserDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "keyview.db"), nil
}
