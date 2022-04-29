package take16

import l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"

func FromCodeToHandIndex(code byte, convertAlternativeValue *func(code byte) l03.HandIdx) l03.HandIdx {

	switch code {
	case 'R':
		return l03.HAND_R1
	case 'B':
		return l03.HAND_B1
	case 'G':
		return l03.HAND_G1
	case 'S':
		return l03.HAND_S1
	case 'N':
		return l03.HAND_N1
	case 'L':
		return l03.HAND_L1
	case 'P':
		return l03.HAND_P1
	case 'r':
		return l03.HAND_R2
	case 'b':
		return l03.HAND_B2
	case 'g':
		return l03.HAND_G2
	case 's':
		return l03.HAND_S2
	case 'n':
		return l03.HAND_N2
	case 'l':
		return l03.HAND_L2
	case 'p':
		return l03.HAND_P2
	default:
		return (*convertAlternativeValue)(code)
	}
}
