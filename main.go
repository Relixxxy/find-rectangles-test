package main

import (
	. "find-rectangles-test/services"
	"fmt"
)

func main() {
	maxValue := 10
	initialPoints := GeneratePoints(20, maxValue)

	PrintCoordinateSys(initialPoints, maxValue)

	lines := CreateLines(initialPoints)

	rects := FindRects(lines)

	for _, value := range rects {
		fmt.Println(value.ToString())
	}
}