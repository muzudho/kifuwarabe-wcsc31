package take10 // not same take7

const (
	// 持ち駒を打つ 0～13 (Index)
	HAND_K1 = iota // 0: 先手玉
	HAND_R1
	HAND_B1
	HAND_G1
	HAND_S1
	HAND_N1
	HAND_L1
	HAND_P1
	HAND_K2
	HAND_R2
	HAND_B2
	HAND_G2
	HAND_S2
	HAND_N2
	HAND_L2
	HAND_P2
	HAND_IDX_START = HAND_R1
	HAND_IDX_END   = HAND_P2 - 1 // この数を含まない
	HAND_TYPE_SIZE = HAND_K2
	HAND_SIZE      = HAND_P2 + 1
)
