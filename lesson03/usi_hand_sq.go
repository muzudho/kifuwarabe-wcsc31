package lesson03

func FromCodeToHandSq(code byte, convertAlternativeValue *func(code byte) HandSq) HandSq {

	switch code {
	case 'R':
		return HANDSQ_R1
	case 'B':
		return HANDSQ_B1
	case 'G':
		return HANDSQ_G1
	case 'S':
		return HANDSQ_S1
	case 'N':
		return HANDSQ_N1
	case 'L':
		return HANDSQ_L1
	case 'P':
		return HANDSQ_P1
	case 'r':
		return HANDSQ_R2
	case 'b':
		return HANDSQ_B2
	case 'g':
		return HANDSQ_G2
	case 's':
		return HANDSQ_S2
	case 'n':
		return HANDSQ_N2
	case 'l':
		return HANDSQ_L2
	case 'p':
		return HANDSQ_P2
	default:
		return (*convertAlternativeValue)(code)
	}
}
