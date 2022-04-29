package take9

import l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"

func (pPos *Position) GetOffsetMoveIndex() int {
	return pPos.OffsetMovesIndex
}

func (pPos *Position) GetCapturedPieceAtMovesIndex(movesIndex int) l03.Piece {
	return pPos.CapturedList[movesIndex]
}

func (pPos *Position) GetMoveAtMovesIndex(movesIndex int) Move {
	return pPos.Moves[movesIndex]
}
