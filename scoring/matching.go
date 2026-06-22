package scoring

import (
	"hyperlocal-delivery/geo"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/routing"
)

func ScoreMD(courier models.Courier, order models.Order, weight Weight) float64 {
	timeDelta, distDelta := ScorePath(courier, order)
	risk := CVarScore(courier, order)
	default_score := weight.TimeWeight*timeDelta + weight.DistanceWeight*distDelta
	return (1-risk) * default_score + risk
}

func ScorePath(courier models.Courier, order models.Order) (float64, float64)  {
	routeBefore := routing.SortListByETA(courier)
	timeBefore := calculateRouteCost(courier.Location, routeBefore)

	copyCourier := courier
	copyCourier.AddMoreOrder(order)
	routeAfter := routing.SortListByETA(copyCourier)
	timeAfter := calculateRouteCost(copyCourier.Location, routeAfter)

	timeDelta := timeAfter - timeBefore
	distDelta := geo.GetDistance(courier.Location, order.Origin)

	return timeDelta, distDelta

}

func calculateRouteCost(start models.Point, orders []models.Order) float64 {
	totalTime := 0.0
	currentPoint := start
	for _, order := range orders {
		totalTime += geo.GetDistance(currentPoint, order.Origin)
		totalTime += geo.GetDistance(order.Origin, order.Destination)
		currentPoint = order.Destination
	}

	return totalTime
}