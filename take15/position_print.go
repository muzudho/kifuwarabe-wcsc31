package take15

import "fmt"

// Print - 局面出力（＾ｑ＾）
func (pPos *Position) Sprint(phase Phase, startMovesNum int, offsetMovesIndex int, moves_text string) string {
	// pPosSys.StartMovesNum
	// pPosSys.OffsetMovesIndex
	// 	moves_text := pPosSys.createMovesText()

	var phase_str string
	switch phase {
	case FIRST:
		phase_str = "First"
	case SECOND:
		phase_str = "Second"
	default:
		phase_str = "?"
	}

	// 0段目
	zeroRanks := [10]string{"  9", "  8", "  7", "  6", "  5", "  4", "  3", "  2", "  1", "   "}
	// 0筋目
	zeroFiles := [9]string{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h ", " i "}

	// 0段目、0筋目に駒置いてたらそれも表示（＾～＾）
	if !pPos.IsEmptySq(90) {
		zeroRanks[0] = pPos.Board[90].ToCode()
	}
	if !pPos.IsEmptySq(80) {
		zeroRanks[1] = pPos.Board[80].ToCode()
	}
	if !pPos.IsEmptySq(70) {
		zeroRanks[2] = pPos.Board[70].ToCode()
	}
	if !pPos.IsEmptySq(60) {
		zeroRanks[3] = pPos.Board[60].ToCode()
	}
	if !pPos.IsEmptySq(50) {
		zeroRanks[4] = pPos.Board[50].ToCode()
	}
	if !pPos.IsEmptySq(40) {
		zeroRanks[5] = pPos.Board[40].ToCode()
	}
	if !pPos.IsEmptySq(30) {
		zeroRanks[6] = pPos.Board[30].ToCode()
	}
	if !pPos.IsEmptySq(20) {
		zeroRanks[7] = pPos.Board[20].ToCode()
	}
	if !pPos.IsEmptySq(10) {
		zeroRanks[8] = pPos.Board[10].ToCode()
	}
	if !pPos.IsEmptySq(0) {
		zeroRanks[9] = pPos.Board[0].ToCode()
	}
	if !pPos.IsEmptySq(1) {
		zeroFiles[0] = pPos.Board[1].ToCode()
	}
	if !pPos.IsEmptySq(2) {
		zeroFiles[1] = pPos.Board[2].ToCode()
	}
	if !pPos.IsEmptySq(3) {
		zeroFiles[2] = pPos.Board[3].ToCode()
	}
	if !pPos.IsEmptySq(4) {
		zeroFiles[3] = pPos.Board[4].ToCode()
	}
	if !pPos.IsEmptySq(5) {
		zeroFiles[4] = pPos.Board[5].ToCode()
	}
	if !pPos.IsEmptySq(6) {
		zeroFiles[5] = pPos.Board[6].ToCode()
	}
	if !pPos.IsEmptySq(7) {
		zeroFiles[6] = pPos.Board[7].ToCode()
	}
	if !pPos.IsEmptySq(8) {
		zeroFiles[7] = pPos.Board[8].ToCode()
	}
	if !pPos.IsEmptySq(9) {
		zeroFiles[8] = pPos.Board[9].ToCode()
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", startMovesNum, (startMovesNum+offsetMovesIndex), phase_str) +
		//
		"\n" +
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
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[91].ToCode(), pPos.Board[81].ToCode(), pPos.Board[71].ToCode(), pPos.Board[61].ToCode(), pPos.Board[51].ToCode(), pPos.Board[41].ToCode(), pPos.Board[31].ToCode(), pPos.Board[21].ToCode(), pPos.Board[11].ToCode(), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[92].ToCode(), pPos.Board[82].ToCode(), pPos.Board[72].ToCode(), pPos.Board[62].ToCode(), pPos.Board[52].ToCode(), pPos.Board[42].ToCode(), pPos.Board[32].ToCode(), pPos.Board[22].ToCode(), pPos.Board[12].ToCode(), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[93].ToCode(), pPos.Board[83].ToCode(), pPos.Board[73].ToCode(), pPos.Board[63].ToCode(), pPos.Board[53].ToCode(), pPos.Board[43].ToCode(), pPos.Board[33].ToCode(), pPos.Board[23].ToCode(), pPos.Board[13].ToCode(), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[94].ToCode(), pPos.Board[84].ToCode(), pPos.Board[74].ToCode(), pPos.Board[64].ToCode(), pPos.Board[54].ToCode(), pPos.Board[44].ToCode(), pPos.Board[34].ToCode(), pPos.Board[24].ToCode(), pPos.Board[14].ToCode(), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[95].ToCode(), pPos.Board[85].ToCode(), pPos.Board[75].ToCode(), pPos.Board[65].ToCode(), pPos.Board[55].ToCode(), pPos.Board[45].ToCode(), pPos.Board[35].ToCode(), pPos.Board[25].ToCode(), pPos.Board[15].ToCode(), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[96].ToCode(), pPos.Board[86].ToCode(), pPos.Board[76].ToCode(), pPos.Board[66].ToCode(), pPos.Board[56].ToCode(), pPos.Board[46].ToCode(), pPos.Board[36].ToCode(), pPos.Board[26].ToCode(), pPos.Board[16].ToCode(), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[97].ToCode(), pPos.Board[87].ToCode(), pPos.Board[77].ToCode(), pPos.Board[67].ToCode(), pPos.Board[57].ToCode(), pPos.Board[47].ToCode(), pPos.Board[37].ToCode(), pPos.Board[27].ToCode(), pPos.Board[17].ToCode(), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[98].ToCode(), pPos.Board[88].ToCode(), pPos.Board[78].ToCode(), pPos.Board[68].ToCode(), pPos.Board[58].ToCode(), pPos.Board[48].ToCode(), pPos.Board[38].ToCode(), pPos.Board[28].ToCode(), pPos.Board[18].ToCode(), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", pPos.Board[99].ToCode(), pPos.Board[89].ToCode(), pPos.Board[79].ToCode(), pPos.Board[69].ToCode(), pPos.Board[59].ToCode(), pPos.Board[49].ToCode(), pPos.Board[39].ToCode(), pPos.Board[29].ToCode(), pPos.Board[19].ToCode(), zeroFiles[8]) +
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
		"\n" +
		//
		"moves"

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
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
