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

// Cache the templates
var sagreSearchTemplate *template.Template
var sagraTemplate *template.Template
var sagreJanueryTemplate *template.Template
var sagreFebruaryTemplate *template.Template
var sagreOctoberTemplate *template.Template
var sagreNovemberTemplate *template.Template
var sagreDecemberTemplate *template.Template
var sagreAutumnTemplate *template.Template

func init() {
	var errSagreSearchTemplate error
	sagreSearchTemplate, errSagreSearchTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-cerca.html")
	if errSagreSearchTemplate != nil {
		log.Fatal("Error parsing template:", errSagreSearchTemplate)
	}

	var errSagraTemplate error
	sagraTemplate, errSagraTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagra.html")
	if errSagraTemplate != nil {
		log.Fatal("Error parsing template:", errSagraTemplate)
	}

	var errSagreJanueryTemplate error
	sagreJanueryTemplate, errSagreJanueryTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-gennaio.html")
	if errSagreJanueryTemplate != nil {
		log.Fatal("Error parsing template:", errSagreJanueryTemplate)
	}

	var errSagreFebruaryTemplate error
	sagreFebruaryTemplate, errSagreFebruaryTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-febbraio.html")
	if errSagreFebruaryTemplate != nil {
		log.Fatal("Error parsing template:", errSagreFebruaryTemplate)
	}

	var errSagreOctoberTemplate error
	sagreOctoberTemplate, errSagreOctoberTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-ottobre.html")
	if errSagreOctoberTemplate != nil {
		log.Fatal("Error parsing template:", errSagreOctoberTemplate)
	}

	var errNovemberTemplate error
	sagreNovemberTemplate, errNovemberTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-novembre.html")
	if errNovemberTemplate != nil {
		log.Fatal("Error parsing template:", errNovemberTemplate)
	}

	var errDecemberTemplate error
	sagreDecemberTemplate, errDecemberTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-dicembre.html")
	if errDecemberTemplate != nil {
		log.Fatal("Error parsing template:", errDecemberTemplate)
	}

	var errorAutumnTemplate error
	sagreAutumnTemplate, errorAutumnTemplate = template.ParseFiles("./views/templates/base.html", "./views/sagre/sagre-autunno.html")
	if errorAutumnTemplate != nil {
		log.Fatal("Error parsing template:", errorAutumnTemplate)
	}
}

func SagreSearchController(w http.ResponseWriter, r *http.Request) {

	data := SagraData{
		PageTitle:       "Sagre oggi vicino a me, cerca l'evento nella tua zona",
		PageDescription: "Sagre oggi vicino a me, cerca l'evento nella tua zona per tipologia, nome, città, comune, paese e frazione, disponibili le sagre, le fiere e le feste",
		CurrentYear:     time.Now().Year(),
	}

	if r.Method == http.MethodGet {
		tmpl := sagreSearchTemplate
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		urlPath := strings.TrimPrefix(r.URL.Path, "/sagre-cerca/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)
		data.CurrentUrl = currentUrlPath

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

		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func SagraController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl := sagraTemplate

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
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func SagreJanuary(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := sagreJanueryTemplate

		// Get Sagre that are planned in January
		setMonth := 1 // MM January
		getJanuarySagre, err := models.SagreGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		if err != nil {
			fmt.Println("Error getting January's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre gennaio 2025: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre gennaio 2025: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di gennaio, nel pieno della stagione invernale"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getJanuarySagre,
		}

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func SagreFebruary(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := sagreFebruaryTemplate

		// Get Sagre that are planned in February
		setMonth := 2 // MM February
		getFebruarySagre, err := models.SagreGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		if err != nil {
			fmt.Println("Error getting January's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre febbraio 2025: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre febbraio 2025: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di febbraio, nel pieno della stagione invernale"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getFebruarySagre,
		}

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func SagreOctober(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := sagreOctoberTemplate

		// Get Sagre that are planned in October
		setMonth := 10 // MM October
		getOctoberSagre, err := models.SagreGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		if err != nil {
			fmt.Println("Error getting October's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre ottobre 2025: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre ottobre 2025: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di ottobre, nel pieno della stagione autunnale"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getOctoberSagre,
		}

		tmpl.Execute(w, data)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func SagreNovember(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := sagreNovemberTemplate

		// Get Sagre that are planned in November
		setMonth := 11 // MM November
		getNovemberSagre, err := models.SagreGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		if err != nil {
			fmt.Println("Error getting November's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre novembre 2025: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre novembre 2025: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di novembre, nel pieno della stagione autunnale"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getNovemberSagre,
		}

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func SagreDecember(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := sagreDecemberTemplate

		// Get Sagre that are planned in December
		setMonth := 12 // MM December
		getDecemberSagre, err := models.SagreGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		if err != nil {
			fmt.Println("Error getting December's sagre:", err)
		}

		data := SagraData{
			PageTitle:       template.HTML("Sagre dicembre 2025: fiere, feste ed eventi in Italia"),
			PageDescription: template.HTML("Sagre dicembre 2025: fiere, feste ed eventi da non perdere che si svolgono in tutta Italia durante il mese di dicembre, tra la stagione autunno e inverno"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/sagre-cerca",
			Sagre:           getDecemberSagre,
		}

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func SagreAutumn(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl := sagreAutumnTemplate

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
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
