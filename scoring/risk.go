package scoring

import (
	"hyperlocal-delivery/models"
	"math/rand"
	"sort"
	"hyperlocal-delivery/geo"
)

func prepTime() float64 {
	min_prep_time := 60.0
    mean := 600.0
    stdDev := 240.0
    prep_time := mean + (stdDev * rand.NormFloat64())
    if prep_time < min_prep_time {
        return min_prep_time
    }
    return prep_time
}

func routeCostWithPrep(courier models.Courier, order models.Order, prep_time float64) float64 {
	total_time := 0.0
	current_position := courier.Location
	all_orders := append(courier.OnHandOrders, order)

	for _, o := range all_orders {
		total_time += geo.GetDistance(current_position, o.Origin)
		total_time += prep_time
		total_time += geo.GetDistance(o.Origin, o.Destination)
		current_position = o.Destination
	}

	return total_time
}

func CVarScore(courier models.Courier, order models.Order) float64 {
	iterations := 20
	cost := make([]float64, iterations)

	for i := range iterations {
		prep_time := prepTime()
		cost[i] = routeCostWithPrep(courier, order, prep_time)
	}

	sort.Float64s(cost)
	iterations_twenty_precent := int(float64(iterations) * 0.8)
	worst_twenty_percent := cost[iterations_twenty_precent:]

	total_time := 0.0
	for _, c := range worst_twenty_percent {
		total_time += c
	}

	return total_time / float64(len(worst_twenty_percent))
}

