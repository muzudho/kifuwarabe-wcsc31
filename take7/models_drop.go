package take7

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	DROP_R1        = Square(100)
	DROP_B1        = Square(101)
	DROP_G1        = Square(102)
	DROP_S1        = Square(103)
	DROP_N1        = Square(104)
	DROP_L1        = Square(105)
	DROP_P1        = Square(106)
	DROP_R2        = Square(107)
	DROP_B2        = Square(108)
	DROP_G2        = Square(109)
	DROP_S2        = Square(110)
	DROP_N2        = Square(111)
	DROP_L2        = Square(112)
	DROP_P2        = Square(113)
	DROP_ORIGIN    = DROP_R1
	DROP_TYPE_SIZE = DROP_P1 - DROP_ORIGIN
)
