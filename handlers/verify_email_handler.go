package handlers

import (
	"CryptogameBackTs/models"
	"CryptogameBackTs/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// Функция для проверки формата email
func isValidEmail(email string) bool {
	// Простое регулярное выражение для проверки email
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func VerifyEmailChange(c *gin.Context) {
	// Извлекаем текущего пользователя из контекста
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(models.User)

	// Получаем код подтверждения и новый email из запроса
	var input models.VerifyEmailCodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Валидация нового email
	if input.NewEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email cannot be empty"})
		return
	}

	if !isValidEmail(input.NewEmail) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Вызываем сервис для проверки кода и обновления email
	if err := services.VerifyEmailChange(currentUser, input.VerificationCode, input.NewEmail); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid verification code"})
		return
	}

	// Успешное обновление email
	c.JSON(http.StatusOK, gin.H{"message": "Email updated successfully"})
}
