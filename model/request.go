package model

import (
	"io"

	"mime/multipart"

	"go_grpc/lib"
)

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password" validate:"required"`
	CompanyID       uint   `json:"company_id"`
	InvitationToken string `json:"invitation_token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GetVerificationTokenRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyRequest struct {
	Token string `json:"token"`
	Code  string `json:"code"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResendOTPRequest struct {
	Token string `json:"token"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type LogoutRequest struct {
	SessionID uint `json:"-"`
}

type UpdateFirebaseTokenRequest struct {
	SessionID uint   `json:"-"`
	Token     string `json:"token"`
}

type DeviceAuthRequest struct {
	DeviceScannerID string `json:"device_scanner_id" validate:"required"`
	ClientID        string `json:"client_id" validate:"required"`
	ClientSecret    string `json:"client_secret" validate:"required"`
}

type UpdateUserProfileRequest struct {
	Name                   string `json:"name"`
	Email                  string `json:"email"`
	PhoneNumber            string `json:"phone_number"`
	InviteNotificationFlag *bool  `json:"invite_notification_flag"`
}

type UpdateDisplayPictureRequest struct {
	UserID                   uint `json:"-"`
	DisplayPictureFile       multipart.File
	DisplayPictureFileHeader *multipart.FileHeader
}

type UploadIdentityRequest struct {
	IdentityNumber string    `json:"identity_number"`
	IdentityImage  io.Reader `json:"identity_image"`
	ContentType    string    `json:"content_type"`
	Extension      string    `json:"extension"`
}

func (r UpdateUserProfileRequest) Validate() error {
	if r.Name == "" {
		return lib.InvalidParameterError("name", "nama harus diisi")
	}

	if r.Email == "" {
		return lib.InvalidParameterError("email", "email harus diisi")
	}

	if r.PhoneNumber == "" {
		return lib.InvalidParameterError("phone_number", "nomor telepon harus diisi")
	}

	return nil
}

type AdminGetUsersRequest struct {
	AdminID uint
	Offset  uint
	Page    uint
	Limit   uint
	Sorts   []string
	Roles   []string
	Keyword string
}

type SuperAdminGetUsersRequest struct {
	Offset  uint
	Page    uint
	Limit   uint
	Sorts   []string
	Roles   []string
	Keyword string
}

type AdminSearchUsersRequest struct {
	AdminID uint
	Email   string
}

type AdminCreateUserRequest struct {
	AdminID  uint   `json:"-"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	RoleName string `json:"role"`
}

type SuperAdminCreateAdminRequest struct {
	SuperAdminID uint   `json:"-"`
	BuildingID   uint   `json:"building_id"`
	CompanyID    uint   `json:"company_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	RoleName     string `json:"role"`
}

type SuperAdminUpdateAdminRequest struct {
	SuperAdminID uint   `json:"-"`
	UserID       uint   `json:"-"`
	CompanyID    uint   `json:"company_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	RoleName     string `json:"role"`
}

type SuperAdminCreateUserRequest struct {
	SuperAdminID uint   `json:"-"`
	BuildingID   uint   `json:"building_id"`
	CompanyID    uint   `json:"company_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	RoleName     string `json:"role"`
}

type SuperAdminUpdateUserRequest struct {
	SuperAdminID uint   `json:"-"`
	UserID       uint   `json:"-"`
	BuildingID   uint   `json:"building_id"`
	CompanyID    uint   `json:"company_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	RoleName     string `json:"role"`
}

type AdminAssignUserCompanyRequest struct {
	AdminID   uint `json:"-"`
	UserID    uint `json:"user_id" validate:"required"`
	CompanyID uint `json:"company_id" validate:"required"`
}

func (r AdminAssignUserCompanyRequest) Validate() error {
	if r.UserID == 0 {
		return lib.InvalidParameterError("user_id", "user_id harus diisi")
	}

	if r.CompanyID == 0 {
		return lib.InvalidParameterError("company_id", "company_id harus diisi")
	}

	return nil
}

type AdminGetUserRequest struct {
	AdminID uint `json:"-"`
	UserID  uint `json:"-"`
}

type AdminDeleteUserRequest struct {
	AdminID uint `json:"-"`
	UserID  uint `json:"-"`
}

type AdminUpdateUserRoleRequest struct {
	AdminID  uint   `json:"-" validate:"required"`
	UserID   uint   `json:"-" validate:"required"`
	RoleName string `json:"role" validate:"required"`
}

type AdminUpdateUserRequest struct {
	AdminID  uint    `json:"-"`
	UserID   uint    `json:"-"`
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UploadAssetRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadAssetResult struct {
	URL string `json:"url"`
}

type SocialAuthRequest struct {
	Type  string `json:"type" validate:"required"`
	Token string `json:"token" validate:"required"`
	Env   string `json:"env"`
}

type AuthInvitationRequest struct {
	Token string `json:"invitation_token" validate:"required"`
}

type CalendarDateTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}
