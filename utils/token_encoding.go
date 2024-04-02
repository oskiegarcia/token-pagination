package utils

import (
	"encoding/base64"
	"strconv"
)

// DecodeToken is a function to decode a token
func DecodeToken(token string) (int, error) {
	if token == "" {
		return 0, nil
	}

	decodedToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	offset, err := strconv.Atoi(string(decodedToken))
	if err != nil {
		return 0, err
	}

	return offset, nil
}

// EncodeToken is a function to encode a token
func EncodeToken(offset int) string {
	return base64.URLEncoding.EncodeToString([]byte(strconv.Itoa(offset)))
}
