package take12

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

// GetLongPiece - 長い利きの駒の場所を取得
func (pPos *Position) GetLocationOfLongPiece(index int) l04.Square {
	return pPos.PieceLocations[index]
}
