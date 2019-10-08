package gamelogic

// A point on the board, defined by row and column.
type Point struct {
	Row, Col uint16
}

// Returns a list of the four adjacent neighbors of the point.
func (p Point) Neighbors() [4]Point {
	r := [4]Point{}
	r[0] = Point{p.Row, p.Col - 1}
	r[1] = Point{p.Row + 1, p.Col}
	r[2] = Point{p.Row, p.Col + 1}
	r[3] = Point{p.Row - 1, p.Col}
	return r
}