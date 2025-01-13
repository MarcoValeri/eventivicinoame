package controllers

import (
	admincontrollers "eventivicinoame/controllers/adminControllers"
	"net/http"
)

func AdminController(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/admin/dashboard":
		admincontrollers.AdminDashboard(w, r)
	case "/admin/admin-users":
		admincontrollers.AdminUsers(w, r)
	case "/admin/admin-user-add":
		admincontrollers.AdminUserAdd(w, r)
	case "/admin/admin-sagre/":
		admincontrollers.AdminSagre(w, r)
	case "/admin/admin-sagra-add":
		admincontrollers.AdminSagraAdd(w, r)
	case "/admin/admin-sagra-edit/":
		admincontrollers.AdminSagraEdit(w, r)
	case "/admin/admin-sagra-delete/":
		admincontrollers.AdminSagraDelete(w, r)
	case "/admin/admin-sagre-checker/":
		admincontrollers.AdminSagreChecker(w, r)
	case "/admin/admin-sagre-search/":
		admincontrollers.AdminSagreSearch(w, r)
	default:
		http.NotFound(w, r)
	}

	// admincontrollers.AdminEvents()
	// admincontrollers.AdminEventAdd()
	// admincontrollers.AdminEventEdit()
	// admincontrollers.AdminEventDelete()
	// admincontrollers.AdminEventsSearch()
	// admincontrollers.AdminEventsChecker()

	// admincontrollers.AdminNews()
	// admincontrollers.AdminNewsAdd()
	// admincontrollers.AdminNewsEdit()
	// admincontrollers.AdminNewsDelete()

	// admincontrollers.AdminImages()
	// admincontrollers.AdminImageAdd()
	// admincontrollers.AdminImageEdit()
	// admincontrollers.AdminImageDelete()
	// admincontrollers.AdminImageAddOnlyFile()
}
