package main

import (
	"fmt"
	"ticktack/response/engine"
	"ticktack/response/engine/player"
)

func main() {
	p1 := player.NewPlayer(1, "John")
	p2 := player.NewPlayer(2, "Jane")
	g := engine.NewGame(&p1, &p2)

	fmt.Println(g.State.IsFinished())
}
