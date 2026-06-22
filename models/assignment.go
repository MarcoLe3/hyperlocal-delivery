package models

import (
	"time"
)

type Assignment struct {
	CourierID 	string
	Order 		Order
	AssignedAt 	time.Time
	MDscore		float64		
}