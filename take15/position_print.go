package take15

import "fmt"

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
		zeroRanks[0] = ToPcCode(pPos.Board[90])
	}
	if !pPos.IsEmptySq(80) {
		zeroRanks[1] = ToPcCode(pPos.Board[80])
	}
	if !pPos.IsEmptySq(70) {
		zeroRanks[2] = ToPcCode(pPos.Board[70])
	}
	if !pPos.IsEmptySq(60) {
		zeroRanks[3] = ToPcCode(pPos.Board[60])
	}
	if !pPos.IsEmptySq(50) {
		zeroRanks[4] = ToPcCode(pPos.Board[50])
	}
	if !pPos.IsEmptySq(40) {
		zeroRanks[5] = ToPcCode(pPos.Board[40])
	}
	if !pPos.IsEmptySq(30) {
		zeroRanks[6] = ToPcCode(pPos.Board[30])
	}
	if !pPos.IsEmptySq(20) {
		zeroRanks[7] = ToPcCode(pPos.Board[20])
	}
	if !pPos.IsEmptySq(10) {
		zeroRanks[8] = ToPcCode(pPos.Board[10])
	}
	if !pPos.IsEmptySq(0) {
		zeroRanks[9] = ToPcCode(pPos.Board[0])
	}
	if !pPos.IsEmptySq(1) {
		zeroFiles[0] = ToPcCode(pPos.Board[1])
	}
	if !pPos.IsEmptySq(2) {
		zeroFiles[1] = ToPcCode(pPos.Board[2])
	}
	if !pPos.IsEmptySq(3) {
		zeroFiles[2] = ToPcCode(pPos.Board[3])
	}
	if !pPos.IsEmptySq(4) {
		zeroFiles[3] = ToPcCode(pPos.Board[4])
	}
	if !pPos.IsEmptySq(5) {
		zeroFiles[4] = ToPcCode(pPos.Board[5])
	}
	if !pPos.IsEmptySq(6) {
		zeroFiles[5] = ToPcCode(pPos.Board[6])
	}
	if !pPos.IsEmptySq(7) {
		zeroFiles[6] = ToPcCode(pPos.Board[7])
	}
	if !pPos.IsEmptySq(8) {
		zeroFiles[7] = ToPcCode(pPos.Board[8])
	}
	if !pPos.IsEmptySq(9) {
		zeroFiles[8] = ToPcCode(pPos.Board[9])
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
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[91]), ToPcCode(pPos.Board[81]), ToPcCode(pPos.Board[71]), ToPcCode(pPos.Board[61]), ToPcCode(pPos.Board[51]), ToPcCode(pPos.Board[41]), ToPcCode(pPos.Board[31]), ToPcCode(pPos.Board[21]), ToPcCode(pPos.Board[11]), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[92]), ToPcCode(pPos.Board[82]), ToPcCode(pPos.Board[72]), ToPcCode(pPos.Board[62]), ToPcCode(pPos.Board[52]), ToPcCode(pPos.Board[42]), ToPcCode(pPos.Board[32]), ToPcCode(pPos.Board[22]), ToPcCode(pPos.Board[12]), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[93]), ToPcCode(pPos.Board[83]), ToPcCode(pPos.Board[73]), ToPcCode(pPos.Board[63]), ToPcCode(pPos.Board[53]), ToPcCode(pPos.Board[43]), ToPcCode(pPos.Board[33]), ToPcCode(pPos.Board[23]), ToPcCode(pPos.Board[13]), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[94]), ToPcCode(pPos.Board[84]), ToPcCode(pPos.Board[74]), ToPcCode(pPos.Board[64]), ToPcCode(pPos.Board[54]), ToPcCode(pPos.Board[44]), ToPcCode(pPos.Board[34]), ToPcCode(pPos.Board[24]), ToPcCode(pPos.Board[14]), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[95]), ToPcCode(pPos.Board[85]), ToPcCode(pPos.Board[75]), ToPcCode(pPos.Board[65]), ToPcCode(pPos.Board[55]), ToPcCode(pPos.Board[45]), ToPcCode(pPos.Board[35]), ToPcCode(pPos.Board[25]), ToPcCode(pPos.Board[15]), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[96]), ToPcCode(pPos.Board[86]), ToPcCode(pPos.Board[76]), ToPcCode(pPos.Board[66]), ToPcCode(pPos.Board[56]), ToPcCode(pPos.Board[46]), ToPcCode(pPos.Board[36]), ToPcCode(pPos.Board[26]), ToPcCode(pPos.Board[16]), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[97]), ToPcCode(pPos.Board[87]), ToPcCode(pPos.Board[77]), ToPcCode(pPos.Board[67]), ToPcCode(pPos.Board[57]), ToPcCode(pPos.Board[47]), ToPcCode(pPos.Board[37]), ToPcCode(pPos.Board[27]), ToPcCode(pPos.Board[17]), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[98]), ToPcCode(pPos.Board[88]), ToPcCode(pPos.Board[78]), ToPcCode(pPos.Board[68]), ToPcCode(pPos.Board[58]), ToPcCode(pPos.Board[48]), ToPcCode(pPos.Board[38]), ToPcCode(pPos.Board[28]), ToPcCode(pPos.Board[18]), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPcCode(pPos.Board[99]), ToPcCode(pPos.Board[89]), ToPcCode(pPos.Board[79]), ToPcCode(pPos.Board[69]), ToPcCode(pPos.Board[59]), ToPcCode(pPos.Board[49]), ToPcCode(pPos.Board[39]), ToPcCode(pPos.Board[29]), ToPcCode(pPos.Board[19]), zeroFiles[8]) +
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
