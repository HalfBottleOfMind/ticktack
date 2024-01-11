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

	fmt.Print(g.GetStatus())
	c := commands.StartGame{}
	c.SetTargets(g)
	g.AddCommandToQueue(&c)
	g.ExecuteNextCommand()
	fmt.Print(g.GetStatus())

	//g.Start()
	//for g.Status == status.InProgress {
	//	g.NextTurn()
	//	g.DealDamageToPlayer(g.PlayerOne)
	//}
}
