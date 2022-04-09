package take16

import l10 "github.com/muzudho/kifuwarabe-wcsc31/take10"

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq Square) {
	if !OnBoard(sq) && !OnHands(sq) {
		panic(App.LogNotEcho.Fatal("ValidateSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *Position, sq Square) {
	piece := pPos.Board[sq]
	if piece == l10.PIECE_EMPTY {
		panic(App.LogNotEcho.Fatal("LogicalError: There are not piece in sq=%d", sq))
	}
}
