package controllers

import (
	"eventivicinoame/models"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)

type HomepageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
	CurrentUrl      string
	Sagre           []models.SagraWithRelatedFields
}

func Home() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Redirect to 404 page if route does not exist
		if r.URL.Path != "/" {
			http.Redirect(w, r, "/error/error-404", http.StatusSeeOther)
		}

		// Get last three published sagre
		getLastPublishedSagre, err := models.SagraGetLimitPublishedSagre(10)
		if err != nil {
			fmt.Println("Error getting last ten sagre:", err)
		}

		// Get current path
		currentUrlPath := path.Clean(r.URL.Path)

		data := HomepageData{
			PageTitle:       "Eventi vicino a me: oggi, domani e nel fine settimana",
			PageDescription: "Eventi vicino a me: sagre, feste, fiere, mercatini, mostre e musei oggi, domani e nel fine settimana, più gli eventi da non perdere il prossimo weekend",
			CurrentYear:     time.Now().Year(),
			CurrentUrl:      currentUrlPath,
			Sagre:           getLastPublishedSagre,
		}
		tmpl.Execute(w, data)
	})
}
