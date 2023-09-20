package lib

import (
	"context"

	fcm "firebase.google.com/go/messaging"
	"go_grpc/lib/logger"
)

type Messaging struct {
	VerihubsClient VerihubsClient
	FirebaseClient *fcm.Client
}

type PushNotifRequest struct {
	Token    string            `json:"token"`
	Title    string            `json:"title"`
	Body     string            `json:"body"`
	ImageURL string            `json:"image_url"`
	Data     map[string]string `json:"data"`
}

func (messaging *Messaging) SendWhatsapp(ctx context.Context, destination string, otp string) error {
	return messaging.VerihubsClient.SendWhatsapp(ctx, destination, otp)
}

func (messaging *Messaging) SendWhatsappInvitationVip(ctx context.Context, destination string, content []string) error {
	return messaging.VerihubsClient.SendWhatsappInvitationVip(ctx, destination, content)
}

func (messaging *Messaging) SendPushNotification(ctx context.Context, req PushNotifRequest) error {
	notification := &fcm.Notification{
		Title:    req.Title,
		Body:     req.Body,
		ImageURL: req.ImageURL,
	}

	res, err := messaging.FirebaseClient.Send(ctx, &fcm.Message{Token: req.Token, Data: req.Data, Notification: notification, Android: &fcm.AndroidConfig{Priority: "high"}})
	if err != nil {
		logger.Error(ctx, "error send firebase push notification", map[string]interface{}{
			"error": err,
			"tags":  []string{"firebase", "push-notif"},
		})

		return err
	}

	logger.Info(ctx, "success send firebase push notification", map[string]interface{}{
		"token":        req.Token,
		"notification": notification,
		"data":         req.Data,
		"result":       res,
		"tags":         []string{"firebase", "push-notif"},
	})

	return nil
}
