package model

import "fmt"

type Rect struct {
	Points PointSlice
}

func NewRect(l1, l2 Line) Rect {
	rect := Rect{
		Points: PointSlice{*l1.A, *l1.B, *l2.A, *l2.B},
	}

	return rect
}

func (rect *Rect) ToString() string {
	str := ""

	for _, value := range rect.Points {
		str += fmt.Sprintf("%s ", value.ToString())
	}

	return str
}

func (rect *Rect) Equals(otherRect *Rect) bool {
	counter := 0
	for _, p := range otherRect.Points{
		if (rect.Points.Contains(&p)) {
			counter++
		}
	}
	return counter == 4
}