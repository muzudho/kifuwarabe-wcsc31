package lesson03

func FromCodeToHandIndex(code byte, convertAlternativeValue *func(code byte) HandIdx) HandIdx {

	switch code {
	case 'R':
		return HAND_R1
	case 'B':
		return HAND_B1
	case 'G':
		return HAND_G1
	case 'S':
		return HAND_S1
	case 'N':
		return HAND_N1
	case 'L':
		return HAND_L1
	case 'P':
		return HAND_P1
	case 'r':
		return HAND_R2
	case 'b':
		return HAND_B2
	case 'g':
		return HAND_G2
	case 's':
		return HAND_S2
	case 'n':
		return HAND_N2
	case 'l':
		return HAND_L2
	case 'p':
		return HAND_P2
	default:
		return (*convertAlternativeValue)(code)
	}
}
