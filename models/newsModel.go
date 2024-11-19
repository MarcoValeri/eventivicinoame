package models

import (
	"eventivicinoame/database"
	"fmt"
)

type News struct {
	Id          int
	Title       string
	Description string
	Url         string
	Published   string
	Updated     string
	ImageId     int
	AuthorId    int
}

type NewsWithRelatedFields struct {
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
}

func NewsNew(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getAuthorId int) News {
	newNews := News{
		Id:          getId,
		Title:       getTitle,
		Description: getDescription,
		Url:         getUrl,
		Published:   getPublished,
		Updated:     getUpdated,
		ImageId:     getImageId,
		AuthorId:    getAuthorId,
	}
	return newNews
}

func NewsNewWithRelatedFileds(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getImageUrl string, getImageAlt string, getAuthorId int, getAuthorName string, getAuthorSurname string, getAuthorUrl string, getAuthorImageUrl string, getAuthorDescription string) NewsWithRelatedFields {
	newNewsWithRelatedFields := NewsWithRelatedFields{
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
	}
	return newNewsWithRelatedFields
}

func NewsGetLimitAndPagination(getLimit int, getPageNumber int) ([]NewsWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	mySqlQuery := "SELECT news.id, news.title, news.description, news.url, news.published, news.updated, news.image_id, images.url, images.description, news.author_id, authors.surname, authors.url, authors.image_url, authors.description FROM news JOIN images ON news.image_id = images.id JOIN authors ON news.author_id = authors.id ORDER BY news.published DESC LIMIT ? OFFSET ?"
	rows, err := db.Query(mySqlQuery, getLimit, getPageNumber)
	if err != nil {
		fmt.Println("Error getting news by limit and pagination:", err)
		return nil, err
	}
	defer rows.Close()

	var allNews []NewsWithRelatedFields
	for rows.Next() {
		var newsId int
		var newsTitle string
		var newsDescription string
		var newsUrl string
		var newsPublished string
		var newUpdated string
		var newsImageId int
		var newsImageUrl string
		var newsImageAlt string
		var newsAuthorId int
		var newsAuthorName string
		var newsAuthorSurname string
		var newsAuthorUrl string
		var newsAuthorImageUrl string
		var newsAuthorDescription string
		err = rows.Scan(&newsId, &newsTitle, &newsDescription, &newsUrl, &newsPublished, &newUpdated, &newsImageId, &newsImageUrl, &newsImageAlt, &newsAuthorId, &newsAuthorName, &newsAuthorSurname, &newsAuthorUrl, &newsAuthorImageUrl, &newsAuthorDescription)
		if err != nil {
			return allNews, err
		}

		newsDetails := NewsNewWithRelatedFileds(
			newsId,
			newsTitle,
			newsDescription,
			newsUrl,
			newsPublished,
			newUpdated,
			newsImageId,
			newsImageUrl,
			newsImageAlt,
			newsAuthorId,
			newsAuthorName,
			newsAuthorSurname,
			newsAuthorUrl,
			newsAuthorImageUrl,
			newsAuthorDescription,
		)
		allNews = append(allNews, newsDetails)
	}
	return allNews, nil
}
