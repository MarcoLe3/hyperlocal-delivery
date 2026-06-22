package dispatch

import (
	"time"
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/scoring"
	"github.com/carsonfeng/KMatch"
)

func Assign(couriers []models.Courier, orders []models.Order, weight scoring.Weight) []models.Assignment {
	var assignments []models.Assignment

	matrix := make([][]float64, len(couriers))
	for i, courier := range couriers {
		matrix[i] = make([]float64, len(orders))
		for j, order := range orders {
			MDscore := scoring.ScoreMD(courier, order, weight)
			matrix[i][j] = -MDscore
		}
	}

	result := hungarian.SolveMax(matrix)

	for courierIdx, orderMap := range result {
		for orderIdx := range orderMap {
			assignments = append(assignments, models.Assignment{
				CourierID: couriers[courierIdx].ID,
				Order: orders[orderIdx],
				AssignedAt: time.Now(),
				MDscore: scoring.ScoreMD(couriers[courierIdx], orders[orderIdx], weight),
			})
		}
	}
	
	return assignments
}
