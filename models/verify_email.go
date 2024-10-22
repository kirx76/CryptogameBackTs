package models

type VerifyEmailCodeInput struct {
	VerificationCode string `json:"verification_code"`
	NewEmail         string `json:"new_email"`
}
