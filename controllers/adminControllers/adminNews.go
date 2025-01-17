package admincontrollers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type newsData struct {
	PageTitle                      string
	TitleError                     string
	DescriptionError               string
	UrlError                       string
	PublishedError                 string
	UpdatedError                   string
	ImageError                     string
	AuthorError                    string
	ContentError                   string
	PreviousButton                 bool
	NextButton                     bool
	PreviousPage                   string
	NextPage                       string
	Images                         []models.Image
	Authors                        []models.Author
	GetNewsWithRelatedFields       []models.NewsWithRelatedFields
	GetSingleNewsWithRelatedFields models.NewsWithRelatedFields
}

func AdminNews(w http.ResponseWriter, r *http.Request) {

	data := newsData{
		PageTitle: "News Admin",
	}

	if r.Method == http.MethodPost {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-news.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-news/")
		urlPath = util.FormSanitizeStringInput(urlPath)

		pageNumber, err := strconv.Atoi(urlPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		// Redirect to /admin/admin-news/1 if pageNumber is 0
		if pageNumber == 0 {
			http.Redirect(w, r, "/admin/admin-news/1", http.StatusSeeOther)
		}

		// Set limit and offset for MySQL query
		limit := 10
		offset := (pageNumber - 1) * limit

		newsDate, err := models.NewsGetLimitAndPagination(limit, offset)
		if err != nil {
			fmt.Println("Error getting newsDate:", err)
		}

		// The previous and next button
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
		if len(newsDate) >= 10 {
			setNextButton = true
			setNextPage = pageNumber + 1
			setNextPageStr = strconv.Itoa(setNextPage)
		}

		data.PreviousButton = setPreviousButton
		data.NextButton = setNextButton
		data.PreviousPage = setPreviousPageStr
		data.NextPage = setNextPageStr
		data.GetNewsWithRelatedFields = newsDate

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminNewsAdd(w http.ResponseWriter, r *http.Request) {

	data := newsData{
		PageTitle: "Admin News Add",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-news-add.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		imagesData, errImagesData := models.ImageShowImagesByUpdated()
		if errImagesData != nil {
			fmt.Println("Error getting imagesData:", errImagesData)
		}

		authorsData, errAuthorsData := models.AuthorShowAuthors()
		if errAuthorsData != nil {
			fmt.Println("Error getting authorsData:", errAuthorsData)
		}

		data.Images = imagesData
		data.Authors = authorsData

		// Flag validation
		var areAdminNewsAddInputsValid [8]bool
		isFormSubmittionValid := false

		// Get the value from the form
		getAdminNewsTitle := r.FormValue("news-title")
		getAdminNewsDescription := r.FormValue("news-description")
		getAdminNewsUrl := r.FormValue("news-url")
		getAdminNewsPublished := r.FormValue("news-published")
		getAdminNewsUpdated := r.FormValue("news-updated")
		getAdminNewsImage := r.FormValue("news-image")
		getAdminNewsAuthor := r.FormValue("news-author")
		getAdminNewsContent := r.FormValue("news-content")
		getAdminNewsAdd := r.FormValue("news-add")

		// Sanitize form inputs
		getAdminNewsTitle = util.FormSanitizeStringInput(getAdminNewsTitle)
		getAdminNewsDescription = util.FormSanitizeStringInput(getAdminNewsDescription)
		getAdminNewsUrl = util.FormSanitizeStringInput(getAdminNewsUrl)
		getAdminNewsPublished = util.FormSanitizeStringInput(getAdminNewsPublished)
		getAdminNewsUpdated = util.FormSanitizeStringInput(getAdminNewsUpdated)
		getAdminNewsImage = util.FormSanitizeStringInput(getAdminNewsImage)
		getAdminNewsAuthor = util.FormSanitizeStringInput(getAdminNewsAuthor)
		getAdminNewsAdd = util.FormSanitizeStringInput(getAdminNewsAdd)

		// Check if the form has been submitted
		if getAdminNewsAdd == "Add new news" {
			// Title validation
			if len(getAdminNewsTitle) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[0] = true
			} else {
				data.TitleError = "Title should be longer than 0"
				areAdminNewsAddInputsValid[0] = false
			}

			// Description validation
			if len(getAdminNewsDescription) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[1] = true
			} else {
				data.TitleError = "Description should be longer than 0"
				areAdminNewsAddInputsValid[1] = false
			}

			// Url validation
			if len(getAdminNewsUrl) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[2] = true
			} else {
				data.TitleError = "Url should be longer than 0"
				areAdminNewsAddInputsValid[2] = false
			}

			// Published validation
			if len(getAdminNewsPublished) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[3] = true
			} else {
				data.TitleError = "Add a date"
				areAdminNewsAddInputsValid[3] = false
			}

			// Updated validation
			if len(getAdminNewsUpdated) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[4] = true
			} else {
				data.TitleError = "Add a date"
				areAdminNewsAddInputsValid[4] = false
			}

			// Image validation
			if len(getAdminNewsImage) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[5] = true
			} else {
				data.TitleError = "An image is required"
				areAdminNewsAddInputsValid[5] = false
			}

			// Author validation
			if len(getAdminNewsAuthor) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[6] = true
			} else {
				data.TitleError = "An author is required"
				areAdminNewsAddInputsValid[6] = false
			}

			// Content validation
			if len(getAdminNewsContent) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[7] = true
			} else {
				data.TitleError = "Content should be longer than 0"
				areAdminNewsAddInputsValid[7] = false
			}

			// Check if all fields are valid
			for i := 0; i < len(areAdminNewsAddInputsValid); i++ {
				isFormSubmittionValid = true
				if !areAdminNewsAddInputsValid[i] {
					isFormSubmittionValid = false
					break
				}
			}

			// Create a new news if all inputs are valid
			if isFormSubmittionValid {
				// Get image id for the relationship one-to-many between events and images
				getAdminNewsImageId, _ := models.ImageFindByUrlReturnItsId(getAdminNewsImage)

				// Get author id for the relationship one-to-many between events and authors
				getAdminNewsAuthorId, _ := models.AuthorFindByUrlReturnItsId(getAdminNewsAuthor)

				createNewNews := models.NewsNew(
					1,
					getAdminNewsTitle,
					getAdminNewsDescription,
					getAdminNewsUrl,
					getAdminNewsPublished,
					getAdminNewsUpdated,
					getAdminNewsContent,
					getAdminNewsImageId,
					getAdminNewsAuthorId,
				)
				models.NewsAddNewToDB(createNewNews)
				http.Redirect(w, r, "/admin/admin-news/1", http.StatusSeeOther)
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func AdminNewsEdit(w http.ResponseWriter, r *http.Request) {

	data := newsData{
		PageTitle: "Admin News Edit",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-news-edit.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		idPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-news-edit/")
		idPath = util.FormSanitizeStringInput(idPath)

		newsId, err := strconv.Atoi(idPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		getNewsEdit, err := models.NewsWithRelatedFieldsFindById(newsId)
		if err != nil {
			fmt.Println("Error to find this news:", err)
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

		data.GetSingleNewsWithRelatedFields = getNewsEdit
		data.Images = imagesData
		data.Authors = authorsData

		/**
		* Check if the form for editing the news has been submitted
		* and
		* validate the inputs
		 */
		// Flag validation
		var areAdminNewsAddInputsValid [8]bool
		isFormSubmittionValid := false

		// Get the values from the form
		getAdminEditNewsTitle := r.FormValue("news-edit-title")
		getAdminEditNewsDescription := r.FormValue("news-edit-description")
		getAdminEditNewsUrl := r.FormValue("news-edit-url")
		getAdminEditNewsPublished := r.FormValue("news-edit-published")
		getAdminEditNewsUpdated := r.FormValue("news-edit-updated")
		getAdminEditNewsImage := r.FormValue("news-edit-image")
		getAdminEditNewsAuthor := r.FormValue("news-edit-author")
		getAdminEditNewsContent := r.FormValue("news-edit-content")
		getAdminEditNews := r.FormValue("news-edit")
		getAdminEditNewsAndExit := r.FormValue("news-edit-and-exit")

		// Sanitize form inputs
		getAdminEditNewsTitle = util.FormSanitizeStringInput(getAdminEditNewsTitle)
		getAdminEditNewsDescription = util.FormSanitizeStringInput(getAdminEditNewsDescription)
		getAdminEditNewsUrl = util.FormSanitizeStringInput(getAdminEditNewsUrl)
		getAdminEditNewsPublished = util.FormSanitizeStringInput(getAdminEditNewsPublished)
		getAdminEditNewsUpdated = util.FormSanitizeStringInput(getAdminEditNewsUpdated)
		getAdminEditNewsImage = util.FormSanitizeStringInput(getAdminEditNewsImage)
		getAdminEditNewsAuthor = util.FormSanitizeStringInput(getAdminEditNewsAuthor)
		getAdminEditNews = util.FormSanitizeStringInput(getAdminEditNews)
		getAdminEditNewsAndExit = util.FormSanitizeStringInput(getAdminEditNewsAndExit)

		if getAdminEditNews == "Edit this news" || getAdminEditNewsAndExit == "Edit this news and exit" {
			// Title validation
			if len(getAdminEditNewsTitle) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[0] = true
			} else {
				data.TitleError = "Title should be longer than 0"
				areAdminNewsAddInputsValid[0] = false
			}

			// Description validation
			if len(getAdminEditNewsDescription) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[1] = true
			} else {
				data.TitleError = "Description should be longer than 0"
				areAdminNewsAddInputsValid[1] = false
			}

			// Url validation
			if len(getAdminEditNewsUrl) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[2] = true
			} else {
				data.TitleError = "Url should be longer than 0"
				areAdminNewsAddInputsValid[2] = false
			}

			// Published validation
			if len(getAdminEditNewsPublished) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[3] = true
			} else {
				data.TitleError = "Add a date"
				areAdminNewsAddInputsValid[3] = false
			}

			// Updated validation
			if len(getAdminEditNewsUpdated) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[4] = true
			} else {
				data.TitleError = "Add a date"
				areAdminNewsAddInputsValid[4] = false
			}

			// Image validation
			if len(getAdminEditNewsImage) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[5] = true
			} else {
				data.TitleError = "An image is required"
				areAdminNewsAddInputsValid[5] = false
			}

			// Author validation
			if len(getAdminEditNewsAuthor) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[6] = true
			} else {
				data.TitleError = "An author is required"
				areAdminNewsAddInputsValid[6] = false
			}

			// Content validation
			if len(getAdminEditNewsContent) > 0 {
				data.TitleError = ""
				areAdminNewsAddInputsValid[7] = true
			} else {
				data.TitleError = "Content should be longer than 0"
				areAdminNewsAddInputsValid[7] = false
			}

			// Check if the all inputs are valid
			for i := 0; i < len(areAdminNewsAddInputsValid); i++ {
				isFormSubmittionValid = true
				if !areAdminNewsAddInputsValid[i] {
					isFormSubmittionValid = false
					break
				}
			}

			// Edit current news if all the inputs are valid and redirect to all news list
			if isFormSubmittionValid {
				// Get the image id for the relationship one-to-many between news and images
				getAdminNewsImageIdEdit, _ := models.ImageFindByUrlReturnItsId(getAdminEditNewsImage)

				// Get the author id for the relationship one-to-many between news and images
				getAdminNewsAuthorIdEdit, _ := models.AuthorFindByUrlReturnItsId(getAdminEditNewsAuthor)

				editNews := models.NewsNew(
					newsId,
					getAdminEditNewsTitle,
					getAdminEditNewsDescription,
					getAdminEditNewsUrl,
					getAdminEditNewsPublished,
					getAdminEditNewsUpdated,
					getAdminEditNewsContent,
					getAdminNewsImageIdEdit,
					getAdminNewsAuthorIdEdit,
				)

				models.NewsEdit(editNews)

				if getAdminEditNews == "Edit this news" {
					http.Redirect(w, r, "/admin/admin-news-edit/"+idPath, http.StatusSeeOther)
				} else if getAdminEditNewsAndExit == "Edit this news and exit" {
					http.Redirect(w, r, "/admin/admin-news/1", http.StatusSeeOther)
				} else {
					http.Redirect(w, r, "/adin/admin-news/1", http.StatusSeeOther)
				}
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func AdminNewsDelete(w http.ResponseWriter, r *http.Request) {

	data := newsData{
		PageTitle: "Admin Delete News",
	}

	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-news-delete.html"))
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		idPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-news-delete/")
		idPath = util.FormSanitizeStringInput(idPath)

		newsId, err := strconv.Atoi(idPath)
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		getNewsDelete, err := models.NewsWithRelatedFieldsFindById(newsId)
		if err != nil {
			fmt.Println("Error to find news by id:", err)
		}

		data.GetSingleNewsWithRelatedFields = getNewsDelete

		/**
		* Check if the form for deleting event
		* has been submitted
		* and
		* delete the selected event
		 */
		isFormSubmittionValid := false

		// Get the value from the form
		getAdminNewsDeleteSubmit := r.FormValue("admin-news-delete")

		// Santize the form input
		getAdminNewsDeleteSubmit = util.FormSanitizeStringInput(getAdminNewsDeleteSubmit)

		// Check if the form has been submitted
		if getAdminNewsDeleteSubmit == "Delete this news" {
			isFormSubmittionValid = true
		}

		if isFormSubmittionValid {
			models.NewsDelete(newsId)
			http.Redirect(w, r, "/admin/admin-news/1", http.StatusSeeOther)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
