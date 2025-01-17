package controllers

import (
	"html/template"
	"log"
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

var aboutUsTemplate *template.Template
var contactTemplate *template.Template
var cookiePolicyTemplate *template.Template
var privacyPolicyTemplate *template.Template

func init() {
	var errAboutUs error
	aboutUsTemplate, errAboutUs = template.ParseFiles("./views/templates/base.html", "./views/pages/about-us.html")
	if errAboutUs != nil {
		log.Fatal("Error parsing template:", errAboutUs)
	}

	var errContact error
	contactTemplate, errContact = template.ParseFiles("./views/templates/base.html", "./views/pages/cookie-policy.html")
	if errContact != nil {
		log.Fatal("Error parsing template:", errContact)
	}

	var errCookiePolicyTemplate error
	cookiePolicyTemplate, errCookiePolicyTemplate = template.ParseFiles("./views/templates/base.html", "./views/pages/cookie-policy.html")
	if errCookiePolicyTemplate != nil {
		log.Fatal("Error parsing template:", errCookiePolicyTemplate)
	}

	var errPrivacyPolicyTemplate error
	privacyPolicyTemplate, errPrivacyPolicyTemplate = template.ParseFiles("./views/templates/base.html", "./views/pages/privacy-policy.html")
	if errPrivacyPolicyTemplate != nil {
		log.Fatal("Error parsing template:", errPrivacyPolicyTemplate)
	}
}

func AboutUs(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := aboutUsTemplate

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Eventi Vicino a Me, pagina chi siamo",
			PageDescription: "Eventi Vicino a Me, pagina chi siamo, con gli obiettivi del progetto",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func Contact(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := contactTemplate

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Eventi Vicino a Me, pagina contatti",
			PageDescription: "Eventi Vicino a Me, pagina ufficiale per contattare la redazione",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func CookiePolicy(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := cookiePolicyTemplate

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Cookie Policy di Eventi Vicino a Me",
			PageDescription: "Cookie Policy di Eventi Vicino a Me per gli utenti",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func PrivacyPolicy(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := privacyPolicyTemplate

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := PageData{
			PageTitle:       "Privacy Policy di Eventi Vicino a Me",
			PageDescription: "Privacy Policy di Eventi Vicino a Me per gli utenti",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}
		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
