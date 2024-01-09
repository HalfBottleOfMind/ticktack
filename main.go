package main

import (
	"ticktack/engine"
	"ticktack/player"
)

func main() {
	players := []player.Player{
		{
			Id:   1,
			Name: "Foo",
		},
		{
			Id:   2,
			Name: "Bar",
		},
	}
	g := engine.NewGame(players)
	g.Start()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.Win(players[0])
}
