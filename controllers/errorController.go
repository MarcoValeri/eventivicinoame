package controllers

import (
	"net/http"
	"path"
	"text/template"
	"time"
)

type ErrorPageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
	CurrentUrl      string
}

func Error404() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/errors/error-404.html"))
	http.HandleFunc("/error/error-404", func(w http.ResponseWriter, r *http.Request) {

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := ErrorPageData{
			PageTitle:       "Error 404, pagina non trovata",
			PageDescription: "Error 404, pagina non trovata su Eventi Vicino a Me",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		tmpl.Execute(w, data)
	})
}
