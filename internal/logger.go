package internal

import (
	"fmt"
	"github.com/google/uuid"
	"ticktack/player"
)

type Logger struct {
	GameId uuid.UUID
}

func (l *Logger) GameStarted() {
	fmt.Printf("%s: Game started\n", l.GameId.String())
}

func (l *Logger) NewTick(tick uint) {
	fmt.Printf("%s: New tick: %d\n", l.GameId.String(), tick)
}

func (l *Logger) GameFinished() {
	fmt.Printf("%s: Game finished\n", l.GameId.String())
}

func (l *Logger) Win(player player.Player) {
	fmt.Printf("%s: Winner is %s (ID: %d)\n", l.GameId.String(), player.Name, player.Id)
}

func (l *Logger) Error(message string) {
	fmt.Printf("%s: Game stopped with error: %s\n", l.GameId.String(), message)
}
