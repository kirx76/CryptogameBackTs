package models

import (
	"time"
)

type User struct {
	ID               uint      `gorm:"column:id;primaryKey" json:"id"`
	Username         string    `gorm:"column:username;unique" json:"username"`
	PasswordHash     string    `gorm:"column:password_hash" json:"-"`
	RegistrationDate time.Time `gorm:"column:registration_date;autoCreateTime" json:"registration_date"`
	Name             string    `gorm:"column:name" json:"name,omitempty"`
	BirthDate        time.Time `gorm:"column:birth_date;type:date" json:"birth_date,omitempty"` // изменено
	Email            string    `gorm:"column:email" json:"email,omitempty"`
	Phone            string    `gorm:"column:phone" json:"phone,omitempty"`
	AccessLevelID    uint      `gorm:"column:access_level_id" json:"access_level_id"`
	CurrentLevel     uint      `gorm:"column:current_level;default:1" json:"current_level"`
}

func (User) TableName() string {
	return "users"
}
