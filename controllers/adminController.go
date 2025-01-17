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
	case "/admin/admin-events/":
		admincontrollers.AdminEvents(w, r)
	case "/admin/admin-event-add":
		admincontrollers.AdminEventAdd(w, r)
	case "/admin/admin-event-edit/":
		admincontrollers.AdminEventEdit(w, r)
	case "/admin/admin-event-delete/":
		admincontrollers.AdminEventDelete(w, r)
	case "/admin/admin-events-checker/":
		admincontrollers.AdminEventsChecker(w, r)
	case "/admin/admin-events-search/":
		admincontrollers.AdminEventsSearch(w, r)
	case "/admin/admin-news/":
		admincontrollers.AdminNews(w, r)
	case "/admin/admin-news-add":
		admincontrollers.AdminNewsAdd(w, r)
	case "/admin/admin-news-edit/":
		admincontrollers.AdminNewsEdit(w, r)
	case "/admin/admin-news-delete/":
		admincontrollers.AdminNewsDelete(w, r)
	case "/admin/admin-images/":
		admincontrollers.AdminImages(w, r)
	case "/admin/admin-image-add":
		admincontrollers.AdminImageAdd(w, r)
	case "/admin/admin-image-edit/":
		admincontrollers.AdminImageEdit(w, r)
	case "/admin/admin-image-delete/":
		admincontrollers.AdminImageDelete(w, r)
	case "/admin/admin-image-add-only-file":
		admincontrollers.AdminImageAddOnlyFile(w, r)
	default:
		http.NotFound(w, r)
	}

}
