package game

const (
	smallGameSize uint64 = 9
)

type Printer interface {
	Print()
}

type PlayerBoard interface {
	IsSmallWin(boardZone uint8) bool
	IsWin() bool
	SetWinMetadata(boardZone uint8)
	GetSmallBoard(boardZone uint8) uint64
	PlaySmallBoard(boardZone, position uint8)
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

func (pb *Player) PlaySmallBoard(boardZone, position uint8) {
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

func CheckWin(check uint64) bool {
	toCheck := uint16(check) & 0b111111111

	const (
		row    uint16 = 0b111000000
		col    uint16 = 0b100100100
		ltToRb uint16 = 0b100010001
		rtToLb uint16 = 0b001010100
	)

	for i := range 3 {
		shiftedRow := (row >> (i * 3))
		shiftedCol := (col >> i)

		if (toCheck&shiftedRow) == shiftedRow || (toCheck&shiftedCol) == shiftedCol {
			return true
		}
	}

	if toCheck&ltToRb == ltToRb || toCheck&rtToLb == rtToLb {
		return true
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
