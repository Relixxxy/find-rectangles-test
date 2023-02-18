package model

import "fmt"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x,y}
}

func (p  *Point) ToString() string {
	return fmt.Sprintf("(X:%d, Y:%d)", p.X, p.Y)
}

func (p  *Point) Equals(point *Point) bool {
	return p.X == point.X && p.Y == point.Y
}
