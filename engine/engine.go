package engine

import (
	"sync"

	"github.com/SirusCodes/ultimate-tic-tac-toe/game"
)

const inf = 1_000_000

type MoveScore struct {
	Board    uint8
	Position uint8
	Score    int
	Count    uint
}

type MovesResult struct {
	StateChecks uint
	AllMoves    []MoveScore
	BestMove    MoveScore
}

func RunEngine(currentGame game.Game, depth uint8) MovesResult {
	wg := sync.WaitGroup{}
	moveScoreChan := make(chan MoveScore)

	for move := range currentGame.GetNextValidMovesSeq() {
		wg.Go(func() {
			game := currentGame.Clone()
			game.PlayMove(move)
			score, count := miniMax(game, move, depth, -inf, inf, false)
			moveScoreChan <- MoveScore{
				Score:    score,
				Board:    move.BoardZone,
				Position: move.Position,
				Count:    count,
			}
		})
	}

	go func() {
		wg.Wait()
		close(moveScoreChan)
	}()

	ans := MovesResult{
		StateChecks: 0,
		AllMoves:    []MoveScore{},
		BestMove:    MoveScore{Score: -inf},
	}

	for result := range moveScoreChan {
		if result.Score > ans.BestMove.Score {
			ans.BestMove = result
		}
		ans.AllMoves = append(ans.AllMoves, result)

		ans.StateChecks += result.Count
	}

	return ans
}

func miniMax(game game.Game, move game.Move, depth uint8, alpha, beta int, maximizing bool) (score int, count uint) {
	if depth == 0 {
		return game.Evaluate(move), 0
	}

	var totalCount uint = 0

	if maximizing {
		val := -inf

		for m := range game.GetNextValidMovesSeq() {
			newGame := game.Clone()
			newGame.PlayMove(m)

			score, count := miniMax(newGame, m, depth-1, alpha, beta, !maximizing)
			totalCount += count + 1

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

		return val, totalCount
	} else {
		val := inf

		for m := range game.GetNextValidMovesSeq() {
			newGame := game.Clone()
			newGame.PlayMove(m)

			score, count := miniMax(newGame, m, depth-1, alpha, beta, !maximizing)
			totalCount += count + 1

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

		return val, totalCount
	}
}
