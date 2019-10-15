package main

import (
	"fmt"
	gl "zetago/gamelogic"
)

// simply testing modules as they are built
func main() {
	a, b, c := gl.Point{2, 2}, gl.Point{2, 3}, gl.Point{3, 3}
	m1, m2, m3 := gl.NewPlay(a), gl.NewPlay(b), gl.NewPlay(c)
	game := gl.NewGame()
	game, _ = game.ApplyMove(gl.Black, m1)
	game, _ = game.ApplyMove(gl.White, m2)
	game, _ = game.ApplyMove(gl.Black, m3)
	fmt.Println(game)

	fmt.Println("")

	d, e, f := gl.Point{2, 4}, gl.Point{3, 4}, gl.Point{2, 5}
	m4, m5, m6 := gl.NewPlay(d), gl.NewPlay(e), gl.NewPlay(f)
	game, _ = game.ApplyMove(gl.White, m4)
	game, _ = game.ApplyMove(gl.Black, m5)
	game, _ = game.ApplyMove(gl.White, m6)
	fmt.Println(game)
}