package services

import (
	. "find-rectangles-test/model"
	"math"
)

func FindRects(lines LineSlice) RectSlice {
	rects := RectSlice{}

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if (isRect(&lines[i], &lines[j])) {
				rect := NewRect(lines[i], lines[j])

				if (!rects.Contains(&rect)) {
					rects = append(rects, rect)
				}
			}
		}
	}

	return rects
}

func isRect(line1, line2 *Line) bool {
	if line1.GetLength() != line2.GetLength() {
		return false
	}

	return isParallel(line1, line2) && hasSameDiagonals(line1, line2)
}

func isParallel(line1, line2 *Line) bool {
	y1 := line1.A.Y - line1.B.Y
	x1 := line1.A.X - line1.B.X

	y2 := line2.A.Y - line2.B.Y
	x2 := line2.A.X - line2.B.X

	if x1 == 0 {
		return false;
	} 

	if x2 == 0 {
		return false;
	} 

	return y1 / x1 == y2 / x2
}

func hasSameDiagonals(line1, line2 *Line) bool {
	AA := NewLine(line1.A, line2.A) 
	AB := NewLine(line1.A, line2.B)

	BA := NewLine(line1.B, line2.A)
	BB := NewLine(line1.B, line2.B)

	hasSameDiagonals := math.Max(AA.GetLength(), AB.GetLength()) == math.Max(BA.GetLength(), BB.GetLength())
	hasSameSide := math.Min(AA.GetLength(), AB.GetLength()) == math.Min(BA.GetLength(), BB.GetLength())

	return hasSameDiagonals && hasSameSide
}

