package dispatch

import (
	"context"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/scoring"
	"time"
)

func RunEngine(ctx context.Context, orderCh <- chan models.Order, courierPool *models.CourierPool, weight scoring.Weight) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			orders := drainUnassignedOrders(orderCh)
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

func drainUnassignedOrders(orderCh <-chan models.Order) []models.Order {
	var batch []models.Order
	for {
		select {
		case order := <-orderCh:
			batch = append(batch,order)
		default:
			return batch
		}
	}
}