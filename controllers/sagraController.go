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

type SagraData struct {
	PageTitle           template.HTML
	PageDescription     template.HTML
	ParameterTitleError string
	ParameterTitle      string
	Sagra               models.SagraWithRelatedImage
	Sagre               []models.SagraWithRelatedImage
	SagraContentRaw     template.HTML
	CurrentUrl          string
	CurrentYear         int
}

func SagreSearchController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre.html"))
	http.HandleFunc("/sagre/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/sagre/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := SagraData{
			PageTitle:       "Sagre oggi vicino a me, cerca l'evento nella tua zona",
			PageDescription: "Sagre oggi vicino a me, cerca l'evento nella tua zona per tipologia, nome, città, comune, paese e frazione, disponibili le sagre, le fiere e le feste",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		/**
		* Check if the form for searching has been submitted
		* and
		* validate the inputs
		 */
		var areSagraSearchInputsValid [1]bool
		isFormSubmmitionValid := false

		// Get values from the form
		getSagraSearchParameterTitle := r.FormValue("sagre-search-title")
		getSagraSearchButton := r.FormValue("sagre-search-button")

		// Sanitize form inputs
		getSagraSearchParameterTitle = util.FormSanitizeStringInput(getSagraSearchParameterTitle)
		getSagraSearchButton = util.FormSanitizeStringInput(getSagraSearchButton)

		if getSagraSearchButton == "Cerca" {
			// Parameter title validation
			if len(getSagraSearchParameterTitle) > 0 {
				data.ParameterTitleError = ""
				areSagraSearchInputsValid[0] = true
			} else {
				data.ParameterTitleError = "La tua ricerca dovrebbe essere più lunga di zero caratteri"
				areSagraSearchInputsValid[0] = false
			}

			for i := 0; i < len(areSagraSearchInputsValid); i++ {
				isFormSubmmitionValid = true
				if !areSagraSearchInputsValid[i] {
					isFormSubmmitionValid = false
					http.Redirect(w, r, "/sagre/", http.StatusSeeOther)
					break
				}
			}

			if isFormSubmmitionValid {
				// Get sagre by search parameter
				redirectURL := "/sagre/" + getSagraSearchParameterTitle
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)

			}
		} else {
			getSagre, err := models.SagraFindByParameter(urlPath)
			if err != nil {
				fmt.Println("Error getting the sagre by parameter:", err)
			}

			// Add data for the page
			data.ParameterTitle = urlPath
			data.Sagre = getSagre

			tmpl.Execute(w, data)
		}

	})
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
		sagraRowTitle := template.HTML(getSagra.Title)
		sagraRowDescription := template.HTML(getSagra.Description)
		sagraRawContent := template.HTML(getSagra.Content)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := SagraData{
			PageTitle:       sagraRowTitle,
			PageDescription: sagraRowDescription,
			Sagra:           getSagra,
			SagraContentRaw: sagraRawContent,
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		tmpl.Execute(w, data)

	})
}
