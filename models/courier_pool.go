package models

import (
	"sync"
	"errors"
)

var ErrCourierNotFound = errors.New("courier not found in the pool")

type CourierPool struct {
	Pool map[string]*Courier
	mu   sync.Mutex
}

func CreateNewPool() *CourierPool {
	return &CourierPool{
		Pool: make(map[string]*Courier),
	}
}

func (c *CourierPool) addCourier(id string, courier Courier) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Pool[id] = &courier
}

func (c *CourierPool) GetAvaliableCourier() []Courier {
	c.mu.Lock()
	defer c.mu.Unlock()
	var avaliable_couriers []Courier
	for _, value := range c.Pool {
		if value.Status == "Avaliable" && value.CanAcceptMoreOrders() {
			avaliable_couriers = append(avaliable_couriers, *value)
		}
	}

	return avaliable_couriers
}

func (c *CourierPool) AddOrderToCourier(id string, order Order) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	courier := c.Pool[id]

	if courier == nil {
		return ErrCourierNotFound
	}

	courier.AddMoreOrder(order)
	return nil
}

func (c *CourierPool) UpdateCourierStatus(id, status string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	courier := c.Pool[id]

	if courier == nil {
		return ErrCourierNotFound
	}

	courier.Status = status
	return nil
}