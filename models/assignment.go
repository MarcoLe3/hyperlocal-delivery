package models

import (
	"time"
)

type Assignment struct {
	CourierID 	string
	OrderIDS 	[]string
	AssignedAt 	time.Time
	Route		[]string
	MDscore		float64		
}