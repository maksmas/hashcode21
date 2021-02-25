package main

import "math"

func CalculateStreetRatio(street Street, carCount int) float64 {
	return float64(carCount) / float64(street.length)
}

func CalculateTrafficRate(carCount map[Street]int) map[Street]float64 {
	result := make(map[Street]float64)
	for street, count := range carCount {
		result[street] = CalculateStreetRatio(street, count)
	}
	return result
}

func NormalizeRatio(ratios map[Street]float64) map[Street]int {
	min := math.MaxFloat64
	result := make(map[Street]int)
	for _, ratio := range ratios {
		if min > ratio {
			min = ratio
		}
	}
	for street, ratio := range ratios {
		seconds := int(ratio / min)
		if seconds < 1 {
			seconds = 1
		}
		if seconds > 10 {
			seconds = 10
		}
		result[street] = seconds
	}
	return result
}
