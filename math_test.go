package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			if got := CalculateStreetRatio(tt.street, tt.carCount); got != tt.want {
				t.Errorf("CalculateStreetRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeRatio(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		street1 := Street{title: "1", length: 2}
		street2 := Street{title: "2", length: 3}
		street3 := Street{title: "3", length: 6}
		ratios := map[Street]float64{street1: 2, street2: 6, street3: 3}

		got := NormalizeRatio(ratios)
		assert.Equal(t, 1, got[street1])
		assert.Equal(t, 3, got[street2])
		assert.Equal(t, 1, got[street3])
	})
}

func TestCalculateTrafficRate(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		street1 := Street{title: "1", length: 2}
		street2 := Street{title: "2", length: 3}
		street3 := Street{title: "3", length: 6}
		ratios := map[Street]int{street1: 2, street2: 6, street3: 3}

		got := CalculateTrafficRate(ratios)
		assert.Equal(t, 1.0, got[street1])
		assert.Equal(t, 2.0, got[street2])
		assert.Equal(t, .5, got[street3])
	})
}
