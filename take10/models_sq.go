package take10 // not same take11

// マス番号 00～99,100～113
type Square uint32

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	SQ_R1         = Square(100)
	SQ_B1         = Square(101)
	SQ_G1         = Square(102)
	SQ_S1         = Square(103)
	SQ_N1         = Square(104)
	SQ_L1         = Square(105)
	SQ_P1         = Square(106)
	SQ_R2         = Square(107)
	SQ_B2         = Square(108)
	SQ_G2         = Square(109)
	SQ_S2         = Square(110)
	SQ_N2         = Square(111)
	SQ_L2         = Square(112)
	SQ_P2         = Square(113)
	SQ_HAND_START = SQ_R1
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
