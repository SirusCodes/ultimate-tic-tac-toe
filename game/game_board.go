package game

import (
	"iter"

	"github.com/SirusCodes/9x9-analysis/player"
)

const (
	NextPlayerMetaPos uint32 = 9
	FilledBoard       uint64 = 0b111111111
	EmptyLowerBits    uint16 = 0b1111111000000000
)

const (
	smallEdge               = .5
	smallCorner             = 1
	smallCenter             = 2
	bigEdgeWin              = 3
	smallPartialWin         = 4
	sendOpponentToFreeBoard = -4
	bigCornerWin            = 5
	smallOpponentDefend     = 5
	smallWin                = 8
	bigCenter               = 10
	bigPartialWin           = 50
	bigOpponentDefend       = 60
	bigWin                  = 1000
)

type Game struct {
	X player.Player
	O player.Player

	// Stores
	// 9 bits - next small game zone
	// 1 bit - next player
	Metadata uint16
}

type NextMove struct {
	BoardZone, Position uint8
}

func NewGame(X, O player.Player, Metadata uint16) Game {
	return Game{
		X:        player.NewPlayer(X.Lo, X.Hi),
		O:        player.NewPlayer(O.Lo, O.Hi),
		Metadata: Metadata,
	}
}

func (g *Game) GetNextValidMovesSeq() iter.Seq[NextMove] {
	currSmallGameZone := g.GetNextSmallGame()
	canPlayAnywhere := currSmallGameZone != 9 && !g.IsSmallGameWin(currSmallGameZone)

	if canPlayAnywhere {
		return g.getValidMovesInBoardZoneSeq(currSmallGameZone)
	}

	return func(yield func(NextMove) bool) {
		for i := range uint8(9) {
			if g.IsSmallGameWin(currSmallGameZone) {
				continue
			}

			for move := range g.getValidMovesInBoardZoneSeq(i) {
				if !yield(move) {
					return
				}
			}
		}
	}
}

func (g *Game) getValidMovesInBoardZoneSeq(boardZone uint8) iter.Seq[NextMove] {
	smallBoard := g.X.GetSmallBoard(boardZone) | g.O.GetSmallBoard(boardZone)

	// Get blank spaces
	blanks := smallBoard ^ FilledBoard

	// Play next moves
	return func(yield func(NextMove) bool) {
		for i := range uint8(9) {
			isBlank := ((blanks >> i) & 1) == 1
			if !isBlank {
				continue
			}

			if !yield(NextMove{BoardZone: boardZone, Position: i}) {
				return
			}
		}
	}
}

func (g *Game) PlayMove(plyr *player.Player, boardZone, position uint8) {
	plyr.Play(boardZone, position)
	g.UpdateNextGameZone(position)
	g.ChangePlayer()
}

func (g *Game) Evaluation(plyr *player.Player, boardZone uint8) int {
	score := 0

	// TODO: can check Partial Wins

	// Check Wins
	if plyr.IsSmallWin(boardZone) {
		plyr.SetWinMetadata(boardZone)
		if plyr.IsWin() {
		}
	}

	return score
}

func (g *Game) GetNextSmallGame() uint8 {
	for i := range 9 {
		if (g.Metadata>>i)&1 == 1 {
			return uint8(i)
		}
	}

	return 9
}

func (g *Game) UpdateNextGameZone(smallGameZone uint8) {
	// Reset all next zone
	g.Metadata &= EmptyLowerBits
	// Set next zone
	g.Metadata |= 0b1 << smallGameZone
}

func (g *Game) IsSmallGameWin(boardZone uint8) bool {
	allWins := (g.O.Hi | g.X.Hi) >> 18
	return (allWins>>boardZone)&1 == 1
}

func (g *Game) GetPlayers() (plyr *player.Player, oppo *player.Player) {
	isX := (g.Metadata>>NextPlayerMetaPos)&1 == 1

	if isX {
		return &g.X, &g.O
	}

	return &g.O, &g.X
}

func (g *Game) ChangePlayer() {
	const flipNextPlayer uint16 = 1 << NextPlayerMetaPos
	g.Metadata ^= flipNextPlayer
}
