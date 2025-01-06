package models

import (
	"eventivicinoame/database"
	"fmt"
)

type SitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

func SitemapAllURL() ([]SitemapURL, error) {
	var setURLsList []SitemapURL

	// Set URLs that are not stored in the db
	url_0 := SitemapURL{"https://www.eventivicinoame.com/", "2024-10-30"}
	url_1 := SitemapURL{"https://www.eventivicinoame.com/sagre-cerca/", "2024-09-26"}
	url_2 := SitemapURL{"https://www.eventivicinoame.com/page/chi-siamo", "2024-09-20"}
	url_3 := SitemapURL{"https://www.eventivicinoame.com/page/contatti", "2024-09-20"}
	url_4 := SitemapURL{"https://www.eventivicinoame.com/page/cookie-policy", "2024-09-20"}
	url_5 := SitemapURL{"https://www.eventivicinoame.com/page/privacy-policy", "2024-09-20"}
	url_6 := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-gennaio", "2024-12-15"}
	url_7 := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-febbraio", "2025-01-06"}
	url_8 := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-ottobre", "2024-12-15"}
	url_9 := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-novembre", "2024-12-15"}
	url_10 := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-dicembre", "2025-01-06"}
	url_11 := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-autunno", "2024-10-21"}
	url_12 := SitemapURL{"https://www.eventivicinoame.com/author/marco-valeri", "2024-10-11"}
	url_13 := SitemapURL{"https://www.eventivicinoame.com/eventi-cerca/", "2024-10-18"}
	url_14 := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-gennaio", "2024-12-15"}
	url_15 := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-febbraio", "2025-01-06"}
	url_16 := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-novembre", "2024-12-15"}
	url_17 := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-dicembre", "2025-01-06"}
	url_18 := SitemapURL{"https://www.eventivicinoame.com/eventi/mercatini-di-natale", "2025-01-06"}
	url_19 := SitemapURL{"https://www.eventivicinoame.com/news-cerca/", "2024-11-25"}

	setURLsList = append(setURLsList, url_0)
	setURLsList = append(setURLsList, url_1)
	setURLsList = append(setURLsList, url_2)
	setURLsList = append(setURLsList, url_3)
	setURLsList = append(setURLsList, url_4)
	setURLsList = append(setURLsList, url_5)
	setURLsList = append(setURLsList, url_6)
	setURLsList = append(setURLsList, url_7)
	setURLsList = append(setURLsList, url_8)
	setURLsList = append(setURLsList, url_9)
	setURLsList = append(setURLsList, url_10)
	setURLsList = append(setURLsList, url_11)
	setURLsList = append(setURLsList, url_12)
	setURLsList = append(setURLsList, url_13)
	setURLsList = append(setURLsList, url_14)
	setURLsList = append(setURLsList, url_15)
	setURLsList = append(setURLsList, url_16)
	setURLsList = append(setURLsList, url_17)
	setURLsList = append(setURLsList, url_18)
	setURLsList = append(setURLsList, url_19)

	// Get all sagre URLs
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT url, updated FROM sagre WHERE published < NOW()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urlSagra SitemapURL
	for rows.Next() {
		var sagraUrl string
		var sagraUpdated string
		err = rows.Scan(&sagraUrl, &sagraUpdated)
		if err != nil {
			return nil, err
		}
		urlSagra.Loc = "https://www.eventivicinoame.com/sagra/" + sagraUrl
		urlSagra.LastMod = sagraUpdated[:10]
		setURLsList = append(setURLsList, urlSagra)
	}

	// Get all events URLs
	rowsEvents, errEvents := db.Query("SELECT url, updated FROM events WHERE published < NOW()")
	if errEvents != nil {
		return nil, err
	}
	defer rowsEvents.Close()

	var urlEvent SitemapURL
	for rowsEvents.Next() {
		var eventUrl string
		var eventUpdated string
		err = rowsEvents.Scan(&eventUrl, &eventUpdated)
		if err != nil {
			return nil, err
		}
		urlEvent.Loc = "https://www.eventivicinoame.com/evento/" + eventUrl
		urlEvent.LastMod = eventUpdated[:10]
		setURLsList = append(setURLsList, urlEvent)
	}

	// Get all news URLs
	rowsNews, errNews := db.Query("SELECT url, updated FROM news WHERE published < NOW()")
	if errNews != nil {
		return nil, err
	}
	defer rowsNews.Close()

	var urlNews SitemapURL
	for rowsNews.Next() {
		var newsUrl string
		var newsUpdated string
		err = rowsNews.Scan(&newsUrl, &newsUpdated)
		if err != nil {
			return nil, err
		}
		urlNews.Loc = "https://www.eventivicinoame.com/news/" + newsUrl
		urlNews.LastMod = newsUpdated[:10]
		setURLsList = append(setURLsList, urlNews)
	}

	// Get all the images
	rowsImage, errImage := db.Query("SELECT url, updated FROM images WHERE published < NOW()")
	if errImage != nil {
		fmt.Println("Error to query images for sitemapModel:", errImage)
		return nil, errImage
	}
	defer rowsImage.Close()

	var urlImage SitemapURL
	for rowsImage.Next() {
		var imageUrl string
		var imageUpdated string
		errImage = rowsImage.Scan(&imageUrl, &imageUpdated)
		if errImage != nil {
			fmt.Println("Error saveing image data firn the sitemap:", errImage)
			return nil, errImage
		}
		urlImage.Loc = "https://www.devwithgo.dev/public/images/" + imageUrl
		urlImage.LastMod = imageUpdated[:10]
		setURLsList = append(setURLsList, urlImage)
	}

	return setURLsList, nil
}
