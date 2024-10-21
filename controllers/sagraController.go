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
	Sagra               models.SagraWithRelatedFields
	Sagre               []models.SagraWithRelatedFields
	SagraContentRaw     template.HTML
	CurrentUrl          string
	CurrentYear         int
}

func SagreSearchController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-cerca.html"))
	http.HandleFunc("/sagre-cerca/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/sagre-cerca/")
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
					http.Redirect(w, r, "/sagre-cerca/", http.StatusSeeOther)
					break
				}
			}

			if isFormSubmmitionValid {
				// Get sagre by search parameter
				redirectURL := "/sagre-cerca/" + getSagraSearchParameterTitle
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
		* Redirect to 404 page if the
		* sagra does not exist
		* or
		* if it is not published yet
		 */
		if getSagra.Id == 0 {
			http.Redirect(w, r, "/error/error-404", http.StatusSeeOther)
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

func SagreOctober() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-ottobre.html"))
	http.HandleFunc("/sagre/sagre-ottobre", func(w http.ResponseWriter, r *http.Request) {
		// Get Sagre that are planned in October
		getOctoberSagre, err := models.SagreGetThemByPeriodOfTime("2024-10-01 00:00:00", "2024-10-31 23:59:59", 50)
		if err != nil {
			fmt.Println("Error getting October's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre ottobre 2024: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre ottobre 2024: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di ottobre, nel pieno della stagione autunnale"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getOctoberSagre,
		}

		tmpl.Execute(w, data)

	})
}

func SagreNovember() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-novembre.html"))
	http.HandleFunc("/sagre/sagre-novembre", func(w http.ResponseWriter, r *http.Request) {
		// Get Sagre that are planned in November
		getNovemberSagre, err := models.SagreGetThemByPeriodOfTime("2024-11-01 00:00:00", "2024-11-30 23:59:59", 50)
		if err != nil {
			fmt.Println("Error getting November's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre novembre 2024: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre novembre 2024: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di novembre, nel pieno della stagione autunnale"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getNovemberSagre,
		}

		tmpl.Execute(w, data)
	})
}

func SagreDecember() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-dicembre.html"))
	http.HandleFunc("/sagre/sagre-dicembre", func(w http.ResponseWriter, r *http.Request) {
		// Get Sagre that are planned in December
		getDecemberSagre, err := models.SagreGetThemByPeriodOfTime("2024-12-01 00:00:00", "2024-12-31 23:59:59", 50)
		if err != nil {
			fmt.Println("Error getting December's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre dicembre 2024: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre dicembre 2024: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di dicembre, tra la stagione autunno e inverno"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getDecemberSagre,
		}

		tmpl.Execute(w, data)
	})
}

func SagreAutumn() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-autunno.html"))
	http.HandleFunc("/sagre/sagre-autunno", func(w http.ResponseWriter, r *http.Request) {
		// Get Sagre that are planned in Autumn
		getAutumnSagre, err := models.SagreGetThemByPeriodOfTime("2024-09-22 00-00-00", "2024-12-21 23-59-59", 50)
		if err != nil {
			fmt.Println("Error getting December's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre autunno 2024: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre autunno 2024: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante la stagione autunnale, tra il mese di settembre e di dicembre"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getAutumnSagre,
		}

		tmpl.Execute(w, data)

	})
}
