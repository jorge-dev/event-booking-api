package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jorge-dev/ev-book/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int       `json:"userId"`
}

var events []Event = []Event{}

func (e *Event) Save() error {
	// save event to database
	query := `INSERT INTO events (name, description, location, dateTime, userId) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAll() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		errorMessage := "Error getting events from db: " + err.Error()
		return nil, errors.New(errorMessage)
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		event := Event{}
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			errorMessage := "Error scanning events from db: " + err.Error()
			return nil, errors.New(errorMessage)
		}
		events = append(events, event)
	}
	return events, nil
}

func GetByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		errorMessage := fmt.Sprintf("Error preparing query to get event by id: %d : error %s", id, err.Error())
		return nil, errors.New(errorMessage)
	}
	defer stmt.Close()

	event := Event{}
	err = stmt.QueryRow(id).Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		errorMessage := fmt.Sprintf("Error getting event by id: %d : error %s", id, err.Error())
		return nil, errors.New(errorMessage)
	}
	return &event, nil
}
