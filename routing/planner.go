package routing

import (
	"time"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/geo"
)

func SortListByETA(courier models.Courier) []models.Order {
	courierList := courier.OnHandOrders
	for i := 1; i < len(courierList); i++ {
		key := courierList[i]
		j := i - 1

		for j >= 0 && courierList[j].ETA.After(key.ETA) {
			courierList[j + 1] = courierList[j]
			j--
		}

		courierList[j + 1] = key
	}

	return courierList
}

func sortListBySlackTime(orders []models.Order) []models.Order {
	calculateSlackTime := func (order models.Order) time.Duration {
		return order.Deadline.Sub(order.ETA)
	}

	for i := 1; i < len(orders); i++ {
		key := orders[i]
		j := i -1
		slack_time_i := calculateSlackTime(orders[i])
		for j >= 0 {
			slack_time_j := calculateSlackTime(orders[j])
			if slack_time_j <= slack_time_i {
				break
			}
			orders[j+1] = orders[j]
			j--
		}

		orders[j+1] = key
	}

	return orders
}
