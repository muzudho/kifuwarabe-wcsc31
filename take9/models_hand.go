package take9

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	HAND_R1        = l04.Square(100)
	HAND_B1        = l04.Square(101)
	HAND_G1        = l04.Square(102)
	HAND_S1        = l04.Square(103)
	HAND_N1        = l04.Square(104)
	HAND_L1        = l04.Square(105)
	HAND_P1        = l04.Square(106)
	HAND_R2        = l04.Square(107)
	HAND_B2        = l04.Square(108)
	HAND_G2        = l04.Square(109)
	HAND_S2        = l04.Square(110)
	HAND_N2        = l04.Square(111)
	HAND_L2        = l04.Square(112)
	HAND_P2        = l04.Square(113)
	HAND_ORIGIN    = HAND_R1
	HAND_TYPE_SIZE = HAND_P1 - HAND_ORIGIN + 1
)
