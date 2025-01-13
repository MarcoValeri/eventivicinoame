package controllers

import (
	"eventivicinoame/session"
	"fmt"
	"net/http"
)

func AuthMiddlewareController(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, errSession := session.Store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error AuthMiddlewareController on session-authentication:", errSession)
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
			return
		}

		if session.Values["admin-user-authentication"] == true {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
