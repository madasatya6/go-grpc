package social

import (
	"context"
	"encoding/json"
	"net/http"

	"go_grpc/lib/logger"
)

type FacebookUserDetails struct {
	ID    string
	Name  string
	Email string
}

func AuthFacebook(token string) (FacebookUserDetails, error) {
	var userDetails FacebookUserDetails
	request, _ := http.NewRequest("GET", "https://graph.facebook.com/me?fields=id,name,email&access_token="+token, nil)
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		logger.Error(context.Background(), "error validate facebook account", map[string]interface{}{
			"error": err,
			"tags":  []string{"socialauth"},
		})
		return FacebookUserDetails{}, err
	}

	decoder := json.NewDecoder(response.Body)
	decoderErr := decoder.Decode(&userDetails)
	defer response.Body.Close()

	if decoderErr != nil {
		logger.Error(context.Background(), "Error occurred while getting information from Facebook", map[string]interface{}{
			"error": err,
			"tags":  []string{"socialauth"},
		})
		return FacebookUserDetails{}, err
	}

	return userDetails, nil
}
