package model

import (
	"time"
)

type Verification struct {
	VerificationID   uint      `json:"verification_id" gorm:"primaryKey"`
	VerificationType int8      `json:"verification_type"`
	UserID           uint      `json:"user_id"`
	Token            string    `json:"token"`
	Code             string    `json:"code"`
	DeliveryAddress  string    `json:"delivery_address"`
	DeliveryType     int8      `json:"delivery_type"`
	State            int8      `json:"state"`
	ExpiresAt        time.Time `json:"expires_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
