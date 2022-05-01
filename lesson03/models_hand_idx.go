package lesson03

type HandIdx uint

// Hand piece type (先後付きの持ち駒の種類)
//
// * 0～15
// * 持ち駒を打つときに利用
// * 駒台に玉も置ける
const (
	HAND_K1 HandIdx = iota // 0: 先手玉
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
	HAND_SIZE
	HAND_TYPE_SIZE = HAND_SIZE / 2 // 割り切れる
	HAND_IDX_BEGIN = HAND_K1
	HAND_IDX_END   = HAND_SIZE // この数を含まない
)

const (
	HAND_TYPE_SIZE_SQ Square = Square(HAND_TYPE_SIZE)
)

func FromSqToHandIndex(sq Square) HandIdx {
	return HandIdx(sq)
}
