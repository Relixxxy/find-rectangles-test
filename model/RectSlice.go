package model

type RectSlice []Rect

func (rects *RectSlice) Contains(rect *Rect) bool {
	for _, element := range *rects {
		if element.Equals(rect) {
			return true
		}
	}

	return false
}