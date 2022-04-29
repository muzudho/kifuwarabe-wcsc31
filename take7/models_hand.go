package take7 // not same take6

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

type HandIdx l04.Square

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	HAND_R1 HandIdx = 100 + iota
	HAND_B1
	HAND_G1
	HAND_S1
	HAND_N1
	HAND_L1
	HAND_P1
	HAND_R2
	HAND_B2
	HAND_G2
	HAND_S2
	HAND_N2
	HAND_L2
	HAND_P2
	HAND_ORIGIN    = HAND_R1
	HAND_TYPE_SIZE = HAND_P1 - HAND_ORIGIN + 1
)

func (h HandIdx) ToSq() l04.Square {
	return l04.Square(h)
}
