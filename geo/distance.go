package geo

import (
	"errors"
	"hyperlocal-delivery/models"
	"time"

	"github.com/umahmood/haversine"
)

const (
	WalkingSpeed = 3.0
	BikingSpeed = 12.0
	DrivingSpeed = 60.0
)

func GetDistance(origin, destination models.Point) float64 {
	originCoord := haversine.Coord{Lat: origin.Lat, Lon: origin.Lng}
	destinationCoord := haversine.Coord{Lat: destination.Lat, Lon: destination.Lng}
	_, km := haversine.Distance(originCoord,destinationCoord)
	return km
}

func GetSpeedOfTravel(method_of_travel string) (float64, error) {
	var ErrNotValidMethod = errors.New("Not a valid method of travel")
	switch method_of_travel {
	case "car":
		return DrivingSpeed, nil
	case "bike":
		return BikingSpeed, nil
	case "walk":
		return WalkingSpeed, nil
	default:
		return -1.0, ErrNotValidMethod
	}
}

/*
	Probably switch to an api to get routes
 */
func GetEstimatedTimeOfArrival(km, speed float64) time.Duration {
	hours := km/speed
	return time.Duration(hours * float64(time.Hour))
}


