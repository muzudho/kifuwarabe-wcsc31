package take15

import l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"

// GetPieceOnBoardAtSq - 盤上の駒を取得
func (pPos *Position) GetPieceOnBoardAtSq(sq l03.Square) l03.Piece {
	return pPos.Board[sq]
}

func (pPos *Position) GetPieceAtIndex(idx int) l03.Piece {
	return pPos.Board[idx]
}
