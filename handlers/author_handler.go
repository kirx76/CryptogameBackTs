package handlers

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/models"
	"CryptogameBackTs/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthorByID(c *gin.Context) {
	// Извлечение пользователя из контекста
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var author models.Author

	if err := database.DB.Preload("Quotes.Level").First(&author, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	// Маппинг модели на кастомную структуру
	var quotes []responses.QuoteResponse
	for _, quote := range author.Quotes {
		quotes = append(quotes, responses.QuoteResponse{
			ID:            quote.ID,
			Text:          quote.Text,
			CreatedAt:     quote.CreatedAt,
			LevelID:       quote.LevelID,
			Level:         responses.LevelResponse{ID: quote.Level.ID, Level: quote.Level.Level},
			OpenedIndexes: quote.OpenedIndexes,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":     author.ID,
		"Name":   author.Name,
		"Quotes": quotes,
	})
}
