package take8

import (
	"fmt"

	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

type positionForLocation interface {
	GetLocationOfLongPiece(int) l04.Square
}

// SprintLocation2 - あの駒どこにいんの？を表示
func SprintLocation2(pPos positionForLocation) string {
	return "\n" +
		//
		" K   k      R          B          L\n" +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		// 持ち駒は３桁になるぜ（＾～＾）
		fmt.Sprintf("|%3d|%3d|  |%3d|%3d|  |%3d|%3d|  |%3d|%3d|%3d|%3d|\n",
			pPos.GetLocationOfLongPiece(PCLOC_K1),
			pPos.GetLocationOfLongPiece(PCLOC_K2),
			pPos.GetLocationOfLongPiece(PCLOC_R1),
			pPos.GetLocationOfLongPiece(PCLOC_R2),
			pPos.GetLocationOfLongPiece(PCLOC_B1),
			pPos.GetLocationOfLongPiece(PCLOC_B2),
			pPos.GetLocationOfLongPiece(PCLOC_L1),
			pPos.GetLocationOfLongPiece(PCLOC_L2),
			pPos.GetLocationOfLongPiece(PCLOC_L3),
			pPos.GetLocationOfLongPiece(PCLOC_L4)) +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}
