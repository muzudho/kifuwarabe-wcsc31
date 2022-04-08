package take16

const (
	// 持ち駒を打つ 100～115
	// 先手飛打
	SQ_K1         = Square(100)
	SQ_R1         = Square(101)
	SQ_B1         = Square(102)
	SQ_G1         = Square(103)
	SQ_S1         = Square(104)
	SQ_N1         = Square(105)
	SQ_L1         = Square(106)
	SQ_P1         = Square(107)
	SQ_K2         = Square(108)
	SQ_R2         = Square(109)
	SQ_B2         = Square(110)
	SQ_G2         = Square(111)
	SQ_S2         = Square(112)
	SQ_N2         = Square(113)
	SQ_L2         = Square(114)
	SQ_P2         = Square(115)
	SQ_HAND_START = SQ_K1
	SQ_HAND_END   = SQ_P2 + 1 // この数を含まない
)
