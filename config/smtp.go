package config

import (
	"os"

	"go_grpc/lib"
)

func NewSMTPClient() lib.SMTPClient {
	return lib.SMTPClient{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Username: os.Getenv("SMTP_USERNAME"),
		Sender:   os.Getenv("SMTP_SENDER"),
	}
}
