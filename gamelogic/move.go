package gamelogic

import "fmt"

type Move struct {
	Pnt *Point
	IsPass bool
	IsResign bool
}

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

func NewPlay(p Point) *Move {
	return &Move{&p, false, false}
}

func NewPass() *Move {
	return &Move{nil, true, false}
}

func NewResign() *Move {
	return &Move{nil, false, true}
}