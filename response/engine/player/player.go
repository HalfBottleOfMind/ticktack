package player

type Player struct {
	Id         uint
	Name       string
	Health     uint
	Initiative bool
}

func (p *Player) GetInitiative() bool {
	return p.Initiative
}

func (p *Player) SetInitiative(i bool) {
	p.Initiative = i
}

func NewPlayer(id uint, name string) Player {
	return Player{
		Id:     id,
		Name:   name,
		Health: 5,
	}
}
