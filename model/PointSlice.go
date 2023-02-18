package model

type PointSlice []Point

func (points *PointSlice) Contains(point *Point) bool {
	for _, element := range *points {
		if element.Equals(point) {
			return true
		}
	}

	return false
}