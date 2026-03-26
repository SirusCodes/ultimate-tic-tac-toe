package utils

import "math/bits"

var winMasks = []uint16{
	0b000000111, // row 0
	0b000111000, // row 1
	0b111000000, // row 2
	0b001001001, // col 0
	0b010010010, // col 1
	0b100100100, // col 2
	0b100010001, // diag
	0b001010100, // anti
}

func PartialWins(player uint16, opponent uint16) uint8 {
	var wins uint8 = 0

	for _, mask := range winMasks {
		p := player & mask
		o := opponent & mask

		if o == 0 && bits.OnesCount16(p) == 2 {
			wins++
		}
	}
	return wins
}

func CheckWin(check uint16) bool {
	for _, mask := range winMasks {
		if (check & mask) == mask {
			return true
		}
	}

	return false
}
