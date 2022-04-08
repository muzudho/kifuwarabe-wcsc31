package take11

// Hand piece type (先後付きの持ち駒の種類)
// 持ち駒を打つときに利用。 0～15
const (
	HAND_K1        = 0
	HAND_R1        = 1 // 先手飛打
	HAND_B1        = 2
	HAND_G1        = 3
	HAND_S1        = 4
	HAND_N1        = 5
	HAND_L1        = 6
	HAND_P1        = 7
	HAND_K2        = 8
	HAND_R2        = 9
	HAND_B2        = 10
	HAND_G2        = 11
	HAND_S2        = 12
	HAND_N2        = 13
	HAND_L2        = 14
	HAND_P2        = 15
	HAND_SIZE      = 16
	HAND_TYPE_SIZE = 8
	HAND_IDX_START = HAND_K1
	HAND_IDX_END   = HAND_SIZE // この数を含まない
)
