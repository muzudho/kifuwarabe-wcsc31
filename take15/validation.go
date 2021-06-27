package take15

import p "github.com/muzudho/kifuwarabe-wcsc31/take16position"

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq p.Square) {
	if !p.OnBoard(sq) && !p.OnHands(sq) {
		panic(G.Log.Fatal("ValidateSq: sq=%d", sq))
	}
}

func ValidateThereArePieceIn(pPos *Position, sq p.Square) {
	piece := pPos.Board[sq]
	if piece == PIECE_EMPTY {
		panic(G.Log.Fatal("LogicalError: There are not piece in sq=%d", sq))
	}
}
