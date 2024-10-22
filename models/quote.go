package models

import (
	"time"
)

type Quote struct {
	ID            uint      `gorm:"column:id;primaryKey" json:"id"`
	Text          string    `gorm:"column:text" json:"text"`
	CreatedAt     time.Time `gorm:"column:createdAt;autoCreateTime" json:"createdAt"`
	LevelID       uint      `gorm:"column:levelId" json:"levelId"`
	Level         Level     `gorm:"foreignKey:LevelID" json:"level"`
	AuthorID      uint      `gorm:"column:authorId" json:"authorId"`
	Author        Author    `gorm:"foreignKey:AuthorID" json:"author"`
	OpenedIndexes string    `gorm:"column:openedIndexes" json:"openedIndexes"`
}

func (Quote) TableName() string {
	return "quote"
}
