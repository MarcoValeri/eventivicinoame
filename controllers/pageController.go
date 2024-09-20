package controllers

import (
	"html/template"
	"net/http"
	"path"
	"time"
)

type PageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
	CurrentUrl      string
}

func AboutUs() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/pages/about-us.html"))
	http.HandleFunc("/page/chi-siamo", func(w http.ResponseWriter, r *http.Request) {

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Eventi Vicino a Me, pagina chi siamo",
			PageDescription: "Eventi Vicino a Me, pagina chi siamo, con gli obiettivi del progetto",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	})
}

func Contact() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/pages/contact.html"))
	http.HandleFunc("/page/contatti", func(w http.ResponseWriter, r *http.Request) {

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Eventi Vicino a Me, pagina contatti",
			PageDescription: "Eventi Vicino a Me, pagina ufficiale per contattare la redazione",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	})
}

func CookiePolicy() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/pages/cookie-policy.html"))
	http.HandleFunc("/page/cookie-policy", func(w http.ResponseWriter, r *http.Request) {

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Cookie Policy di Eventi Vicino a Me",
			PageDescription: "Cookie Policy di Eventi Vicino a Me per gli utenti",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	})
}

func PrivacyPolicy() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/pages/privacy-policy.html"))
	http.HandleFunc("/page/privacy-policy", func(w http.ResponseWriter, r *http.Request) {

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Privacy Policy di Eventi Vicino a Me",
			PageDescription: "Privacy Policy di Eventi Vicino a Me per gli utenti",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	})
}
