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
	AuthorId       int
	Country        string
	Region         string
	City           string
	Town           string
	Fraction       string
	SagraStartDate string
	SagraEndDate   string
}

type SagraWithRelatedFields struct {
	Id                int
	Title             string
	Description       string
	Url               string
	Published         string
	Updated           string
	ImageId           int
	ImageUrl          string
	ImageAlt          string
	AuthorId          int
	AuthorName        string
	AuthorSurname     string
	AuthorUrl         string
	AuthorImageUrl    string
	AuthorDescription string
	Content           string
	Country           string
	Region            string
	City              string
	Town              string
	Fraction          string
	SagraStartDate    string
	SagraEndDate      string
}

func SagraNew(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getAuthorId int, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getSagraStartDate string, getSagraEndDate string) Sagra {
	newSagra := Sagra{
		Id:             getId,
		Title:          getTitle,
		Description:    getDescription,
		Url:            getUrl,
		Published:      getPublished,
		Updated:        getUpdated,
		ImageId:        getImageId,
		AuthorId:       getAuthorId,
		Content:        getContent,
		Country:        getCountry,
		Region:         getRegion,
		City:           getCity,
		Town:           getTown,
		Fraction:       getFraction,
		SagraStartDate: getSagraStartDate,
		SagraEndDate:   getSagraEndDate,
	}
	return newSagra
}

func SagraNewWithRelatedFields(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getImageUrl string, getImageAlt string, getAuthorId int, getAuthorName string, getAuthorSurname string, getAuthorUrl string, getAuthorImageUrl string, getAuthorDescription string, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getSagraStartDate string, getSagraEndDate string) SagraWithRelatedFields {
	newSagraWithRelatedImage := SagraWithRelatedFields{
		Id:                getId,
		Title:             getTitle,
		Description:       getDescription,
		Url:               getUrl,
		Published:         getPublished,
		Updated:           getUpdated,
		ImageId:           getImageId,
		ImageUrl:          getImageUrl,
		ImageAlt:          getImageAlt,
		AuthorId:          getAuthorId,
		AuthorName:        getAuthorName,
		AuthorSurname:     getAuthorSurname,
		AuthorUrl:         getAuthorUrl,
		AuthorImageUrl:    getAuthorImageUrl,
		AuthorDescription: getAuthorDescription,
		Content:           getContent,
		Country:           getCountry,
		Region:            getRegion,
		City:              getCity,
		Town:              getTown,
		Fraction:          getFraction,
		SagraStartDate:    getSagraStartDate,
		SagraEndDate:      getSagraEndDate,
	}
	return newSagraWithRelatedImage
}

func SagraAddNewToDB(getSagra Sagra) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"INSERT INTO sagre (title, description, url, published, updated, image_id, author_id, content, country, region, city, town, fraction, sagra_start_date, sagra_end_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		getSagra.Title, getSagra.Description, getSagra.Url, getSagra.Published, getSagra.Updated, getSagra.ImageId, getSagra.AuthorId, getSagra.Content, getSagra.Country, getSagra.Region, getSagra.City, getSagra.Town, getSagra.Fraction, getSagra.SagraStartDate, getSagra.SagraEndDate,
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
		"UPDATE sagre SET title = ?, description = ?, url = ?, published = ?, updated = ?, content = ?, image_id = ?, author_id = ?, country = ?, region = ?, city = ?, town = ?, fraction = ?, sagra_start_date = ?, sagra_end_date = ? WHERE id = ?",
		getSagra.Title, getSagra.Description, getSagra.Url, getSagra.Published, getSagra.Updated, getSagra.Content, getSagra.ImageId, getSagra.AuthorId, getSagra.Country, getSagra.Region, getSagra.City, getSagra.Town, getSagra.Fraction, getSagra.SagraStartDate, getSagra.SagraEndDate, getSagra.Id,
	)
	if err != nil {
		fmt.Println("Error on editing sagra:", err)
		return err
	}
	defer query.Close()

	return nil
}

func SagraFindByUrl(getSagraUrl string) (SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	var getSagraData SagraWithRelatedFields

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id WHERE sagre.url=? AND sagre.published < NOW()", getSagraUrl)
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return getSagraData, err
		}

		getSagraData = SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
	}

	return getSagraData, nil
}

func SagraDelete(getSagraId int) error {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("DELETE FROM sagre WHERE id = ? LIMIT 1", getSagraId)
	if err != nil {
		fmt.Println("Error, no able to delete this sagra:", err)
		return err
	}
	defer rows.Close()

	return nil
}

func SagraShowSagre() ([]SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id")
	if err != nil {
		fmt.Println("Error getting sagre from the db:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedFields
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		allSagre = append(allSagre, sagraDetails)
	}

	return allSagre, nil
}

func SagraWithRelatedImageFindById(getSagraId int) (SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	var getSagraData SagraWithRelatedFields

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id WHERE sagre.id=?", getSagraId)
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return getSagraData, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		getSagraData = sagraDetails
	}
	return getSagraData, nil
}

func SagraGetLimitPublishedSagre(getLimit int) ([]SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id WHERE sagre.published < NOW() ORDER BY sagre.updated DESC LIMIT ?", getLimit)
	if err != nil {
		fmt.Println("Error getting published sagre with limit:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedFields
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		allSagre = append(allSagre, sagraDetails)
	}

	return allSagre, nil
}

func SagreGetLimitAndPagination(getLimit, getPageNumber int) ([]SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id ORDER BY sagre.published DESC LIMIT ? OFFSET ?", getLimit, getPageNumber)
	if err != nil {
		fmt.Println("Error getting LimitAndPagination sagre:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedFields
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		allSagre = append(allSagre, sagraDetails)
	}

	return allSagre, nil
}

func SagraFindByParameter(getParameter string) ([]SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	query := "SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id WHERE (sagre.title LIKE ? OR sagre.description LIKE ? OR sagre.content LIKE ?) AND sagre.published < NOW() ORDER BY sagre.updated DESC LIMIT ?"
	likePattern := "%" + getParameter + "%"

	rows, err := db.Query(query, likePattern, likePattern, likePattern, 10)
	if err != nil {
		fmt.Println("Error on the sagre query:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedFields
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		allSagre = append(allSagre, sagraDetails)
	}
	return allSagre, nil
}

func SagreGetThemByPeriodOfTime(getStartDate string, getEndDate string, getLimit int) ([]SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id WHERE ((sagre.sagra_start_date >= ? AND sagre.sagra_start_date <= ?) OR (sagre.sagra_end_date >= ? AND sagre.sagra_end_date <= ?) OR (sagre.sagra_start_date <= ? AND sagre.sagra_end_date >= ?)) AND sagre.published < NOW() ORDER BY sagre.updated DESC LIMIT ?", getStartDate, getEndDate, getStartDate, getEndDate, getStartDate, getEndDate, getLimit)
	if err != nil {
		fmt.Println("Error getting sagre by period of time:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedFields
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		allSagre = append(allSagre, sagraDetails)
	}
	return allSagre, nil
}

func SagreGetAllPassed(getCurrentDate string, getLimit int, getOffset int) ([]SagraWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	mySqlQuery := "SELECT sagre.id, sagre.title, sagre.description, sagre.url, sagre.published, sagre.updated, sagre.image_id, images.url, images.description, sagre.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, sagre.content, sagre.country, sagre.region, sagre.city, sagre.town, sagre.fraction, sagre.sagra_start_date, sagre.sagra_end_date FROM sagre JOIN images ON sagre.image_id = images.id JOIN authors ON sagre.author_id = authors.id WHERE sagre.sagra_end_date <= ? AND sagre.published < NOW() ORDER BY sagre.sagra_end_date DESC LIMIT ? OFFSET ?"
	rows, err := db.Query(mySqlQuery, getCurrentDate, getLimit, getOffset)
	if err != nil {
		fmt.Println("Error getting passed sagre:", err)
		return nil, err
	}
	defer rows.Close()

	var allSagre []SagraWithRelatedFields
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
		var sagraAuthorId int
		var sagraAuthorName string
		var sagraAuthorSurname string
		var sagraAuthorUrl string
		var sagraAuthorImageUrl string
		var sagraAuthorDescription string
		var sagraContent string
		var sagraCountry string
		var sagraRegion string
		var sagraCity string
		var sagraTown string
		var sagraFraction string
		var sagraStartDate string
		var sagraEndDate string
		err = rows.Scan(&sagraId, &sagraTitle, &sagraDescription, &sagraUrl, &sagraPublished, &sagraUpdated, &sagraImageId, &sagraImageUrl, &sagraImageAlt, &sagraAuthorId, &sagraAuthorName, &sagraAuthorSurname, &sagraAuthorUrl, &sagraAuthorImageUrl, &sagraAuthorDescription, &sagraContent, &sagraCountry, &sagraRegion, &sagraCity, &sagraTown, &sagraFraction, &sagraStartDate, &sagraEndDate)
		if err != nil {
			return allSagre, err
		}

		sagraDetails := SagraNewWithRelatedFields(
			sagraId,
			sagraTitle,
			sagraDescription,
			sagraUrl,
			sagraPublished,
			sagraUpdated,
			sagraImageId,
			sagraImageUrl,
			sagraImageAlt,
			sagraAuthorId,
			sagraAuthorName,
			sagraAuthorSurname,
			sagraAuthorUrl,
			sagraAuthorImageUrl,
			sagraAuthorDescription,
			sagraContent,
			sagraCountry,
			sagraRegion,
			sagraCity,
			sagraTown,
			sagraFraction,
			sagraStartDate,
			sagraEndDate,
		)
		allSagre = append(allSagre, sagraDetails)
	}
	return allSagre, nil
}
