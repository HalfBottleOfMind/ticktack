package player

import (
	"ticktack/src/engine/hero"
)

type Player struct {
	Id         uint
	Name       string
	Health     uint
	Initiative bool
	Hero       *hero.Hero
}

func NewPlayer(id uint, name string) Player {
	return Player{
		Id:     id,
		Name:   name,
		Health: 5,
	}
}
