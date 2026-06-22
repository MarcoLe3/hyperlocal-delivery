package models

import (
	"sync"
)

type OrderPool struct {
	Pool	map[string]*Order
	mu 		sync.Mutex
}

func (o *OrderPool) AddOrderToPool(id string, order Order) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.Pool[id] = &order
}

func (o *OrderPool) GetDeliveredOrders() {

}

func (o *OrderPool) GetPendingOrders() {

}

func (o *OrderPool) GetPickedUpOrders() []Order {
	o.mu.Lock()
	defer o.mu.Unlock()

	var picked_up_order []Order

	for _, value := range o.Pool {
		if value.Status == "Picked Up" {
			picked_up_order = append(picked_up_order, *value)
		}
	}

	return picked_up_order
}