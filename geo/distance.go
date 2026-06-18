package geo

import (
	"time"
	"hyperlocal-delivery/models"
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

/*
	Probably switch to an api to get routes
 */
func GetEstimatedTimeOfArrival(speed, miles float64) time.Duration {
	hours := miles/speed
	return time.Duration(hours * float64(time.Hour))
}


