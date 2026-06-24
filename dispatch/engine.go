package dispatch

import (
	"context"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/scoring"
	"time"
	"log"
)

func RunEngine(ctx context.Context, courierPool *models.CourierPool, orderPool *models.OrderPool, weight scoring.Weight) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	log.Println("Engine is starting")
	
	for {
		select {
		case <-ticker.C:
			orders := orderPool.GetPendingOrders()
			couriers := courierPool.GetAvaliableCourier()
			assignments := Assign(couriers, orders, weight)

			for _, assignment := range assignments {
				courierPool.AddOrderToCourier(assignment.CourierID, assignment.Order)
				courierPool.UpdateCourierStatus(assignment.CourierID, "In Use")
			}
		case <-ctx.Done():
			return
		}
	}
}