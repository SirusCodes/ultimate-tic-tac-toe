package game

const (
	smallGameSize uint64 = 9
)

type Printer interface {
	Print()
}

type PlayerBoard interface {
	// Should be between 0 to 8
	IsSmallWin(boardZone uint8) bool
	IsWin() bool
	UpdateWinMetadata(boardZone uint8)
}

type Player struct {
	Lo uint64
	Hi uint64
}

func NewPlayerBoard(player uint8) PlayerBoard {
	return &Player{}
}

func (pb *Player) UpdateWinMetadata(boardZone uint8) {
	pb.Hi |= 0b1 << ((smallGameSize * 2) + uint64(boardZone))
}

func (pb *Player) IsSmallWin(boardZone uint8) bool {
	var forCheck uint64
	// Low for if between 0 and 6 else High
	if boardZone <= 6 {
		forCheck = pb.Lo >> uint64(boardZone*uint8(smallGameSize))
	} else {
		forCheck = pb.Hi >> uint64((boardZone-7)*uint8(smallGameSize))
	}

	return CheckWin(forCheck)
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
