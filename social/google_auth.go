package social

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go_grpc/lib"
	"go_grpc/lib/logger"
)

// GoogleClaims -
type GoogleUserDetails struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

// ValidateGoogleJWT -
func AuthGoogle(tokenString string) (GoogleUserDetails, error) {
	claimsStruct := GoogleUserDetails{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return GoogleUserDetails{}, err
	}

	claims, ok := token.Claims.(*GoogleUserDetails)

	if !ok {
		logger.Error(context.Background(), "error claim google account", map[string]interface{}{
			"error": err,
			"tags":  []string{"socialauth"},
		})
		return GoogleUserDetails{}, lib.ErrorInvalidToken
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return GoogleUserDetails{}, lib.ErrorInvalidIssuer
	}

	if claims.Audience != os.Getenv("GOOGLE_AUDIENCE_ID") {
		return GoogleUserDetails{}, lib.ErrorInvalidAudience
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return GoogleUserDetails{}, lib.ErrorExpiredToken
	}

	return *claims, nil
}

// ValidateGoogleJWTWeb -
func AuthGoogleWeb(tokenString string, env string) (GoogleUserDetails, error) {
	claimsStruct := GoogleUserDetails{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return GoogleUserDetails{}, err
	}

	claims, ok := token.Claims.(*GoogleUserDetails)

	if !ok {
		logger.Error(context.Background(), "error claim google account", map[string]interface{}{
			"error": err,
			"tags":  []string{"socialauth"},
		})
		return GoogleUserDetails{}, lib.ErrorInvalidToken
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return GoogleUserDetails{}, lib.ErrorInvalidIssuer
	}

	if env == "local" {
		if claims.Audience != os.Getenv("GOOGLE_AUDIENCE_WEB_LOCAL_ID") {
			return GoogleUserDetails{}, lib.ErrorInvalidAudience
		}
	} else {
		if claims.Audience != os.Getenv("GOOGLE_AUDIENCE_WEB_PRODUCTION_ID") {
			return GoogleUserDetails{}, lib.ErrorInvalidAudience
		}
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return GoogleUserDetails{}, lib.ErrorExpiredToken
	}

	return *claims, nil
}
