package main

import (
	"fmt"
	gl "zetago/gamelogic"
)

// simply testing modules as they are built
func main() {
	var b = gl.Point{2, 2}
	fmt.Println(b)
	fmt.Println(b.Neighbors())

	var p = gl.Black
	fmt.Println(p)
	fmt.Println(p.Other())
}