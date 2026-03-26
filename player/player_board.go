package player

const (
	smallGameSize uint64 = 9
)

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

func (pb *Player) SetWinMetadata(boardZone uint8) {
	pb.Hi |= 0b1 << ((smallGameSize * 2) + uint64(boardZone))
}

func (pb *Player) GetWinMetadata() uint16 {
	return uint16(pb.Hi >> (smallGameSize * 2))
}
