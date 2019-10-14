package main

import (
	"fmt"
	gl "zetago/gamelogic"
)

// simply testing modules as they are built
func main() {
	board := gl.NewBoard(19, 19)
	board.PlaceStone(gl.Black, gl.Point{2, 2})
	board.PlaceStone(gl.White, gl.Point{2, 3})
	board.PlaceStone(gl.Black, gl.Point{3, 3})
	board.PlaceStone(gl.White, gl.Point{2, 4})
	fmt.Println(*board)
}