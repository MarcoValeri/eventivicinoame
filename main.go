package main

import (
	"eventivicinoame/controllers"
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
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Controllers
	controllers.Home()
	controllers.AboutUs()
	controllers.Contact()
	controllers.CookiePolicy()
	controllers.PrivacyPolicy()

	controllers.SagreSearchController()
	controllers.SagraController()
	controllers.SagreOctober()
	controllers.SagreNovember()
	controllers.SagreDecember()
	controllers.SagreAutumn()

	controllers.EventsSearchController()
	controllers.EventController()

	controllers.AuthorController()
	controllers.AdminController()

	controllers.SitemapController()
	controllers.RobotController()

	controllers.Error404()

	/**
	* DB connection
	* parameter "platform" connect to Platform.sh
	* parameter "local" connect to local db
	 */
	database.DatabaseConnection()

	// Local env
	http.ListenAndServe(":80", nil)

	// Platform SH env
	// http.ListenAndServe(":"+platformSH.Port, nil)
}
