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

// Print - 局面出力（＾ｑ＾）
func SprintBoard(pPos *l15.Position) string {
	// pPosSys.StartMovesNum
	// pPosSys.OffsetMovesIndex
	// 	moves_text := pPosSys.createMovesText()

	// 0段目
	zeroRanks := [10]string{"  9", "  8", "  7", "  6", "  5", "  4", "  3", "  2", "  1", "   "}
	// 0筋目
	zeroFiles := [9]string{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h ", " i "}

	// 0段目、0筋目に駒置いてたらそれも表示（＾～＾）
	if !pPos.IsEmptySq(90) {
		zeroRanks[0] = pPos.Board[90].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(80) {
		zeroRanks[1] = pPos.Board[80].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(70) {
		zeroRanks[2] = pPos.Board[70].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(60) {
		zeroRanks[3] = pPos.Board[60].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(50) {
		zeroRanks[4] = pPos.Board[50].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(40) {
		zeroRanks[5] = pPos.Board[40].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(30) {
		zeroRanks[6] = pPos.Board[30].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(20) {
		zeroRanks[7] = pPos.Board[20].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(10) {
		zeroRanks[8] = pPos.Board[10].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(0) {
		zeroRanks[9] = pPos.Board[0].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(1) {
		zeroFiles[0] = pPos.Board[1].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(2) {
		zeroFiles[1] = pPos.Board[2].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(3) {
		zeroFiles[2] = pPos.Board[3].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(4) {
		zeroFiles[3] = pPos.Board[4].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(5) {
		zeroFiles[4] = pPos.Board[5].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(6) {
		zeroFiles[5] = pPos.Board[6].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(7) {
		zeroFiles[6] = pPos.Board[7].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(8) {
		zeroFiles[7] = pPos.Board[8].ToCodeOfPc()
	}
	if !pPos.IsEmptySq(9) {
		zeroFiles[8] = pPos.Board[9].ToCodeOfPc()
	}

	var s1 = "\n" +
		//
		"  k  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands1[8], pPos.Hands1[9], pPos.Hands1[10], pPos.Hands1[11], pPos.Hands1[12], pPos.Hands1[13], pPos.Hands1[14], pPos.Hands1[15]) +
		//
		"+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf("%3s%3s%3s%3s%3s%3s%3s%3s%3s%3s\n", zeroRanks[0], zeroRanks[1], zeroRanks[2], zeroRanks[3], zeroRanks[4], zeroRanks[5], zeroRanks[6], zeroRanks[7], zeroRanks[8], zeroRanks[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[91].ToCodeOfPc(), pPos.Board[81].ToCodeOfPc(), pPos.Board[71].ToCodeOfPc(), pPos.Board[61].ToCodeOfPc(), pPos.Board[51].ToCodeOfPc(), pPos.Board[41].ToCodeOfPc(), pPos.Board[31].ToCodeOfPc(), pPos.Board[21].ToCodeOfPc(), pPos.Board[11].ToCodeOfPc(), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[92].ToCodeOfPc(), pPos.Board[82].ToCodeOfPc(), pPos.Board[72].ToCodeOfPc(), pPos.Board[62].ToCodeOfPc(), pPos.Board[52].ToCodeOfPc(), pPos.Board[42].ToCodeOfPc(), pPos.Board[32].ToCodeOfPc(), pPos.Board[22].ToCodeOfPc(), pPos.Board[12].ToCodeOfPc(), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[93].ToCodeOfPc(), pPos.Board[83].ToCodeOfPc(), pPos.Board[73].ToCodeOfPc(), pPos.Board[63].ToCodeOfPc(), pPos.Board[53].ToCodeOfPc(), pPos.Board[43].ToCodeOfPc(), pPos.Board[33].ToCodeOfPc(), pPos.Board[23].ToCodeOfPc(), pPos.Board[13].ToCodeOfPc(), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[94].ToCodeOfPc(), pPos.Board[84].ToCodeOfPc(), pPos.Board[74].ToCodeOfPc(), pPos.Board[64].ToCodeOfPc(), pPos.Board[54].ToCodeOfPc(), pPos.Board[44].ToCodeOfPc(), pPos.Board[34].ToCodeOfPc(), pPos.Board[24].ToCodeOfPc(), pPos.Board[14].ToCodeOfPc(), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[95].ToCodeOfPc(), pPos.Board[85].ToCodeOfPc(), pPos.Board[75].ToCodeOfPc(), pPos.Board[65].ToCodeOfPc(), pPos.Board[55].ToCodeOfPc(), pPos.Board[45].ToCodeOfPc(), pPos.Board[35].ToCodeOfPc(), pPos.Board[25].ToCodeOfPc(), pPos.Board[15].ToCodeOfPc(), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[96].ToCodeOfPc(), pPos.Board[86].ToCodeOfPc(), pPos.Board[76].ToCodeOfPc(), pPos.Board[66].ToCodeOfPc(), pPos.Board[56].ToCodeOfPc(), pPos.Board[46].ToCodeOfPc(), pPos.Board[36].ToCodeOfPc(), pPos.Board[26].ToCodeOfPc(), pPos.Board[16].ToCodeOfPc(), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[97].ToCodeOfPc(), pPos.Board[87].ToCodeOfPc(), pPos.Board[77].ToCodeOfPc(), pPos.Board[67].ToCodeOfPc(), pPos.Board[57].ToCodeOfPc(), pPos.Board[47].ToCodeOfPc(), pPos.Board[37].ToCodeOfPc(), pPos.Board[27].ToCodeOfPc(), pPos.Board[17].ToCodeOfPc(), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[98].ToCodeOfPc(), pPos.Board[88].ToCodeOfPc(), pPos.Board[78].ToCodeOfPc(), pPos.Board[68].ToCodeOfPc(), pPos.Board[58].ToCodeOfPc(), pPos.Board[48].ToCodeOfPc(), pPos.Board[38].ToCodeOfPc(), pPos.Board[28].ToCodeOfPc(), pPos.Board[18].ToCodeOfPc(), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[99].ToCodeOfPc(), pPos.Board[89].ToCodeOfPc(), pPos.Board[79].ToCodeOfPc(), pPos.Board[69].ToCodeOfPc(), pPos.Board[59].ToCodeOfPc(), pPos.Board[49].ToCodeOfPc(), pPos.Board[39].ToCodeOfPc(), pPos.Board[29].ToCodeOfPc(), pPos.Board[19].ToCodeOfPc(), zeroFiles[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"     K  R  B  G  S  N  L  P\n" +
		"   +--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("   |%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands1[0], pPos.Hands1[1], pPos.Hands1[2], pPos.Hands1[3], pPos.Hands1[4], pPos.Hands1[5], pPos.Hands1[6], pPos.Hands1[7]) +
		//
		"   +--+--+--+--+--+--+--+--+\n" +
		//
		"\n"
		//

	return s1
}
