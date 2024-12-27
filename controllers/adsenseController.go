package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

func AdsenseController() {
	http.HandleFunc("/ads.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		output := []string{
			"google.com, pub-9306947071363993, DIRECT, f08c47fec0942fa0",
		}

		fmt.Fprint(w, strings.Join(output, "\n"))
	})
}
