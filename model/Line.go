package model

import (
	"fmt"
	"math"
)

type Line struct {
	A *Point
	B *Point
}

func NewLine(p1, p2 *Point) Line {
	return Line{A:p1, B:p2}
}

func (line  *Line) ToString() string {
	return fmt.Sprintf("P1:%s, P2:%s, Len:%f", line.A.ToString(), line.B.ToString(), line.GetLength())
}

func (line *Line) GetLength() float64 {
	d := math.Sqrt(math.Pow(float64(line.B.X - line.A.X),2) + math.Pow(float64(line.B.Y - line.A.Y),2))
	return d
}
