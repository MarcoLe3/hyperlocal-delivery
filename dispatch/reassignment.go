package dispatch

import (
	"context"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/scoring"
	"math"
	"time"
)

func RunReassignment(ctx context.Context, courierPool *models.CourierPool, orderPool *models.OrderPool, weight scoring.Weight) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			inProgressOrders := orderPool.GetInProgressOrder()

			for orderIdx := range inProgressOrders {
				order := inProgressOrders[orderIdx]
				if orderIsLate(order) {
					bestCourier := findBestCourierForOrder(order, courierPool, weight)
					if bestCourier == nil {
						continue
					}
					courierPool.RemoveOrderFromCourier(order.AssignedID, order.ID)
					courierPool.AddOrderToCourier(bestCourier.ID, order)
					orderPool.UpdateAssignedCourier(order.ID, bestCourier.ID)
				} 
			}
		case <- ctx.Done():
			return
		}
	}
}

func orderIsLate(order models.Order) bool {
	threshold := 5 * time.Minute
	delay_time := order.ETA.Sub(order.Deadline)
	return delay_time > threshold
}

func findBestCourierForOrder(order models.Order, courierPool *models.CourierPool, weight scoring.Weight) *models.Courier {
	avaliableCouriers := courierPool.GetAvaliableCourier()

	bestMDscore := math.MaxFloat64
	var bestCourier *models.Courier

	for courier := range avaliableCouriers {
		courierScore := scoring.ScoreMD(avaliableCouriers[courier],order,weight)
		if courierScore < bestMDscore {
			bestMDscore = courierScore
			bestCourier = &avaliableCouriers[courier]
		}
	}
	return bestCourier
}
