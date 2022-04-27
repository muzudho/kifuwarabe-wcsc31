package take8

import (
	"fmt"

	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

type positionForLocation interface {
	GetKingLocation(int) l04.Square
	GetRookLocation(int) l04.Square
	GetBishopLocation(int) l04.Square
	GetLanceLocation(int) l04.Square
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
			pPos.GetKingLocation(0),
			pPos.GetKingLocation(1),
			pPos.GetRookLocation(0),
			pPos.GetRookLocation(1),
			pPos.GetBishopLocation(0),
			pPos.GetBishopLocation(1),
			pPos.GetLanceLocation(0),
			pPos.GetLanceLocation(1),
			pPos.GetLanceLocation(2),
			pPos.GetLanceLocation(3)) +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}
