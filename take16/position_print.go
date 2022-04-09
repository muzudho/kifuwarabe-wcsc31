package take16

import (
	"fmt"

	l12 "github.com/muzudho/kifuwarabe-wcsc31/take12"
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
		zeroRanks[0] = l12.ToCodeOfPc(pPos.Board[90])
	}
	if !pPos.IsEmptySq(80) {
		zeroRanks[1] = l12.ToCodeOfPc(pPos.Board[80])
	}
	if !pPos.IsEmptySq(70) {
		zeroRanks[2] = l12.ToCodeOfPc(pPos.Board[70])
	}
	if !pPos.IsEmptySq(60) {
		zeroRanks[3] = l12.ToCodeOfPc(pPos.Board[60])
	}
	if !pPos.IsEmptySq(50) {
		zeroRanks[4] = l12.ToCodeOfPc(pPos.Board[50])
	}
	if !pPos.IsEmptySq(40) {
		zeroRanks[5] = l12.ToCodeOfPc(pPos.Board[40])
	}
	if !pPos.IsEmptySq(30) {
		zeroRanks[6] = l12.ToCodeOfPc(pPos.Board[30])
	}
	if !pPos.IsEmptySq(20) {
		zeroRanks[7] = l12.ToCodeOfPc(pPos.Board[20])
	}
	if !pPos.IsEmptySq(10) {
		zeroRanks[8] = l12.ToCodeOfPc(pPos.Board[10])
	}
	if !pPos.IsEmptySq(0) {
		zeroRanks[9] = l12.ToCodeOfPc(pPos.Board[0])
	}
	if !pPos.IsEmptySq(1) {
		zeroFiles[0] = l12.ToCodeOfPc(pPos.Board[1])
	}
	if !pPos.IsEmptySq(2) {
		zeroFiles[1] = l12.ToCodeOfPc(pPos.Board[2])
	}
	if !pPos.IsEmptySq(3) {
		zeroFiles[2] = l12.ToCodeOfPc(pPos.Board[3])
	}
	if !pPos.IsEmptySq(4) {
		zeroFiles[3] = l12.ToCodeOfPc(pPos.Board[4])
	}
	if !pPos.IsEmptySq(5) {
		zeroFiles[4] = l12.ToCodeOfPc(pPos.Board[5])
	}
	if !pPos.IsEmptySq(6) {
		zeroFiles[5] = l12.ToCodeOfPc(pPos.Board[6])
	}
	if !pPos.IsEmptySq(7) {
		zeroFiles[6] = l12.ToCodeOfPc(pPos.Board[7])
	}
	if !pPos.IsEmptySq(8) {
		zeroFiles[7] = l12.ToCodeOfPc(pPos.Board[8])
	}
	if !pPos.IsEmptySq(9) {
		zeroFiles[8] = l12.ToCodeOfPc(pPos.Board[9])
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
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[91]), l12.ToCodeOfPc(pPos.Board[81]), l12.ToCodeOfPc(pPos.Board[71]), l12.ToCodeOfPc(pPos.Board[61]), l12.ToCodeOfPc(pPos.Board[51]), l12.ToCodeOfPc(pPos.Board[41]), l12.ToCodeOfPc(pPos.Board[31]), l12.ToCodeOfPc(pPos.Board[21]), l12.ToCodeOfPc(pPos.Board[11]), zeroFiles[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[92]), l12.ToCodeOfPc(pPos.Board[82]), l12.ToCodeOfPc(pPos.Board[72]), l12.ToCodeOfPc(pPos.Board[62]), l12.ToCodeOfPc(pPos.Board[52]), l12.ToCodeOfPc(pPos.Board[42]), l12.ToCodeOfPc(pPos.Board[32]), l12.ToCodeOfPc(pPos.Board[22]), l12.ToCodeOfPc(pPos.Board[12]), zeroFiles[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[93]), l12.ToCodeOfPc(pPos.Board[83]), l12.ToCodeOfPc(pPos.Board[73]), l12.ToCodeOfPc(pPos.Board[63]), l12.ToCodeOfPc(pPos.Board[53]), l12.ToCodeOfPc(pPos.Board[43]), l12.ToCodeOfPc(pPos.Board[33]), l12.ToCodeOfPc(pPos.Board[23]), l12.ToCodeOfPc(pPos.Board[13]), zeroFiles[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[94]), l12.ToCodeOfPc(pPos.Board[84]), l12.ToCodeOfPc(pPos.Board[74]), l12.ToCodeOfPc(pPos.Board[64]), l12.ToCodeOfPc(pPos.Board[54]), l12.ToCodeOfPc(pPos.Board[44]), l12.ToCodeOfPc(pPos.Board[34]), l12.ToCodeOfPc(pPos.Board[24]), l12.ToCodeOfPc(pPos.Board[14]), zeroFiles[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[95]), l12.ToCodeOfPc(pPos.Board[85]), l12.ToCodeOfPc(pPos.Board[75]), l12.ToCodeOfPc(pPos.Board[65]), l12.ToCodeOfPc(pPos.Board[55]), l12.ToCodeOfPc(pPos.Board[45]), l12.ToCodeOfPc(pPos.Board[35]), l12.ToCodeOfPc(pPos.Board[25]), l12.ToCodeOfPc(pPos.Board[15]), zeroFiles[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[96]), l12.ToCodeOfPc(pPos.Board[86]), l12.ToCodeOfPc(pPos.Board[76]), l12.ToCodeOfPc(pPos.Board[66]), l12.ToCodeOfPc(pPos.Board[56]), l12.ToCodeOfPc(pPos.Board[46]), l12.ToCodeOfPc(pPos.Board[36]), l12.ToCodeOfPc(pPos.Board[26]), l12.ToCodeOfPc(pPos.Board[16]), zeroFiles[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[97]), l12.ToCodeOfPc(pPos.Board[87]), l12.ToCodeOfPc(pPos.Board[77]), l12.ToCodeOfPc(pPos.Board[67]), l12.ToCodeOfPc(pPos.Board[57]), l12.ToCodeOfPc(pPos.Board[47]), l12.ToCodeOfPc(pPos.Board[37]), l12.ToCodeOfPc(pPos.Board[27]), l12.ToCodeOfPc(pPos.Board[17]), zeroFiles[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[98]), l12.ToCodeOfPc(pPos.Board[88]), l12.ToCodeOfPc(pPos.Board[78]), l12.ToCodeOfPc(pPos.Board[68]), l12.ToCodeOfPc(pPos.Board[58]), l12.ToCodeOfPc(pPos.Board[48]), l12.ToCodeOfPc(pPos.Board[38]), l12.ToCodeOfPc(pPos.Board[28]), l12.ToCodeOfPc(pPos.Board[18]), zeroFiles[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%3s\n", l12.ToCodeOfPc(pPos.Board[99]), l12.ToCodeOfPc(pPos.Board[89]), l12.ToCodeOfPc(pPos.Board[79]), l12.ToCodeOfPc(pPos.Board[69]), l12.ToCodeOfPc(pPos.Board[59]), l12.ToCodeOfPc(pPos.Board[49]), l12.ToCodeOfPc(pPos.Board[39]), l12.ToCodeOfPc(pPos.Board[29]), l12.ToCodeOfPc(pPos.Board[19]), zeroFiles[8]) +
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
