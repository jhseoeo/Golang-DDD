package chapter3

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

const (
	directionUnkonwn = iota
	directionNorth
	directionSouth
	directionEast
	directionWest
)

func TrackPlayer() {
	currLocation := NewPoint(0, 0)
	currLocation = move(currLocation, directionNorth)
}

func move(currLocation Point, direction int) Point {
	switch direction {
	case directionNorth:
		return NewPoint(currLocation.x, currLocation.y+1)
	case directionSouth:
		return NewPoint(currLocation.x, currLocation.y-1)
	case directionEast:
		return NewPoint(currLocation.x+1, currLocation.y)
	case directionWest:
		return NewPoint(currLocation.x-1, currLocation.y)
	default:
		// :D
	}

	return currLocation
}
