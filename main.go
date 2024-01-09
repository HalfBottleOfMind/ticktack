package main

import (
	"ticktack/response/engine"
	"ticktack/response/logger"
	"ticktack/response/player"
)

func main() {
	p1 := player.Player{
		Id:   1,
		Name: "Foo",
	}
	p2 := player.Player{
		Id:   2,
		Name: "Bar",
	}
	g := engine.NewGame(p1, p2, logger.NewFmtLogger())
	g.Start()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.Win(p1)
}
