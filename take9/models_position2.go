package take9

func (pPos *Position) GetOffsetMoveIndex() int {
	return pPos.OffsetMovesIndex
}

func (pPos *Position) GetCapturedPieceAtMovesIndex(movesIndex int) Piece {
	return pPos.CapturedList[movesIndex]
}

func (pPos *Position) GetMoveAtMovesIndex(movesIndex int) Move {
	return pPos.Moves[movesIndex]
}
