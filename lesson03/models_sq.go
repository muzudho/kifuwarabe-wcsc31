package lesson03

// マス番号 00～99,100～113
type Square uint32

const (
	SQ_EMPTY Square = 0          // マス番号を指定しないことを意味するマス番号
	SQ_K1    Square = 100 + iota // 持ち駒を打つ 100～115
	SQ_R1                        // 先手飛打
	SQ_B1
	SQ_G1
	SQ_S1
	SQ_N1
	SQ_L1
	SQ_P1
	SQ_K2
	SQ_R2
	SQ_B2
	SQ_G2
	SQ_S2
	SQ_N2
	SQ_L2
	SQ_P2
	SQ_HAND_START = SQ_K1
	SQ_HAND_END   = SQ_P2 + 1 // この数を含まない
)

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}
