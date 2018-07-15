package http

import (
	"encoding/base64"
)

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetBasicAuthHeaderValue(user, passwod string) string {
	return "Basic " + BasicAuth(user, passwod)
}
