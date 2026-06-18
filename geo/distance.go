package geo

import (
	"fmt"
	"time"
    "github.com/umahmood/haversine"
)

const (
	WalkingSpeed = 3.0
	BikingSpeed = 12.0
	DrivingSpeed = 60.0
)

func GetDistance(origin, destination string) float64 {
	mi, km := haversine.Distance(origin,destination)
	return mi
}

/*
	Probably switch to an api to get routes
 */
func GetEstimatedTimeOfArrival(speed, miles float64) time.Duration {
	hours := miles/speed
	return time.Duration(hours * float64(time.Hour))
}


