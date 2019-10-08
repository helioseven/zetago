package main

import (
	"fmt"
	gl "zetago/gamelogic"
)

// simply testing modules as they are built
func main() {
	// test board.Point
	var b = gl.Point{2, 2}
	fmt.Println(b)
	fmt.Println(b.Neighbors())

	// test player
	var p = gl.Black
	fmt.Println(p)
	fmt.Println(p.Other())

	// test move.Move
	var m1 = gl.NewPlay(b)
	var m2 = gl.NewPass()
	fmt.Println(*m1)
	fmt.Println(*m2)
}