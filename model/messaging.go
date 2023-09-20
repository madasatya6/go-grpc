package model

type PushNotifRequest struct {
	UserID   uint              `json:"user_id"`
	Title    string            `json:"title"`
	Body     string            `json:"body"`
	ImageURL string            `json:"image_url"`
	Data     map[string]string `json:"data"`
}

type AdminPushNotifRequest struct {
	Email    string `json:"email"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageURL string `json:"image_url"`
}
