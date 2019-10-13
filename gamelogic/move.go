package gamelogic

import (
	"errors"
	"fmt"
)

// A move: either play (at given point), pass, or resign.
type Move struct {
	Pnt *Point
	IsPass bool
	IsResign bool
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

// Method implements Stringer interface for Move struct.
func (m Move) String() (string, error) {
	if m.Pnt != nil && !(m.IsPass || m.IsResign) {
		return fmt.Sprint("Move.Play{", m.Pnt.Row, m.Pnt.Col, "}"), nil
	} else if m.IsPass && !m.IsResign {
		return fmt.Sprint("Move.Pass"), nil
	} else if m.IsResign {
		return fmt.Sprint("Move.Resign"), nil
	} else {
		return "", errors.New("Invalid move properties.")
	}
}