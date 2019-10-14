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
	m := make(map[Point]*StoneGroup, w*h)
	b := Board{h, w, m}
	return &b
}

// Method implements Stringer interface for Board struct.
func (b Board) String() string {
	s := fmt.Sprint("Board (", b.Width, "x", b.Height, ") {\n")
	visited := []*StoneGroup{}
	for _, e := range b.StoneMap {
		v := false
		for _, vsg := range visited {
			if vsg == e {
				v = true
			}
		}
		if !v {
			s += fmt.Sprintln(*e)
			visited = append(visited, e)
		}
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
	// error checking
	if !isOnBoard(*b, p) {
		return errors.New("Given point is not within the board.")
	}
	if b.GetStoneGroup(p) != nil {
		return errors.New("Given point on the board is already occupied.")
	}

	// initialize utilities
	addGroup := func(list []*StoneGroup, item *StoneGroup) []*StoneGroup {
		for _, e := range list {
			if e == item {
				return list
			}
		}
		return append(list, item)
	}
	adj_same := []*StoneGroup{}
	adj_opp := []*StoneGroup{}
	adj_lib := []Point{}

	// collect information about neighbors
	for _, e := range p.Neighbors() {
		if !isOnBoard(*b, e) {
			continue
		}
		nsg := b.GetStoneGroup(e)
		if nsg == nil {
			adj_lib = append(adj_lib, e)
			continue
		}
		if nsg.Color == c {
			adj_same = addGroup(adj_same, nsg)
		} else {
			adj_opp = addGroup(adj_opp, nsg)
		}
	}

	// apply game logic
	newsg := &StoneGroup{c, []Point{p}, adj_lib}
	for _, e := range adj_same {
		newsg.MergeWith(*e)
	}
	for _, e := range newsg.Stones {
		b.StoneMap[e] = newsg
	}
	for _, e := range adj_opp {
		e.RemoveLiberty(p)
	}
	for _, e := range adj_opp {
		if e.NumLiberties() == 0 {
			b.removeStones(e)
		}
	}
	return nil
}

// Method removes a StoneGroup from the board.
func (b *Board) removeStones(sg *StoneGroup) error {
	for _, e := range sg.Stones {
		if b.StoneMap[e] != sg {
			return errors.New("StoneGroup to be removed does not match board state.")
		}
	}
	for _, e := range sg.Stones {
		for _, p := range e.Neighbors() {
			if !isOnBoard(*b, p) {
				continue
			}
			nsg := b.GetStoneGroup(p)
			if nsg != nil && nsg != sg {
				nsg.AddLiberty(e)
			}
		}
		b.StoneMap[e] = nil
	}
	return nil
}

// Utility function checks whether a point is within the board.
func isOnBoard(b Board, p Point) bool {
	b1 := (1 <= p.Row) && (p.Row <= b.Height)
	b2 := (1 <= p.Col) && (p.Col <= b.Width)
	return b1 && b2
}