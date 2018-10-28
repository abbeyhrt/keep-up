package cookie

import (
	"net/http"
	"time"
)

const SessionKey = "keepup.sid"
const StateKey = "keepup.state"

func Create(name string, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(time.Minute * 30),
		Secure:   true,
		HttpOnly: true,
	}
}
