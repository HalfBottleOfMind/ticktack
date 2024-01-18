package player

type Player struct {
	Id         uint
	Name       string
	Health     uint
	Initiative bool
}

func NewPlayer(id uint, name string) Player {
	return Player{
		Id:     id,
		Name:   name,
		Health: 5,
	}
}
