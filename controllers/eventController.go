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

type EventData struct {
	PageTitle           template.HTML
	PageDescription     template.HTML
	ParameterTitleError string
	ParameterTitle      string
	Event               models.EventWithRelatedFields
	Events              []models.EventWithRelatedFields
	EventContentRaw     template.HTML
	CurrentUrl          string
	CurrentYear         int
}

func EventsSearchController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/events-search.html"))
	http.HandleFunc("/eventi-cerca/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/eventi-cerca/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := EventData{
			PageTitle:       "Eventi oggi vicino a me, cerca l'evento nella tua zona",
			PageDescription: "Eventi oggi vicino a me, cerca l'evento nella tua zona per tipologia, nome, città, comune, paese e frazione, disponibili mercatini, musei, mostre e altro",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		/**
		* Check if the form for searching has been submitted
		* and
		* validate the inputs
		 */
		var areEventsSerachInputValid [1]bool
		isFormSubmitionValid := false

		// Get values from the form
		getEventsSearchParameterTitle := r.FormValue("event-search-title")
		getEventsSearchButtons := r.FormValue("event-search-button")

		// Sanitize the inputs
		getEventsSearchParameterTitle = util.FormSanitizeStringInput(getEventsSearchParameterTitle)
		getEventsSearchButtons = util.FormSanitizeStringInput(getEventsSearchButtons)

		// Check if the form has been submitted and validate the inputs
		if getEventsSearchButtons == "Cerca" {
			// Parameter title validation
			if len(getEventsSearchParameterTitle) > 0 {
				data.ParameterTitleError = ""
				areEventsSerachInputValid[0] = true
			} else {
				data.ParameterTitleError = "La tua ricerca dovrebbe essere più lunga di zero caratteri"
				areEventsSerachInputValid[0] = false
			}

			for i := 0; i < len(areEventsSerachInputValid); i++ {
				isFormSubmitionValid = true
				if !areEventsSerachInputValid[i] {
					isFormSubmitionValid = false
					http.Redirect(w, r, "/eventi-cerca/", http.StatusSeeOther)
					break
				}
			}

			if isFormSubmitionValid {
				// Get events by search parameter
				redirectURL := "/eventi-cerca/" + getEventsSearchParameterTitle
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			}
		} else {
			getEvents, err := models.EventsFindByParameter(urlPath)
			if err != nil {
				fmt.Println("Error getting the events by parameter:", err)
			}

			// Add data for the page
			data.ParameterTitle = urlPath
			data.Events = getEvents

			tmpl.Execute(w, data)
		}

	})
}

func EventController() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/event.html"))
	http.HandleFunc("/evento/", func(w http.ResponseWriter, r *http.Request) {

		urlPath := strings.TrimPrefix(r.URL.Path, "/evento/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		// Get Event by URL
		getEvent, err := models.EventWithRelatedFieldsFindByUrl(urlPath)
		if err != nil {
			fmt.Println("Error finding event by URL:", err)
		}

		/**
		* Redirect to 404 page if the
		* event does not exist
		* or
		* if it is not published yet
		 */
		if getEvent.Id == 0 {
			http.Redirect(w, r, "/error/error-404", http.StatusSeeOther)
		}

		// Create raw content for html template
		eventRowTitle := template.HTML(getEvent.Title)
		eventRowDescription := template.HTML(getEvent.Description)
		eventRowContent := template.HTML(getEvent.Content)

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := EventData{
			PageTitle:       eventRowTitle,
			PageDescription: eventRowDescription,
			Event:           getEvent,
			EventContentRaw: eventRowContent,
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
		}

		tmpl.Execute(w, data)
	})
}

func EventsMercatiniDiNatale() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/events-mercatini-di-natale.html"))
	http.HandleFunc("/eventi/mercatini-di-natale", func(w http.ResponseWriter, r *http.Request) {
		// Get all the events that hava event_type = mercatini di natale
		getEventsTypeMercatiniDiNatale, err := models.EventsGetByEventType("mercatini di natale", 50)
		if err != nil {
			fmt.Println("Error getting event type Marcatini di Natale:", err)
		}

		data := EventData{
			PageTitle:       template.HTML("Mercatini di Natale 2025: mercati, fiere e feste natalizie"),
			PageDescription: template.HTML("Mercatini di Natale 2025: mercati, fiere e feste natalizie in Italia, Europa e nel resto del mondo, dal più bel mercato di Natale a quello con più bancarelle"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/eventi-cerca",
			Events:          getEventsTypeMercatiniDiNatale,
		}

		tmpl.Execute(w, data)
	})

}

func EventsJanuary() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/events-gennaio.html"))
	http.HandleFunc("/eventi/eventi-gennaio", func(w http.ResponseWriter, r *http.Request) {
		// Get all the events in January
		setMonth := 1 // MM January
		getEventsJanuary, err := models.EventsGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		// getEventsNovember, err := models.EventsGetThemByPeriodOfTime("2025-01-01 00:00:00", "2025-01-31 23:59:59", 50)
		if err != nil {
			fmt.Println("Error getting November's events:", err)
		}

		data := EventData{
			PageTitle:       template.HTML("Eventi gennaio 2025: le cose più belle da fare a gennaio"),
			PageDescription: template.HTML("Eventi gennaio 2025: le cose più belle da fare in questo periodo dell'anno in Italia, Europa e resto del mondo, pianifica il tuo evento in inverno"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/eventi-cerca",
			Events:          getEventsJanuary,
		}

		tmpl.Execute(w, data)
	})
}

func EventsFebruary() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/events-febbraio.html"))
	http.HandleFunc("/eventi/eventi-febbraio", func(w http.ResponseWriter, r *http.Request) {
		// Get all the events in February
		setMonth := 2 // MM February
		getEventsFebruary, err := models.EventsGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		if err != nil {
			fmt.Println("Error getting February's events:", err)
		}

		data := EventData{
			PageTitle:       template.HTML("Eventi febbraio 2025: le cose più belle da fare a febbraio"),
			PageDescription: template.HTML("Eventi febbraio 2025: le cose più belle da fare in questo periodo dell'anno in Italia, Europa e resto del mondo, pianifica il tuo evento in inverno"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/eventi-cerca",
			Events:          getEventsFebruary,
		}

		tmpl.Execute(w, data)
	})
}

func EventsNovember() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/events-novembre.html"))
	http.HandleFunc("/eventi/eventi-novembre", func(w http.ResponseWriter, r *http.Request) {
		// Get all the events in November
		setMonth := 11 // MM November
		getEventsNovember, err := models.EventsGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		// getEventsNovember, err := models.EventsGetThemByPeriodOfTime("2024-11-01 00:00:00", "2024-11-30 23:59:59", 50)
		if err != nil {
			fmt.Println("Error getting November's events:", err)
		}

		data := EventData{
			PageTitle:       template.HTML("Eventi novembre 2024: le cose più belle da fare a novembre"),
			PageDescription: template.HTML("Eventi novembre 2024: le cose più belle da fare in questo periodo dell'anno in Italia, Europa e resto del mondo, pianifica il tuo evento in autunno"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/eventi-cerca",
			Events:          getEventsNovember,
		}

		tmpl.Execute(w, data)
	})
}

func EventsDecember() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/events/events-dicembre.html"))
	http.HandleFunc("/eventi/eventi-dicembre", func(w http.ResponseWriter, r *http.Request) {
		// Get all the events in December
		setMonth := 12 // MM November
		getEventsDecember, err := models.EventsGetThemByPeriodOfTimeWithoutYear(setMonth, 50)
		// getEventsDecember, err := models.EventsGetThemByPeriodOfTime("2024-12-01 00:00:00", "2024-12-31 23:59:59", 50)
		if err != nil {
			fmt.Println("Error getting December's events:", err)
		}

		data := EventData{
			PageTitle:       template.HTML("Eventi dicembre 2025: le cose più belle da fare a dicembre"),
			PageDescription: template.HTML("Eventi dicembre 2025: le cose più belle da fare in questo periodo dell'anno in Italia, Europa e resto del mondo, pianifica il tuo evento in autunno e inverno"),
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      "/eventi-cerca",
			Events:          getEventsDecember,
		}

		tmpl.Execute(w, data)
	})

}
