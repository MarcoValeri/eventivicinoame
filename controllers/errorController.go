package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

type ErrorPageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
	CurrentUrl      string
}

var error404Template *template.Template

func init() {
	var errError404Template error
	error404Template, errError404Template = template.ParseFiles("./views/templates/base.html", "./views/news/news-search.html")
	if errError404Template != nil {
		log.Fatal("Error parsing template:", errError404Template)
	}
}

func Error404(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := error404Template

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := ErrorPageData{
			PageTitle:       "Error 404, pagina non trovata",
			PageDescription: "Error 404, pagina non trovata su Eventi Vicino a Me",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

// func Error404() {
// 	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/errors/error-404.html"))
// 	http.HandleFunc("/error/error-404", func(w http.ResponseWriter, r *http.Request) {

// 		// Get current path
// 		currentUrlPath := path.Clean(r.URL.Path)

// 		data := ErrorPageData{
// 			PageTitle:       "Error 404, pagina non trovata",
// 			PageDescription: "Error 404, pagina non trovata su Eventi Vicino a Me",
// 			CurrentYear:     time.Now().Year(),
// 			CurrentUrl:      currentUrlPath,
// 		}

// 		tmpl.Execute(w, data)
// 	})
// }
