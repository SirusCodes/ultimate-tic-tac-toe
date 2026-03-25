package game_test

import (
	"slices"
	"testing"

	"github.com/SirusCodes/9x9-analysis/game"
	"github.com/SirusCodes/9x9-analysis/player"
)

const xTurnMetadata uint16 = 1 << game.NextPlayerMetaPos

func TestGetNextValidMovesSeq(t *testing.T) {
	g := getGameWithMetadata(0)

	g.PlayMove(&g.O, 0, 1)
	g.PlayMove(&g.X, 0, 3)
	g.PlayMove(&g.O, 0, 4)
	g.PlayMove(&g.X, 0, 0)

	moves := slices.Collect(g.GetNextValidMovesSeq())

	expected := []game.NextMove{
		{BoardZone: 0, Position: 2},
		{BoardZone: 0, Position: 5},
		{BoardZone: 0, Position: 6},
		{BoardZone: 0, Position: 7},
		{BoardZone: 0, Position: 8},
	}

	if len(moves) != len(expected) {
		t.Fatalf("move count mismatch. expected: %d, got: %d", len(expected), len(moves))
	}

	for _, move := range moves {
		if !slices.Contains(expected, move) {
			t.Fatalf("couldn't find %v in the result", move)
		}
	}

	// Play a win condition and check if the win board is part of the next moves
	g.PlayMove(&g.O, 0, 7)
	if !g.O.IsSmallWin(0) {
		t.Fatal("not win for O verify tests!")
	}

	g.O.SetWinMetadata(0)

	g.PlayMove(&g.X, 7, 0)

	for move := range g.GetNextValidMovesSeq() {
		if move.BoardZone == 0 {
			t.Fatalf("boardZone 0 is won so shouldn't be part of the legal moves but found: %+v", move)
		}
	}
}

func TestPlayMove(t *testing.T) {
	tt := []struct {
		boardZone, position uint8
	}{
		{boardZone: 0, position: 2},
		{boardZone: 1, position: 2},
		{boardZone: 5, position: 5},
	}

	for i, test := range tt {
		g := getGameWithMetadata(0)

		g.PlayMove(&g.O, test.boardZone, test.position)

		if !g.O.HasPlayed(test.boardZone, test.position) {
			t.Fatalf("didn't play the proper position for %d index", i)
		}

		if g.GetNextSmallGame() != test.position {
			t.Fatalf("didn't update the next play zone for %d index", i)
		}

		if g.GetPlayer() != &g.X {
			t.Fatalf("didn't update the next player for %d index", i)
		}
	}
}

func TestEvaluation(t *testing.T) {
	// TODO: Need to decide on metrics first
}

func TestGetNextSmallGame(t *testing.T) {
	tt := []struct {
		metadata uint16
		expected uint8
	}{
		{
			metadata: 0b000000000,
			expected: 9,
		},
		{
			metadata: 0b000001000,
			expected: 3,
		},
		{
			metadata: 0b100000000,
			expected: 8,
		},
	}

	for i, test := range tt {
		g := getGameWithMetadata(test.metadata)
		if val := g.GetNextSmallGame(); val != test.expected {
			t.Fatalf("next zones not correct (%d) expected: %v, got: %v", i, test.expected, val)
		}
	}
}

func TestUpdateNextGameZone(t *testing.T) {
	tt := []struct {
		metadata  uint16
		boardZone uint8
		expected  uint16
	}{
		{
			metadata:  0b000000000,
			boardZone: 0,
			expected:  0b000000001,
		},
		{
			metadata:  0b000001000,
			boardZone: 4,
			expected:  0b000010000,
		},
	}

	for i, test := range tt {
		g := getGameWithMetadata(test.metadata)
		if g.UpdateNextGameZone(test.boardZone); g.Metadata != test.expected {
			t.Fatalf("next zones not updated (%d) expected: %b, got: %b", i, test.expected, g.Metadata)
		}
	}
}

func TestIsSmallGameWin(t *testing.T) {
	const playerWinMetadata int = 18

	tt := []struct {
		playerMetadata uint64
		boardZone      uint8
		expected       bool
	}{
		{
			playerMetadata: 0b1 << playerWinMetadata,
			boardZone:      0,
			expected:       true,
		},
		{
			playerMetadata: 0b1 << playerWinMetadata,
			boardZone:      3,
			expected:       false,
		},
		{
			playerMetadata: 0b100000000 << playerWinMetadata,
			boardZone:      8,
			expected:       true,
		},
		{
			playerMetadata: 0b011111111 << playerWinMetadata,
			boardZone:      8,
			expected:       false,
		},
	}

	for i, test := range tt {
		g := getGameWithPlayers(player.NewPlayer(0, test.playerMetadata), player.NewPlayer(0, 0))
		if val := g.IsSmallGameWin(test.boardZone); val != test.expected {
			t.Fatalf("win zones not correct (%d) expected: %v, got: %v", i, test.expected, val)
		}
	}
}

func TestGetPlayer(t *testing.T) {
	x := player.NewPlayer(1, 0)
	o := player.NewPlayer(0, 0)
	g := game.NewGame(x, o, xTurnMetadata)

	if val := g.GetPlayer(); *val != x {
		t.Fatalf("incorrect user; expected: %v, got: %v", &x, val)
	}
}

func TestChangePlayer(t *testing.T) {
	g := getGame()

	if val := g.Metadata >> game.NextPlayerMetaPos; val != 1 {
		t.Fatalf("incorrect starting position, expected to be X's turn, got %b", val)
	}

	g.ChangePlayer()

	if val := g.Metadata >> game.NextPlayerMetaPos; val == 1 {
		t.Fatalf("didn't change the metadata player property, expected: 0, got: %b", val)
	}
}

func getGameWithPlayers(x, o player.Player) game.Game {
	return game.NewGame(x, o, 0)
}

func getGame() game.Game {
	x := player.NewPlayer(0, 0)
	o := player.NewPlayer(0, 0)

	x.Play(4, 4)
	o.Play(4, 0)
	x.Play(0, 1)
	o.Play(1, 2)

	return game.NewGame(x, o, xTurnMetadata)
}

func getGameWithMetadata(metadata uint16) game.Game {
	x := player.NewPlayer(0, 0)
	o := player.NewPlayer(0, 0)

	return game.NewGame(x, o, metadata)
}
