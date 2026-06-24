package models

import (
	"log"
)

type Courier struct {
	Capacity 		int			`json:"capacity"`
	Location 		Point		`json:"location"`
	OnHandOrders 	[]Order		`json:"on_hand_orders"`
	Status 			string		`json:"status"`
	MethodOfTravel	string		`json:"method_of_travel"`
	ID				string		`json:"id"`
}

func (courier Courier) CanAcceptMoreOrders() bool {
	return len(courier.OnHandOrders) < courier.Capacity
}

func (courier Courier) AddMoreOrder(order Order) {
	if !courier.CanAcceptMoreOrders() {
		log.Println("Cannot add order")
		return
	}

	courier.OnHandOrders = append(courier.OnHandOrders, order)
	log.Println("Order: " + order.ID + " was assigned to Courier: " + courier.ID)
}

func (courier Courier) RemoveOrder(order_id string) {
	for i, order := range courier.OnHandOrders {
		if order.ID == order_id {
			courier.OnHandOrders = append(courier.OnHandOrders[:i], courier.OnHandOrders[i+1:]...)
			log.Println("Removed Order: " + order.ID)
			return
		}
	}
}