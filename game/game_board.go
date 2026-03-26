package game

import (
	"iter"

	"github.com/SirusCodes/9x9-analysis/player"
	"github.com/SirusCodes/9x9-analysis/utils"
)

const (
	NextPlayerMetaPos uint32 = 9
	FilledBoard       uint64 = 0b111111111
	EmptyLowerBits    uint16 = 0b1111111000000000
)

const (
	edge                    = 1
	corner                  = 2
	center                  = 3
	smallPartialWin         = 5
	smallWin                = 10
	smallOpponentWin        = -12
	sendOpponentToFreeBoard = -10
	bigPartialWin           = 50
	bigOpponentWin          = -60
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

type Move struct {
	BoardZone, Position uint8
}

func NewGame(X, O player.Player, Metadata uint16) Game {
	return Game{
		X:        player.NewPlayer(X.Lo, X.Hi),
		O:        player.NewPlayer(O.Lo, O.Hi),
		Metadata: Metadata,
	}
}

func (g *Game) GetNextValidMovesSeq() iter.Seq[Move] {
	currSmallGameZone := g.GetNextSmallGame()

	if !g.canPlayAnywhere(currSmallGameZone) {
		return g.getValidMovesInBoardZoneSeq(currSmallGameZone)
	}

	return func(yield func(Move) bool) {
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

func (g *Game) canPlayAnywhere(currSmallGameZone uint8) bool {
	if currSmallGameZone == 9 || g.IsSmallGameWin(currSmallGameZone) {
		return true
	}

	board := g.O.GetSmallBoard(currSmallGameZone) | g.X.GetSmallBoard(currSmallGameZone)
	isFilled := (board & FilledBoard) == FilledBoard

	return isFilled
}

func (g *Game) getValidMovesInBoardZoneSeq(boardZone uint8) iter.Seq[Move] {
	smallBoard := g.X.GetSmallBoard(boardZone) | g.O.GetSmallBoard(boardZone)

	// Get blank spaces
	blanks := smallBoard ^ FilledBoard

	// Play next moves
	return func(yield func(Move) bool) {
		for i := range uint8(9) {
			isBlank := ((blanks >> i) & 1) == 1
			if !isBlank {
				continue
			}

			if !yield(Move{BoardZone: boardZone, Position: i}) {
				return
			}
		}
	}
}

func (g *Game) PlayMove(boardZone, position uint8) {
	plyr, _ := g.GetPlayers()

	plyr.Play(boardZone, position)
	if utils.CheckWin(uint16(plyr.GetSmallBoard(boardZone))) {
		plyr.SetWinMetadata(boardZone)
	}

	g.UpdateNextGameZone(position)
	g.ChangePlayer()
}

func (g *Game) Evaluate(move Move) int {
	plyr, oppo := g.GetPlayers()

	score := 0

	//
	// Big Win
	//
	// Check big win
	plyrWinMetadata := plyr.GetWinMetadata()
	oppoWinMetadata := oppo.GetWinMetadata()
	if utils.CheckWin(plyrWinMetadata) {
		score += bigWin
	}
	// If any big wins are possible
	if wins := utils.PartialWins(plyrWinMetadata, oppoWinMetadata); wins > 0 {
		score += bigPartialWin * wins
	}
	// If opponent can have big win
	if wins := utils.PartialWins(oppoWinMetadata, plyrWinMetadata); wins > 0 {
		score += bigOpponentWin * wins
	}

	//
	// Small wins
	//
	// Check small win
	plyrBoard := uint16(plyr.GetSmallBoard(move.BoardZone))
	oppoBoard := uint16(oppo.GetSmallBoard(move.BoardZone))
	if utils.CheckWin(plyrBoard) {
		score += smallWin
		score += moveScore(&move.BoardZone) * 5 // as big wins are also important
	}
	// If any small wins are possible
	if wins := utils.PartialWins(plyrBoard, oppoBoard); wins > 0 {
		score += smallPartialWin * wins
	}
	// If opponent can have small win
	if wins := utils.PartialWins(oppoBoard, plyrBoard); wins > 0 {
		score += smallOpponentWin * wins
	}

	// If opponent gets a free move
	if g.canPlayAnywhere(move.BoardZone) {
		score += sendOpponentToFreeBoard
	}
	// for center, corner or edge
	score += moveScore(&move.Position)

	return score
}

func moveScore(move *uint8) int {
	// 0 1 2
	// 3 4 5
	// 6 7 8
	switch *move {
	case 4:
		return center
	case 0 | 2 | 6 | 8:
		return corner
	default:
		return edge
	}
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
