package engine_test

import (
	"testing"

	"github.com/SirusCodes/9x9-analysis/engine"
	"github.com/SirusCodes/9x9-analysis/game"
	"github.com/SirusCodes/9x9-analysis/player"
)

func BenchmarkRunEngine_Depth10(b *testing.B) {
	b.ReportAllocs()
	depth := uint8(10)

	for b.Loop() {
		g := game.NewGame(player.NewPlayer(0, 0), player.NewPlayer(0, 0), 0)
		engine.RunEngine(g, depth)
	}
}
