package take14

import l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"

// GetLongPiece - 長い利きの駒の場所を取得
func (pPos *Position) GetLocationOfLongPiece(index int) l03.Square {
	return pPos.PieceLocations[index]
}
