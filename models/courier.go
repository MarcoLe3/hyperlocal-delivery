package models

type Courier struct {
	Capacity 		int			`json:"capacity"`
	Location 		string		`json:"location"`
	OnHandOrders 	[]Order		`json:"on_hand_orders"`
	Status 			string		`json:"status"`
	MethodOfTravel	string		`json:"method_of_travel"`
}

func (courier Courier) CanAcceptMoreOrders() bool {
	return len(courier.OnHandOrders) < courier.Capacity
}