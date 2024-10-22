package models

type Level struct {
	ID     uint    `gorm:"column:id;primaryKey" json:"id"`
	Level  int     `gorm:"column:level" json:"level"`
	Quotes []Quote `gorm:"foreignKey:LevelID" json:"-"` // Убираем из JSON, чтобы избежать дублирования
}

func (Level) TableName() string {
	return "level"
}
