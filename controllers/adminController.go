package controllers

import admincontrollers "eventivicinoame/controllers/adminControllers"

func AdminController() {
	admincontrollers.AdminLogin()

	admincontrollers.AdminDashboard()

	admincontrollers.AdminUsers()
	admincontrollers.AdminUserAdd()

	admincontrollers.AdminSagre()
	admincontrollers.AdminSagraAdd()
	admincontrollers.AdminSagraEdit()
	admincontrollers.AdminSagraDelete()

	admincontrollers.AdminEvents()
	admincontrollers.AdminEventAdd()
}
