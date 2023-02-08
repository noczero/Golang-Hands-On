package LeetCode

import (
	"fmt"
	"testing"
)

// [1,2], [3,-1], [2,1], [2.3]
func TestDistanceCompute(t *testing.T) {
	points := []point{
		{
			x: 1,
			y: 2,
		},
		{
			x: 3,
			y: -1,
		},
		{
			x: 2,
			y: 1,
		},
		{
			x: 2,
			y: 3,
		},
	}
	vertex := point{
		x: 2,
		y: 2,
	}

	result := getDistanceToVertex(points, vertex)
	fmt.Println(result)
}
