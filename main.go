package main

import (
	"ticktack/internal/engine"
	"ticktack/internal/logger"
	"ticktack/internal/player"
)

func main() {
	p := []player.Player{
		{
			Id:   1,
			Name: "Foo",
		},
		{
			Id:   2,
			Name: "Bar",
		},
	}
	l := &logger.FmtLogger{}
	g := engine.NewGame(p, l)
	g.Start()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.Win(p[0])
}
