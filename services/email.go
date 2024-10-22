package services

import (
	"CryptogameBackTs/database"
	"CryptogameBackTs/models"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

var emailVerificationCodes = make(map[string]string)

func generateVerificationCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	code := rand.Intn(900000) + 100000 // Генерируем число в диапазоне от 100000 до 999999
	return fmt.Sprintf("%d", code)     // Форматируем как строку без ведущих нулей
}

func sendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Создаем сообщение
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Настройки аутентификации
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Отправка письма
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func RequestEmailChange(user models.User, newEmail string) error {
	// Генерируем код
	verificationCode := generateVerificationCode()

	log.Println(verificationCode, "verificationCode")

	// Отправляем email с кодом
	subject := "Email Change Verification"
	body := fmt.Sprintf("Your verification code is: %s", verificationCode)
	err := sendEmail(newEmail, subject, body)
	if err != nil {
		return err
	}

	// Сохраняем код и email в мапе для дальнейшей проверки
	emailVerificationCodes[user.Username] = verificationCode

	return nil
}

func VerifyEmailChange(user models.User, code string, newEmail string) error {
	if storedCode, exists := emailVerificationCodes[user.Username]; exists {
		if storedCode == code {
			// Код совпал, обновляем email пользователя
			user.Email = newEmail
			if err := database.DB.Save(&user).Error; err != nil {
				return err
			}
			// Удаляем код после успешного подтверждения
			delete(emailVerificationCodes, user.Username)
			return nil
		}
	}
	return fmt.Errorf("invalid verification code")
}
