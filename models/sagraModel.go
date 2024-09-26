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
	ImageId        int
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

func SagraNew(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getSagraStartDate string) Sagra {
	newSagra := Sagra{
		Id:             getId,
		Title:          getTitle,
		Description:    getDescription,
		Url:            getUrl,
		Published:      getPublished,
		Updated:        getUpdated,
		ImageId:        getImageId,
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

func SagraAddNewToDB(getSagra Sagra) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"INSERT INTO sagre (title, description, url, published, updated, image_id, content, country, region, city, town, fraction, sagra_start_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		getSagra.Title, getSagra.Description, getSagra.Url, getSagra.Published, getSagra.Updated, getSagra.ImageId, getSagra.Content, getSagra.Country, getSagra.Region, getSagra.City, getSagra.Town, getSagra.Fraction, getSagra.SagraStartDate,
	)
	if err != nil {
		fmt.Println("Error adding a new sagra:", err)
		return err
	}
	defer query.Close()

	return nil
}

func SagraEdit(getSagra Sagra) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"UPDATE sagre SET title = ?, description = ?, url = ?, published = ?, updated = ?, content = ?, image_id = ?, country = ?, region = ?, city = ?, town = ?, fraction = ?, sagra_start_date = ? WHERE id = ?",
		getSagra.Title, getSagra.Description, getSagra.Url, getSagra.Published, getSagra.Updated, getSagra.Content, getSagra.ImageId, getSagra.Country, getSagra.Region, getSagra.City, getSagra.Town, getSagra.Fraction, getSagra.SagraStartDate, getSagra.Id,
	)
	if err != nil {
		fmt.Println("Error on editing sagra:", err)
		return err
	}
	defer query.Close()

	return nil
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

func SagraDelete(getSagraId int) error {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("DELETE FROM sagre WHERE id = ?", getSagraId)
	if err != nil {
		fmt.Println("Error, no able to delete this sagra:", err)
		return err
	}
	defer rows.Close()

	return nil
}

func SagraShowSagre() ([]SagraWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date FROM sagre JOIN images ON sagre.image_id = images.id")
	if err != nil {
		fmt.Println("Error getting sagre from the db:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedImage
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
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedImage(
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
		allSagre = append(allSagre, sagraDetails)
	}

	return allSagre, nil
}

func SagraWithRelatedImageFindById(getSagraId int) (SagraWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	var getSagraData SagraWithRelatedImage

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date FROM sagre JOIN images ON sagre.image_id = images.id WHERE sagre.id=?", getSagraId)
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

		sagraDetails := SagraNewWithRelatedImage(
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
		getSagraData = sagraDetails
	}
	return getSagraData, nil
}

func SagraGetLimitPublishedSagre(getLimit int) ([]SagraWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date FROM sagre JOIN images ON sagre.image_id = images.id WHERE sagre.published < NOW() ORDER BY sagre.updated DESC LIMIT ?", getLimit)
	if err != nil {
		fmt.Println("Error getting published sagre with limit:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedImage
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
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedImage(
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
		allSagre = append(allSagre, sagraDetails)
	}

	return allSagre, nil
}

func SagraFindByParameter(getParameter string) ([]SagraWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	query := "SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date FROM sagre JOIN images ON sagre.image_id = images.id WHERE (sagre.title LIKE ? OR sagre.description LIKE ? OR sagre.content LIKE ?) AND sagre.published < NOW() ORDER BY sagre.updated DESC LIMIT ?"
	likePattern := "%" + getParameter + "%"

	rows, err := db.Query(query, likePattern, likePattern, likePattern, 10)
	if err != nil {
		fmt.Println("Error on the sagre query:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedImage
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
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedImage(
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
		allSagre = append(allSagre, sagraDetails)
	}
	return allSagre, nil
}

func SagreGetThemByMonth(getMonth string, getLimit int) ([]SagraWithRelatedImage, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date FROM sagre JOIN images ON sagre.image_id = images.id WHERE sagre.published < NOW() AND MONTH(sagre.sagra_start_date) = ? ORDER BY sagre.updated DESC LIMIT ?", getMonth, getLimit)
	if err != nil {
		fmt.Println("Error getting sagre by their month:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedImage
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
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedImage(
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
		allSagre = append(allSagre, sagraDetails)
	}

	return allSagre, nil
}
