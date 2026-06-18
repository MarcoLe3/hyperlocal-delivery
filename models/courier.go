package models

type Courier struct {
	Capacity 		int			`json:"capacity"`
	Location 		Point		`json:"location"`
	OnHandOrders 	[]Order		`json:"on_hand_orders"`
	Status 			string		`json:"status"`
	MethodOfTravel	string		`json:"method_of_travel"`
}

func (courier Courier) CanAcceptMoreOrders() bool {
	return len(courier.OnHandOrders) < courier.Capacity
}

func (courier Courier) AddMoreOrder(order Order) {
	if !courier.CanAcceptMoreOrders() {
		return
	}

	courier.OnHandOrders = append(courier.OnHandOrders, order)
}