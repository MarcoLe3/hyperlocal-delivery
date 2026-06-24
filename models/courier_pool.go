package models

import (
	"sync"
	"errors"
	"log"
)

var ErrCourierNotFound = errors.New("courier not found in the pool")

type CourierPool struct {
	Pool map[string]*Courier
	mu   sync.Mutex
}

func CreateNewCourierPool() *CourierPool {
	return &CourierPool{
		Pool: make(map[string]*Courier),
	}
}

func (c *CourierPool) AddCourierToPool(id string, courier Courier) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Pool[id] = &courier
	log.Println("Added Courier: " + courier.ID)
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

func (c *CourierPool) RemoveOrderFromCourier(id, order_id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	courier := c.Pool[id]

	if courier == nil {
		return ErrCourierNotFound
	}

	courier.RemoveOrder(order_id)
	return nil
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