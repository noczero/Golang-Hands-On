package LeetCode

import "math"

/*
Given point [[1,2], [3,-1], [2,1], [2.3]]
vertex = [2,2]
calculate distance over the points to vertex
*/

type point struct {
	x, y int
}

func getDistanceToVertex(points []point, vertex point) []float64 {

	var result []float64
	for _, p := range points {
		result = append(result, calculateDistance(p, vertex))
	}

	return result
}

func calculateDistance(p point, vertex point) float64 {
	return math.Sqrt(float64((p.x-vertex.x)*(p.x-vertex.x) + (p.y-vertex.y)*(p.y-vertex.y)))
}
