package gamelogic

type Player byte

// Players are represented by one of two constant values.
const (
	Black Player = iota
	White
)

// Returns the constant associated with the other player.
func (p Player) Other() Player {
	if p == Black {
		return White
	} else {
		return Black
	}
}