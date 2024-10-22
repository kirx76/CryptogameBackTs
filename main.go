package main

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/handlers"
	"CryptogameBackTs/middlewares"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Функция для проверки, что все переменные окружения заданы
func checkEnvVariables(vars []string) {
	for _, v := range vars {
		if value := os.Getenv(v); value == "" {
			log.Fatalf("Environment variable %s is not set or empty", v)
		}
	}
}

func main() {
	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Проверка обязательных переменных окружения
	requiredEnvVars := []string{
		"EMAIL_USER",
		"EMAIL_PASSWORD",
		"SMTP_HOST",
		"SMTP_PORT",
	}

	// Вызываем функцию валидации
	checkEnvVariables(requiredEnvVars)

	r := gin.Default()

	// Инициализация базы данных
	database.InitDB()

	// Маршруты авторизации
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.POST("/logout", handlers.LogoutUser)

	// Маршруты с авторизацией
	authorized := r.Group("/")
	authorized.Use(middlewares.AuthRequired)
	{
		authorized.GET("/quote/:id", handlers.GetQuoteByID)
		authorized.GET("/author/:id", handlers.GetAuthorByID)
		authorized.GET("/quote/current", handlers.GetQuoteForCurrentLevel)

		// Маршруты с изменением пользовательских данных
		authorized.PUT("/user/update", handlers.UpdateUserData)
		authorized.PUT("/user/verify-email", handlers.VerifyEmailChange)
	}

	// Запуск сервера
	r.Run(":8080")
}
