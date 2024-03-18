package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
}

var events = []Event{}

func (e Event) Save() {
	// save to database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
