package scoring

import (
	"time"
	"errors"
	"hyperlocal-delivery/geo"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/routing"
)

func ScoreMD(courier models.Courier, order models.Order, weight Weight) float64 {
	timeDelta, distDelta := ScorePath(courier, order)
	risk := CVarScore(courier, order)
	default_score := weight.TimeWeight*float64(timeDelta) + weight.DistanceWeight*float64(distDelta)
	return (1-risk) * default_score + risk
}

func ScorePath(courier models.Courier, order models.Order) (time.Duration, float64)  {
	routeBefore := routing.SortListByETA(courier)
	timeBefore, _ := calculateRouteCost(courier.Location, courier.MethodOfTravel, routeBefore)

	copyCourier := courier
	copyCourier.AddMoreOrder(order)
	routeAfter := routing.SortListByETA(copyCourier)
	timeAfter, _ := calculateRouteCost(copyCourier.Location, courier.MethodOfTravel, routeAfter)

	timeDelta := timeAfter - timeBefore
	distDelta := geo.GetDistance(courier.Location, order.Origin)

	return timeDelta, distDelta

}

func calculateRouteCost(start models.Point, method_of_travel string, orders []models.Order) (time.Duration, error) {
	var total_time time.Duration
	var ErrCalculateRouteCost = errors.New("error with caculating route cost")
	currentPoint := start
	speed, error := geo.GetSpeedOfTravel(method_of_travel)
	if error != nil {
		return -1.0, ErrCalculateRouteCost 
	}

	for _, order := range orders {
		distance_curr_origin := geo.GetDistance(currentPoint, order.Origin)
		total_time += geo.GetEstimatedTimeOfArrival(distance_curr_origin, speed)
		distance_origin_destination := geo.GetDistance(order.Origin, order.Destination)
		total_time += geo.GetEstimatedTimeOfArrival(distance_origin_destination, speed)
		currentPoint = order.Destination
	}

	return total_time, nil
}