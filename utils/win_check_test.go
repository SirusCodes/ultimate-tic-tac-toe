package utils_test

import (
	"testing"

	"github.com/SirusCodes/ultimate-tic-tac-toe/utils"
)

func TestPartialWins(t *testing.T) {
	tt := []struct {
		player   uint16
		opponent uint16
		answer   int
	}{
		{
			player: 0b000000000,
			answer: 0,
		},
		{
			player: 0b001110000,
			answer: 2,
		},
		{
			player: 0b010010001,
			answer: 2,
		},
		{
			player: 0b010101010,
			answer: 2,
		},
		{
			player: 0b010101010,
			answer: 2,
		},
		{
			player:   0b010001010,
			opponent: 0b000010000,
			answer:   0,
		},
		{
			player: 0b000100001,
			answer: 0,
		},
		{
			player: 0b010000100,
			answer: 0,
		},
		{
			player: 0b000000101,
			answer: 1,
		},
		{
			player: 0b001000001,
			answer: 1,
		},
		{
			player: 0b100000001,
			answer: 1,
		},
		{
			player: 0b001000100,
			answer: 1,
		},
		{
			player:   0b001000000,
			opponent: 0b000000100,
			answer:   0,
		},
	}

	for i, test := range tt {
		if val := utils.PartialWins(test.player, test.opponent); val != test.answer {
			t.Fatalf("test failed for player(%016b) and opponent(%016b) (%d), expected: %v got: %v", test.player, test.opponent, i, test.answer, val)
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
		if val := utils.CheckWin(test.test); val != test.answer {
			t.Fatalf("test failed for %016b, expected: %v got: %v", test.test, test.answer, val)
		}
	}
}
