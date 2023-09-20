package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go_grpc/lib/logger"
)

type VerihubsClient struct {
	VerihubsApplicationID string
	VerihubsAPIKey        string
}

type SendWhatsappPayload struct {
	PhoneNumber string `json:"msisdn"`
	OTP         string `json:"otp"`
}

type SendWhatsappInvitationVipPayload struct {
	Content     []string `json:"content"`
	PhoneNumber string   `json:"msisdn"`
}

func (client *VerihubsClient) SendWhatsapp(ctx context.Context, destination string, otp string) error {
	url := "https://verihubs.com/api/v1/whatsapp/otp/send"

	payload := SendWhatsappPayload{
		PhoneNumber: destination,
		OTP:         otp,
	}
	encoded, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(encoded))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("App-ID", client.VerihubsApplicationID)
	req.Header.Add("API-Key", client.VerihubsAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(ctx, "error send whatsapp", map[string]interface{}{
			"error":   err,
			"payload": encoded,
			"tags":    []string{"whatsapp"},
		})

		return err
	}

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		logger.Error(ctx, "error send whatsapp", map[string]interface{}{
			"error":    err,
			"request":  encoded,
			"response": string(body),
			"tags":     []string{"whatsapp"},
		})

		return ErrorSendOTP
	}

	return nil
}

func (client *VerihubsClient) SendWhatsappInvitationVip(ctx context.Context, destination string, content []string) error {
	url := "https://verihubs.com/api/v1/whatsapp/message/send"

	payload := SendWhatsappInvitationVipPayload{
		Content:     content,
		PhoneNumber: destination,
	}
	encoded, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(encoded))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("App-ID", client.VerihubsApplicationID)
	req.Header.Add("API-Key", client.VerihubsAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(ctx, "error send whatsapp invitaion vip", map[string]interface{}{
			"error":   err,
			"payload": encoded,
			"tags":    []string{"whatsapp"},
		})

		return err
	}

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		logger.Error(ctx, "error send whatsapp invitaion vip", map[string]interface{}{
			"error":    err,
			"request":  encoded,
			"response": string(body),
			"tags":     []string{"whatsapp"},
		})

		return ErrorInvalidOTPCode
	}

	return nil
}
