package take13

import (
	"fmt"

	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
)

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq Square) {
	if !OnBoard(sq) && !OnHands(sq) {
		panic(fmt.Errorf("TestSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *Position, sq Square) {
	piece := pPos.Board[sq]
	if piece == l11.PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: There are not piece in sq=%d", sq))
	}
}
