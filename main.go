package main

import (
	"fmt"
	gl "zetago/gamelogic"
)

// simply testing modules as they are built
func main() {
	// test board.Point
	b := gl.Point{2, 2}
	fmt.Println(b)
	fmt.Println(b.Neighbors())

	// test player.Player
	p := gl.Black
	fmt.Println(p)
	fmt.Println(p.Other())

	// test move.Move
	m1 := gl.NewPlay(b)
	m2 := gl.NewPass()
	fmt.Println(m1)
	fmt.Println(m2)

	// test stones.StoneGroup
	stones := []gl.Point{b, gl.Point{2, 3}}
	libs := []gl.Point{gl.Point{1, 2}, gl.Point{1, 3}}
	sg1 := gl.StoneGroup{gl.Black, stones, libs}
	fmt.Println(sg1)
	fmt.Println(sg1.NumStones())
	sg1.AddLiberty(gl.Point{2, 4})
	fmt.Println(sg1.NumLiberties())
	sg1.RemoveLiberty(gl.Point{1, 2})
	fmt.Println(sg1)
	sg2 := gl.StoneGroup{gl.Black, []gl.Point{gl.Point{3, 2}, gl.Point{3, 3}}, libs}
	sg2.AddLiberty(gl.Point{1, 4})
	fmt.Println(sg2)
	sg1.MergeWith(sg2)
	fmt.Println(sg1)
}