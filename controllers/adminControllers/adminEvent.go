package admincontrollers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type eventData struct {
	PageTitle               string
	PreviusButton           bool
	TitleError              string
	DescriptionError        string
	UrlError                string
	PublishedError          string
	UpdatedError            string
	ImageError              string
	AuthorError             string
	EventTypeError          string
	ContentError            string
	CountryError            string
	RegioneError            string
	CityError               string
	TownError               string
	FractionError           string
	EventStartDateError     string
	EventEndDateError       string
	EventsSearchInput       string
	EventsSearchInputError  string
	NextButton              bool
	PreviousPage            string
	NextPage                string
	Images                  []models.Image
	Authors                 []models.Author
	EventsWithRelatedFields []models.EventWithRelatedFields
	EventWithRelatedFields  models.EventWithRelatedFields
}

func AdminEvents() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-events.html"))
	http.HandleFunc("/admin/admin-events/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-events/")
			urlPath = util.FormSanitizeStringInput(urlPath)

			pageNumber, err := strconv.Atoi(urlPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			// Redirect to /admin/admin-events/1 if pageNumber is 0
			if pageNumber == 0 {
				http.Redirect(w, r, "/admin/admin-events/1", http.StatusSeeOther)
			}

			// Set limit and offset for MySQL query
			limit := 10
			offset := (pageNumber - 1) * limit

			eventsDate, err := models.EventGetLimitAndPagination(limit, offset)
			if err != nil {
				fmt.Println("Error getting eventsDate:", err)
			}

			// The previous and next buttons
			setPreviousButton := false
			var setPreviousPage int
			var setPreviousPageStr string
			if (pageNumber - 1) > 0 {
				setPreviousButton = true
				setPreviousPage = pageNumber - 1
				setPreviousPageStr = strconv.Itoa(setPreviousPage)
			}

			setNextButton := false
			var setNextPage int
			var setNextPageStr string
			if len(eventsDate) >= 10 {
				setNextButton = true
				setNextPage = pageNumber + 1
				setNextPageStr = strconv.Itoa(setNextPage)
			}

			data := eventData{
				PageTitle:               "Events Admin",
				PreviusButton:           setPreviousButton,
				NextButton:              setNextButton,
				PreviousPage:            setPreviousPageStr,
				NextPage:                setNextPageStr,
				EventsWithRelatedFields: eventsDate,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminEventAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-event-add.html"))
	http.HandleFunc("/admin/admin-event-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			imagesData, errImagesData := models.ImageShowImagesByUpdated()
			if errImagesData != nil {
				fmt.Println("Error getting imagesData:", errImagesData)
			}

			authorsData, errAuthorsData := models.AuthorShowAuthors()
			if errAuthorsData != nil {
				fmt.Println("Error getting authorsData:", errAuthorsData)
			}

			data := eventData{
				PageTitle: "Admin Event Add",
				Images:    imagesData,
				Authors:   authorsData,
			}

			// Flag validation
			var areAdminEventInputsValid [16]bool
			isFormSubmittionValid := false

			// Get the value from the form
			getAdminEventTitle := r.FormValue("event-title")
			getAdminEventDescription := r.FormValue("event-description")
			getAdminEventUrl := r.FormValue("event-url")
			getAdminEventPublished := r.FormValue("event-published")
			getAdminEventUpdated := r.FormValue("event-updated")
			getAdminEventImage := r.FormValue("event-image")
			getAdminEventAuthor := r.FormValue("event-author")
			getAdminEventType := r.FormValue("event-type")
			getAdminEventContent := r.FormValue("event-content")
			getAdminEventCountry := r.FormValue("event-country")
			getAdminEventRegion := r.FormValue("event-region")
			getAdminEventCity := r.FormValue("event-city")
			getAdminEventTown := r.FormValue("event-town")
			getAdminEventFraction := r.FormValue("event-fraction")
			getAdminEventStartDate := r.FormValue("event-start-date")
			getAdminEventEndDate := r.FormValue("event-end-date")
			getAdminEventAdd := r.FormValue("event-add")

			// Sanitize form inputs
			getAdminEventTitle = util.FormSanitizeStringInput(getAdminEventTitle)
			getAdminEventDescription = util.FormSanitizeStringInput(getAdminEventDescription)
			getAdminEventUrl = util.FormSanitizeStringInput(getAdminEventUrl)
			getAdminEventPublished = util.FormSanitizeStringInput(getAdminEventPublished)
			getAdminEventUpdated = util.FormSanitizeStringInput(getAdminEventUpdated)
			getAdminEventImage = util.FormSanitizeStringInput(getAdminEventImage)
			getAdminEventAuthor = util.FormSanitizeStringInput(getAdminEventAuthor)
			getAdminEventType = util.FormSanitizeStringInput(getAdminEventType)
			getAdminEventCountry = util.FormSanitizeStringInput(getAdminEventCountry)
			getAdminEventRegion = util.FormSanitizeStringInput(getAdminEventRegion)
			getAdminEventCity = util.FormSanitizeStringInput(getAdminEventCity)
			getAdminEventTown = util.FormSanitizeStringInput(getAdminEventTown)
			getAdminEventFraction = util.FormSanitizeStringInput(getAdminEventFraction)
			getAdminEventStartDate = util.FormSanitizeStringInput(getAdminEventStartDate)
			getAdminEventEndDate = util.FormSanitizeStringInput(getAdminEventEndDate)
			getAdminEventAdd = util.FormSanitizeStringInput(getAdminEventAdd)

			// Check if the form has been submitted
			if getAdminEventAdd == "Add new event" {

				// Title validation
				if len(getAdminEventTitle) > 0 {
					data.TitleError = ""
					areAdminEventInputsValid[0] = true
				} else {
					data.TitleError = "Title should be longer than 0"
					areAdminEventInputsValid[0] = false
				}

				// Description validation
				if len(getAdminEventDescription) > 0 {
					data.DescriptionError = ""
					areAdminEventInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be longer than 0"
					areAdminEventInputsValid[1] = false
				}

				// Url validation
				if len(getAdminEventUrl) > 0 {
					data.UrlError = ""
					areAdminEventInputsValid[2] = true
				} else {
					data.UrlError = "Url should be longer than 0"
					areAdminEventInputsValid[2] = false
				}

				// Published validation
				if len(getAdminEventPublished) > 0 {
					data.PublishedError = ""
					areAdminEventInputsValid[3] = true
				} else {
					data.PublishedError = "Add a date"
					areAdminEventInputsValid[3] = false
				}

				// Updated validation
				if len(getAdminEventUpdated) > 0 {
					data.UpdatedError = ""
					areAdminEventInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a date"
					areAdminEventInputsValid[4] = false
				}

				// Image validation
				if len(getAdminEventImage) > 0 {
					data.ImageError = ""
					areAdminEventInputsValid[5] = true
				} else {
					data.ImageError = "An image is required"
					areAdminEventInputsValid[5] = false
				}

				// Author validation
				if len(getAdminEventAuthor) > 0 {
					data.AuthorError = ""
					areAdminEventInputsValid[6] = true
				} else {
					data.AuthorError = "An author is required"
					areAdminEventInputsValid[6] = false
				}

				// Event Type validation
				if len(getAdminEventType) > 0 {
					data.EventTypeError = ""
					areAdminEventInputsValid[7] = true
				} else {
					data.EventTypeError = "An event type is required"
					areAdminEventInputsValid[7] = false
				}

				// Content validation
				if len(getAdminEventContent) > 0 {
					data.ContentError = ""
					areAdminEventInputsValid[8] = true
				} else {
					data.ContentError = "Content should be longer than 0"
					areAdminEventInputsValid[8] = false
				}

				// Country validation
				if len(getAdminEventCountry) > 0 {
					data.CountryError = ""
					areAdminEventInputsValid[9] = true
				} else {
					data.CountryError = "Country should be longer than 0"
					areAdminEventInputsValid[9] = false
				}

				// Region validation
				if len(getAdminEventRegion) > 0 {
					data.RegioneError = ""
					areAdminEventInputsValid[10] = true
				} else {
					data.RegioneError = "Region should be longer than 0"
					areAdminEventInputsValid[10] = false
				}

				// City validation
				if len(getAdminEventCity) > 0 {
					data.CityError = ""
					areAdminEventInputsValid[11] = true
				} else {
					data.CityError = "City should be longer than 0"
					areAdminEventInputsValid[11] = false
				}

				// Town validation
				if len(getAdminEventTown) > 0 {
					data.TownError = ""
					areAdminEventInputsValid[12] = true
				} else {
					data.TownError = "Town should be longer than 0"
					areAdminEventInputsValid[12] = false
				}

				// Fraction validation
				if len(getAdminEventFraction) > 0 {
					data.FractionError = ""
					areAdminEventInputsValid[13] = true
				} else {
					data.FractionError = "Fraction should be longer than 0"
					areAdminEventInputsValid[13] = false
				}

				// Event Start Date validation
				if len(getAdminEventStartDate) > 0 {
					data.EventStartDateError = ""
					areAdminEventInputsValid[14] = true
				} else {
					data.EventStartDateError = "A date is required"
					areAdminEventInputsValid[14] = false
				}

				// Event End Date validation
				if len(getAdminEventEndDate) > 0 {
					data.EventEndDateError = ""
					areAdminEventInputsValid[15] = true
				} else {
					data.EventEndDateError = "A date is required"
					areAdminEventInputsValid[15] = false
				}

				// Check if all fields are valid
				for i := 0; i < len(areAdminEventInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminEventInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Create a new event if all inputs are valid
				if isFormSubmittionValid {

					// Get image id for the relationship one-to-many between events and images
					getAdminEventImageId, _ := models.ImageFindByUrlReturnItsId(getAdminEventImage)

					// Get author id for the relationship one-to-many between events and authors
					getAdminEventAuthorId, _ := models.AuthorFindByUrlReturnItsId(getAdminEventAuthor)

					createNewEvent := models.EventNew(
						1,
						getAdminEventTitle,
						getAdminEventDescription,
						getAdminEventUrl,
						getAdminEventPublished,
						getAdminEventUpdated,
						getAdminEventImageId,
						getAdminEventAuthorId,
						getAdminEventType,
						getAdminEventContent,
						getAdminEventCountry,
						getAdminEventRegion,
						getAdminEventCity,
						getAdminEventTown,
						getAdminEventFraction,
						getAdminEventStartDate,
						getAdminEventEndDate,
					)
					models.EventAddNewToDB(createNewEvent)
					http.Redirect(w, r, "/admin/admin-events/1", http.StatusSeeOther)
				}
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminEventEdit() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-event-edit.html"))
	http.HandleFunc("/admin/admin-event-edit/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			idPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-event-edit/")
			idPath = util.FormSanitizeStringInput(idPath)

			eventId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getEventEdit, err := models.EventWithRelatedFieldsFindById(eventId)
			if err != nil {
				fmt.Println("Error to find this event:", err)
				return
			}

			imagesData, errImagesData := models.ImageShowImagesByUpdated()
			if errImagesData != nil {
				fmt.Println("Error getting imagesData:", errImagesData)
			}

			authorsData, errAuthorsData := models.AuthorShowAuthors()
			if errAuthorsData != nil {
				fmt.Println("Error getting authorsData:", errAuthorsData)
			}

			// Create data for the page
			data := eventData{
				PageTitle:              "Admin Event Edit",
				EventWithRelatedFields: getEventEdit,
				Images:                 imagesData,
				Authors:                authorsData,
			}

			/**
			* Check if the form for editing the event has been submitted
			* and
			* validate the inputs
			 */
			var areAdminEventEditInputsValid [16]bool
			isFormSubmittionValid := false

			// Get the value from the form
			getAdminEventEditTitle := r.FormValue("event-edit-title")
			getAdminEventEditDescription := r.FormValue("event-edit-description")
			getAdminEventEditUrl := r.FormValue("event-edit-url")
			getAdminEventEditPublished := r.FormValue("event-edit-published")
			getAdminEventEditUpdated := r.FormValue("event-edit-updated")
			getAdminEventEditImage := r.FormValue("event-edit-image")
			getAdminEventEditAuthor := r.FormValue("event-edit-author")
			getAdminEventEditType := r.FormValue("event-edit-type")
			getAdminEventEditContent := r.FormValue("event-edit-content")
			getAdminEventEditCountry := r.FormValue("event-edit-country")
			getAdminEventEditRegion := r.FormValue("event-edit-region")
			getAdminEventEditCity := r.FormValue("event-edit-city")
			getAdminEventEditTown := r.FormValue("event-edit-town")
			getAdminEventEditFraction := r.FormValue("event-edit-fraction")
			getAdminEventEditStartDate := r.FormValue("event-edit-start-date")
			getAdminEventEditEndDate := r.FormValue("event-edit-end-date")
			getAdminEventEditAdd := r.FormValue("event-edit")
			getAdminEventEditAndExit := r.FormValue("event-edit-and-exit")

			// Sanitize form inputs
			getAdminEventEditTitle = util.FormSanitizeStringInput(getAdminEventEditTitle)
			getAdminEventEditDescription = util.FormSanitizeStringInput(getAdminEventEditDescription)
			getAdminEventEditUrl = util.FormSanitizeStringInput(getAdminEventEditUrl)
			getAdminEventEditPublished = util.FormSanitizeStringInput(getAdminEventEditPublished)
			getAdminEventEditUpdated = util.FormSanitizeStringInput(getAdminEventEditUpdated)
			getAdminEventEditImage = util.FormSanitizeStringInput(getAdminEventEditImage)
			getAdminEventEditAuthor = util.FormSanitizeStringInput(getAdminEventEditAuthor)
			getAdminEventEditType = util.FormSanitizeStringInput(getAdminEventEditType)
			getAdminEventEditCountry = util.FormSanitizeStringInput(getAdminEventEditCountry)
			getAdminEventEditRegion = util.FormSanitizeStringInput(getAdminEventEditRegion)
			getAdminEventEditCity = util.FormSanitizeStringInput(getAdminEventEditCity)
			getAdminEventEditTown = util.FormSanitizeStringInput(getAdminEventEditTown)
			getAdminEventEditFraction = util.FormSanitizeStringInput(getAdminEventEditFraction)
			getAdminEventEditStartDate = util.FormSanitizeStringInput(getAdminEventEditStartDate)
			getAdminEventEditEndDate = util.FormSanitizeStringInput(getAdminEventEditEndDate)
			getAdminEventEditAdd = util.FormSanitizeStringInput(getAdminEventEditAdd)
			getAdminEventEditAndExit = util.FormSanitizeStringInput(getAdminEventEditAndExit)

			// Check if the form has been submitted
			if getAdminEventEditAdd == "Edit this event" || getAdminEventEditAndExit == "Edit this event and exit" {

				// Title validation
				if len(getAdminEventEditTitle) > 0 {
					data.TitleError = ""
					areAdminEventEditInputsValid[0] = true
				} else {
					data.TitleError = "Title should be longer than 0"
					areAdminEventEditInputsValid[0] = false
				}

				// Description validation
				if len(getAdminEventEditDescription) > 0 {
					data.DescriptionError = ""
					areAdminEventEditInputsValid[1] = true
				} else {
					data.DescriptionError = "Description should be longer than 0"
					areAdminEventEditInputsValid[1] = false
				}

				// Url validation
				if len(getAdminEventEditUrl) > 0 {
					data.UrlError = ""
					areAdminEventEditInputsValid[2] = true
				} else {
					data.UrlError = "Url should be longer than 0"
					areAdminEventEditInputsValid[2] = false
				}

				// Published validation
				if len(getAdminEventEditPublished) > 0 {
					data.PublishedError = ""
					areAdminEventEditInputsValid[3] = true
				} else {
					data.PublishedError = "Add a date"
					areAdminEventEditInputsValid[3] = false
				}

				// Updated validation
				if len(getAdminEventEditUpdated) > 0 {
					data.UpdatedError = ""
					areAdminEventEditInputsValid[4] = true
				} else {
					data.UpdatedError = "Add a date"
					areAdminEventEditInputsValid[4] = false
				}

				// Image validation
				if len(getAdminEventEditImage) > 0 {
					data.ImageError = ""
					areAdminEventEditInputsValid[5] = true
				} else {
					data.ImageError = "An image is required"
					areAdminEventEditInputsValid[5] = false
				}

				// Author validation
				if len(getAdminEventEditAuthor) > 0 {
					data.AuthorError = ""
					areAdminEventEditInputsValid[6] = true
				} else {
					data.AuthorError = "An author is required"
					areAdminEventEditInputsValid[6] = false
				}

				// Event Type validation
				if len(getAdminEventEditType) > 0 {
					data.EventTypeError = ""
					areAdminEventEditInputsValid[7] = true
				} else {
					data.EventTypeError = "An event type is required"
					areAdminEventEditInputsValid[7] = false
				}

				// Content validation
				if len(getAdminEventEditContent) > 0 {
					data.ContentError = ""
					areAdminEventEditInputsValid[8] = true
				} else {
					data.ContentError = "Content should be longer than 0"
					areAdminEventEditInputsValid[8] = false
				}

				// Country validation
				if len(getAdminEventEditCountry) > 0 {
					data.CountryError = ""
					areAdminEventEditInputsValid[9] = true
				} else {
					data.CountryError = "Country should be longer than 0"
					areAdminEventEditInputsValid[9] = false
				}

				// Region validation
				if len(getAdminEventEditRegion) > 0 {
					data.RegioneError = ""
					areAdminEventEditInputsValid[10] = true
				} else {
					data.RegioneError = "Region should be longer than 0"
					areAdminEventEditInputsValid[10] = false
				}

				// City validation
				if len(getAdminEventEditCity) > 0 {
					data.CityError = ""
					areAdminEventEditInputsValid[11] = true
				} else {
					data.CityError = "City should be longer than 0"
					areAdminEventEditInputsValid[11] = false
				}

				// Town validation
				if len(getAdminEventEditTown) > 0 {
					data.TownError = ""
					areAdminEventEditInputsValid[12] = true
				} else {
					data.TownError = "Town should be longer than 0"
					areAdminEventEditInputsValid[12] = false
				}

				// Fraction validation
				if len(getAdminEventEditFraction) > 0 {
					data.FractionError = ""
					areAdminEventEditInputsValid[13] = true
				} else {
					data.FractionError = "Fraction should be longer than 0"
					areAdminEventEditInputsValid[13] = false
				}

				// Event Start Date validation
				if len(getAdminEventEditStartDate) > 0 {
					data.EventStartDateError = ""
					areAdminEventEditInputsValid[14] = true
				} else {
					data.EventStartDateError = "A date is required"
					areAdminEventEditInputsValid[14] = false
				}

				// Event End Date validation
				if len(getAdminEventEditEndDate) > 0 {
					data.EventEndDateError = ""
					areAdminEventEditInputsValid[15] = true
				} else {
					data.EventEndDateError = "A date is required"
					areAdminEventEditInputsValid[15] = false
				}

				// Check if the all inputs are valid
				for i := 0; i < len(areAdminEventEditInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminEventEditInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Edit current event if all the inputs are valid and redirect to all event list
				if isFormSubmittionValid {

					// Get the image id for the relationship one-to-many between events and images
					getAdminEventImageIdEdit, _ := models.ImageFindByUrlReturnItsId(getAdminEventEditImage)

					// Get the author id for the relationship one-to-many between events and images
					getAdminEventAuthorIdEdit, _ := models.AuthorFindByUrlReturnItsId(getAdminEventEditAuthor)

					editEvent := models.EventNew(
						eventId,
						getAdminEventEditTitle,
						getAdminEventEditDescription,
						getAdminEventEditUrl,
						getAdminEventEditPublished,
						getAdminEventEditUpdated,
						getAdminEventImageIdEdit,
						getAdminEventAuthorIdEdit,
						getAdminEventEditType,
						getAdminEventEditContent,
						getAdminEventEditCountry,
						getAdminEventEditRegion,
						getAdminEventEditCity,
						getAdminEventEditTown,
						getAdminEventEditFraction,
						getAdminEventEditStartDate,
						getAdminEventEditEndDate,
					)
					models.EventEdit(editEvent)

					if getAdminEventEditAdd == "Edit this event" {
						http.Redirect(w, r, "/admin/admin-event-edit/"+idPath, http.StatusSeeOther)
					} else if getAdminEventEditAndExit == "Edit this event and exit" {
						http.Redirect(w, r, "/admin/admin-events/1", http.StatusSeeOther)
					} else {
						http.Redirect(w, r, "/admin/admin-events/1", http.StatusSeeOther)
					}
				}
			}
			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}

func AdminEventDelete() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-event-delete.html"))
	http.HandleFunc("/admin/admin-event-delete/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {
			idPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-event-delete/")
			idPath = util.FormSanitizeStringInput(idPath)

			eventId, err := strconv.Atoi(idPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			getEventDelete, err := models.EventWithRelatedFieldsFindById(eventId)
			if err != nil {
				fmt.Println("Error to find event by id:", err)
			}

			data := eventData{
				PageTitle:              "Admin Delete Event",
				EventWithRelatedFields: getEventDelete,
			}

			/**
			* Check if the form for deleting event
			* has been submitted
			* and
			* delete the selected event
			 */
			isFormSubmittionValid := false

			// Get the value from the form
			getAdminEventDeleteSubmit := r.FormValue("admin-event-delete")

			// Sanitize the form input
			getAdminEventDeleteSubmit = util.FormSanitizeStringInput(getAdminEventDeleteSubmit)

			// Check if the form has been submitte
			if getAdminEventDeleteSubmit == "Delete this event" {
				isFormSubmittionValid = true
			}

			// Delete the event
			if isFormSubmittionValid {
				models.EventDelete(eventId)
				http.Redirect(w, r, "/admin/admin-events/1", http.StatusSeeOther)
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}

func AdminEventsChecker() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-events-checker.html"))
	http.HandleFunc("/admin/admin-events-checker/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-events-checker/")
			urlPath = util.FormSanitizeStringInput(urlPath)

			pageNumber, err := strconv.Atoi(urlPath)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}

			// Redirect to /admin/admin-events-checker/1 if pageNumber is 0
			if pageNumber == 0 {
				http.Redirect(w, r, "/admin/admin-events-checker/1", http.StatusSeeOther)
			}

			// Set limit and offset for MySQL query
			limit := 10
			offset := (pageNumber - 1) * limit

			// Set current date
			getCurrentDate := time.Now()
			setCurrentDate := getCurrentDate.Format("2006-01-02 15:04:05")

			eventsPassed, err := models.EventsGetAllPassed(setCurrentDate, 10, offset)
			if err != nil {
				fmt.Println("Error getting eventsPassed:", err)
			}

			// The previous and next buttons
			setPreviousButton := false
			var setPreviousPage int
			var setPreviousPageStr string
			if (pageNumber - 1) > 0 {
				setPreviousButton = true
				setPreviousPage = pageNumber - 1
				setPreviousPageStr = strconv.Itoa(setPreviousPage)
			}

			setNextButton := false
			var setNextPage int
			var setNextPageStr string
			if len(eventsPassed) >= 10 {
				setNextButton = true
				setNextPage = pageNumber + 1
				setNextPageStr = strconv.Itoa(setNextPage)
			}

			data := eventData{
				PageTitle:               "Events Admin",
				PreviusButton:           setPreviousButton,
				NextButton:              setNextButton,
				PreviousPage:            setPreviousPageStr,
				NextPage:                setNextPageStr,
				EventsWithRelatedFields: eventsPassed,
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminEventsSearch() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-events-search.html"))
	http.HandleFunc("/admin/admin-events-search/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {
			urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-events-search/")
			urlPath = util.FormSanitizeStringInput(urlPath)

			data := eventData{
				PageTitle: "Admin Events Search",
			}

			/**
			* Check if the form for searching has been submitted
			* and
			* validate the inputs
			 */
			var areAdminEventsSerachInputValid [1]bool
			isFormSubmittionValid := false

			// Get values from the form
			getAdminEventsSearchInput := r.FormValue("admin-events-search-input")
			getAdminEventsSearchButton := r.FormValue("admin-events-search-button")

			// Sanitize form inputs
			getAdminEventsSearchInput = util.FormSanitizeStringInput(getAdminEventsSearchInput)
			getAdminEventsSearchButton = util.FormSanitizeStringInput(getAdminEventsSearchButton)

			if getAdminEventsSearchButton == "Search" {
				// Input validation
				if len(getAdminEventsSearchInput) > 0 {
					data.EventsSearchInputError = ""
					areAdminEventsSerachInputValid[0] = true
				} else {
					data.EventsSearchInputError = "Add a valid input"
					areAdminEventsSerachInputValid[0] = false
				}

				for i := 0; i < len(areAdminEventsSerachInputValid); i++ {
					isFormSubmittionValid = true
					if !areAdminEventsSerachInputValid[i] {
						isFormSubmittionValid = false
						http.Redirect(w, r, "/admin/admin-events-search/", http.StatusSeeOther)
						break
					}
				}

				if isFormSubmittionValid {
					// Get events by search parameter
					redirectURL := "/admin/admin-events-search/" + getAdminEventsSearchInput
					http.Redirect(w, r, redirectURL, http.StatusSeeOther)
				}
			} else {
				getEvents, err := models.EventsFindByParameterAlsoNotPublished(urlPath)
				if err != nil {
					fmt.Println("Error getting the events by search input:", err)
				}

				// Add data for the page
				data.EventsSearchInput = urlPath
				data.EventsWithRelatedFields = getEvents

				tmpl.Execute(w, data)
			}
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}

	})
}
