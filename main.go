package main

import (
	"eventivicinoame/controllers"
	admincontrollers "eventivicinoame/controllers/adminControllers"
	"net/http"
	// psh "github.com/platformsh/gohelper"
)

func main() {
	// PlatformSH
	// platformSH, err := psh.NewPlatformInfo()
	// if err != nil {
	// 	panic("Not in a Platform.sh environment")
	// }

	// Static files
	fileServer := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()

	// Handle static files
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	mux.Handle("/admin/", controllers.AuthMiddlewareController(http.HandlerFunc(controllers.AdminController)))

	mux.HandleFunc("/admin/login", admincontrollers.AdminLogin)

	mux.HandleFunc("/", controllers.Home)

	// if r.MethodGet == http.MethodGet {

	// } else {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	// OLD
	// Static files
	// fs := http.FileServer(http.Dir("./public"))
	// http.Handle("/public/", http.StripPrefix("/public/", fs))

	// // Controllers
	// controllers.Home()
	// controllers.AboutUs()
	// controllers.Contact()
	// controllers.CookiePolicy()
	// controllers.PrivacyPolicy()

	// controllers.SagreSearchController()
	// controllers.SagraController()
	// controllers.SagreJanuary()
	// controllers.SagreFebruary()
	// controllers.SagreOctober()
	// controllers.SagreNovember()
	// controllers.SagreDecember()
	// controllers.SagreAutumn()

	// controllers.EventsSearchController()
	// controllers.EventController()
	// controllers.EventsJanuary()
	// controllers.EventsFebruary()
	// controllers.EventsNovember()
	// controllers.EventsDecember()
	// controllers.EventsMercatiniDiNatale()

	// controllers.NewsSearchController()
	// controllers.NewsController()

	// controllers.AuthorController()
	// controllers.AdminController()

	// controllers.SitemapController()
	// controllers.RobotController()
	// controllers.AdsenseController()

	// controllers.Error404()

	// /**
	// * DB connection
	// * parameter "platform" connect to Platform.sh
	// * parameter "local" connect to local db
	//  */
	// database.DatabaseConnection()

	// // Local env
	// http.ListenAndServe(":80", nil)

	// Platform SH env
	// http.ListenAndServe(":"+platformSH.Port, nil)
}
