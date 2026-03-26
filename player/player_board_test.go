package player_test

import (
	"testing"

	"github.com/SirusCodes/9x9-analysis/player"
)

func getPlayerBoard(lo, hi uint64) *player.Player {
	return &player.Player{
		Lo: lo,
		Hi: hi,
	}
}

func TestGetSmallBoard(t *testing.T) {
	tt := []struct {
		lo        uint64
		hi        uint64
		boardZone uint8
		expected  uint64
	}{
		{
			lo:        0b111000,
			boardZone: 0,
			expected:  0b111000,
		},
		{
			lo:        0b111000 << 9,
			boardZone: 1,
			expected:  0b111000,
		},
		{
			hi:        0b111000,
			boardZone: 7,
			expected:  0b111000,
		},
		{
			hi:        0b111 << 9,
			boardZone: 8,
			expected:  0b111,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(test.lo, test.hi)
		if board := player.GetSmallBoard(test.boardZone); board != test.expected {
			t.Fatalf("test failed for %+v (%d), expected: %064b got: %064b", player, i, test.expected, player.Hi)
		}
	}
}

func TestHasPlayed(t *testing.T) {
	tt := []struct {
		lo        uint64
		hi        uint64
		boardZone uint8
		position  uint8
		expected  bool
	}{
		{
			lo:        0b111001,
			boardZone: 0,
			position:  0,
			expected:  true,
		},
		{
			lo:        0b111000000 << 9,
			boardZone: 1,
			position:  2,
			expected:  false,
		},
		{
			hi:        0b111000000,
			boardZone: 7,
			position:  1,
			expected:  false,
		},
		{
			hi:        0b111000010 << 9,
			boardZone: 8,
			position:  1,
			expected:  true,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(test.lo, test.hi)
		if val := player.HasPlayed(test.boardZone, test.position); val != test.expected {
			t.Fatalf("test failed for %+v (%d), expected: %v got: %v", player, i, test.expected, val)
		}
	}
}

func TestPlay(t *testing.T) {
	tt := []struct {
		lo         uint64
		hi         uint64
		boardZone  uint8
		position   uint8
		expectedLo uint64
		expectedHi uint64
	}{
		{
			lo:         0b111000,
			boardZone:  0,
			position:   0,
			expectedLo: 0b111001,
		},
		{
			lo:         0b111000000 << 9,
			boardZone:  1,
			position:   1,
			expectedLo: 0b111000010 << 9,
		},
		{
			hi:         0b111000000,
			boardZone:  7,
			position:   0,
			expectedHi: 0b111000001,
		},
		{
			hi:         0b111000000 << 9,
			boardZone:  8,
			position:   1,
			expectedHi: 0b111000010 << 9,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(test.lo, test.hi)
		player.Play(test.boardZone, test.position)
		if player.Lo != test.expectedLo {
			t.Fatalf("test failed for %+v (%d), expectedLo: %064b got: %064b", player, i, test.expectedLo, player.Lo)
		}

		if player.Hi != test.expectedHi {
			t.Fatalf("test failed for %+v (%d), expectedHi: %064b got: %064b", player, i, test.expectedHi, player.Hi)
		}
	}
}

func TestSetWinMetadata(t *testing.T) {
	tt := []struct {
		Hi         uint64
		boardZone  uint8
		expectedHi uint64
	}{
		{
			Hi:         0b111000000 << (9 * 2),
			boardZone:  4,
			expectedHi: 0b111010000 << (9 * 2),
		},
		{
			Hi:         0b110100000 << (9 * 2),
			boardZone:  0,
			expectedHi: 0b110100001 << (9 * 2),
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(0, test.Hi)
		player.SetWinMetadata(test.boardZone)

		if player.Hi != test.expectedHi {
			t.Fatalf("test failed for %+v (%d), expected: %064b got: %064b", player, i, test.expectedHi, player.Hi)
		}
	}
}

func TestGetWinMetadata(t *testing.T) {
	tt := []struct {
		Hi               uint64
		expectedMetadata uint16
	}{
		{
			Hi:               0b111000000 << (9 * 2),
			expectedMetadata: 0b111000000,
		},
		{
			Hi:               0b110100000 << (9 * 2),
			expectedMetadata: 0b110100000,
		},
	}

	for i, test := range tt {
		player := getPlayerBoard(0, test.Hi)

		if val := player.GetWinMetadata(); val != test.expectedMetadata {
			t.Fatalf("test failed for %+v (%d), expected: %016b got: %016b", player, i, test.expectedMetadata, val)
		}
	}
}
