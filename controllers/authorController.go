package controllers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

type AuthorData struct {
	PageTitle            template.HTML
	PageDescription      template.HTML
	Author               models.Author
	AuthorRowName        template.HTML
	AuthorRowSurname     template.HTML
	AuthorRowDescription template.HTML
	CurrentUrl           string
	CurrentYear          int
}

var authorTemplate *template.Template

func init() {
	var errAuthorTemplate error
	authorTemplate, errAuthorTemplate = template.ParseFiles("./views/templates/base.html", "./views/authors/author.html")
	if errAuthorTemplate != nil {
		log.Fatal("Error parsing template:", errAuthorTemplate)
	}
}

func AuthorController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := authorTemplate

		urlPath := strings.TrimPrefix(r.URL.Path, "/author/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get author by URL
		getAuthor, err := models.AuthorFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding author by URL:", err)
		}

		/**
		* Redirect to 404 page it the
		* author does not exists
		 */
		if getAuthor.Id == 0 {
			http.Redirect(w, r, "/error/error-404", http.StatusSeeOther)
		}

		// Create raw content for html template
		authorRowName := template.HTML(getAuthor.Name)
		authorRowSurname := template.HTML(getAuthor.Surname)
		authorRowDescription := template.HTML(getAuthor.Description)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := AuthorData{
			PageTitle:            "Autore " + authorRowName + authorRowSurname,
			PageDescription:      "Autore " + authorRowName + authorRowSurname + ": biografia e lista degli ultimi articoli scritti su Eventi Vicino a Me",
			Author:               getAuthor,
			AuthorRowName:        authorRowName,
			AuthorRowSurname:     authorRowSurname,
			AuthorRowDescription: authorRowDescription,
			CurrentYear:          time.Now().Year(),
			CurrentUrl:           currentUrlPath,
		}

		tmpl.Execute(w, data)
	}

}
