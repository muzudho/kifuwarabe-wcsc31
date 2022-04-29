package lesson03

type HandSq Square

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
	HANDSQ_OVER
	HANDSQ_SIZE      = (HANDSQ_OVER - HANDSQ_ORIGIN)
	HANDSQ_TYPE_SIZE = HANDSQ_SIZE / 2 // 割り切れる
	HANDSQ_ORIGIN    = HANDSQ_R1
)

const (
	HANDSQ_TYPE_SIZE_SQ Square = Square(HANDSQ_TYPE_SIZE)
)

func (h HandSq) ToSq() Square {
	return Square(h)
}

func FromSqToHandSq(sq Square) HandSq {
	return HandSq(sq)
}
