package controllers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type SagraData struct {
	PageTitle       string
	PageDescription string
	Sagra           models.SagraWithRelatedImage
	SagraContentRaw template.HTML
	CurrentYear     int
}

func SagraController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagra.html"))
	http.HandleFunc("/sagra/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/sagra/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get Sagra by URL
		getSagra, err := models.SagraFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding sagra by URL:", err)
		}

		/**
		* Redirect to home if the
		* sagra does not exist
		* or
		* if it is not published yet
		 */
		if getSagra.Id == 0 {
			fmt.Println("Statement 1")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		// Create raw content for html template
		sagraRawContent := template.HTML(getSagra.Content)

		data := SagraData{
			PageTitle:       getSagra.Title,
			PageDescription: getSagra.Description,
			Sagra:           getSagra,
			SagraContentRaw: sagraRawContent,
			CurrentYear:     time.Now().Year(),
		}

		tmpl.Execute(w, data)

	})
}
