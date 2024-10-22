package middlewares

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthRequired проверяет, что пользователь авторизован
func AuthRequired(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Поиск сессии в базе данных
	var session models.Session
	if err := database.DB.Where("session_id = ?", sessionID).First(&session).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		c.Abort()
		return
	}

	// Поиск пользователя по ID из сессии
	var user models.User
	if err := database.DB.Where("id = ?", session.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	// Сохранение пользователя в контексте для дальнейшего использования в обработчиках
	c.Set("user", user)

	c.Next()
}

// ModeratorRequired проверяет, что пользователь имеет уровень доступа "модератор" (2) или выше
func ModeratorRequired(c *gin.Context) {
	AuthRequired(c)

	// Проверка, не был ли предыдущий middleware прерван
	if c.IsAborted() {
		return
	}

	// Извлечение пользователя из контекста
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	currentUser := user.(models.User)

	// Проверка уровня доступа: уровень 2 — модератор или выше
	if currentUser.AccessLevelID < 2 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
		c.Abort()
		return
	}

	c.Next()
}

// AdminRequired проверяет, что пользователь имеет уровень доступа "администратор" (3)
func AdminRequired(c *gin.Context) {
	AuthRequired(c)

	// Проверка, не был ли предыдущий middleware прерван
	if c.IsAborted() {
		return
	}

	// Извлечение пользователя из контекста
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	currentUser := user.(models.User)

	// Проверка уровня доступа: уровень 3 — администратор
	if currentUser.AccessLevelID < 3 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
		c.Abort()
		return
	}

	c.Next()
}
