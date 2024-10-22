package handlers

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Проверка длины логина
	if len(user.Username) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be at least 5 characters long"})
		return
	}

	// Проверка на уникальность логина
	var existingUser models.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	// Хешируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.PasswordHash = string(hash)

	// Сохраняем пользователя
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Генерация уникального идентификатора сессии
	sessionID := uuid.New().String()

	// Сохраняем сессию в базе данных
	session := models.Session{
		UserID:    user.ID,
		SessionID: sessionID,
	}
	if err := database.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating session"})
		return
	}

	// Установка куки
	c.SetCookie("session_id", sessionID, 604800, "/", "", false, true)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "session_id": sessionID})
}

func LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Ищем пользователя по имени пользователя (username)
	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Проверяем пароль, сравнивая введенный пароль с сохраненным хешем
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Генерация уникального идентификатора сессии
	sessionID := uuid.New().String()

	// Сохраняем сессию в базе данных
	session := models.Session{
		UserID:    user.ID,
		SessionID: sessionID,
	}
	if err := database.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating session"})
		return
	}

	// Установка куки на одну неделю
	c.SetCookie("session_id", sessionID, 604800, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "session_id": sessionID})
}

func LogoutUser(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session not found"})
		return
	}

	// Удаление сессии из базы данных
	if err := database.DB.Where("session_id = ?", sessionID).Delete(&models.Session{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error logging out"})
		return
	}

	// Удаление куки
	c.SetCookie("session_id", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
