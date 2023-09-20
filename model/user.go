package model

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	UserID                   uint      `json:"user_id" gorm:"primaryKey"`
	CompanyID                *uint     `json:"company_id"`
	Name                     *string   `json:"name"`
	Email                    string    `json:"email"`
	InvitationToken          string    `json:"invitation_token"`
	Phone                    *string   `json:"phone"`
	EncryptedPassword        string    `json:"encrypted_password"`
	DisplayPictureURL        *string   `json:"display_picture_url"`
	IdentityID               *string   `json:"identity_id"`
	IdentityImageURL         *string   `json:"identity_image_url"`
	InviteNotificationFlag   bool      `json:"invite_notification_flag"`
	IsActiveFlag             bool      `json:"is_active_flag"`
	IsEmailVerifiedFlag      bool      `json:"is_email_verified_flag"`
	IsRegisteredFlag         bool      `json:"is_registered_flag"`
	InvitationTokenExpiredAt time.Time `json:"invitation_token_expired_at"`
	CreatedAt                time.Time `json:"created_at"`
	CreatedBy                uint      `json:"created_by"`
	UpdatedAt                time.Time `json:"updated_at"`
	UpdatedBy                uint      `json:"updated_by"`
}

type UserResponse struct {
	UserID            uint    `json:"user_id" gorm:"primaryKey"`
	Name              *string `json:"name"`
	Email             string  `json:"email"`
	Phone             *string `json:"phone"`
	DisplayPictureURL *string `json:"display_picture_url"`
}

func (u *User) SerializePublic() interface{} {
	return struct {
		UserID                 uint      `json:"user_id" gorm:"primaryKey"`
		CompanyID              *uint     `json:"company_id"`
		Name                   *string   `json:"name"`
		Email                  string    `json:"email"`
		Phone                  *string   `json:"phone"`
		DisplayPictureURL      *string   `json:"display_picture_url"`
		IdentityID             *string   `json:"identity_id"`
		IdentityImageURL       *string   `json:"identity_image_url"`
		InviteNotificationFlag bool      `json:"invite_notification_flag"`
		CreatedAt              time.Time `json:"created_at"`
		UpdatedAt              time.Time `json:"updated_at"`
	}{
		UserID:                 u.UserID,
		CompanyID:              u.CompanyID,
		Name:                   u.Name,
		Email:                  u.Email,
		Phone:                  u.Phone,
		DisplayPictureURL:      u.DisplayPictureURL,
		IdentityID:             u.IdentityID,
		IdentityImageURL:       u.IdentityImageURL,
		InviteNotificationFlag: u.InviteNotificationFlag,
		CreatedAt:              u.CreatedAt,
		UpdatedAt:              u.UpdatedAt,
	}
}

type UsersQuery struct {
	CompanyID uint
	Offset    uint
	Page      uint
	Limit     uint
	Sort      []string
	Role      []string
	Keyword   string
}

func (query UsersQuery) GetOrderQuery() string {
	fieldMap := map[string]string{
		"user_id":      "users.user_id",
		"name":         "users.name",
		"company_name": "companies.name",
	}

	result := []string{}
	for _, s := range query.Sort {
		if len(s) == 0 {
			continue
		}

		order, key := "ASC", s
		if s[:1] == "-" {
			order, key = "DESC", s[1:]
		}

		fieldName, ok := fieldMap[key]
		if !ok {
			continue
		}

		result = append(result, fmt.Sprintf("%s %s", fieldName, order))
	}

	return strings.Join(result, ",")
}

type SearchUsersQuery struct {
	CompanyID uint
	Email     string
}
