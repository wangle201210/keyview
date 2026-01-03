package app

import "time"

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

// KeyStats 按键统计
type KeyStats struct {
	KeyName string `json:"key_name"`
	Count   int64  `json:"count"`
}
