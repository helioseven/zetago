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
	Width, Height uint16
	StoneMap map[Point]*StoneGroup
}

// Constructor function builds a new Board with passed width and height.
func NewBoard(w uint16, h uint16) *Board {
	m := make(map[Point]*StoneGroup, w*h)
	b := Board{w, h, m}
	return &b
}

// Method implements Stringer interface for Board struct.
func (b *Board) String() string {
	s := fmt.Sprint("Board (", b.Width, "x", b.Height, ") {\n")
	for _, e := range b.GetAllStoneGroups() {
		s += fmt.Sprintln(e)
	}
	return s + "}"
}

// Method returns a deep copy of a Board struct.
func (b *Board) Copy() *Board {
	nb := NewBoard(b.Width, b.Height)
	for _, e := range b.GetAllStoneGroups() {
		csg := e.Copy()
		for _, p := range csg.Stones {
			nb.StoneMap[p] = csg
		}
	}
	return nb
}

// Method returns a pointer to the StoneGroup occupying a point.
func (b *Board) GetStoneGroup(p Point) *StoneGroup {
	if sg, e := b.StoneMap[p]; e {
		return sg
	}
	return nil
}

// Method returns a slice of pointers to all unique StoneGroups.
func (b *Board) GetAllStoneGroups() []*StoneGroup {
	v := []*StoneGroup{}
	for _, e := range b.StoneMap {
		f := false
		for _, vsg := range v {
			if vsg == e {
				f = true
				break
			}
		}
		if !f {
			v = append(v, e)
		}
	}
	return v
}

// Method returns a new Board with a Player stone played at a Point.
func (b *Board) PlaceStone(turn Player, p Point) (*Board, error) {
	// error checking
	if !isOnBoard(b, p) {
		return nil, errors.New("Given point is not within the board.")
	}
	if b.GetStoneGroup(p) != nil {
		return nil, errors.New("Given point on the board is already occupied.")
	}

	// initialize utilities
	add_group := func(list []*StoneGroup, item *StoneGroup) []*StoneGroup {
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
	new_board := b.Copy()

	// collect information about neighbors
	for _, e := range p.Neighbors() {
		if !isOnBoard(b, e) {
			continue
		}
		nsg := new_board.GetStoneGroup(e)
		if nsg == nil {
			adj_lib = append(adj_lib, e)
			continue
		}
		if nsg.Color == turn {
			adj_same = add_group(adj_same, nsg)
		} else {
			adj_opp = add_group(adj_opp, nsg)
		}
	}

	// apply game logic
	newsg := StoneGroup{turn, []Point{p}, adj_lib}
	for _, e := range adj_same {
		err := newsg.MergeIn(e)
		if err != nil {
			return nil, err
		}
	}
	for _, e := range newsg.Stones {
		new_board.StoneMap[e] = &newsg
	}
	for _, e := range adj_opp {
		err := e.RemoveLiberty(p)
		if err != nil {
			return new_board, err
		}
	}
	for _, e := range adj_opp {
		if e.NumLiberties() == 0 {
			err := new_board.removeStones(e)
			if err != nil {
				return new_board, err
			}
		}
	}
	return new_board, nil
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
			if !isOnBoard(b, p) {
				continue
			}
			nsg := b.GetStoneGroup(p)
			if nsg != nil && nsg != sg {
				err := nsg.AddLiberty(e)
				if err != nil {
					return err
				}
			}
		}
		b.StoneMap[e] = nil
	}
	return nil
}

// Utility function checks whether a point is within the board.
func isOnBoard(b *Board, p Point) bool {
	b1 := (1 <= p.Row) && (p.Row <= b.Height)
	b2 := (1 <= p.Col) && (p.Col <= b.Width)
	return b1 && b2
}