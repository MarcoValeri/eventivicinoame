package controllers

import (
	"html/template"
	"net/http"
	"time"
)

type PageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
	CurrentUrl      string
}

func CookiePolicy() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/pages/cookie-policy.html"))
	http.HandleFunc("/pages/cookie-policy", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle:       "Cookie Policy di Eventi Vicino a Me",
			PageDescription: "Cookie Policy di Eventi Vicino a Me per gli utenti",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/pages/cookie-policy",
		}
		tmpl.Execute(w, data)
	})
}

func PrivacyPolicy() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/pages/privacy-policy.html"))
	http.HandleFunc("/pages/privacy-policy", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle:       "Privacy Policy di Eventi Vicino a Me",
			PageDescription: "Privacy Policy di Eventi Vicino a Me per gli utenti",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/pages/privacy-policy",
		}
		tmpl.Execute(w, data)
	})
}
