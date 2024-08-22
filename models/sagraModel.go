package models

import (
	"eventivicinoame/database"
	"fmt"
)

type Sagra struct {
	Id             int
	Title          string
	Description    string
	Url            string
	Published      string
	Updated        string
	Content        string
	Country        string
	Region         string
	City           string
	Town           string
	Fraction       string
	SagraStartDate string
}

type SagraWithRelatedImage struct {
	Id             int
	Title          string
	Description    string
	Url            string
	Published      string
	Updated        string
	ImageId        int
	ImageUrl       string
	ImageAlt       string
	Content        string
	Country        string
	Region         string
	City           string
	Town           string
	Fraction       string
	SagraStartDate string
}

func SagraNew(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getSagraStartDate string) Sagra {
	newSagra := Sagra{
		Id:             getId,
		Title:          getTitle,
		Description:    getDescription,
		Url:            getUrl,
		Published:      getPublished,
		Updated:        getUpdated,
		Content:        getContent,
		Country:        getCountry,
		Region:         getRegion,
		City:           getCity,
		Town:           getTown,
		Fraction:       getFraction,
		SagraStartDate: getSagraStartDate,
	}
	return newSagra
}

func SagraNewWithRelatedImage(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getImageUrl string, getImageAlt string, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getSagraStartDate string) SagraWithRelatedImage {
	newSagraWithRelatedImage := SagraWithRelatedImage{
		Id:             getId,
		Title:          getTitle,
		Description:    getDescription,
		Url:            getUrl,
		Published:      getPublished,
		Updated:        getUpdated,
		ImageId:        getImageId,
		ImageUrl:       getImageUrl,
		ImageAlt:       getImageAlt,
		Content:        getContent,
		Country:        getCountry,
		Region:         getRegion,
		City:           getCity,
		Town:           getTown,
		Fraction:       getFraction,
		SagraStartDate: getSagraStartDate,
	}
	return newSagraWithRelatedImage
}

func SagraFindByUrl(getSagraUrl string) (SagraWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	var getSagraData SagraWithRelatedImage

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date FROM sagre JOIN images ON sagre.image_id = images.id WHERE sagre.url=? AND sagre.published < NOW()", getSagraUrl)
	if err != nil {
		fmt.Println("Error on the sagra query:", err)
		return getSagraData, err
	}
	defer rows.Close()

	for rows.Next() {
		var sagraId int
		var sagraTitle string
		var sagraDescription string
		var sagraUrl string
		var sagraPublished string
		var sagraUpdated string
		var sagraImageId int
		var sagraImageUrl string
		var sagraImageAlt string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate)
		if err != nil {
			return getSagraData, err
		}

		getSagraData = SagraNewWithRelatedImage(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
		)
	}

	return getSagraData, nil
}
