package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"os"
)

type Session struct {
	Id    uint
	Name  string
	Email string
	Token string
}

func StoreSession() {
	store := cookie.NewStore(([]byte(os.Getenv("APP_SECRET"))))
	sessions.Sessions("coresession", store)

	return
}
