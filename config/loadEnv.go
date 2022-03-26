package config

import (
	"goproject/pkg/constants"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("urban-auto"))

func GetSession(r *http.Request, cookieName string) (*sessions.Session, error) {

	return store.Get(r, constants.USER_COOKIE_NAME)
}

func UnsetSession() {

	store.MaxAge(-1)
}

func SetSession() {

	store.MaxAge(0)
}
