package models

import (
	"sync"
	"errors"
)

var ErrNotFoundOrder = errors.New("Order Not Found")

type OrderPool struct {
	Pool	map[string]*Order
	mu 		sync.Mutex
}

func CreateNewOrderPool() *OrderPool {
	return &OrderPool{
		Pool: make(map[string]*Order),
	}
}

func (o *OrderPool) AddOrderToPool(id string, order Order) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.Pool[id] = &order
}

func (o *OrderPool) GetDeliveredOrders() []Order {
	o.mu.Lock()
	defer o.mu.Unlock()

	var delievered_order []Order

	for _, value := range o.Pool {
		if value.Status == "Delivered" {
			delievered_order = append(delievered_order, *value)
		}
	}

	return delievered_order
}

func (o *OrderPool) GetPendingOrders() []Order {
	o.mu.Lock()
	defer o.mu.Unlock()

	var pending_order []Order

	for _, value := range o.Pool {
		if value.Status == "Pending" {
			pending_order = append(pending_order, *value)
		}
	}

	return pending_order
}

func (o *OrderPool) GetInProgressOrder() []Order {
	o.mu.Lock()
	defer o.mu.Unlock()

	var in_progress_order []Order

	for _, value := range o.Pool {
		if value.Status == "In Progress" {
			in_progress_order = append(in_progress_order, *value)
		}
	}

	return in_progress_order
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

func (o *OrderPool) UpdateAssignedCourier(order_id, courier_id string) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	order_to_update := o.Pool[order_id]

	if order_to_update == nil {
		return ErrNotFoundOrder
	}

	order_to_update.AssignedID = courier_id
	return nil
}