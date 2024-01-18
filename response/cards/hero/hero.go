package hero

type Hero struct {
	BaseInfluence    uint
	CurrentInfluence uint
}

func NewHero(baseInfluence uint) Hero {
	return Hero{
		BaseInfluence:    baseInfluence,
		CurrentInfluence: baseInfluence,
	}
}
