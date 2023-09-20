package model

import (
	"time"
)

type RegisterResult struct {
	VerificationToken string `json:"verification_token"`
}

type ForgotPasswordResult struct {
	VerificationToken string    `json:"verification_token"`
	ExpiresAt         time.Time `json:"expires_at"`
}
