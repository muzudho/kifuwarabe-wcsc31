package take8

func (pPos *Position) GetOffsetMoveIndex() int {
	return pPos.OffsetMovesIndex
}

func (pPos *Position) GetNameOfCapturedPieceAtMovesIndex(movesIndex int) string {
	return pPos.CapturedList[movesIndex]
}

func (pPos *Position) GetMoveAtMovesIndex(movesIndex int) Move {
	return pPos.Moves[movesIndex]
}
