package geo

import (
	"hyperlocal-delivery/models"
	"github.com/mmcloughlin/geohash"
)

func geoEncodeLatAndLog(coordinates models.Point) string {
	percision := uint(9)
	return geohash.EncodeWithPrecision(coordinates.Lat,coordinates.Lng,percision)
}

/* TODO: needs to implement */
func geoDecodeLatAndLog(geoHash string) (float64, float64) {
	return geohash.Decode(geoHash)
}