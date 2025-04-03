package initializers

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func InitOAuth() {

	sessionKey := os.Getenv("SESSION_KEY")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackUrl := os.Getenv("CALLBACK_URL")
	// This is our JWT
	store := sessions.NewCookieStore([]byte(sessionKey))
	store.Options.Path = "/"
	store.MaxAge(86400 * 30) // 30 days
	store.Options.Secure = true
	store.Options.HttpOnly = true
	gothic.Store = store
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, callbackUrl, "email"),
	)
}
