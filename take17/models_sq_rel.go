package take17

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// GetSqSouthOf - １つ南のマス。無ければ空マス
func GetSqSouthOf(turn l03.Phase, src l03.Square) l03.Square {
	var relative int

	switch turn {
	case l03.FIRST:
		relative = 1
	case l03.SECOND:
		relative = -1
	default:
		panic(App.Log.Fatal(fmt.Sprintf("turn=[%d]", turn)))
	}

	var newRank = l03.Square(int(l03.Rank(src)) + relative)

	if 1 <= newRank && newRank < 10 {
		// 盤内
		var newFile = l03.File(src)
		// if App.IsDebug {
		// 	App.Out.Print("# newFile=%d newRank=%d\n", newFile, newRank)
		// }

		// 移動先マスの南の座標
		return l03.FromFileRankToSq(newFile, newRank)
	}

	return l03.SQ_EMPTY
}
