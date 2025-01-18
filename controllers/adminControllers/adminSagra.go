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

type sagraData struct {
	PageTitle              string
	TitleError             string
	DescriptionError       string
	UrlError               string
	PublishedError         string
	UpdatedError           string
	ImageError             string
	AuthorError            string
	ContentError           string
	CountryError           string
	RegioneError           string
	CityError              string
	TownError              string
	FractionError          string
	SagraStartDateError    string
	SagraEndDateError      string
	SagreSearchInput       string
	SagreSearchInputError  string
	PreviusButton          bool
	NextButton             bool
	PreviousPage           string
	NextPage               string
	Images                 []models.Image
	Authors                []models.Author
	Sagre                  []models.Sagra
	SagreWithRelatedFields []models.SagraWithRelatedFields
	SagraWithRelatedFields models.SagraWithRelatedFields
}

func AdminSagre(w http.ResponseWriter, r *http.Request) {

	data := sagraData{
		PageTitle: "Sagre Admin",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-sagre.html"))

		urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-sagre/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		pageNumber, err := strconv.Atoi(urlPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		// Redirect to /admin/admin-sagre/1 if pageNumber is 0
		if pageNumber == 0 {
			http.Redirect(w, r, "/admin/admin-sagre/1", http.StatusSeeOther)
		}

		// Set limit and offset for MySQL query
		limit := 10
		offset := (pageNumber - 1) * limit

		sagreDate, err := models.SagreGetLimitAndPagination(10, offset)
		if err != nil {
			fmt.Println("Error getting sagreData:", err)
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
		if len(sagreDate) >= 10 {
			setNextButton = true
			setNextPage = pageNumber + 1
			setNextPageStr = strconv.Itoa(setNextPage)
		}

		data.PreviusButton = setPreviousButton
		data.NextButton = setNextButton
		data.PreviousPage = setPreviousPageStr
		data.NextPage = setNextPageStr
		data.SagreWithRelatedFields = sagreDate

		tmpl.Execute(w, data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminSagraAdd(w http.ResponseWriter, r *http.Request) {

	data := sagraData{
		PageTitle: "Admin Sagra Add",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-sagra-add.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		imagesData, errImagesData := models.ImageShowImagesByUpdated()
		if errImagesData != nil {
			fmt.Println("Error getting imagesData:", errImagesData)
		}

		authorsData, errAuthorsData := models.AuthorShowAuthors()
		if errAuthorsData != nil {
			fmt.Println("Error getting authorsData:", authorsData)
		}

		data.Images = imagesData
		data.Authors = authorsData
	}

	// Flag validation
	var areAdminSagraInputsValid [15]bool
	isFormSubmittionValid := false

	// Get the value from the form
	getAdminSagraTitle := r.FormValue("sagra-title")
	getAdminSagraDescription := r.FormValue("sagra-description")
	getAdminSagraUrl := r.FormValue("sagra-url")
	getAdminSagraPublished := r.FormValue("sagra-published")
	getAdminSagraUpdated := r.FormValue("sagra-updated")
	getAdminSagraImage := r.FormValue("sagra-image")
	getAdminSagraAuthor := r.FormValue("sagra-author")
	getAdminSagraContent := r.FormValue("sagra-content")
	getAdminSagraCountry := r.FormValue("sagra-country")
	getAdminSagraRegion := r.FormValue("sagra-region")
	getAdminSagraCity := r.FormValue("sagra-city")
	getAdminSagraTown := r.FormValue("sagra-town")
	getAdminSagraFraction := r.FormValue("sagra-fraction")
	getAdminSagraStartDate := r.FormValue("sagra-start-date")
	getAdminSagraEndDate := r.FormValue("sagra-end-date")
	getAdminSagraAdd := r.FormValue("sagra-add")

	// Sanitize form inputs
	getAdminSagraTitle = util.FormSanitizeStringInput(getAdminSagraTitle)
	getAdminSagraDescription = util.FormSanitizeStringInput(getAdminSagraDescription)
	getAdminSagraUrl = util.FormSanitizeStringInput(getAdminSagraUrl)
	getAdminSagraPublished = util.FormSanitizeStringInput(getAdminSagraPublished)
	getAdminSagraUpdated = util.FormSanitizeStringInput(getAdminSagraUpdated)
	getAdminSagraImage = util.FormSanitizeStringInput(getAdminSagraImage)
	getAdminSagraAuthor = util.FormSanitizeStringInput(getAdminSagraAuthor)
	getAdminSagraCountry = util.FormSanitizeStringInput(getAdminSagraCountry)
	getAdminSagraRegion = util.FormSanitizeStringInput(getAdminSagraRegion)
	getAdminSagraCity = util.FormSanitizeStringInput(getAdminSagraCity)
	getAdminSagraTown = util.FormSanitizeStringInput(getAdminSagraTown)
	getAdminSagraFraction = util.FormSanitizeStringInput(getAdminSagraFraction)
	getAdminSagraStartDate = util.FormSanitizeStringInput(getAdminSagraStartDate)
	getAdminSagraEndDate = util.FormSanitizeStringInput(getAdminSagraEndDate)
	getAdminSagraAdd = util.FormSanitizeStringInput(getAdminSagraAdd)

	// Check if the form has been submitted
	if getAdminSagraAdd == "Add new sagra" {
		// Title validation
		if len(getAdminSagraTitle) > 0 {
			data.TitleError = ""
			areAdminSagraInputsValid[0] = true
		} else {
			data.TitleError = "Title should be longer than 0"
			areAdminSagraInputsValid[0] = false
		}

		// Description validation
		if len(getAdminSagraDescription) > 0 {
			data.DescriptionError = ""
			areAdminSagraInputsValid[1] = true
		} else {
			data.DescriptionError = "Description should be longer than 0"
			areAdminSagraInputsValid[1] = false
		}

		// URL validation
		if len(getAdminSagraUrl) > 0 {
			data.UrlError = ""
			areAdminSagraInputsValid[2] = true
		} else {
			data.UrlError = "Url should be longer than 0"
			areAdminSagraInputsValid[2] = false
		}

		// Published validation
		if len(getAdminSagraPublished) > 0 {
			data.PublishedError = ""
			areAdminSagraInputsValid[3] = true
		} else {
			data.PublishedError = "Add a date"
			areAdminSagraInputsValid[3] = false
		}

		// Updated validation
		if len(getAdminSagraUpdated) > 0 {
			data.UpdatedError = ""
			areAdminSagraInputsValid[4] = true
		} else {
			data.UpdatedError = "Add a date"
			areAdminSagraInputsValid[4] = false
		}

		// Image validation
		if len(getAdminSagraImage) > 0 {
			data.ImageError = ""
			areAdminSagraInputsValid[5] = true
		} else {
			data.ImageError = "An image is required"
			areAdminSagraInputsValid[5] = false
		}

		// Author validation
		if len(getAdminSagraAuthor) > 0 {
			data.AuthorError = ""
			areAdminSagraInputsValid[6] = true
		} else {
			data.AuthorError = "An author is required"
			areAdminSagraInputsValid[6] = false
		}

		// Content validation
		if len(getAdminSagraContent) > 0 {
			data.ContentError = ""
			areAdminSagraInputsValid[7] = true
		} else {
			data.ContentError = "Content should be longer than 0"
			areAdminSagraInputsValid[7] = false
		}

		// Country validation
		if len(getAdminSagraCountry) > 0 {
			data.CountryError = ""
			areAdminSagraInputsValid[8] = true
		} else {
			data.CountryError = "Country should be longer than 0"
			areAdminSagraInputsValid[8] = false
		}

		// Region validation
		if len(getAdminSagraRegion) > 0 {
			data.RegioneError = ""
			areAdminSagraInputsValid[9] = true
		} else {
			data.RegioneError = "Region should be longer than 0"
			areAdminSagraInputsValid[9] = false
		}

		// City validation
		if len(getAdminSagraCity) > 0 {
			data.CityError = ""
			areAdminSagraInputsValid[10] = true
		} else {
			data.CityError = "City should be longer than 0"
			areAdminSagraInputsValid[10] = false
		}

		// Town validation
		if len(getAdminSagraTown) > 0 {
			data.TownError = ""
			areAdminSagraInputsValid[11] = true
		} else {
			data.TownError = "Town should be longer than 0"
			areAdminSagraInputsValid[11] = false
		}

		// Fraction validation
		if len(getAdminSagraFraction) > 0 {
			data.FractionError = ""
			areAdminSagraInputsValid[12] = true
		} else {
			data.FractionError = "Fraction should be longer than 0"
			areAdminSagraInputsValid[12] = false
		}

		// Sagra Start date validation
		if len(getAdminSagraStartDate) > 0 {
			data.SagraStartDateError = ""
			areAdminSagraInputsValid[13] = true
		} else {
			data.SagraStartDateError = "Add a date"
			areAdminSagraInputsValid[13] = false
		}

		// Sagra End Date validation
		if len(getAdminSagraEndDate) > 0 {
			data.SagraEndDateError = ""
			areAdminSagraInputsValid[14] = true
		} else {
			data.SagraEndDateError = "Add a date"
			areAdminSagraInputsValid[14] = false
		}

		for i := 0; i < len(areAdminSagraInputsValid); i++ {
			isFormSubmittionValid = true
			if !areAdminSagraInputsValid[i] {
				isFormSubmittionValid = false
				break
			}
		}

		// Create a new sagra if all inputs are valid
		if isFormSubmittionValid {
			// Get image id for the relationship one-to-many between sagre and images
			getAdminSagraImageId, _ := models.ImageFindByUrlReturnItsId(getAdminSagraImage)

			// Get author id for the relationship one-to-many between sagre and authors
			getAdminSagraAuthorId, _ := models.AuthorFindByUrlReturnItsId(getAdminSagraAuthor)

			createNewSagra := models.SagraNew(
				1,
				getAdminSagraTitle,
				getAdminSagraDescription,
				getAdminSagraUrl,
				getAdminSagraPublished,
				getAdminSagraUpdated,
				getAdminSagraImageId,
				getAdminSagraAuthorId,
				getAdminSagraContent,
				getAdminSagraCountry,
				getAdminSagraRegion,
				getAdminSagraCity,
				getAdminSagraTown,
				getAdminSagraFraction,
				getAdminSagraStartDate,
				getAdminSagraEndDate,
			)
			models.SagraAddNewToDB(createNewSagra)
			http.Redirect(w, r, "/admin/admin-sagre/1", http.StatusSeeOther)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminSagraEdit(w http.ResponseWriter, r *http.Request) {

	data := sagraData{
		PageTitle: "Admin Sagra Edit",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-sagra-edit.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		idPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-sagra-edit/")
		idPath = util.FormSanitizeStringInput(idPath)

		sagraId, err := strconv.Atoi(idPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		getSagraEdit, err := models.SagraWithRelatedImageFindById(sagraId)
		if err != nil {
			fmt.Println("Error to find sagra:", err)
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
		data.SagraWithRelatedFields = getSagraEdit
		data.Images = imagesData
		data.Authors = authorsData

		/**
		* Check if the form for editing the sagra has been submitted
		* and
		* validate the inputs
		 */
		var areAdminSagraInputsValid [15]bool
		isFormSubmittionValid := false

		// Get the value from the form
		getAdminSagraTitleEdit := r.FormValue("sagra-edit-title")
		getAdminSagraDescriptionEdit := r.FormValue("sagra-edit-description")
		getAdminSagraUrlEdit := r.FormValue("sagra-edit-url")
		getAdminSagraPublishedEdit := r.FormValue("sagra-edit-published")
		getAdminSagraUpdatedEdit := r.FormValue("sagra-edit-updated")
		getAdminSagraImageEdit := r.FormValue("sagra-edit-image")
		getAdminSagraAuthorEdit := r.FormValue("sagra-edit-author")
		getAdminSagraContentEdit := r.FormValue("sagra-edit-content")
		getAdminSagraCountryEdit := r.FormValue("sagra-edit-country")
		getAdminSagraRegionEdit := r.FormValue("sagra-edit-region")
		getAdminSagraCityEdit := r.FormValue("sagra-edit-city")
		getAdminSagraTownEdit := r.FormValue("sagra-edit-town")
		getAdminSagraFractionEdit := r.FormValue("sagra-edit-fraction")
		getAdminSagraStartDateEdit := r.FormValue("sagra-edit-start-date")
		getAdminSagraEndDateEdit := r.FormValue("sagra-edit-end-date")
		getAdminSagraAddEdit := r.FormValue("sagra-edit")
		getAdminSagraAddEditAndExit := r.FormValue("sagra-edit-and-exit")

		// Sanitize form inputs
		getAdminSagraTitleEdit = util.FormSanitizeStringInput(getAdminSagraTitleEdit)
		getAdminSagraDescriptionEdit = util.FormSanitizeStringInput(getAdminSagraDescriptionEdit)
		getAdminSagraUrlEdit = util.FormSanitizeStringInput(getAdminSagraUrlEdit)
		getAdminSagraPublishedEdit = util.FormSanitizeStringInput(getAdminSagraPublishedEdit)
		getAdminSagraUpdatedEdit = util.FormSanitizeStringInput(getAdminSagraUpdatedEdit)
		getAdminSagraImageEdit = util.FormSanitizeStringInput(getAdminSagraImageEdit)
		getAdminSagraAuthorEdit = util.FormSanitizeStringInput(getAdminSagraAuthorEdit)
		getAdminSagraCountryEdit = util.FormSanitizeStringInput(getAdminSagraCountryEdit)
		getAdminSagraRegionEdit = util.FormSanitizeStringInput(getAdminSagraRegionEdit)
		getAdminSagraCityEdit = util.FormSanitizeStringInput(getAdminSagraCityEdit)
		getAdminSagraTownEdit = util.FormSanitizeStringInput(getAdminSagraTownEdit)
		getAdminSagraFractionEdit = util.FormSanitizeStringInput(getAdminSagraFractionEdit)
		getAdminSagraStartDateEdit = util.FormSanitizeStringInput(getAdminSagraStartDateEdit)
		getAdminSagraEndDateEdit = util.FormSanitizeStringInput(getAdminSagraEndDateEdit)
		getAdminSagraAddEdit = util.FormSanitizeStringInput(getAdminSagraAddEdit)
		getAdminSagraAddEditAndExit = util.FormSanitizeStringInput(getAdminSagraAddEditAndExit)

		// Check if the form has been submitted
		if getAdminSagraAddEdit == "Edit this sagra" || getAdminSagraAddEditAndExit == "Edit this sagra and exit" {
			// Title validation
			if len(getAdminSagraTitleEdit) > 0 {
				data.TitleError = ""
				areAdminSagraInputsValid[0] = true
			} else {
				data.TitleError = "Title should be longer than 0"
				areAdminSagraInputsValid[0] = false
			}

			// Description validation
			if len(getAdminSagraDescriptionEdit) > 0 {
				data.DescriptionError = ""
				areAdminSagraInputsValid[1] = true
			} else {
				data.DescriptionError = "Description should be longer than 0"
				areAdminSagraInputsValid[1] = false
			}

			// URL validation
			if len(getAdminSagraUrlEdit) > 0 {
				data.UrlError = ""
				areAdminSagraInputsValid[2] = true
			} else {
				data.UrlError = "Url should be longer than 0"
				areAdminSagraInputsValid[2] = false
			}

			// Published validation
			if len(getAdminSagraPublishedEdit) > 0 {
				data.PublishedError = ""
				areAdminSagraInputsValid[3] = true
			} else {
				data.PublishedError = "Add a date"
				areAdminSagraInputsValid[3] = false
			}

			// Updated validation
			if len(getAdminSagraUpdatedEdit) > 0 {
				data.UpdatedError = ""
				areAdminSagraInputsValid[4] = true
			} else {
				data.UpdatedError = "Add a date"
				areAdminSagraInputsValid[4] = false
			}

			// Image validation
			if len(getAdminSagraImageEdit) > 0 {
				data.ImageError = ""
				areAdminSagraInputsValid[5] = true
			} else {
				data.ImageError = "An image is required"
				areAdminSagraInputsValid[5] = false
			}

			// Author validation
			if len(getAdminSagraAuthorEdit) > 0 {
				data.AuthorError = ""
				areAdminSagraInputsValid[6] = true
			} else {
				data.AuthorError = "An author is required"
				areAdminSagraInputsValid[6] = false
			}

			// Content validation
			if len(getAdminSagraContentEdit) > 0 {
				data.ContentError = ""
				areAdminSagraInputsValid[7] = true
			} else {
				data.ContentError = "Content should be longer than 0"
				areAdminSagraInputsValid[7] = false
			}

			// Country validation
			if len(getAdminSagraCountryEdit) > 0 {
				data.CountryError = ""
				areAdminSagraInputsValid[8] = true
			} else {
				data.CountryError = "Country should be longer than 0"
				areAdminSagraInputsValid[8] = false
			}

			// Region validation
			if len(getAdminSagraRegionEdit) > 0 {
				data.RegioneError = ""
				areAdminSagraInputsValid[9] = true
			} else {
				data.RegioneError = "Region should be longer than 0"
				areAdminSagraInputsValid[9] = false
			}

			// City validation
			if len(getAdminSagraCityEdit) > 0 {
				data.CityError = ""
				areAdminSagraInputsValid[10] = true
			} else {
				data.CityError = "City should be longer than 0"
				areAdminSagraInputsValid[10] = false
			}

			// Town validation
			if len(getAdminSagraTownEdit) > 0 {
				data.TownError = ""
				areAdminSagraInputsValid[11] = true
			} else {
				data.TownError = "Town should be longer than 0"
				areAdminSagraInputsValid[11] = false
			}

			// Fraction validation
			if len(getAdminSagraFractionEdit) > 0 {
				data.FractionError = ""
				areAdminSagraInputsValid[12] = true
			} else {
				data.FractionError = "Fraction should be longer than 0"
				areAdminSagraInputsValid[12] = false
			}

			// Sagra Start date validation
			if len(getAdminSagraStartDateEdit) > 0 {
				data.SagraStartDateError = ""
				areAdminSagraInputsValid[13] = true
			} else {
				data.SagraStartDateError = "Add a date"
				areAdminSagraInputsValid[13] = false
			}

			if len(getAdminSagraEndDateEdit) > 0 {
				data.SagraEndDateError = ""
				areAdminSagraInputsValid[14] = true
			} else {
				data.SagraEndDateError = ""
				areAdminSagraInputsValid[14] = false
			}

			for i := 0; i < len(areAdminSagraInputsValid); i++ {
				isFormSubmittionValid = true
				if !areAdminSagraInputsValid[i] {
					isFormSubmittionValid = false
					break
				}
			}

			// Edit current sagra if all the inputs are valid and redirect to all sagre list
			if isFormSubmittionValid {
				// Get the image id for the relationship one-to-many between sagre and images
				getAdminSagraImageIdEdit, _ := models.ImageFindByUrlReturnItsId(getAdminSagraImageEdit)

				// Get the author id for the relationship one-to-many between sagre and images
				getAdminSagraAuthorIdEdit, _ := models.AuthorFindByUrlReturnItsId(getAdminSagraAuthorEdit)

				editSagra := models.SagraNew(
					sagraId,
					getAdminSagraTitleEdit,
					getAdminSagraDescriptionEdit,
					getAdminSagraUrlEdit,
					getAdminSagraPublishedEdit,
					getAdminSagraUpdatedEdit,
					getAdminSagraImageIdEdit,
					getAdminSagraAuthorIdEdit,
					getAdminSagraContentEdit,
					getAdminSagraCountryEdit,
					getAdminSagraRegionEdit,
					getAdminSagraCityEdit,
					getAdminSagraTownEdit,
					getAdminSagraFractionEdit,
					getAdminSagraStartDateEdit,
					getAdminSagraEndDateEdit,
				)
				models.SagraEdit(editSagra)

				if getAdminSagraAddEdit == "Edit this sagra" {
					http.Redirect(w, r, "/admin/admin-sagra-edit/"+idPath, http.StatusSeeOther)
				} else if getAdminSagraAddEditAndExit == "Edit this sagra and exit" {
					http.Redirect(w, r, "/admin/admin-sagre/1", http.StatusSeeOther)
				} else {
					http.Redirect(w, r, "/admin/admin-sagre/1", http.StatusSeeOther)
				}
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminSagraDelete(w http.ResponseWriter, r *http.Request) {

	data := sagraData{
		PageTitle: "Admin Delete Sagra",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-sagra-delete.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		idPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-sagra-delete/")
		idPath = util.FormSanitizeStringInput(idPath)

		sagraId, err := strconv.Atoi(idPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		getSagraDelete, err := models.SagraWithRelatedImageFindById(sagraId)
		if err != nil {
			fmt.Println("Error to find sagra by id:", err)
		}

		data.SagraWithRelatedFields = getSagraDelete

		/**
		* Check if the form for deleting sagra
		* has been submitted
		* and
		* delete the selected sagra
		 */
		isFormSubmittionValid := false

		// Get the value from the form
		getAdminSagraDeleteSubmit := r.FormValue("admin-sagra-delete")

		// Sanitize the form input
		getAdminSagraDeleteSubmit = util.FormSanitizeStringInput(getAdminSagraDeleteSubmit)

		// Check if the form has been submitted
		if getAdminSagraDeleteSubmit == "Delete this sagra" {
			isFormSubmittionValid = true
		}

		if isFormSubmittionValid {
			models.SagraDelete(sagraId)
			http.Redirect(w, r, "/admin/admin-sagre/1", http.StatusSeeOther)
		}

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminSagreChecker(w http.ResponseWriter, r *http.Request) {

	data := sagraData{
		PageTitle: "Sagre Checker",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-sagre-checker.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-sagre-checker/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		pageNumber, err := strconv.Atoi(urlPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		// Redirect to /admin/admin-sagre-checker/1 if pageNumber is 0
		if pageNumber == 0 {
			http.Redirect(w, r, "/admin/admin-sagre-checker/1", http.StatusSeeOther)
		}

		// Set limit and offset for MySQL query
		limit := 10
		offset := (pageNumber - 1) * limit

		// Set current date
		getCurrentDate := time.Now()
		setCurrentDate := getCurrentDate.Format("2006-01-02 15:04:05")

		sagrePassed, err := models.SagreGetAllPassed(setCurrentDate, 10, offset)
		if err != nil {
			fmt.Println("Error getting sagrePassed:", err)
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
		if len(sagrePassed) >= 10 {
			setNextButton = true
			setNextPage = pageNumber + 1
			setNextPageStr = strconv.Itoa(setNextPage)
		}

		data.PreviusButton = setPreviousButton
		data.NextButton = setNextButton
		data.PreviousPage = setPreviousPageStr
		data.NextPage = setNextPageStr
		data.SagreWithRelatedFields = sagrePassed

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminSagreSearch(w http.ResponseWriter, r *http.Request) {

	data := sagraData{
		PageTitle: "Admin Sagre Search",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-sagre-search.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-sagre-search/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		/**
		* Check if the form for searching has been submitted
		* and
		* validate the inputs
		 */
		var areAdminSagreSerachInputValid [1]bool
		isFormSubmittionValid := false

		// Get values from the form
		getAdminSagreSearchInput := r.FormValue("admin-sagre-search-input")
		getAdminSagreSearchButton := r.FormValue("admin-sagre-search-button")

		// Sanitize form inputs
		getAdminSagreSearchInput = util.FormSanitizeStringInput(getAdminSagreSearchInput)
		getAdminSagreSearchButton = util.FormSanitizeStringInput(getAdminSagreSearchButton)

		if getAdminSagreSearchButton == "Search" {
			// Input validation
			if len(getAdminSagreSearchInput) > 0 {
				data.SagreSearchInputError = ""
				areAdminSagreSerachInputValid[0] = true
			} else {
				data.SagreSearchInputError = "Add a valid input"
				areAdminSagreSerachInputValid[0] = false
			}

			for i := 0; i < len(areAdminSagreSerachInputValid); i++ {
				isFormSubmittionValid = true
				if !areAdminSagreSerachInputValid[i] {
					isFormSubmittionValid = false
					http.Redirect(w, r, "/admin/admin-sagre-search/", http.StatusSeeOther)
					break
				}
			}

			if isFormSubmittionValid {
				// Get sagre by search parameter
				redirectURL := "/admin/admin-sagre-search/" + getAdminSagreSearchInput
				http.Redirect(w, r, redirectURL, http.StatusSeeOther)
			}
		} else {
			getSagre, err := models.SagraFindByParameterAlsoNotPublished(urlPath)
			if err != nil {
				fmt.Println("Error getting the sagre by search input:", err)
			}

			// Add data for the page
			data.SagreSearchInput = urlPath
			data.SagreWithRelatedFields = getSagre

		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
