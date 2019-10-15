package gamelogic

import (
	"errors"
	"fmt"
)

// A game state, defined by board position, next turn, previous game state, and last move made.
type GameState struct {
	BoardPosition *Board
	PlayerTurn Player
	PreviousState *GameState
	LastMove *Move
}

// Constructor function builds a new GameState with Board of default width and height.
func NewGame() *GameState {
	return NewGameOfSize(19, 19)
}

// Constructor function builds a new GameState with Board of passed width and height.
func NewGameOfSize(h uint16, w uint16) *GameState {
	b := NewBoard(h, w)
	return &GameState{b, Black, nil, nil}
}

// Method implements Stringer interface for GameState struct.
func (gs *GameState) String() string {
	s := fmt.Sprintln("Next turn: ", gs.PlayerTurn, "\nLast move: ", gs.LastMove)
	return fmt.Sprint(s, gs.BoardPosition)
}

// Method returns a new GameState created from applying the given move.
func (gs *GameState) ApplyMove(p Player, m Move) (*GameState, error) {
	if (gs == nil && p != Black) || p != gs.PlayerTurn {
		return nil, errors.New("Not this player's turn to play.")
	}
	var next *Board
	if m.IsPlay {
		var err error
		next, err = gs.BoardPosition.PlaceStone(p, m.Pnt)
		if err != nil {
			return nil, err
		}
		// ** check to see if move is valid **
	} else {
		next = gs.BoardPosition.Copy()
	}
	return &GameState{next, p.Other(), gs, &m}, nil
}

func (gs *GameState) IsOver() bool {
	lm := gs.LastMove
	if lm == nil {
		return false
	}
	if lm.IsResign {
		return true
	}
	slm := gs.PreviousState.LastMove
	if slm == nil {
		return false
	}
	return lm.IsPass && slm.IsPass
}