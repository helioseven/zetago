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

	d, e, f, g := gl.Point{3, 4}, gl.Point{2, 4}, gl.Point{4, 3}, gl.Point{1, 3}
	m4, m5, m6, m7 := gl.NewPlay(d), gl.NewPlay(e), gl.NewPlay(f), gl.NewPlay(g)
	game, _ = game.ApplyMove(gl.White, m4)
	game, _ = game.ApplyMove(gl.Black, m5)
	game, _ = game.ApplyMove(gl.White, m6)
	game, _ = game.ApplyMove(gl.Black, m7)
	fmt.Println(game)
	fmt.Println("")


	h, i, j := gl.Point{3, 2}, gl.Point{3, 5}, gl.Point{2, 3}
	m8, m9, m10 := gl.NewPlay(h), gl.NewPlay(i), gl.NewPlay(j)

	game, _ = game.ApplyMove(gl.White, m8)
	game, _ = game.ApplyMove(gl.Black, m9)
	game, _ = game.ApplyMove(gl.White, m10)
	fmt.Println(game)
	fmt.Println("")

	k := gl.Point{3, 3}
	m11 := gl.NewPlay(k)
	_, err := game.ApplyMove(gl.Black, m11)
	fmt.Println(m11, err)
}