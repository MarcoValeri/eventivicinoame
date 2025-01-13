package session

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
