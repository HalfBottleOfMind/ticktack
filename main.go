package main

import (
	"fmt"
	"ticktack/response/engine"
	"ticktack/response/engine/commands"
	"ticktack/response/log"
	"ticktack/response/player"
)

func main() {
	p1 := player.NewPlayer(1, "John")
	p2 := player.NewPlayer(2, "Jane")
	g := engine.NewGame(p1, p2, log.NewFmtLogger())

	g.AddCommandToQueue(&commands.StartGame{g})
	g.ExecuteNextCommand()

	fmt.Println(g.GetStatus())
}
