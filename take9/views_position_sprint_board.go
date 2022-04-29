package take9

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// Print - 局面出力（＾ｑ＾）
func SprintBoard(pPos *Position) string {
	var phase_str = "?"
	if pPos.Phase == l03.FIRST {
		phase_str = "First"
	} else if pPos.Phase == l03.SECOND {
		phase_str = "Second"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPos.StartMovesNum, (pPos.StartMovesNum+pPos.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[7], pPos.Hands[8], pPos.Hands[9], pPos.Hands[10], pPos.Hands[11], pPos.Hands[12], pPos.Hands[13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pPos.Board[90].ToCodeOfPc(), pPos.Board[80].ToCodeOfPc(), pPos.Board[70].ToCodeOfPc(), pPos.Board[60].ToCodeOfPc(), pPos.Board[50].ToCodeOfPc(), pPos.Board[40].ToCodeOfPc(), pPos.Board[30].ToCodeOfPc(), pPos.Board[20].ToCodeOfPc(), pPos.Board[10].ToCodeOfPc(), pPos.Board[0].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[91].ToCodeOfPc(), pPos.Board[81].ToCodeOfPc(), pPos.Board[71].ToCodeOfPc(), pPos.Board[61].ToCodeOfPc(), pPos.Board[51].ToCodeOfPc(), pPos.Board[41].ToCodeOfPc(), pPos.Board[31].ToCodeOfPc(), pPos.Board[21].ToCodeOfPc(), pPos.Board[11].ToCodeOfPc(), pPos.Board[1].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[92].ToCodeOfPc(), pPos.Board[82].ToCodeOfPc(), pPos.Board[72].ToCodeOfPc(), pPos.Board[62].ToCodeOfPc(), pPos.Board[52].ToCodeOfPc(), pPos.Board[42].ToCodeOfPc(), pPos.Board[32].ToCodeOfPc(), pPos.Board[22].ToCodeOfPc(), pPos.Board[12].ToCodeOfPc(), pPos.Board[2].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[93].ToCodeOfPc(), pPos.Board[83].ToCodeOfPc(), pPos.Board[73].ToCodeOfPc(), pPos.Board[63].ToCodeOfPc(), pPos.Board[53].ToCodeOfPc(), pPos.Board[43].ToCodeOfPc(), pPos.Board[33].ToCodeOfPc(), pPos.Board[23].ToCodeOfPc(), pPos.Board[13].ToCodeOfPc(), pPos.Board[3].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[94].ToCodeOfPc(), pPos.Board[84].ToCodeOfPc(), pPos.Board[74].ToCodeOfPc(), pPos.Board[64].ToCodeOfPc(), pPos.Board[54].ToCodeOfPc(), pPos.Board[44].ToCodeOfPc(), pPos.Board[34].ToCodeOfPc(), pPos.Board[24].ToCodeOfPc(), pPos.Board[14].ToCodeOfPc(), pPos.Board[4].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[95].ToCodeOfPc(), pPos.Board[85].ToCodeOfPc(), pPos.Board[75].ToCodeOfPc(), pPos.Board[65].ToCodeOfPc(), pPos.Board[55].ToCodeOfPc(), pPos.Board[45].ToCodeOfPc(), pPos.Board[35].ToCodeOfPc(), pPos.Board[25].ToCodeOfPc(), pPos.Board[15].ToCodeOfPc(), pPos.Board[5].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[96].ToCodeOfPc(), pPos.Board[86].ToCodeOfPc(), pPos.Board[76].ToCodeOfPc(), pPos.Board[66].ToCodeOfPc(), pPos.Board[56].ToCodeOfPc(), pPos.Board[46].ToCodeOfPc(), pPos.Board[36].ToCodeOfPc(), pPos.Board[26].ToCodeOfPc(), pPos.Board[16].ToCodeOfPc(), pPos.Board[6].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[97].ToCodeOfPc(), pPos.Board[87].ToCodeOfPc(), pPos.Board[77].ToCodeOfPc(), pPos.Board[67].ToCodeOfPc(), pPos.Board[57].ToCodeOfPc(), pPos.Board[47].ToCodeOfPc(), pPos.Board[37].ToCodeOfPc(), pPos.Board[27].ToCodeOfPc(), pPos.Board[17].ToCodeOfPc(), pPos.Board[7].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[98].ToCodeOfPc(), pPos.Board[88].ToCodeOfPc(), pPos.Board[78].ToCodeOfPc(), pPos.Board[68].ToCodeOfPc(), pPos.Board[58].ToCodeOfPc(), pPos.Board[48].ToCodeOfPc(), pPos.Board[38].ToCodeOfPc(), pPos.Board[28].ToCodeOfPc(), pPos.Board[18].ToCodeOfPc(), pPos.Board[8].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.Board[99].ToCodeOfPc(), pPos.Board[89].ToCodeOfPc(), pPos.Board[79].ToCodeOfPc(), pPos.Board[69].ToCodeOfPc(), pPos.Board[59].ToCodeOfPc(), pPos.Board[49].ToCodeOfPc(), pPos.Board[39].ToCodeOfPc(), pPos.Board[29].ToCodeOfPc(), pPos.Board[19].ToCodeOfPc(), pPos.Board[9].ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[0], pPos.Hands[1], pPos.Hands[2], pPos.Hands[3], pPos.Hands[4], pPos.Hands[5], pPos.Hands[6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	moves_text := pPos.createMovesText()

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
}
