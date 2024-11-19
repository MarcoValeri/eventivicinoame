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

func AdminNews() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-news.html"))
	http.HandleFunc("/admin/admin-news/", func(w http.ResponseWriter, r *http.Request) {
		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {
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

			data := newsData{
				PageTitle:                "News Admin",
				PreviousButton:           setPreviousButton,
				NextButton:               setNextButton,
				PreviousPage:             setPreviousPageStr,
				NextPage:                 setNextPageStr,
				GetNewsWithRelatedFields: newsDate,
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}

func AdminNewsAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-news-add.html"))
	http.HandleFunc("/admin/admin-news-add", func(w http.ResponseWriter, r *http.Request) {
		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			imagesData, errImagesData := models.ImageShowImages()
			if errImagesData != nil {
				fmt.Println("Error getting imagesData:", errImagesData)
			}

			authorsData, errAuthorsData := models.AuthorShowAuthors()
			if errAuthorsData != nil {
				fmt.Println("Error getting authorsData:", errAuthorsData)
			}

			data := newsData{
				PageTitle: "Admin News Add",
				Images:    imagesData,
				Authors:   authorsData,
			}

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
			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
