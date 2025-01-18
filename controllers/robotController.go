package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

func RobotController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain")

		output := []string{
			"User-agent: GPTBot",
			"Disallow: /",
			"",
			"User-agent: ChatGPT-User",
			"Disallow: /",
			"",
			"User-agent: Applebot-Extended",
			"Disallow: /",
			"",
			"User-agent: *",
			"Disallow:",
			"",
			"Disallow: /admin/",
			"",
			"Sitemap: https://www.devwithgo.dev/sitemap.xml",
		}

		fmt.Fprint(w, strings.Join(output, "\n"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// func RobotController() {
// 	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain")

// 		output := []string{
// 			"User-agent: GPTBot",
// 			"Disallow: /",
// 			"",
// 			"User-agent: ChatGPT-User",
// 			"Disallow: /",
// 			"",
// 			"User-agent: Applebot-Extended",
// 			"Disallow: /",
// 			"",
// 			"User-agent: *",
// 			"Disallow:",
// 			"",
// 			"Disallow: /admin/",
// 			"",
// 			"Sitemap: https://www.devwithgo.dev/sitemap.xml",
// 		}

// 		fmt.Fprint(w, strings.Join(output, "\n"))
// 	})
// }
