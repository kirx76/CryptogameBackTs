package handlers

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/models"
	"CryptogameBackTs/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

// Регулярное выражение для валидации формата даты
var dateRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

func UpdateUserData(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	currentUser := user.(models.User)

	var input models.UpdateUserDataInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Обновляем необязательные данные, если они были отправлены
	if input.Name != "" {
		currentUser.Name = input.Name
	}

	// Парсим и валидируем дату рождения
	if input.BirthDate != "" {
		if !dateRegex.MatchString(input.BirthDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birth date format. Expected YYYY-MM-DD"})
			return
		}

		// Парсим дату в формате "YYYY-MM-DD"
		parsedDate, err := time.Parse("2006-01-02", input.BirthDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse birth date"})
			return
		}
		currentUser.BirthDate = parsedDate
	}

	if input.Phone != "" {
		currentUser.Phone = input.Phone
	}

	// Обновление email через сервис
	if input.Email != "" && input.Email != currentUser.Email {
		err := services.RequestEmailChange(currentUser, input.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error requesting email change"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Verification email sent. Please verify to complete the email change."})
		return
	}

	if err := database.DB.Save(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user data"})
		return
	}

	// Сохраняем изменения в базе данных
	if err := database.DB.Save(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User data updated successfully"})
}
