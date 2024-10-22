package handlers

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQuoteByID(c *gin.Context) {
	// Извлечение пользователя из контекста
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Остальной код без изменений
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quote ID"})
		return
	}

	var quote models.Quote
	if result := database.DB.Preload("Author").Preload("Level").First(&quote, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quote not found"})
		return
	}

	c.JSON(http.StatusOK, quote)
}

func GetQuoteForCurrentLevel(c *gin.Context) {
	// Извлечение пользователя из контекста
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Преобразование user в структуру модели User
	currentUser := user.(models.User)

	// Поиск цитаты по текущему уровню пользователя
	var quote models.Quote
	if result := database.DB.Preload("Author").Preload("Level").Where("id = ?", currentUser.CurrentLevel).First(&quote); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quote for current level not found"})
		return
	}

	c.JSON(http.StatusOK, quote)
}
