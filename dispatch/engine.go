package dispatch

import (
	"hyperlocal-delivery/models"
)

func run(orderCh <- chan models.Order, )

func drainOrders(orderCh <-chan models.Order) []models.Order {
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
