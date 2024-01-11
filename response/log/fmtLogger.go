package log

import (
	"fmt"
	"github.com/google/uuid"
	"ticktack/response/player"
)

type FmtLogger struct {
	GameId uuid.UUID
}

func NewFmtLogger() *FmtLogger {
	return &FmtLogger{}
}

func (l *FmtLogger) SetGameId(gameId uuid.UUID) {
	l.GameId = gameId
}

func (l *FmtLogger) GameStarted() {
	fmt.Printf("%s: Game started\n", l.GameId.String())
}

func (l *FmtLogger) NewTick(tick uint) {
	fmt.Printf("%s: New tick: %d\n", l.GameId.String(), tick)
}

func (l *FmtLogger) NewTurn(turn uint) {
	fmt.Printf("%s: New turn: %d\n", l.GameId.String(), turn)
}

func (l *FmtLogger) GameFinished() {
	fmt.Printf("%s: Game finished\n", l.GameId.String())
}

func (l *FmtLogger) Win(player player.Player) {
	fmt.Printf("%s: Winner is %s (ID: %d)\n", l.GameId.String(), player.Name, player.Id)
}

func (l *FmtLogger) Error(message string) {
	fmt.Printf("%s: Game stopped with error: %s\n", l.GameId.String(), message)
}

func (l *FmtLogger) DamageDoneToPlayer(player player.Player) {
	fmt.Printf("%s: Damage done to player %s (ID: %d)\n", l.GameId.String(), player.Name, player.Id)
}
