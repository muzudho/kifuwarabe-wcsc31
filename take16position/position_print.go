package take16position

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
		zeroRanks[0] = ToPieceCode(pPos.Board[90])
	}
	if !pPos.IsEmptySq(80) {
		zeroRanks[1] = ToPieceCode(pPos.Board[80])
	}
	if !pPos.IsEmptySq(70) {
		zeroRanks[2] = ToPieceCode(pPos.Board[70])
	}
	if !pPos.IsEmptySq(60) {
		zeroRanks[3] = ToPieceCode(pPos.Board[60])
	}
	if !pPos.IsEmptySq(50) {
		zeroRanks[4] = ToPieceCode(pPos.Board[50])
	}
	if !pPos.IsEmptySq(40) {
		zeroRanks[5] = ToPieceCode(pPos.Board[40])
	}
	if !pPos.IsEmptySq(30) {
		zeroRanks[6] = ToPieceCode(pPos.Board[30])
	}
	if !pPos.IsEmptySq(20) {
		zeroRanks[7] = ToPieceCode(pPos.Board[20])
	}
	if !pPos.IsEmptySq(10) {
		zeroRanks[8] = ToPieceCode(pPos.Board[10])
	}
	if !pPos.IsEmptySq(0) {
		zeroRanks[9] = ToPieceCode(pPos.Board[0])
	}
	if !pPos.IsEmptySq(1) {
		zeroFiles[0] = ToPieceCode(pPos.Board[1])
	}
	if !pPos.IsEmptySq(2) {
		zeroFiles[1] = ToPieceCode(pPos.Board[2])
	}
	if !pPos.IsEmptySq(3) {
		zeroFiles[2] = ToPieceCode(pPos.Board[3])
	}
	if !pPos.IsEmptySq(4) {
		zeroFiles[3] = ToPieceCode(pPos.Board[4])
	}
	if !pPos.IsEmptySq(5) {
		zeroFiles[4] = ToPieceCode(pPos.Board[5])
	}
	if !pPos.IsEmptySq(6) {
		zeroFiles[5] = ToPieceCode(pPos.Board[6])
	}
	if !pPos.IsEmptySq(7) {
		zeroFiles[6] = ToPieceCode(pPos.Board[7])
	}
	if !pPos.IsEmptySq(8) {
		zeroFiles[7] = ToPieceCode(pPos.Board[8])
	}
	if !pPos.IsEmptySq(9) {
		zeroFiles[8] = ToPieceCode(pPos.Board[9])
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
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[91]), ToPieceCode(pPos.Board[81]), ToPieceCode(pPos.Board[71]), ToPieceCode(pPos.Board[61]), ToPieceCode(pPos.Board[51]), ToPieceCode(pPos.Board[41]), ToPieceCode(pPos.Board[31]), ToPieceCode(pPos.Board[21]), ToPieceCode(pPos.Board[11]), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[92]), ToPieceCode(pPos.Board[82]), ToPieceCode(pPos.Board[72]), ToPieceCode(pPos.Board[62]), ToPieceCode(pPos.Board[52]), ToPieceCode(pPos.Board[42]), ToPieceCode(pPos.Board[32]), ToPieceCode(pPos.Board[22]), ToPieceCode(pPos.Board[12]), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[93]), ToPieceCode(pPos.Board[83]), ToPieceCode(pPos.Board[73]), ToPieceCode(pPos.Board[63]), ToPieceCode(pPos.Board[53]), ToPieceCode(pPos.Board[43]), ToPieceCode(pPos.Board[33]), ToPieceCode(pPos.Board[23]), ToPieceCode(pPos.Board[13]), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[94]), ToPieceCode(pPos.Board[84]), ToPieceCode(pPos.Board[74]), ToPieceCode(pPos.Board[64]), ToPieceCode(pPos.Board[54]), ToPieceCode(pPos.Board[44]), ToPieceCode(pPos.Board[34]), ToPieceCode(pPos.Board[24]), ToPieceCode(pPos.Board[14]), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[95]), ToPieceCode(pPos.Board[85]), ToPieceCode(pPos.Board[75]), ToPieceCode(pPos.Board[65]), ToPieceCode(pPos.Board[55]), ToPieceCode(pPos.Board[45]), ToPieceCode(pPos.Board[35]), ToPieceCode(pPos.Board[25]), ToPieceCode(pPos.Board[15]), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[96]), ToPieceCode(pPos.Board[86]), ToPieceCode(pPos.Board[76]), ToPieceCode(pPos.Board[66]), ToPieceCode(pPos.Board[56]), ToPieceCode(pPos.Board[46]), ToPieceCode(pPos.Board[36]), ToPieceCode(pPos.Board[26]), ToPieceCode(pPos.Board[16]), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[97]), ToPieceCode(pPos.Board[87]), ToPieceCode(pPos.Board[77]), ToPieceCode(pPos.Board[67]), ToPieceCode(pPos.Board[57]), ToPieceCode(pPos.Board[47]), ToPieceCode(pPos.Board[37]), ToPieceCode(pPos.Board[27]), ToPieceCode(pPos.Board[17]), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[98]), ToPieceCode(pPos.Board[88]), ToPieceCode(pPos.Board[78]), ToPieceCode(pPos.Board[68]), ToPieceCode(pPos.Board[58]), ToPieceCode(pPos.Board[48]), ToPieceCode(pPos.Board[38]), ToPieceCode(pPos.Board[28]), ToPieceCode(pPos.Board[18]), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", ToPieceCode(pPos.Board[99]), ToPieceCode(pPos.Board[89]), ToPieceCode(pPos.Board[79]), ToPieceCode(pPos.Board[69]), ToPieceCode(pPos.Board[59]), ToPieceCode(pPos.Board[49]), ToPieceCode(pPos.Board[39]), ToPieceCode(pPos.Board[29]), ToPieceCode(pPos.Board[19]), zeroFiles[8]) +
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

// SprintLocation - あの駒どこにいんの？を表示
func (pPos *Position) SprintLocation() string {
	return "\n" +
		//
		" K   k      R          B          L\n" +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		// 持ち駒は３桁になるぜ（＾～＾）
		fmt.Sprintf("|%3d|%3d|  |%3d|%3d|  |%3d|%3d|  |%3d|%3d|%3d|%3d|\n",
			pPos.PieceLocations[PCLOC_K1], pPos.PieceLocations[PCLOC_K2],
			pPos.PieceLocations[PCLOC_R1], pPos.PieceLocations[PCLOC_R2],
			pPos.PieceLocations[PCLOC_B1], pPos.PieceLocations[PCLOC_B2],
			pPos.PieceLocations[PCLOC_L1], pPos.PieceLocations[PCLOC_L2],
			pPos.PieceLocations[PCLOC_L3], pPos.PieceLocations[PCLOC_L4]) +
		//
		"+---+---+  +---+---+  +---+---+  +---+---+---+---+\n" +
		//
		"\n"
}
