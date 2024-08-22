package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle       string
	PageDescription string
}

func Home() {
	tmpl := template.Must(template.ParseFiles("./views/templates/base.html", "./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// TEST
		fmt.Println("homeController")

		data := PageData{
			PageTitle:       "Eventi vicino a me",
			PageDescription: "Add a description",
		}
		tmpl.Execute(w, data)
	})
}
