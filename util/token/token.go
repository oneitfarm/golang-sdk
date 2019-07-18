package token

import (
	"fmt"
	"net/http"
	"regexp"
)

var regex = regexp.MustCompile("^Bearer\\s+(\\S+)$")

func GetBearerToken(r *http.Request) (string, bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", false
	}

	return GetBearerFromAuth(auth)
}

func GetBearerFromAuth(auth string) (string, bool) {
	result := regex.FindStringSubmatch(auth)
	fmt.Println(result)
	if len(result) < 2 {
		return "", false
	}

	return result[1], true
}
