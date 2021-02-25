package main

func CalculateTrafficRatio(street Street, carCount int) float64 {
	return float64(carCount) / float64(street.length)
}
