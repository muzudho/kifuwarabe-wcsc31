package take15

import l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"

func (pPos *Position) GetPieceAtSq(sq l03.Square) l03.Piece {
	return pPos.Board[sq]
}

func (pPos *Position) GetPieceAtIndex(idx int) l03.Piece {
	return pPos.Board[idx]
}
