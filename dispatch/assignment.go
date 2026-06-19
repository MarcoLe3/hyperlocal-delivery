package dispatch

import (
	"hyperlocal-delivery/models"
	"hyperlocal-delivery/scoring"
	"github.com/carsonfeng/KMatch/hungarian"
)

func Assign(couriers []models.Courier, orders []models.Order) []models.Assignment {
	var assignments []models.Assignment

	matrix := make([][]float64, len(couriers))
	for i, courier := range couriers {
		matrix[i] = make([]float64, len(orders))
		for j, order := range orders {
			timeDelta, distDelta := scoring.ScorePath(courier, order)
			matrix[i][j] = (timeDelta + distDelta)
		}
	}

	result := hungarian.SolveMax(matrix)
	
	return assignments
}
