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
	Content     string
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
	Content           string
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

func NewsNew(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getContent string, getImageId int, getAuthorId int) News {
	newNews := News{
		Id:          getId,
		Title:       getTitle,
		Description: getDescription,
		Url:         getUrl,
		Published:   getPublished,
		Updated:     getUpdated,
		Content:     getContent,
		ImageId:     getImageId,
		AuthorId:    getAuthorId,
	}
	return newNews
}

func NewsNewWithRelatedFileds(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getContent string, getImageId int, getImageUrl string, getImageAlt string, getAuthorId int, getAuthorName string, getAuthorSurname string, getAuthorUrl string, getAuthorImageUrl string, getAuthorDescription string) NewsWithRelatedFields {
	newNewsWithRelatedFields := NewsWithRelatedFields{
		Id:                getId,
		Title:             getTitle,
		Description:       getDescription,
		Url:               getUrl,
		Published:         getPublished,
		Updated:           getUpdated,
		Content:           getContent,
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

func NewsAddNewToDB(getSingleNews News) error {
	db := database.DatabaseConnection()
	defer db.Close()

	mySqlQuery := "INSERT INTO news (title, description, url, published, updated, content, image_id, author_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	query, err := db.Query(mySqlQuery, getSingleNews.Title, getSingleNews.Description, getSingleNews.Url, getSingleNews.Published, getSingleNews.Updated, getSingleNews.Content, getSingleNews.ImageId, getSingleNews.AuthorId)
	if err != nil {
		fmt.Println("Error adding a new event:", err)
		return err
	}
	defer query.Close()

	return nil
}

func NewsEdit(getNews News) error {
	db := database.DatabaseConnection()
	defer db.Close()

	mySqlQuery := "UPDATE news SET title = ?, description = ?, url = ?, published = ?, updated = ?, content = ?, image_id = ?, author_id = ? WHERE id = ?"
	query, err := db.Query(mySqlQuery, getNews.Title, getNews.Description, getNews.Url, getNews.Published, getNews.Updated, getNews.Content, getNews.ImageId, getNews.AuthorId, getNews.Id)
	if err != nil {
		fmt.Println("Error on editing event:", err)
		return err
	}
	defer query.Close()

	return nil
}

func NewsGetLimitAndPagination(getLimit int, getPageNumber int) ([]NewsWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	mySqlQuery := "SELECT news.id, news.title, news.description, news.url, news.published, news.updated, news.content, news.image_id, images.url, images.description, news.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description FROM news JOIN images ON news.image_id = images.id JOIN authors ON news.author_id = authors.id ORDER BY news.published DESC LIMIT ? OFFSET ?"
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
		var newsUpdated string
		var newsContent string
		var newsImageId int
		var newsImageUrl string
		var newsImageAlt string
		var newsAuthorId int
		var newsAuthorName string
		var newsAuthorSurname string
		var newsAuthorUrl string
		var newsAuthorImageUrl string
		var newsAuthorDescription string
		err = rows.Scan(&newsId, &newsTitle, &newsDescription, &newsUrl, &newsPublished, &newsUpdated, &newsContent, &newsImageId, &newsImageUrl, &newsImageAlt, &newsAuthorId, &newsAuthorName, &newsAuthorSurname, &newsAuthorUrl, &newsAuthorImageUrl, &newsAuthorDescription)
		if err != nil {
			return allNews, err
		}

		newsDetails := NewsNewWithRelatedFileds(
			newsId,
			newsTitle,
			newsDescription,
			newsUrl,
			newsPublished,
			newsUpdated,
			newsContent,
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

func NewsWithRelatedFieldsFindById(getNewsId int) (NewsWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	var getNewsDate NewsWithRelatedFields

	mySqlQuery := "SELECT news.id, news.title, news.description, news.url, news.published, news.updated, news.content, news.image_id, images.url, images.description, news.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description FROM news JOIN images ON news.image_id = images.id JOIN authors ON news.author_id = authors.id WHERE news.id = ?"
	rows, err := db.Query(mySqlQuery, getNewsId)
	if err != nil {
		fmt.Println("Error on the event query NewsWithRelatedFieldsFindById:", err)
		return getNewsDate, err
	}
	defer rows.Close()

	for rows.Next() {
		var newsId int
		var newsTitle string
		var newsDescription string
		var newsUrl string
		var newsPublished string
		var newsUpdated string
		var newsContent string
		var newsImageId int
		var newsImageUrl string
		var newsImageAlt string
		var newsAuthorId int
		var newsAuthorName string
		var newsAuthorSurname string
		var newsAuthorUrl string
		var newsAuthorImageUrl string
		var newsAuthorDescription string
		err = rows.Scan(&newsId, &newsTitle, &newsDescription, &newsUrl, &newsPublished, &newsUpdated, &newsContent, &newsImageId, &newsImageUrl, &newsImageAlt, &newsAuthorId, &newsAuthorName, &newsAuthorSurname, &newsAuthorUrl, &newsAuthorImageUrl, &newsAuthorDescription)
		if err != nil {
			return getNewsDate, err
		}

		getNewsDate = NewsNewWithRelatedFileds(
			newsId,
			newsTitle,
			newsDescription,
			newsUrl,
			newsPublished,
			newsUpdated,
			newsContent,
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
	}
	return getNewsDate, nil
}
