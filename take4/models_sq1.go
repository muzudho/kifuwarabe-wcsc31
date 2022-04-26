package take4

// マス番号 00～99,100～113
type Square uint32

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}
