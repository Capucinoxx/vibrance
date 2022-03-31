package model

import "time"

type Event struct {
	At   time.Time `json:"at"`
	From string    `json:"from"`
	To   string    `json:"to"`
}

type Events struct {
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Events []Event   `json:"events"`
}

type ListEventsInput struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
