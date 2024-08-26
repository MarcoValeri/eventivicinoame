package controllers

import (
	"html/template"
	"net/http"
	"time"
)

type PageData struct {
	PageTitle       string
	PageDescription string
	CurrentYear     int
}

func Home() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			PageTitle:       "Eventi vicino a me: oggi, domani e nel fine settimana",
			PageDescription: "Eventi vicino a me: sagre, feste, fiere, mercatini, mostre e musei oggi, domani e nel fine settimana, più gli eventi da non perdere il prossimo weekend",
			CurrentYear:     time.Now().Year(),
		}
		tmpl.Execute(w, data)
	})
}
