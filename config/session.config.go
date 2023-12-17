package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore *session.Store
var StateStore *session.Store

func InitSession() {
	SessionStore = session.New(session.Config{
		Expiration:     24 * time.Hour, // 1 days
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookiePath:     "/",
	})
}

func InitStateSession() {
	StateStore = session.New(session.Config{
		KeyLookup:      "cookie:session_state",
		Expiration:     5 * time.Minute, // 5 minutes
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookiePath:     "/",
	})
}
