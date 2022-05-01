package take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l07 "github.com/muzudho/kifuwarabe-wcsc31/take7"
)

type positionForLocation interface {
	GetLocationOfLongPiece(int) l03.Square
}

// SprintLocation - あの駒どこにいんの？を表示
func SprintLocation(pPos positionForLocation) string {
	return "\n" +
		//
		" K   k      R          B          L\n" +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		// 持ち駒は３桁になるぜ（＾～＾）
		fmt.Sprintf("|%3d|%3d|  |%3d|%3d|  |%3d|%3d|  |%3d|%3d|%3d|%3d|\n",
			pPos.GetLocationOfLongPiece(l07.PCLOC_K1),
			pPos.GetLocationOfLongPiece(l07.PCLOC_K2),
			pPos.GetLocationOfLongPiece(l07.PCLOC_R1),
			pPos.GetLocationOfLongPiece(l07.PCLOC_R2),
			pPos.GetLocationOfLongPiece(l07.PCLOC_B1),
			pPos.GetLocationOfLongPiece(l07.PCLOC_B2),
			pPos.GetLocationOfLongPiece(l07.PCLOC_L1),
			pPos.GetLocationOfLongPiece(l07.PCLOC_L2),
			pPos.GetLocationOfLongPiece(l07.PCLOC_L3),
			pPos.GetLocationOfLongPiece(l07.PCLOC_L4)) +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}
