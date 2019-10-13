package gamelogic

import (
	"errors"
	"fmt"
)

// A point on the board, defined by row and column.
type Point struct {
	Row, Col uint16
}

// Method returns a list of the four adjacent neighbors of a point.
func (p Point) Neighbors() [4]Point {
	r := [4]Point{}
	r[0] = Point{p.Row, p.Col - 1}
	r[1] = Point{p.Row + 1, p.Col}
	r[2] = Point{p.Row, p.Col + 1}
	r[3] = Point{p.Row - 1, p.Col}
	return r
}

// A Board object, defined by width, height, and a map of stone groups.
type Board struct {
	Height, Width uint16
	StoneMap map[Point]*StoneGroup
}

// Constructor function builds a new Board with passed width and height.
func NewBoard(h uint16, w uint16) *Board {
	m := make(map[Point]*StoneGroup)
	b := Board{h, w, m}
	return &b
}

// Method implements Stringer interface for Board struct.
func (b Board) String() string {
	s := fmt.Sprint("Board (", b.Width, "x", b.Height, "){\n")
	for _, v := range(b.StoneMap) {
		s += fmt.Sprintln(*v)
	}
	return s + "}"
}

// Method returns a pointer to the StoneGroup occupying a point.
func (b *Board) GetStoneGroup(p Point) *StoneGroup {
	if sg, e := b.StoneMap[p]; e {
		return sg
	}
	return nil
}

// Method adds a stone to the board at a point.
func (b *Board) PlaceStone(c Player, p Point) error {
	if !isOnBoard(*b, p) {
		return errors.New("Given point is not within the board.")
	}
	if b.GetStoneGroup(p) != nil {
		return errors.New("Given point on the board is already occupied.")
	}
	// missing implementation
	return nil
}

// Utility function checks whether a point is within the board.
func isOnBoard(b Board, p Point) bool {
	b1 := (1 <= p.Row) && (p.Row <= b.Height)
	b2 := (1 <= p.Col) && (p.Col <= b.Width)
	return b1 && b2
}