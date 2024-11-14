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
	admincontrollers.AdminSagreSearch()
	admincontrollers.AdminSagreChecker()

	admincontrollers.AdminEvents()
	admincontrollers.AdminEventAdd()
	admincontrollers.AdminEventEdit()
	admincontrollers.AdminEventDelete()
	admincontrollers.AdminEventsSearch()
	admincontrollers.AdminEventsChecker()

	admincontrollers.AdminImages()
	admincontrollers.AdminImageAdd()
	admincontrollers.AdminImageEdit()
}
