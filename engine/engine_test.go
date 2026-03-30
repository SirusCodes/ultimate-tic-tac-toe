package engine_test

import (
	"testing"

	"github.com/SirusCodes/ultimate-tic-tac-toe/engine"
	"github.com/SirusCodes/ultimate-tic-tac-toe/game"
	"github.com/SirusCodes/ultimate-tic-tac-toe/player"
)

func BenchmarkRunEngine_Depth10(b *testing.B) {
	b.ReportAllocs()
	depth := uint8(10)

	for b.Loop() {
		g := game.NewGame(player.NewPlayer(0, 0), player.NewPlayer(0, 0), 0)
		engine.RunEngine(g, depth)
	}
}
