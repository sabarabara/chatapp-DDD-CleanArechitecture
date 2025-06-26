package config

import (
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func InitSessionStore(secretKey string) {
	Store = sessions.NewCookieStore([]byte(secretKey))

	
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 1週間
		HttpOnly: true,
		Secure:   true, // HTTPS時のみ
	}
}
