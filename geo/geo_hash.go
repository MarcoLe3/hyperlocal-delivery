package geo

import (
	"github.com/mmcloughlin/geohash"
)

func geoEncodeLatAndLog(latitude, longitude float64) string {
	percision := uint(9)
	return geohash.EncodeWithPrecision(latitude,longitude,percision)
}

/* TODO: needs to implement */
func geoDecodeLatAndLog(geoHash string) {
	geoHash.DecodeIntWithPrecision()
}