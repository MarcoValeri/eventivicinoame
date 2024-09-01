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
	urlZero := SitemapURL{"https://www.eventivicinoame.com/", "2024-08-29"}
	urlOne := SitemapURL{"https://www.eventivicinoame.com/sagre/", "2024-09-01"}
	setURLsList = append(setURLsList, urlZero)
	setURLsList = append(setURLsList, urlOne)

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
		urlSagra.Loc = "https://www.eventivicinoame.com/" + sagraUrl
		urlSagra.LastMod = sagraUpdated[:10]
		setURLsList = append(setURLsList, urlSagra)
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
