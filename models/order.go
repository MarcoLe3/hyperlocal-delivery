package models

import (
	"time"
)

type Point struct {
	Lat float64
	Lng float64
}

type Order struct {
	ID          string          `json:"id"`
	Status      string          `json:"status"`
	ETA         time.Time       `json:"eta"`
	SceduledFor time.Time       `json:"sceduled_for"`
	Deadline	time.Time		`json:"deadline"`
	Origin		Point			`json:"origin"`
	Destination	Point			`json:"destination"`
}
