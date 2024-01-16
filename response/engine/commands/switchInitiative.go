package commands

import "ticktack/response/engine/interfaces"

type SwitchInitiative struct {
	Target interfaces.HasPlayers
}

func (si *SwitchInitiative) Execute() {
	si.Target.GetPlayerOne().SetInitiative(!si.Target.GetPlayerOne().GetInitiative())
	si.Target.GetPlayerTwo().SetInitiative(!si.Target.GetPlayerTwo().GetInitiative())
}
