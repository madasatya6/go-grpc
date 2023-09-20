package social

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go_grpc/lib/logger"
)

type LinkedinUserDetails struct {
	Email string `json:"emailAddress"`
}

type LinkedinEmailResponse struct {
	Elements []LinkedinEmailResponseDetail `json:"elements"`
}

type LinkedinEmailResponseDetail struct {
	HandleMap LinkedinEmailAddress `json:"handle~"`
	Handle    string               `json:"handle"`
}

type LinkedinEmailAddress struct {
	EmailAddress string `json:"emailAddress"`
}

func AuthLinkedin(token string) (LinkedinUserDetails, error) {
	var userDetails LinkedinUserDetails
	var userEmail LinkedinEmailResponse
	request, _ := http.NewRequest("GET", "https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		logger.Error(context.Background(), "error validate linkedin account", map[string]interface{}{
			"error": err,
			"tags":  []string{"socialauth"},
		})
		return LinkedinUserDetails{}, err
	}

	decoder := json.NewDecoder(response.Body)
	decoderErr := decoder.Decode(&userEmail)
	defer response.Body.Close()

	if decoderErr != nil {
		logger.Error(context.Background(), "Error occurred while getting information from Linkedin", map[string]interface{}{
			"error": decoderErr,
			"tags":  []string{"socialauth"},
		})
		return LinkedinUserDetails{}, decoderErr
	}

	userDetails.Email = userEmail.Elements[0].HandleMap.EmailAddress

	return userDetails, nil
}
