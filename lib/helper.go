package lib

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashAndPassword(str1 string, str2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(str1), []byte(str2))
	if err != nil {
		return err
	}

	return nil
}

func GenerateOTPNumber() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}

	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func SliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
