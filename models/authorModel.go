package models

import (
	"eventivicinoame/database"
	"fmt"
)

type Author struct {
	Id          int
	Email       string
	Password    string
	Name        string
	Surname     string
	Description string
	Url         string
	ImageUrl    string
	Published   string
	Updated     string
}

func AuthorNew(getAuthorId int, getAuthorEmail string, getAuthorPassword string, getAuthorName string, getAuthorSurname string, getAuthorDescription string, getAuthorUrl string, getAuthorImageUrl string, getAuthorPublished string, getAuthorUpdated string) Author {
	setNewAuthor := Author{
		Id:          getAuthorId,
		Email:       getAuthorEmail,
		Password:    getAuthorPassword,
		Name:        getAuthorName,
		Surname:     getAuthorSurname,
		Description: getAuthorDescription,
		Url:         getAuthorUrl,
		ImageUrl:    getAuthorImageUrl,
		Published:   getAuthorPublished,
		Updated:     getAuthorUpdated,
	}
	return setNewAuthor
}

func AuthorShowAuthors() ([]Author, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		fmt.Println("Error retriving authors data from the db:", err)
		return nil, err
	}
	defer rows.Close()

	var allAuthors []Author
	for rows.Next() {
		var authorId int
		var authorEmail string
		var authorPassword string
		var authorName string
		var authorSurname string
		var authorDescription string
		var authorUrl string
		var authorImageUrl string
		var authorPublished string
		var authorUpdated string
		err = rows.Scan(&authorId, &authorEmail, &authorPassword, &authorName, &authorSurname, &authorDescription, &authorUrl, &authorImageUrl, &authorPublished, &authorUpdated)
		if err != nil {
			fmt.Println("Error to create allAuthors:", err)
			return nil, err
		}

		authorDetails := AuthorNew(authorId, authorEmail, authorPassword, authorName, authorSurname, authorDescription, authorUrl, authorImageUrl, authorPublished, authorUpdated)
		allAuthors = append(allAuthors, authorDetails)
	}
	return allAuthors, nil
}

func AuthorFindByUrlReturnItsId(getAuthorUrl string) (int, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM authors WHERE url=?", getAuthorUrl)
	if err != nil {
		fmt.Println("Error on author query:", err)
		return 0, err
	}
	defer rows.Close()

	var getAuthorData Author
	for rows.Next() {
		var authorId int
		var authorEmail string
		var authorPassword string
		var authorName string
		var authorSurname string
		var authorDescription string
		var authorUrl string
		var authorImageUrl string
		var authorPublished string
		var authorUpdated string
		err = rows.Scan(&authorId, &authorEmail, &authorPassword, &authorName, &authorSurname, &authorDescription, &authorUrl, &authorImageUrl, &authorPublished, &authorUpdated)
		if err != nil {
			fmt.Println("Error getting author data:", err)
			return 0, err
		}

		authorDetails := AuthorNew(authorId, authorEmail, authorPassword, authorName, authorSurname, authorDescription, authorUrl, authorImageUrl, authorPublished, authorUpdated)
		getAuthorData = authorDetails

	}
	return getAuthorData.Id, nil
}
