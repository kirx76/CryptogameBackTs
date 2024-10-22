package models

type Author struct {
	ID     uint    `gorm:"column:id;primaryKey" json:"id"`
	Name   string  `gorm:"column:name" json:"name"`
	Quotes []Quote `gorm:"foreignKey:AuthorID" json:"-"` // Убираем это поле из JSON
}

func (Author) TableName() string {
	return "author"
}
