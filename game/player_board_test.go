package game_test

import (
	"testing"

	"github.com/SirusCodes/9x9-analysis/game"
)

func getPlayerBoard(lo, hi uint64) *game.Player {
	return &game.Player{
		Lo: lo,
		Hi: hi,
	}
}

func TestUpdateWinMetadata(t *testing.T) {
	tt := []struct {
		Hi        uint64
		boardZone uint8
		expected  uint64
	}{
		{
			Hi:        0b000000001 << 18,
			boardZone: 1,
			expected:  0b000000011 << 18,
		},
		{
			Hi:        0b000000001 << 18,
			boardZone: 2,
			expected:  0b101 << 18,
		},
		{
			Hi:        0b000000001 << 18,
			boardZone: 8,
			expected:  0b100000001 << 18,
		},
		{
			Hi:        0b000000000,
			boardZone: 0,
			expected:  0b000000001 << 18,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(0, test.Hi)
		if player.UpdateWinMetadata(test.boardZone); player.Hi != test.expected {
			t.Fatalf("test failed for %+v (%d), expected: %v got: %v", player, i, test.expected, player.Hi)
		}
	}
}

func TestIsWin(t *testing.T) {
	tt := []struct {
		Hi     uint64
		answer bool
	}{
		{
			Hi:     uint64(0b111000000 << (9 * 2)),
			answer: true,
		},
		{
			Hi:     uint64(0b110100000 << (9 * 2)),
			answer: false,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(0, test.Hi)
		if val := player.IsWin(); val != test.answer {
			t.Fatalf("test failed for %+v (%d), expected: %v got: %v", player, i, test.answer, val)
		}
	}
}

func TestIsSmallWin(t *testing.T) {
	tt := []struct {
		Lo     uint64
		Hi     uint64
		zone   uint8
		answer bool
	}{
		{
			Lo:     0b000111,
			zone:   0,
			answer: true,
		},
		{
			Hi:     0b10010001,
			zone:   7,
			answer: false,
		},
		{
			Hi:     0b001001001,
			zone:   7,
			answer: true,
		},
		{
			zone:   0,
			answer: false,
		},
		{
			Hi:     0b001001001000000000,
			zone:   8,
			answer: true,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(test.Lo, test.Hi)
		if val := player.IsSmallWin(test.zone); val != test.answer {
			t.Fatalf("test failed for %+v (%d), expected: %v got: %v", player, i, test.answer, val)
		}
	}
}

func TestCheckWin(t *testing.T) {
	tt := []struct {
		test   uint16
		answer bool
	}{
		{
			test:   0,
			answer: false,
		},
		{
			test:   0b001110000,
			answer: false,
		},
		{
			test:   0b0100100100,
			answer: true,
		},
		{
			test:   0b100010001,
			answer: true,
		},
		{
			test:   0b010010001,
			answer: false,
		},
		// All Columns
		{
			test:   0b1001001,
			answer: true,
		},
		{
			test:   0b10010010,
			answer: true,
		},
		{
			test:   0b100100100,
			answer: true,
		},
		// All Rows
		{
			test:   0b111,
			answer: true,
		},
		{
			test:   0b111000000,
			answer: true,
		},
		{
			test:   0b111000,
			answer: true,
		},
	}

	for _, test := range tt {
		if val := game.CheckWin(uint64(test.test)); val != test.answer {
			t.Fatalf("test failed for %b, expected: %v got: %v", test.test, test.answer, val)
		}
	}
}
