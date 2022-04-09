package take7

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	HAND_R1        = Square(100)
	HAND_B1        = Square(101)
	HAND_G1        = Square(102)
	HAND_S1        = Square(103)
	HAND_N1        = Square(104)
	HAND_L1        = Square(105)
	HAND_P1        = Square(106)
	HAND_R2        = Square(107)
	HAND_B2        = Square(108)
	HAND_G2        = Square(109)
	HAND_S2        = Square(110)
	HAND_N2        = Square(111)
	HAND_L2        = Square(112)
	HAND_P2        = Square(113)
	HAND_ORIGIN    = HAND_R1
	HAND_TYPE_SIZE = HAND_P1 - HAND_ORIGIN
)
