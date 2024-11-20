package controllers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
	"time"
)

type NewsData struct {
	PageTitle           template.HTML
	PageDescription     template.HTML
	ParameterTitleError string
	ParameterTitle      string
	SingleNews          models.NewsWithRelatedFields
	News                []models.NewsWithRelatedFields
	NewsContentRaw      template.HTML
	CurrentUrl          string
	CurrentYear         int
}

func NewsSearchController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/news/news-search.html"))
	http.HandleFunc("/news-cerca/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.TrimPrefix(r.URL.Path, "/news-cerca/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := NewsData{
			PageTitle:       "ADD TITLE",
			PageDescription: "ADD DESCRIPTION",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		/**
		* Check if the form for searching has been submitted
		* and
		* validate the inputs
		 */
		var areNewsSearchInputValid [1]bool
		isFormSubmitted := false

		// Get values from the form
		getNewsSearchParameterTitle := r.FormValue("news-search-title")
		getNewsSearchButtons := r.FormValue("news-search-button")

		// Sanitize the inputs
		getNewsSearchParameterTitle = util.FormSanitizeStringInput(getNewsSearchParameterTitle)
		getNewsSearchButtons = util.FormSanitizeStringInput(getNewsSearchButtons)

		// Check if the form has been submitted and validate the inputs
		if getNewsSearchButtons == "Cerca" {
			// Parameter title validation
			if len(getNewsSearchParameterTitle) > 0 {
				data.ParameterTitleError = ""
				areNewsSearchInputValid[0] = true
			} else {
				data.ParameterTitleError = "La tua ricerca dovrebbe essere più lunga di zero caratteri"
				areNewsSearchInputValid[0] = false
			}

			for i := 0; i < len(areNewsSearchInputValid); i++ {
				isFormSubmitted = true
				if !areNewsSearchInputValid[i] {
					isFormSubmitted = false
					http.Redirect(w, r, "/news-cerca/", http.StatusSeeOther)
					break
				}
			}

			if isFormSubmitted {
				// Get news by search parameter
				redirectURL := "/news-cerca/" + getNewsSearchParameterTitle
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			}
		} else {
			getNews, err := models.NewsFindByParameter(urlPath)
			if err != nil {
				fmt.Println("Error getting the news by parameter:", err)
			}

			// Add data for the page
			data.ParameterTitle = urlPath
			data.News = getNews

			tmpl.Execute(w, data)
		}
	})
}
