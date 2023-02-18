package services

import (
	. "find-rectangles-test/model"
	"fmt"
)

func PrintCoordinateSys(points PointSlice, size int) {
	height := size
	width := size

	fmt.Println("↑")
	for i := height - 1; i > 0; i-- {
		fmt.Print(i)

		for j := 0; j < width; j++ {
			if points.Contains(&Point{X:j, Y:i}) {
				fmt.Printf("*(%d,%d)", j, i)
			}

			fmt.Printf("\t")
		}

		fmt.Println("\n|\n|")
	}

	fmt.Print(0)

	for i := 1; i <= width; i++ {
		fmt.Printf("-------%d", i)
	}

	fmt.Println(">")
}