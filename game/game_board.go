package game

import (
	"iter"

	"github.com/SirusCodes/9x9-analysis/player"
)

const (
	GameStateMetaPos  uint32 = 9
	NextPlayerMetaPos uint32 = 18
	FilledBoard       uint64 = 0b111111111
	EmptyLowerBits    uint32 = 0b11111111111111111111111000000000
)

type EvaluationScore int

const (
	_ EvaluationScore = iota
	PartialSmallWin
	SmallWin
	PartialBigWin
	BigWin
)

type GameBoard interface {
	GetNextValidMoves() []GameBoard
	// Value will be between 0 to 8, 9 means anywhere
	GetNextSmallGame() uint8
	UpdateNextGameZone(smallGameZone uint8)

	UpdateSmallGameWin(boardZone uint8)
	GetSmallGameWin(boardZone uint8) bool
	// true is X and false is O
	GetPlayer() bool
	ChangePlayer()
}

type Game struct {
	X player.Player
	O player.Player

	// Stores
	// 9 bits - next small game zone
	// 9 bits - small game wins
	// 1 bit - next player
	Metadata uint32
}

type NextMove struct {
	BoardZone, Position uint8
}

func NewGame(X, O player.Player, Metadata uint32) Game {
	return Game{
		X:        player.NewPlayer(X.Lo, X.Hi),
		O:        player.NewPlayer(O.Lo, O.Hi),
		Metadata: Metadata,
	}
}

func (g *Game) GetNextValidMovesSeq() iter.Seq[NextMove] {
	currSmallGameZone := g.GetNextSmallGame()

	if currSmallGameZone != 9 && !g.IsSmallGameWin(currSmallGameZone) {
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
		g.UpdateSmallGameWin(plyr, boardZone)
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
	gameStateMeta := g.Metadata >> (uint8(GameStateMetaPos) + boardZone)

	return gameStateMeta&1 == 1
}

func (g *Game) UpdateSmallGameWin(plyr *player.Player, boardZone uint8) {
	plyr.SetWinMetadata(boardZone)
	const setWinBoard uint32 = 0b1 << GameStateMetaPos
	g.Metadata |= (setWinBoard << boardZone)
}

func (g *Game) GetPlayer() *player.Player {
	isX := (g.Metadata>>NextPlayerMetaPos)&1 == 1

	var plyr *player.Player

	if isX {
		plyr = &g.X
	} else {
		plyr = &g.O
	}

	return plyr
}

func (g *Game) ChangePlayer() {
	const flipNextPlayer uint32 = 1 << NextPlayerMetaPos
	g.Metadata ^= flipNextPlayer
}
