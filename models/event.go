package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserId      int       `json:"user_id"`
}

var events []Event = []Event{}

func (e *Event) Save() {
	// save event to database
	events = append(events, *e)
}

func GetAll() []Event {
	return events
}
