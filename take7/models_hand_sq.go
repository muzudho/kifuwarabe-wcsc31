package take7 // not same take6

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

type HandSq l04.Square

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	HANDSQ_R1 HandSq = 100 + iota
	HANDSQ_B1
	HANDSQ_G1
	HANDSQ_S1
	HANDSQ_N1
	HANDSQ_L1
	HANDSQ_P1
	HANDSQ_R2
	HANDSQ_B2
	HANDSQ_G2
	HANDSQ_S2
	HANDSQ_N2
	HANDSQ_L2
	HANDSQ_P2
	HANDSQ_ORIGIN    = HANDSQ_R1
	HANDSQ_TYPE_SIZE = HANDSQ_P1 - HANDSQ_ORIGIN + 1
)

func (h HandSq) ToSq() l04.Square {
	return l04.Square(h)
}
