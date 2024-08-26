package admincontrollers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AdminDashboard() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-dashboard.html"))
	http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {
			tmpl.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
