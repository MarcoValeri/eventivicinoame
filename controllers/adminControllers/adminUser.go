package admincontrollers

import (
	"eventivicinoame/models"
	"eventivicinoame/util"
	"fmt"
	"html/template"
	"net/http"
)

type adminDataPage struct {
	PageTitle          string
	EmailError         string
	PasswordError      string
	PassworRepeatError string
	PasswordMath       string
}

func AdminUsers() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-users.html"))
	http.HandleFunc("/admin/admin-users", func(w http.ResponseWriter, r *http.Request) {

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

func AdminUserAdd() {
	tmpl := template.Must(template.ParseFiles("./views/admin/templates/baseAdmin.html", "./views/admin/admin-user-add.html"))
	http.HandleFunc("/admin/admin-user-add", func(w http.ResponseWriter, r *http.Request) {

		session, errSession := store.Get(r, "session-user-admin-authentication")
		if errSession != nil {
			fmt.Println("Error on session-authentication:", errSession)
		}

		if session.Values["admin-user-authentication"] == true {

			data := adminDataPage{
				PageTitle: "Admin User Add",
			}

			// Flag validation
			var areAdminUserInputsValid [5]bool
			isFormSubmittionValid := false

			// Get value from the form
			getAdminUserEmail := r.FormValue("admin-user-email")
			getAdminUserPassword := r.FormValue("admin-user-password")
			getAdminUserPasswordRepeat := r.FormValue("admin-user-password-repeat")
			getAdminUserSubmit := r.FormValue("admin-user-add")

			// Saniteze form inputs
			getAdminUserEmail = util.FormSanitizeStringInput(getAdminUserEmail)
			getAdminUserPassword = util.FormSanitizeStringInput(getAdminUserPassword)
			getAdminUserPasswordRepeat = util.FormSanitizeStringInput(getAdminUserPasswordRepeat)
			getAdminUserSubmit = util.FormSanitizeStringInput(getAdminUserSubmit)

			// Check if the form has been submitted
			if getAdminUserSubmit == "Add new admin user" {
				// Email validation
				if util.FormEmailInput(getAdminUserEmail) {
					data.EmailError = ""
					areAdminUserInputsValid[0] = true
					if util.FormEmailLengthInput(getAdminUserEmail) && areAdminUserInputsValid[0] {
						data.EmailError = ""
						areAdminUserInputsValid[0] = true
					} else {
						data.EmailError = "Email length is not valid"
						areAdminUserInputsValid[0] = false
					}
				} else {
					data.EmailError = "Email format is not valid"
					areAdminUserInputsValid[0] = false
				}

				// Password validation
				if util.FormPasswordInput(getAdminUserPassword) {
					data.PasswordError = ""
					areAdminUserInputsValid[1] = true
				} else {
					data.PasswordError = "Password should be between 8 to 20 characters"
					areAdminUserInputsValid[1] = false
				}

				if util.FormPasswordInput(getAdminUserPasswordRepeat) {
					data.PassworRepeatError = ""
					areAdminUserInputsValid[2] = true
				} else {
					data.PassworRepeatError = "Password should be between 8 to 20 characters"
					areAdminUserInputsValid[2] = false
				}

				if getAdminUserPassword == getAdminUserPasswordRepeat {
					data.PasswordMath = ""
					areAdminUserInputsValid[3] = true
				} else {
					data.PasswordMath = "Password and repeat password do not match"
					areAdminUserInputsValid[3] = false
				}

				// Submit validation
				if getAdminUserSubmit == "Add new admin user" {
					areAdminUserInputsValid[4] = true
				} else {
					areAdminUserInputsValid[4] = false
				}

				for i := 0; i < len(areAdminUserInputsValid); i++ {
					isFormSubmittionValid = true
					if !areAdminUserInputsValid[i] {
						isFormSubmittionValid = false
						break
					}
				}

				// Create a new user if all inputs are valid
				if isFormSubmittionValid {
					createNewUserAdmin := models.UserAdminNew(1, getAdminUserEmail, getAdminUserPassword)
					models.UserAdminAddNewToDB(createNewUserAdmin)
					http.Redirect(w, r, "/admin/admin-users", http.StatusSeeOther)
				}
			}

			tmpl.Execute(w, data)
		} else {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
	})
}
