package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ticktack/response/engine/player"
)

func TestNewGame(t *testing.T) {
	p1 := player.NewPlayer(1, "John")
	p2 := player.NewPlayer(2, "Jane")

	g := NewGame(&p1, &p2)

	assert.Same(t, &p1, g.State.PlayerOne)
	assert.Same(t, &p2, g.State.PlayerTwo)
}
