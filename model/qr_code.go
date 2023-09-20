package model

type QRCodeJWTPayload struct {
	MeetingID  uint  `json:"meeting_id"`
	UserID     uint  `json:"user_id"`
	ExpiryTime int64 `json:"expiry_time"` // unix timestamp
}
