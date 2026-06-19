package scoring

import (
	"hyperlocal-delivery/routing"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/geo"
)

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