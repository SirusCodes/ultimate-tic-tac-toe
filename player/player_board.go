package player

import "math/bits"

const (
	smallGameSize uint64 = 9
)

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

type Player struct {
	Lo uint64
	Hi uint64
}

func NewPlayer(Lo, Hi uint64) Player {
	return Player{Lo: Lo, Hi: Hi}
}

func (pb *Player) GetSmallBoard(boardZone uint8) uint64 {
	if boardZone <= 6 {
		// In low bits
		return pb.Lo >> uint64(boardZone*uint8(smallGameSize))
	}

	// In high bits
	return pb.Hi >> uint64((boardZone-7)*uint8(smallGameSize))
}

func (p *Player) HasPlayed(boardZone, position uint8) bool {
	if boardZone <= 6 {
		// In low bits
		shifted := p.Lo >> (smallGameSize*uint64(boardZone) + uint64(position))
		return (shifted & 1) == 1
	}

	// In high bits
	shifted := p.Hi >> (smallGameSize*uint64(boardZone-7) + uint64(position))
	return (shifted & 1) == 1
}

func (pb *Player) Play(boardZone, position uint8) {
	if boardZone <= 6 {
		// In low bits
		pb.Lo = pb.Lo | (0b1 << (smallGameSize*uint64(boardZone) + uint64(position)))
		return
	}

	// In high bits
	pb.Hi = pb.Hi | (0b1 << (smallGameSize*uint64(boardZone-7) + uint64(position)))
}

func (pb *Player) IsSmallWin(boardZone uint8) bool {
	return CheckWin(pb.GetSmallBoard(boardZone))
}

func PartialWins(player uint64, opponent uint64) uint8 {
	playerBoard := uint16(player) & 0b111111111
	opponentBoard := uint16(opponent) & 0b111111111

	var wins uint8 = 0

	for _, mask := range winMasks {
		p := playerBoard & mask
		o := opponentBoard & mask

		if o == 0 && bits.OnesCount16(p) == 2 {
			wins++
		}
	}
	return wins
}

func CheckWin(check uint64) bool {
	toCheck := uint16(check) & 0b111111111

	for _, mask := range winMasks {
		if (toCheck & mask) == mask {
			return true
		}
	}

	return false
}

func (pb *Player) IsWin() bool {
	toCheck := pb.Hi >> (smallGameSize * 2)

	return CheckWin(toCheck)
}

func (pb *Player) SetWinMetadata(boardZone uint8) {
	pb.Hi |= 0b1 << ((smallGameSize * 2) + uint64(boardZone))
}
