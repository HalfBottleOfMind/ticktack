package state

import (
	"ticktack/response/engine/player"
)

type State struct {
	GameStatus GameStatus
	PlayerOne  *player.Player
	PlayerTwo  *player.Player
}

func (s *State) IsFinished() bool {
	return s.GameStatus == Finished || s.GameStatus == Error
}

func NewState(playerOne, playerTwo *player.Player, gameStatus GameStatus) State {
	return State{
		GameStatus: gameStatus,
		PlayerOne:  playerOne,
		PlayerTwo:  playerTwo,
	}
}
