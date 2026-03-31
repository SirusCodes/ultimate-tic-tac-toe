package main

import (
	"fmt"
	"time"

	"github.com/SirusCodes/ultimate-tic-tac-toe/engine"
	"github.com/SirusCodes/ultimate-tic-tac-toe/game"
	"github.com/SirusCodes/ultimate-tic-tac-toe/player"
)

func main() {

	start := time.Now()
	g := game.NewGame(player.NewPlayer(0, 0), player.NewPlayer(0, 0), 0)
	result := engine.RunEngine(g, 10)
	fmt.Printf("Ran for %fs\n", time.Since(start).Seconds())

	fmt.Printf("moves: %+v\n", len(result.AllMoves))
	fmt.Printf("best: %+v\n", result.BestMove)
	fmt.Printf("generated states: %d\n", result.StateChecks)
}
