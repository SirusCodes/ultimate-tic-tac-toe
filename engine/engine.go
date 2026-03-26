package engine

import (
	"fmt"
	"sync"

	"github.com/SirusCodes/9x9-analysis/game"
)

const inf = 1_000_000

type MoveScore struct {
	Board    uint8
	Position uint8
	Score    int
}

const depth = 6

func RunEngine(currentGame game.Game) {
	wg := sync.WaitGroup{}
	moveScoreChan := make(chan MoveScore)

	for move := range currentGame.GetNextValidMovesSeq() {
		wg.Go(func() {
			game := currentGame.Clone()
			game.PlayMove(move)
			score := miniMax(game, move, depth, -inf, inf, false)
			moveScoreChan <- MoveScore{
				Score:    score,
				Board:    move.BoardZone,
				Position: move.Position,
			}
		})
	}

	go func() {
		wg.Wait()
		close(moveScoreChan)
	}()

	moves := make([]MoveScore, 9)
	bestMove := MoveScore{
		Score: -inf,
	}

	for result := range moveScoreChan {
		if result.Score > bestMove.Score {
			bestMove = result
		}
		moves = append(moves, result)
	}

	fmt.Printf("best: %+v\n", bestMove)
}

func miniMax(game game.Game, move game.Move, depth uint8, alpha, beta int, maximizing bool) int {
	if depth == 0 {
		return game.Evaluate(move)
	}

	if maximizing {
		val := -inf

		for m := range game.GetNextValidMovesSeq() {
			newGame := game.Clone()
			newGame.PlayMove(m)

			score := miniMax(newGame, move, depth-1, alpha, beta, !maximizing)

			if score > val {
				val = score
			}

			if val > alpha {
				alpha = val
			}

			if alpha >= beta {
				break
			}
		}

		return val
	} else {
		val := inf

		for m := range game.GetNextValidMovesSeq() {
			newGame := game.Clone()
			newGame.PlayMove(m)

			score := miniMax(newGame, move, depth-1, alpha, beta, !maximizing)

			if score < val {
				val = score
			}

			if val < beta {
				beta = val
			}

			if beta <= alpha {
				break
			}
		}

		return val
	}
}
