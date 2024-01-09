package main

import "ticktack/engine"

func main() {
	g := engine.NewGame()
	g.Start()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.NextTick()
	g.Finish()
}
