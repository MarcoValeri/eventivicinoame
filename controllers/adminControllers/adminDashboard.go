package admincontrollers

import (
	"html/template"
	"net/http"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-dashboard.html"))
	tmpl.Execute(w, nil)
	// OLD
	// tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-dashboard.html"))
	// http.HandleFunc("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {

	// 	session, errSession := store.Get(r, "session-user-admin-authentication")
	// 	if errSession != nil {
	// 		fmt.Println("Error on session-authentication:", errSession)
	// 	}

	// 	if session.Values["admin-user-authentication"] == true {
	// 		tmpl.Execute(w, nil)
	// 	} else {
	// 		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	// 	}
	// })
}
