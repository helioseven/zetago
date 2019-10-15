package gamelogic

import "fmt"

// A move: either play (at given point), pass, or resign.
type Move struct {
	Pnt Point
	IsPlay bool
	IsPass bool
	IsResign bool
}

// Constructs a play Move with passed point.
func NewPlay(p Point) Move {
	return Move{p, true, false, false}
}

// Constructs a pass Move.
func NewPass() Move {
	return Move{Point{}, false, true, false}
}

// Constructs a resign Move.
func NewResign() Move {
	return Move{Point{}, false, false, true}
}

// Method implements Stringer interface for Move struct.
func (m Move) String() string {
	if m.IsPlay && !(m.IsPass || m.IsResign) {
		return fmt.Sprint("Play{", m.Pnt.Row, m.Pnt.Col, "}")
	} else if m.IsPass && !m.IsResign {
		return "Pass"
	} else if m.IsResign {
		return "Resign"
	} else {
		return "Invalid"
	}
}