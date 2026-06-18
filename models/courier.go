package models

type courier struct {
	Capacity 		int			`json:"capacity"`
	Location 		string		`json:"location"`
	OnHandOrders 	[]string	`json:"on_hand_orders"`
	Status 			string		`json:"status"`
}

func (courier courier) CanAcceptMoreOrders() bool {
	return len(courier.OnHandOrders) < courier.Capacity
}