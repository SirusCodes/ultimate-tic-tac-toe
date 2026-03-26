package main

import (
	"github.com/SirusCodes/9x9-analysis/engine"
	"github.com/SirusCodes/9x9-analysis/game"
	"github.com/SirusCodes/9x9-analysis/player"
)

func main() {
	g := game.NewGame(player.NewPlayer(0, 0), player.NewPlayer(0, 0), 0)
	g.PlayMove(game.Move{BoardZone: 4, Position: 4})
	g.PlayMove(game.Move{BoardZone: 4, Position: 7})
	g.PlayMove(game.Move{BoardZone: 7, Position: 4})
	engine.RunEngine(g)
}
