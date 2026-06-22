package scoring

type Weight struct {
	TimeWeight		float64
	DistanceWeight 	float64
	TargetTime		float64
}

var weight = Weight{
	TimeWeight: 	95.0,
	DistanceWeight:	5.0,
	TargetTime:		95.0,
}