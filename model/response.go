package model

type AuthRegisterResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthLoginResponse struct {
	AccessToken string                `json:"access_token"`
	User        AuthLoginUserResponse `json:"user"`
}

type AuthLoginUserResponse struct {
	UserID                 uint    `json:"user_id" gorm:"primaryKey"`
	Name                   *string `json:"name"`
	Email                  string  `json:"email"`
	Phone                  *string `json:"phone"`
	DisplayPictureURL      *string `json:"display_picture_url"`
	IdentityID             *string `json:"identity_id"`
	IdentityImageURL       *string `json:"identity_image_url"`
	InviteNotificationFlag bool    `json:"invite_notification_flag"`
	IsActiveFlag           bool    `json:"is_active_flag"`
	IsEmailVerifiedFlag    bool    `json:"is_email_verified_flag"`
	RoleID                 uint    `json:"role_id"`
}

type AdminGetUsersResult struct {
	Total         uint          `json:"-"`
	Offset        uint          `json:"-"`
	Page          uint          `json:"-"`
	Limit         uint          `json:"-"`
}

type AdminSearchUsersResult struct {
	Total         uint          `json:"-"`
	Offset        uint          `json:"-"`
	Limit         uint          `json:"-"`
}

type AuthInvitationResponse struct {
	Name            string `json:"name"`
	Emai            string `json:"email"`
	InvitationToken string `json:"invitation_token"`
}
