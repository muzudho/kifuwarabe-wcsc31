package take10 // not same take7

const (
	// 持ち駒を打つ 0～13 (Index)
	HAND_R1        = 0 // 先手飛打
	HAND_B1        = 1
	HAND_G1        = 2
	HAND_S1        = 3
	HAND_N1        = 4
	HAND_L1        = 5
	HAND_P1        = 6
	HAND_R2        = 7
	HAND_B2        = 8
	HAND_G2        = 9
	HAND_S2        = 10
	HAND_N2        = 11
	HAND_L2        = 12
	HAND_P2        = 13
	HAND_IDX_START = HAND_R1
	HAND_IDX_END   = HAND_P2 - 1 // この数を含まない
	HAND_TYPE_SIZE = 7
)
