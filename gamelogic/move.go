package gamelogic

import "fmt"

// A move: either play (at given point), pass, or resign.
type Move struct {
	Pnt *Point
	IsPass bool
	IsResign bool
}

// Method implements Stringer interface for Move struct.
func (m Move) String() string {
	if m.Pnt != nil {
		return fmt.Sprint("Move.Play{", m.Pnt.Row, m.Pnt.Col, "}")
	} else if m.IsPass {
		return fmt.Sprint("Move.Pass")
	} else if m.IsResign {
		return fmt.Sprint("Move.Resign")
	} else {
		return fmt.Sprint("Move.Invalid")
	}
}

// Constructs a play Move with passed point.
func NewPlay(p Point) Move {
	return Move{&p, false, false}
}

// Constructs a pass Move.
func NewPass() Move {
	return Move{nil, true, false}
}

// Constructs a resign Move.
func NewResign() Move {
	return Move{nil, false, true}
}