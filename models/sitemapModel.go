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
	urlZero := SitemapURL{"https://www.eventivicinoame.com/", "2024-10-30"}
	urlOne := SitemapURL{"https://www.eventivicinoame.com/sagre-cerca/", "2024-09-26"}
	urlThree := SitemapURL{"https://www.eventivicinoame.com/page/chi-siamo", "2024-09-20"}
	urlFour := SitemapURL{"https://www.eventivicinoame.com/page/contatti", "2024-09-20"}
	urlFive := SitemapURL{"https://www.eventivicinoame.com/page/cookie-policy", "2024-09-20"}
	urlSix := SitemapURL{"https://www.eventivicinoame.com/page/privacy-policy", "2024-09-20"}
	urlSeven := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-ottobre", "2024-09-26"}
	urlEight := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-gennaio", "2024-11-29"}
	urlNine := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-novembre", "2024-10-05"}
	urlTen := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-dicembre", "2024-10-21"}
	urlEleven := SitemapURL{"https://www.eventivicinoame.com/sagre/sagre-autunno", "2024-10-21"}
	urlTwelve := SitemapURL{"https://www.eventivicinoame.com/author/marco-valeri", "2024-10-11"}
	urlThirteen := SitemapURL{"https://www.eventivicinoame.com/eventi-cerca/", "2024-10-18"}
	urlFourteen := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-gennaio", "2024-11-29"}
	urlFifteen := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-novembre", "2024-10-30"}
	urlSixteen := SitemapURL{"https://www.eventivicinoame.com/eventi/eventi-dicembre", "2024-10-30"}
	urlSeventeen := SitemapURL{"https://www.eventivicinoame.com/eventi/mercatini-di-natale", "2024-10-30"}
	urlEighteen := SitemapURL{"https://www.eventivicinoame.com/news-cerca/", "2024-11-25"}

	setURLsList = append(setURLsList, urlZero)
	setURLsList = append(setURLsList, urlOne)
	setURLsList = append(setURLsList, urlThree)
	setURLsList = append(setURLsList, urlFour)
	setURLsList = append(setURLsList, urlFive)
	setURLsList = append(setURLsList, urlSix)
	setURLsList = append(setURLsList, urlSeven)
	setURLsList = append(setURLsList, urlEight)
	setURLsList = append(setURLsList, urlNine)
	setURLsList = append(setURLsList, urlTen)
	setURLsList = append(setURLsList, urlEleven)
	setURLsList = append(setURLsList, urlTwelve)
	setURLsList = append(setURLsList, urlThirteen)
	setURLsList = append(setURLsList, urlFourteen)
	setURLsList = append(setURLsList, urlFifteen)
	setURLsList = append(setURLsList, urlSixteen)
	setURLsList = append(setURLsList, urlSeventeen)
	setURLsList = append(setURLsList, urlEighteen)

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
