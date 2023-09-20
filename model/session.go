package model

import "time"

type Session struct {
	SessionID     uint       `json:"session_id" gorm:"primaryKey"`
	ResourceType  string     `json:"resource_type"`
	ResourceID    uint       `json:"resource_id"`
	Token         string     `json:"token"`
	FirebaseToken string     `json:"firebase_token"`
	ExpiresAt     *time.Time `json:"expires_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
