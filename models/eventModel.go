package models

import (
	"eventivicinoame/database"
	"fmt"
)

type Event struct {
	Id             int
	Title          string
	Description    string
	Url            string
	Published      string
	Updated        string
	ImageId        int
	AuthorId       int
	EventType      string
	Content        string
	Country        string
	Region         string
	City           string
	Town           string
	Fraction       string
	EventStartDate string
	EventEndDate   string
}

type EventWithRelatedFields struct {
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
	EventType         string
	Content           string
	Country           string
	Region            string
	City              string
	Town              string
	Fraction          string
	EventStartDate    string
	EventEndDate      string
}

func EventNew(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getAuthorId int, getEventType string, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getEventStartDate string, getEventEndDate string) Event {
	newEvent := Event{
		Id:             getId,
		Title:          getTitle,
		Description:    getDescription,
		Url:            getUrl,
		Published:      getPublished,
		Updated:        getUpdated,
		ImageId:        getImageId,
		AuthorId:       getAuthorId,
		EventType:      getEventType,
		Content:        getContent,
		Country:        getCountry,
		Region:         getRegion,
		City:           getCity,
		Town:           getTown,
		Fraction:       getFraction,
		EventStartDate: getEventStartDate,
		EventEndDate:   getEventEndDate,
	}
	return newEvent
}

func EventAddNewToDB(getEvent Event) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"INSERT INTO events (title, description, url, published, updated, image_id, author_id, event_type, content, country, region, city, town, fraction, event_start_date, event_end_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		getEvent.Title, getEvent.Description, getEvent.Url, getEvent.Published, getEvent.Updated, getEvent.ImageId, getEvent.AuthorId, getEvent.EventType, getEvent.Content, getEvent.Country, getEvent.Region, getEvent.City, getEvent.Town, getEvent.Fraction, getEvent.EventStartDate, getEvent.EventEndDate,
	)

	if err != nil {
		fmt.Println("Error adding a new event:", err)
		return err
	}
	defer query.Close()

	return nil
}

func EventEdit(getEvent Event) error {
	db := database.DatabaseConnection()
	defer db.Close()

	query, err := db.Query(
		"UPDATE events SET title = ?, description = ?, url = ?, published = ?, updated = ?, image_id = ?, author_id = ?, event_type = ?, content = ?, country = ?, region = ?, city = ?, town = ?, fraction = ?, event_start_date = ?, event_end_date = ? WHERE id = ?",
		getEvent.Title, getEvent.Description, getEvent.Url, getEvent.Published, getEvent.Updated, getEvent.ImageId, getEvent.AuthorId, getEvent.EventType, getEvent.Content, getEvent.Country, getEvent.Region, getEvent.City, getEvent.Town, getEvent.Fraction, getEvent.EventStartDate, getEvent.EventEndDate, getEvent.Id,
	)
	if err != nil {
		fmt.Println("Error on editing event:", err)
		return err
	}
	defer query.Close()

	return nil
}

func EventNewWithRelatedFields(getId int, getTitle string, getDescription string, getUrl string, getPublished string, getUpdated string, getImageId int, getImageUrl string, getImageAlt string, getAuthorId int, getAuthorName string, getAuthorSurname string, getAuthorUrl string, getAuthorImageUrl string, getAuthorDescription string, getEventType string, getContent string, getCountry string, getRegion string, getCity string, getTown string, getFraction string, getEventStartDate string, getEventEndDate string) EventWithRelatedFields {
	newEventWithRelatedFields := EventWithRelatedFields{
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
		EventType:         getEventType,
		Content:           getContent,
		Country:           getCountry,
		Region:            getRegion,
		City:              getCity,
		Town:              getTown,
		Fraction:          getFraction,
		EventStartDate:    getEventStartDate,
		EventEndDate:      getEventEndDate,
	}
	return newEventWithRelatedFields
}

func EventGetLimitAndPagination(getLimit, getPageNumber int) ([]EventWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT events.id, events.title, events.description, events.url, events.published, events.updated, events.image_id, images.url, images.description, events.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, events.event_type, events.content, events.country, events.region, events.city, events.town, events.fraction, events.event_start_date, events.event_end_date FROM events JOIN images ON events.image_id = images.id JOIN authors ON events.author_id = authors.id ORDER BY events.published DESC LIMIT ? OFFSET ?", getLimit, getPageNumber)
	if err != nil {
		fmt.Println("Error getting events by limit and pagination:", err)
		return nil, err
	}
	defer rows.Close()

	var allEvents []EventWithRelatedFields
	for rows.Next() {
		var eventId int
		var eventTitle string
		var eventDescription string
		var eventUrl string
		var eventPublished string
		var eventUpdated string
		var eventImageId int
		var eventImageUrl string
		var eventImageAlt string
		var eventAuthorId int
		var eventAuthorName string
		var eventAuthorSurname string
		var eventAuthorUrl string
		var eventAuthorImageUrl string
		var eventAuthorDescription string
		var eventType string
		var eventContent string
		var eventCountry string
		var eventRegion string
		var eventCity string
		var eventTown string
		var eventFraction string
		var eventStartDate string
		var eventEndDate string
		err = rows.Scan(&eventId, &eventTitle, &eventDescription, &eventUrl, &eventPublished, &eventUpdated, &eventImageId, &eventImageUrl, &eventImageAlt, &eventAuthorId, &eventAuthorName, &eventAuthorSurname, &eventAuthorUrl, &eventAuthorImageUrl, &eventAuthorDescription, &eventType, &eventContent, &eventCountry, &eventRegion, &eventCity, &eventTown, &eventFraction, &eventStartDate, &eventEndDate)
		if err != nil {
			return allEvents, err
		}

		eventDetails := EventNewWithRelatedFields(
			eventId,
			eventTitle,
			eventDescription,
			eventUrl,
			eventPublished,
			eventUpdated,
			eventImageId,
			eventImageUrl,
			eventImageAlt,
			eventAuthorId,
			eventAuthorName,
			eventAuthorSurname,
			eventAuthorUrl,
			eventAuthorImageUrl,
			eventAuthorDescription,
			eventType,
			eventContent,
			eventCountry,
			eventRegion,
			eventCity,
			eventTown,
			eventFraction,
			eventStartDate,
			eventEndDate,
		)
		allEvents = append(allEvents, eventDetails)
	}

	return allEvents, nil
}

func EventWithRelatedFieldsFindById(getEventId int) (EventWithRelatedFields, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	var getEventDate EventWithRelatedFields

	rows, err := db.Query("SELECT events.id, events.title, events.description, events.url, events.published, events.updated, events.image_id, images.url, images.description, events.author_id, authors.name, authors.surname, authors.url, authors.image_url, authors.description, events.event_type, events.content, events.country, events.region, events.city, events.town, events.fraction, events.event_start_date, events.event_end_date FROM events JOIN images ON events.image_id = images.id JOIN authors ON events.author_id = authors.id WHERE events.id = ?", getEventId)
	if err != nil {
		fmt.Println("Error on the event query EventWithRelatedFieldsFindById:", err)
		return getEventDate, err
	}
	defer rows.Close()

	for rows.Next() {
		var eventId int
		var eventTitle string
		var eventDescription string
		var eventUrl string
		var eventPublished string
		var eventUpdated string
		var eventImageId int
		var eventImageUrl string
		var eventImageAlt string
		var eventAuthorId int
		var eventAuthorName string
		var eventAuthorSurname string
		var eventAuthorUrl string
		var eventAuthorImageUrl string
		var eventAuthorDescription string
		var eventType string
		var eventContent string
		var eventCountry string
		var eventRegion string
		var eventCity string
		var eventTown string
		var eventFraction string
		var eventStartDate string
		var eventEndDate string
		err = rows.Scan(&eventId, &eventTitle, &eventDescription, &eventUrl, &eventPublished, &eventUpdated, &eventImageId, &eventImageUrl, &eventImageAlt, &eventAuthorId, &eventAuthorName, &eventAuthorSurname, &eventAuthorUrl, &eventAuthorImageUrl, &eventAuthorDescription, &eventType, &eventContent, &eventCountry, &eventRegion, &eventCity, &eventTown, &eventFraction, &eventStartDate, &eventEndDate)
		if err != nil {
			return getEventDate, err
		}

		eventDetails := EventNewWithRelatedFields(
			eventId,
			eventTitle,
			eventDescription,
			eventUrl,
			eventPublished,
			eventUpdated,
			eventImageId,
			eventImageUrl,
			eventImageAlt,
			eventAuthorId,
			eventAuthorName,
			eventAuthorSurname,
			eventAuthorUrl,
			eventAuthorImageUrl,
			eventAuthorDescription,
			eventType,
			eventContent,
			eventCountry,
			eventRegion,
			eventCity,
			eventTown,
			eventFraction,
			eventStartDate,
			eventEndDate,
		)
		getEventDate = eventDetails
	}
	return getEventDate, nil
}
