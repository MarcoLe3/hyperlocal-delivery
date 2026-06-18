package models

import (
	"time"
)

type FulfillmentType string

const (
	TypePickup   FulfillmentType = "pickup"
	TypeDelivery FulfillmentType = "delivery"
)

type Order struct {
	ID          string          `json:"id"`
	OrderType   FulfillmentType `json:"order_type"`
	Status      string          `json:"status"`
	ETA         time.Time       `json:"eta"`
	SceduledFor time.Time       `json:"sceduled_for"`
	Deadline	time.Time		`json:"deadline"`
	Destination	string			`json:"destination"`
}
