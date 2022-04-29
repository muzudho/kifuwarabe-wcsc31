package take14

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq l04.Square) {
	if !OnBoard(sq) && !OnHands(sq) {
		panic(fmt.Errorf("TestSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *Position, sq l04.Square) {
	piece := pPos.Board[sq]
	if piece == l03.PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: There are not piece in sq=%d", sq))
	}
}
