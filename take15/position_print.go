package take15

import (
	"fmt"

	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
)

// Print - 局面出力（＾ｑ＾）
func (pPos *Position) SprintBoardHeader(phase Phase, startMovesNum int, offsetMovesIndex int) string {
	var phase_str string
	switch phase {
	case FIRST:
		phase_str = "First"
	case SECOND:
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
func (pPos *Position) SprintBoard() string {
	// pPosSys.StartMovesNum
	// pPosSys.OffsetMovesIndex
	// 	moves_text := pPosSys.createMovesText()

	// 0段目
	zeroRanks := [10]string{"  9", "  8", "  7", "  6", "  5", "  4", "  3", "  2", "  1", "   "}
	// 0筋目
	zeroFiles := [9]string{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h ", " i "}

	// 0段目、0筋目に駒置いてたらそれも表示（＾～＾）
	if !pPos.IsEmptySq(90) {
		zeroRanks[0] = l11.ToCodeOfPc(pPos.Board[90])
	}
	if !pPos.IsEmptySq(80) {
		zeroRanks[1] = l11.ToCodeOfPc(pPos.Board[80])
	}
	if !pPos.IsEmptySq(70) {
		zeroRanks[2] = l11.ToCodeOfPc(pPos.Board[70])
	}
	if !pPos.IsEmptySq(60) {
		zeroRanks[3] = l11.ToCodeOfPc(pPos.Board[60])
	}
	if !pPos.IsEmptySq(50) {
		zeroRanks[4] = l11.ToCodeOfPc(pPos.Board[50])
	}
	if !pPos.IsEmptySq(40) {
		zeroRanks[5] = l11.ToCodeOfPc(pPos.Board[40])
	}
	if !pPos.IsEmptySq(30) {
		zeroRanks[6] = l11.ToCodeOfPc(pPos.Board[30])
	}
	if !pPos.IsEmptySq(20) {
		zeroRanks[7] = l11.ToCodeOfPc(pPos.Board[20])
	}
	if !pPos.IsEmptySq(10) {
		zeroRanks[8] = l11.ToCodeOfPc(pPos.Board[10])
	}
	if !pPos.IsEmptySq(0) {
		zeroRanks[9] = l11.ToCodeOfPc(pPos.Board[0])
	}
	if !pPos.IsEmptySq(1) {
		zeroFiles[0] = l11.ToCodeOfPc(pPos.Board[1])
	}
	if !pPos.IsEmptySq(2) {
		zeroFiles[1] = l11.ToCodeOfPc(pPos.Board[2])
	}
	if !pPos.IsEmptySq(3) {
		zeroFiles[2] = l11.ToCodeOfPc(pPos.Board[3])
	}
	if !pPos.IsEmptySq(4) {
		zeroFiles[3] = l11.ToCodeOfPc(pPos.Board[4])
	}
	if !pPos.IsEmptySq(5) {
		zeroFiles[4] = l11.ToCodeOfPc(pPos.Board[5])
	}
	if !pPos.IsEmptySq(6) {
		zeroFiles[5] = l11.ToCodeOfPc(pPos.Board[6])
	}
	if !pPos.IsEmptySq(7) {
		zeroFiles[6] = l11.ToCodeOfPc(pPos.Board[7])
	}
	if !pPos.IsEmptySq(8) {
		zeroFiles[7] = l11.ToCodeOfPc(pPos.Board[8])
	}
	if !pPos.IsEmptySq(9) {
		zeroFiles[8] = l11.ToCodeOfPc(pPos.Board[9])
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
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[91]), l11.ToCodeOfPc(pPos.Board[81]), l11.ToCodeOfPc(pPos.Board[71]), l11.ToCodeOfPc(pPos.Board[61]), l11.ToCodeOfPc(pPos.Board[51]), l11.ToCodeOfPc(pPos.Board[41]), l11.ToCodeOfPc(pPos.Board[31]), l11.ToCodeOfPc(pPos.Board[21]), l11.ToCodeOfPc(pPos.Board[11]), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[92]), l11.ToCodeOfPc(pPos.Board[82]), l11.ToCodeOfPc(pPos.Board[72]), l11.ToCodeOfPc(pPos.Board[62]), l11.ToCodeOfPc(pPos.Board[52]), l11.ToCodeOfPc(pPos.Board[42]), l11.ToCodeOfPc(pPos.Board[32]), l11.ToCodeOfPc(pPos.Board[22]), l11.ToCodeOfPc(pPos.Board[12]), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[93]), l11.ToCodeOfPc(pPos.Board[83]), l11.ToCodeOfPc(pPos.Board[73]), l11.ToCodeOfPc(pPos.Board[63]), l11.ToCodeOfPc(pPos.Board[53]), l11.ToCodeOfPc(pPos.Board[43]), l11.ToCodeOfPc(pPos.Board[33]), l11.ToCodeOfPc(pPos.Board[23]), l11.ToCodeOfPc(pPos.Board[13]), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[94]), l11.ToCodeOfPc(pPos.Board[84]), l11.ToCodeOfPc(pPos.Board[74]), l11.ToCodeOfPc(pPos.Board[64]), l11.ToCodeOfPc(pPos.Board[54]), l11.ToCodeOfPc(pPos.Board[44]), l11.ToCodeOfPc(pPos.Board[34]), l11.ToCodeOfPc(pPos.Board[24]), l11.ToCodeOfPc(pPos.Board[14]), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[95]), l11.ToCodeOfPc(pPos.Board[85]), l11.ToCodeOfPc(pPos.Board[75]), l11.ToCodeOfPc(pPos.Board[65]), l11.ToCodeOfPc(pPos.Board[55]), l11.ToCodeOfPc(pPos.Board[45]), l11.ToCodeOfPc(pPos.Board[35]), l11.ToCodeOfPc(pPos.Board[25]), l11.ToCodeOfPc(pPos.Board[15]), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[96]), l11.ToCodeOfPc(pPos.Board[86]), l11.ToCodeOfPc(pPos.Board[76]), l11.ToCodeOfPc(pPos.Board[66]), l11.ToCodeOfPc(pPos.Board[56]), l11.ToCodeOfPc(pPos.Board[46]), l11.ToCodeOfPc(pPos.Board[36]), l11.ToCodeOfPc(pPos.Board[26]), l11.ToCodeOfPc(pPos.Board[16]), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[97]), l11.ToCodeOfPc(pPos.Board[87]), l11.ToCodeOfPc(pPos.Board[77]), l11.ToCodeOfPc(pPos.Board[67]), l11.ToCodeOfPc(pPos.Board[57]), l11.ToCodeOfPc(pPos.Board[47]), l11.ToCodeOfPc(pPos.Board[37]), l11.ToCodeOfPc(pPos.Board[27]), l11.ToCodeOfPc(pPos.Board[17]), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[98]), l11.ToCodeOfPc(pPos.Board[88]), l11.ToCodeOfPc(pPos.Board[78]), l11.ToCodeOfPc(pPos.Board[68]), l11.ToCodeOfPc(pPos.Board[58]), l11.ToCodeOfPc(pPos.Board[48]), l11.ToCodeOfPc(pPos.Board[38]), l11.ToCodeOfPc(pPos.Board[28]), l11.ToCodeOfPc(pPos.Board[18]), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l11.ToCodeOfPc(pPos.Board[99]), l11.ToCodeOfPc(pPos.Board[89]), l11.ToCodeOfPc(pPos.Board[79]), l11.ToCodeOfPc(pPos.Board[69]), l11.ToCodeOfPc(pPos.Board[59]), l11.ToCodeOfPc(pPos.Board[49]), l11.ToCodeOfPc(pPos.Board[39]), l11.ToCodeOfPc(pPos.Board[29]), l11.ToCodeOfPc(pPos.Board[19]), zeroFiles[8]) +
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
