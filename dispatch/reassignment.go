package dispatch

import (
	"context"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/scoring"
	"math"
	"time"
)

func RunAssignment(ctx context.Context, orderCh <- chan models.Order, weight scoring.Weight) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		
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

func findBestCourierForOrder(order models.Order, courierPool models.CourierPool, weight scoring.Weight) *models.Courier {
	avaliable_couriers := courierPool.GetAvaliableCourier()

	best_MDscore := math.MaxFloat64
	var best_courier *models.Courier

	for courier := range avaliable_couriers {
		courier_score := scoring.ScoreMD(avaliable_couriers[courier],order,weight)
		if courier_score < best_MDscore {
			best_MDscore = courier_score
			best_courier = &avaliable_couriers[courier]
		}
	}
	return best_courier
}
