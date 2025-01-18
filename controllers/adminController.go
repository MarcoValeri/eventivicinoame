package controllers

import (
	admincontrollers "eventivicinoame/controllers/adminControllers"
	"net/http"
	"strings"
)

func AdminController(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/admin/dashboard":
		admincontrollers.AdminDashboard(w, r)
	case r.URL.Path == "/admin/admin-users":
		admincontrollers.AdminUsers(w, r)
	case r.URL.Path == "/admin/admin-user-add":
		admincontrollers.AdminUserAdd(w, r)
	case strings.HasPrefix(r.URL.Path, "/admin/admin-sagre/"):
		admincontrollers.AdminSagre(w, r)
	case r.URL.Path == "/admin/admin-sagra-add":
		admincontrollers.AdminSagraAdd(w, r)
	case r.URL.Path == "/admin/admin-sagra-edit/":
		admincontrollers.AdminSagraEdit(w, r)
	case r.URL.Path == "/admin/admin-sagra-delete/":
		admincontrollers.AdminSagraDelete(w, r)
	case r.URL.Path == "/admin/admin-sagre-checker/":
		admincontrollers.AdminSagreChecker(w, r)
	case r.URL.Path == "/admin/admin-sagre-search/":
		admincontrollers.AdminSagreSearch(w, r)
	case r.URL.Path == "/admin/admin-events/":
		admincontrollers.AdminEvents(w, r)
	case r.URL.Path == "/admin/admin-event-add":
		admincontrollers.AdminEventAdd(w, r)
	case r.URL.Path == "/admin/admin-event-edit/":
		admincontrollers.AdminEventEdit(w, r)
	case r.URL.Path == "/admin/admin-event-delete/":
		admincontrollers.AdminEventDelete(w, r)
	case r.URL.Path == "/admin/admin-events-checker/":
		admincontrollers.AdminEventsChecker(w, r)
	case r.URL.Path == "/admin/admin-events-search/":
		admincontrollers.AdminEventsSearch(w, r)
	case r.URL.Path == "/admin/admin-news/":
		admincontrollers.AdminNews(w, r)
	case r.URL.Path == "/admin/admin-news-add":
		admincontrollers.AdminNewsAdd(w, r)
	case r.URL.Path == "/admin/admin-news-edit/":
		admincontrollers.AdminNewsEdit(w, r)
	case r.URL.Path == "/admin/admin-news-delete/":
		admincontrollers.AdminNewsDelete(w, r)
	case r.URL.Path == "/admin/admin-images/":
		admincontrollers.AdminImages(w, r)
	case r.URL.Path == "/admin/admin-image-add":
		admincontrollers.AdminImageAdd(w, r)
	case r.URL.Path == "/admin/admin-image-edit/":
		admincontrollers.AdminImageEdit(w, r)
	case r.URL.Path == "/admin/admin-image-delete/":
		admincontrollers.AdminImageDelete(w, r)
	case r.URL.Path == "/admin/admin-image-add-only-file":
		admincontrollers.AdminImageAddOnlyFile(w, r)
	default:
		http.NotFound(w, r)
	}

}
