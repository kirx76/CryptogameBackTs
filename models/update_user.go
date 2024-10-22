package models

type UpdateUserDataInput struct {
	Name      string `json:"name,omitempty"`
	BirthDate string `json:"birth_date,omitempty"` // Дата рождения как строка
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
