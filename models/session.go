package models

import "time"

type Session struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id" json:"user_id"`
	SessionID string    `gorm:"column:session_id;unique" json:"session_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Session) TableName() string {
	return "sessions"
}
