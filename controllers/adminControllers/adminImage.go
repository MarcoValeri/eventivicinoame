package admincontrollers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type imageData struct {
	PageTitle             string
	ImageTitleError       string
	ImageUrlError         string
	ImageDescriptionError string
	ImageCreditError      string
	ImagePublishedError   string
	ImageUpdatedError     string
	PreviusButton         bool
	NextButton            bool
	PreviousPage          string
	NextPage              string
	Images                []models.Image
}

func AdminImages() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-images.html"))
	http.HandleFunc("/admin/admin-images/", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			urlPath := strings.TrimPrefix(r.URL.Path, "/admin/admin-images/")
			urlPath = util.FormSanitizeStringInput(urlPath)

			pageNumber, err := strconv.Atoi(urlPath)
			if err != nil {
				fmt.Println("Error convertinf string to integer:", err)
				return
			}

			// Redirect to /admin/admin-images/1 if pageNumber is 0
			if pageNumber == 0 {
				http.Redirect(w, r, "/admin/admin-images/1", http.StatusSeeOther)
			}

			// Set limit and offset for MySQL query
			limit := 10
			offset := (pageNumber - 1) * limit

			// Get images from db
			getAllImages, err := models.ImagesGetLimitAndPagination(limit, offset)
			if err != nil {
				fmt.Println("Error getting imageData:", err)
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
			if len(getAllImages) >= 10 {
				setNextButton = true
				setNextPage = pageNumber + 1
				setNextPageStr = strconv.Itoa(setNextPage)
			}

			data := imageData{
				PageTitle:     "Images Admin",
				PreviusButton: setPreviousButton,
				NextButton:    setNextButton,
				PreviousPage:  setPreviousPageStr,
				NextPage:      setNextPageStr,
				Images:        getAllImages,
			}

			tmpl.Execute(w, data)

		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
