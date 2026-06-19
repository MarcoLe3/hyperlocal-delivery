package dispatch

import (
	"time"
	"hyperlocal-delivery/models"
)

func Run(orderCh <- chan models.Order, courierCh <- chan models.Courier ) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			orders := drainUnassignedOrders(orderCh)
			couriers := drainAvaliableCouriers(courierCh)
			dispatch.Assign()
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

func drainAvaliableCouriers(courierCh <- chan models.Courier) []models.Courier {
	var batch []models.Courier
	for {
		select {
		case courier := <- courierCh:
			batch = append(batch, courier)
		default:
			return batch
		}
	}
}