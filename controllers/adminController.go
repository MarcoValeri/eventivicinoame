package controllers

import admincontrollers "eventivicinoame/controllers/adminControllers"

func AdminController() {
	admincontrollers.AdminLogin()

	admincontrollers.AdminDashboard()

	admincontrollers.AdminUsers()
	admincontrollers.AdminUserAdd()
}
