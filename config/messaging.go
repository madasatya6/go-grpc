package config

import (
	"os"

	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"go_grpc/lib"
	"go_grpc/lib/logger"
)

func NewMessaging() *lib.Messaging {
	verihubsAppID := os.Getenv("VERIHUBS_APP_ID")
	verihubsApiKey := os.Getenv("VERIHUBS_API_KEY")
	firebaseConfigPath := os.Getenv("FIREBASE_CONFIG_PATH")

	opt := option.WithCredentialsFile(firebaseConfigPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error(context.Background(), "", map[string]interface{}{
			"error": err,
		})
	}

	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		logger.Error(context.Background(), "", map[string]interface{}{
			"error": err,
		})
	}

	return &lib.Messaging{
		VerihubsClient: lib.VerihubsClient{
			VerihubsApplicationID: verihubsAppID,
			VerihubsAPIKey:        verihubsApiKey,
		},
		FirebaseClient: fcmClient,
	}
}
