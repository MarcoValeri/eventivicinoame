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
			PageTitle:       "Eventi Vicino a Me News: novità e notizie su cosa fare",
			PageDescription: "Eventi Vicino a Me News: novità, notizie e aggiornamenti su cosa fare in Italia, in Europa e nel resto del mondo, tra eventi, feste e tempo libero",
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

func NewsController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/news/news.html"))
	http.HandleFunc("/news/", func(w http.ResponseWriter, r *http.Request) {
		urlPath := strings.TrimPrefix(r.URL.Path, "/news/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get News by URL
		getNews, err := models.NewsWithRelatedFieldsFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding news by URL:", err)
		}

		/**
		* Redirect to 404 page if the
		* news does not exist
		* or
		* if it is not published yet
		 */
		if getNews.Id == 0 {
			http.Redirect(w, r, "/error/error-404", http.StatusSeeOther)
		}

		// Create raw content for html template
		newsRowTitle := template.HTML(getNews.Title)
		newsRowDescription := template.HTML(getNews.Description)
		newsRowContent := template.HTML(getNews.Content)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := NewsData{
			PageTitle:       newsRowTitle,
			PageDescription: newsRowDescription,
			SingleNews:      getNews,
			NewsContentRaw:  newsRowContent,
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		tmpl.Execute(w, data)
	})
}
