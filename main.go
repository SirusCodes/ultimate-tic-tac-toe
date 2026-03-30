package main

import (
	"fmt"

	"github.com/SirusCodes/ultimate-tic-tac-toe/engine"
	"github.com/SirusCodes/ultimate-tic-tac-toe/game"
	"github.com/SirusCodes/ultimate-tic-tac-toe/player"
)

func main() {
	g := game.NewGame(player.NewPlayer(0, 0), player.NewPlayer(0, 0), 0)
	g.PlayMove(game.Move{BoardZone: 4, Position: 4})
	g.PlayMove(game.Move{BoardZone: 4, Position: 7})
	g.PlayMove(game.Move{BoardZone: 7, Position: 4})
	result := engine.RunEngine(g, 10)

	fmt.Printf("%+v\n", len(result.AllMoves))
	fmt.Printf("best: %+v\n", result.BestMove)
	fmt.Printf("generated states: %d\n", result.StateChecks)
}
