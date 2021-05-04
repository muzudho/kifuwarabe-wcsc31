package take13

import "fmt"

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq Square) {
	if !OnBoard(sq) && !OnHands(sq) {
		panic(fmt.Errorf("TestSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *Position, sq Square) {
	piece := pPos.Board[sq]
	if piece == PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: There are not piece in sq=%d", sq))
	}
}
