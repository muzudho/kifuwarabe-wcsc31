package take16

import (
	"fmt"

	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Print - 局面出力（＾ｑ＾）
func SprintBoardHeader(pPos *l15.Position, phase l06.Phase, startMovesNum int, offsetMovesIndex int) string {
	var phase_str string
	switch phase {
	case l06.FIRST:
		phase_str = "First"
	case l06.SECOND:
		phase_str = "Second"
	default:
		phase_str = "?"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats / %d value]\n", startMovesNum, (startMovesNum+offsetMovesIndex), phase_str, pPos.MaterialValue)
		//
	return s1
}
