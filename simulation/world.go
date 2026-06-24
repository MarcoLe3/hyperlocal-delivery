package simulation

import (
	"context"
	"time"
	"hyperlocal-delivery/scoring"
	"hyperlocal-delivery/dispatch"
	"hyperlocal-delivery/models"
)

func RunWorld() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	courierPool := models.CreateNewCourierPool()
	orderPool := models.CreateNewOrderPool()
	defaultWeight := scoring.DefaultWeight

    courier := models.Courier{
        ID:             "fakeID",
        Capacity:       1,
        Location:       models.Point{Lat: 12.0, Lng: 12.0},
        OnHandOrders:   []models.Order{},
        Status:         "Available",
        MethodOfTravel: "car",
    }

	orderOne := models.Order{
		ID:          "order-1",
    	Status:      "pending",
		ETA:         time.Now().Add(30 * time.Minute),
		ScheduledFor: time.Now(),
		Deadline:    time.Now().Add(45 * time.Minute),
		Origin:      models.Point{Lat: 12.0, Lng: 12.0},
		Destination: models.Point{Lat: 12.5, Lng: 12.5},
		AssignedID:  "",
	}

    courierPool.AddCourierToPool(courier.ID, courier)
	orderPool.AddOrderToPool(orderOne.ID, orderOne)
	dispatch.RunEngine(ctx, courierPool, orderPool, defaultWeight)
	dispatch.RunReassignment(ctx,courierPool,orderPool,defaultWeight)
}

