package lib

const (
	VerificationRegistrationType   = 0
	VerificationForgotPasswordType = 1

	VerificationDeliveryByEmailType = 0
	VerificationDeliveryByPhoneType = 1

	VerificationPendingState       = 0
	VerificationCodeConfirmedState = 1
	VerificationVerifiedState      = 2

	RoleAdminID        = 1
	RoleHostID         = 2
	RoleGuestID        = 3
	RoleSuperAdminID   = 4

	WebsocketRedisTopic = "upsert_meeting"
)
