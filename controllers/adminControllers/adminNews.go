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
	PreviousButton                 bool
	NextButton                     bool
	PreviousPage                   string
	NextPage                       string
	Images                         []models.Image
	Author                         []models.Author
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
