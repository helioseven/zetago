package gamelogic

type Player byte

// Players are represented by one of two constant values.
const (
	Black Player = iota
	White
)

// Method implements the Stringer interface for Player type.
func (p Player) String() string {
	if p == Black {
		return "Black"
	} else {
		return "White"
	}
}

// Returns the constant associated with the other player.
func (p Player) Other() Player {
	if p == Black {
		return White
	} else {
		return Black
	}
}