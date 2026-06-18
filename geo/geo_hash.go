package geo

import (
	"fmt"
	"github.com/mmcloughlin/geohash"
)

func geoEncodeLatAndLog(latitude, longitude float64) string {
	percision := uint(9)
	return geohash.EncodeWithPrecision(latitude,longitude,percision)
}

func geoDecodeLatAndLog(geoHash string) {
	geoHash.DecodeIntWithPrecision()
}