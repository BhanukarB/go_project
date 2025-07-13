package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts api key from headers of an http request
// example :
// Authorization: ApiKey{insert apikey here}
func GetAPIKey(headers http.Header)(string, error){
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing api key")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2{
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey"{
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}