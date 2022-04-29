package take16

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
)

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq l03.Square) {
	if !l15.OnBoard(sq) && !l15.OnHands(sq) {
		panic(App.LogNotEcho.Fatal("ValidateSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *l15.Position, sq l03.Square) {
	piece := pPos.Board[sq]
	if piece == l03.PIECE_EMPTY {
		panic(App.LogNotEcho.Fatal("LogicalError: There are not piece in sq=%d", sq))
	}
}
