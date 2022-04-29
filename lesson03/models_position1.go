package lesson03

func (pPos *Position) GetPieceAtSq(sq Square) Piece {
	return pPos.Board[sq]
}

func (pPos *Position) GetPieceAtIndex(idx int) Piece {
	return pPos.Board[idx]
}
