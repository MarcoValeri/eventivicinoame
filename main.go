package main

import (
	"eventivicinoame/controllers"
	admincontrollers "eventivicinoame/controllers/adminControllers"
	"eventivicinoame/database"
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

	mux.HandleFunc("/page/chi-siamo", controllers.AboutUs)
	mux.HandleFunc("/page/contatti", controllers.Contact)
	mux.HandleFunc("/page/cookie-policy", controllers.CookiePolicy)
	mux.HandleFunc("/page/privacy-policy", controllers.PrivacyPolicy)

	mux.HandleFunc("/sagre-cerca/", controllers.SagreSearchController)
	mux.HandleFunc("/sagra/", controllers.SagraController)
	mux.HandleFunc("/sagre/sagre-gennaio", controllers.SagreJanuary)
	mux.HandleFunc("/sagre/sagre-febbraio", controllers.SagreFebruary)
	mux.HandleFunc("/sagre/sagre-ottobre", controllers.SagreOctober)
	mux.HandleFunc("/sagre/sagre-novembre", controllers.SagreNovember)
	mux.HandleFunc("/sagre/sagre-dicembre", controllers.SagreDecember)
	mux.HandleFunc("/sagre/sagre-autunno", controllers.SagreAutumn)

	mux.HandleFunc("/eventi-cerca/", controllers.EventsSearchController)
	mux.HandleFunc("/evento/", controllers.EventController)
	mux.HandleFunc("/eventi/mercatini-di-natale", controllers.EventsMercatiniDiNatale)
	mux.HandleFunc("/eventi/eventi-gennaio", controllers.EventsJanuary)
	mux.HandleFunc("/eventi/eventi-febbraio", controllers.EventsFebruary)
	mux.HandleFunc("/eventi/eventi-novembre", controllers.EventsNovember)
	mux.HandleFunc("/eventi/eventi-dicembre", controllers.EventsDecember)

	mux.HandleFunc("/news-cerca/", controllers.NewsSearchController)
	mux.HandleFunc("/news/", controllers.NewsController)

	mux.HandleFunc("/author/", controllers.AuthorController)

	mux.HandleFunc("/sitemap.xml", controllers.SitemapController)
	mux.HandleFunc("/robots.txt", controllers.RobotController)

	mux.HandleFunc("/error/error-404", controllers.Error404)

	/**
	* DB connection
	* parameter "platform" connect to Platform.sh
	* parameter "local" connect to local db
	 */
	database.DatabaseConnection()

	// Local env
	http.ListenAndServe(":80", mux)

	// Platform SH env
	// http.ListenAndServe(":"+platformSH.Port, mux)

	// OLD - OLD
	// if r.Method == http.MethodGet {

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
