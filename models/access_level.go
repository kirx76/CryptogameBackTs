package models

type AccessLevel struct {
	ID   uint   `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (AccessLevel) TableName() string {
	return "access_levels"
}
