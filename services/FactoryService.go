package services

import (
	. "find-rectangles-test/model"
	"math/rand"
)

func GeneratePoints(size, max int) PointSlice {
	points := PointSlice{}

	for i := 0; i < size; i++ {
		var point Point

		for {
			point = NewPoint(rand.Intn(max - 1) + 1, rand.Intn(max - 1) + 1)
			
			if !points.Contains(&point) {
				break
			}
		}

		points = append(points, point)
	}

	return points
}

func CreateLines(points PointSlice) LineSlice {
	lines := LineSlice{}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			line := NewLine(&points[i], &points[j])
			lines = append(lines, line)
		}
	}

	return lines
}
