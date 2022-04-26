package take15

import (
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq l04.Square) {
	if !OnBoard(sq) && !OnHands(sq) {
		panic(App.LogNotEcho.Fatal("ValidateSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *Position, sq l04.Square) {
	piece := pPos.Board[sq]
	if piece == l09.PIECE_EMPTY {
		panic(App.LogNotEcho.Fatal("LogicalError: There are not piece in sq=%d", sq))
	}
}
