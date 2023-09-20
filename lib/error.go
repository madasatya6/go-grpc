package lib

import (
	"net/http"
)

type CustomError struct {
	Message  string
	Field    string
	Code     int
	HTTPCode int
}

func (err CustomError) Error() string {
	return err.Message
}

var (
	ErrorForbidden = CustomError{
		Message:  "Forbidden",
		Code:     1000,
		HTTPCode: http.StatusForbidden,
	}

	ErrorInvalidParameter = CustomError{
		Message:  "Invalid Parameter",
		Code:     1001,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorNotFound = CustomError{
		Message:  "Not Found",
		Code:     1002,
		HTTPCode: http.StatusNotFound,
	}

	ErrorInvalidOTPCode = CustomError{
		Message:  "Invalid OTP",
		Code:     1003,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorExpiredOTPCode = CustomError{
		Message:  "OTP Has Been Expired",
		Code:     1004,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorInvalidVerificationToken = CustomError{
		Message:  "Invalid Verification Token",
		Code:     1005,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorUserAlreadyExist = CustomError{
		Message:  "User already exist",
		Code:     1006,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorPhoneNumberAlreadyExist = CustomError{
		Message:  "Phone number already used",
		Code:     1006,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorIdentityIDAlreadyExist = CustomError{
		Message:  "Identity id already used",
		Code:     1006,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorEmailAlreadyExist = CustomError{
		Message:  "Email already used",
		Code:     1006,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorUserNotFound = CustomError{
		Message:  "User not found",
		Code:     1007,
		HTTPCode: http.StatusNotFound,
	}

	ErrorSendOTP = CustomError{
		Message:  "Failed send OTP",
		Code:     1008,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorSessionExpired = CustomError{
		Message:  "Your session has been expired",
		Code:     1009,
		HTTPCode: http.StatusUnauthorized,
	}

	ErrorWrongPassword = CustomError{
		Message:  "Wrong password",
		Code:     1010,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorInternalServer = CustomError{
		Message:  "Something went wrong",
		Code:     1011,
		HTTPCode: http.StatusInternalServerError,
	}

	ErrorMemberHasBeenRegistered = CustomError{
		Message:  "User has been registered",
		Code:     1012,
		HTTPCode: http.StatusForbidden,
	}

	ErrorExpiredToken = CustomError{
		Message:  "Token Has Been Expired",
		Code:     1013,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorInvalidToken = CustomError{
		Message:  "Token Has Been Invalid",
		Code:     1014,
		HTTPCode: http.StatusUnprocessableEntity,
	}
	ErrorInvalidIssuer = CustomError{
		Message:  "Issuer Has Been Invalid",
		Code:     1015,
		HTTPCode: http.StatusUnprocessableEntity,
	}
	ErrorInvalidAudience = CustomError{
		Message:  "Audience Has Been Invalid",
		Code:     1016,
		HTTPCode: http.StatusUnprocessableEntity,
	}
	ErrorEmailNotVerified = CustomError{
		Message:  "Email/Phone Number Not Verified",
		Code:     1111,
		HTTPCode: http.StatusForbidden,
	}
	ErrorShouldRegisterUsingLink = CustomError{
		Message:  "You should register using link from meeting invitation",
		Code:     1112,
		HTTPCode: http.StatusForbidden,
	}

	ErrorIdentityNotFound = CustomError{
		Message:  "Please upload your identity on profile to get the qr code",
		Code:     1113,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorUnauthorized = CustomError{
		Message:  "Unauthorized",
		Code:     1000,
		HTTPCode: http.StatusUnauthorized,
	}

)

func InvalidParameterError(field, message string) error {
	return CustomError{
		Message:  message,
		Field:    field,
		Code:     1001,
		HTTPCode: http.StatusUnprocessableEntity,
	}
}
