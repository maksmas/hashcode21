package main

import "testing"

func TestCalculateTrafficRatio(t *testing.T) {
	tests := []struct {
		name     string
		street   Street
		carCount int
		want     float64
	}{
		{"1", Street{title: "1", length: 2}, 2, 1.0},
		{"2", Street{title: "2", length: 10}, 5, .5},
		{"3", Street{title: "3", length: 3}, 9, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTrafficRatio(tt.street, tt.carCount); got != tt.want {
				t.Errorf("CalculateTrafficRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
