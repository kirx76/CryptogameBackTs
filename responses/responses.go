package responses

import "time"

// Упрощенная структура цитаты без информации об авторе
type QuoteResponse struct {
	ID            uint          `json:"id"`
	Text          string        `json:"text"`
	CreatedAt     time.Time     `json:"createdAt"`
	LevelID       uint          `json:"levelId"`
	Level         LevelResponse `json:"level"`
	OpenedIndexes string        `json:"openedIndexes"`
}

type LevelResponse struct {
	ID    uint `json:"id"`
	Level int  `json:"level"`
}
