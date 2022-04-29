package take10

import (
	"bytes"
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Print - 局面出力（＾ｑ＾）
func Sprint(pPos *Position) string {
	var phase_str string
	switch pPos.GetPhase() {
	case l06.FIRST:
		phase_str = "First"
	case l06.SECOND:
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

// CreateMovesList - " 7g7f 3c3d" みたいな部分を返します。最初は半角スペースです
func (pPos *Position) createMovesText() string {
	moves_text := make([]byte, 0, MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for i := 0; i < pPos.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pPos.Moves[i].ToCodeOfM()...)
	}
	return string(moves_text)
}

// SprintControl - 利き数ボード出力（＾ｑ＾）
//
// Parameters
// ----------
// * `layer` - 利き数ボードのレイヤー番号（＾～＾）
func (pPos *Position) SprintControl(phase l06.Phase, layer int) string {
	var board [BOARD_SIZE]int8
	var phase_str string
	var title string

	switch phase {
	case l06.FIRST:
		phase_str = "First"
	case l06.SECOND:
		phase_str = "Second"
	default:
		return "\n"
	}

	var ph = phase - 1
	if ph < 2 { // ph >= 0 &&
		title = fmt.Sprintf("Control(%d)%s", layer, GetControlLayerName(layer))
		board = pPos.ControlBoards[ph][layer]
	}

	return "\n" +
		//
		fmt.Sprintf("[%s %s]\n", title, phase_str) +
		//
		"\n" +
		//
		"  9  8  7  6  5  4  3  2  1\n" +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| a\n", board[91], board[81], board[71], board[61], board[51], board[41], board[31], board[21], board[11]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| b\n", board[92], board[82], board[72], board[62], board[52], board[42], board[32], board[22], board[12]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| c\n", board[93], board[83], board[73], board[63], board[53], board[43], board[33], board[23], board[13]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| d\n", board[94], board[84], board[74], board[64], board[54], board[44], board[34], board[24], board[14]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| e\n", board[95], board[85], board[75], board[65], board[55], board[45], board[35], board[25], board[15]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| f\n", board[96], board[86], board[76], board[66], board[56], board[46], board[36], board[26], board[16]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| g\n", board[97], board[87], board[77], board[67], board[57], board[47], board[37], board[27], board[17]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| h\n", board[98], board[88], board[78], board[68], board[58], board[48], board[38], board[28], board[18]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| i\n", board[99], board[89], board[79], board[69], board[59], board[49], board[39], board[29], board[19]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n"
}

// SprintSfen - SFEN文字列返せよ（＾～＾）
func (pPos *Position) SprintSfen() string {
	// 9x9=81 + 8slash = 89 文字 なんだが成り駒で増えるし めんどくさ（＾～＾）多めに取っとくか（＾～＾）
	// 成り駒２文字なんで、byte型だとめんどくさ（＾～＾）
	buf := make([]byte, 0, 200)

	spaces := 0
	for rank := l03.Square(1); rank < 10; rank += 1 {
		for file := l03.Square(9); file > 0; file -= 1 {
			piece := pPos.Board[SquareFrom(file, rank)]

			if piece != l03.PIECE_EMPTY {
				if spaces > 0 {
					buf = append(buf, OneDigitNumbers[spaces])
					spaces = 0
				}

				pieceString := piece.ToCodeOfPc()
				length := len(pieceString)
				switch length {
				case 2:
					buf = append(buf, pieceString[0])
					buf = append(buf, pieceString[1])
				case 1:
					buf = append(buf, pieceString[0])
				default:
					panic(fmt.Errorf("LogicError: length=%d", length))
				}
			} else {
				// Space
				spaces += 1
			}

		}

		if spaces > 0 {
			buf = append(buf, OneDigitNumbers[spaces])
			spaces = 0
		}

		if rank < 9 {
			buf = append(buf, '/')
		}
	}

	// 手番
	var phaseStr string
	switch pPos.GetPhase() {
	case l06.FIRST:
		phaseStr = "b"
	case l06.SECOND:
		phaseStr = "w"
	default:
		panic(fmt.Errorf("LogicalError: Unknows phase=[%d]", pPos.GetPhase()))
	}

	// 持ち駒
	hands := ""
	num := pPos.Hands[0]
	if num == 1 {
		hands += "R"
	} else if num > 1 {
		hands += fmt.Sprintf("R%d", num)
	}

	num = pPos.Hands[1]
	if num == 1 {
		hands += "B"
	} else if num > 1 {
		hands += fmt.Sprintf("B%d", num)
	}

	num = pPos.Hands[2]
	if num == 1 {
		hands += "G"
	} else if num > 1 {
		hands += fmt.Sprintf("G%d", num)
	}

	num = pPos.Hands[3]
	if num == 1 {
		hands += "S"
	} else if num > 1 {
		hands += fmt.Sprintf("S%d", num)
	}

	num = pPos.Hands[4]
	if num == 1 {
		hands += "N"
	} else if num > 1 {
		hands += fmt.Sprintf("N%d", num)
	}

	num = pPos.Hands[5]
	if num == 1 {
		hands += "L"
	} else if num > 1 {
		hands += fmt.Sprintf("L%d", num)
	}

	num = pPos.Hands[6]
	if num == 1 {
		hands += "P"
	} else if num > 1 {
		hands += fmt.Sprintf("P%d", num)
	}

	num = pPos.Hands[7]
	if num == 1 {
		hands += "r"
	} else if num > 1 {
		hands += fmt.Sprintf("r%d", num)
	}

	num = pPos.Hands[8]
	if num == 1 {
		hands += "b"
	} else if num > 1 {
		hands += fmt.Sprintf("b%d", num)
	}

	num = pPos.Hands[9]
	if num == 1 {
		hands += "g"
	} else if num > 1 {
		hands += fmt.Sprintf("g%d", num)
	}

	num = pPos.Hands[10]
	if num == 1 {
		hands += "s"
	} else if num > 1 {
		hands += fmt.Sprintf("s%d", num)
	}

	num = pPos.Hands[11]
	if num == 1 {
		hands += "n"
	} else if num > 1 {
		hands += fmt.Sprintf("n%d", num)
	}

	num = pPos.Hands[12]
	if num == 1 {
		hands += "l"
	} else if num > 1 {
		hands += fmt.Sprintf("l%d", num)
	}

	num = pPos.Hands[13]
	if num == 1 {
		hands += "p"
	} else if num > 1 {
		hands += fmt.Sprintf("p%d", num)
	}

	if hands == "" {
		hands = "-"
	}

	// 手数
	movesNum := pPos.StartMovesNum + pPos.OffsetMovesIndex

	// 指し手
	moves_text := pPos.createMovesText()

	return fmt.Sprintf("position sfen %s %s %s %d moves%s\n", buf, phaseStr, hands, movesNum, moves_text)
}

// SprintRecord - 棋譜表示（＾～＾）
func (pPos *Position) SprintRecord() string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, MOVES_SIZE*9)
	for i := 0; i < pPos.OffsetMovesIndex; i += 1 {
		record_text = append(record_text, pPos.Moves[i].ToCodeOfM()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, pPos.CapturedList[i].ToCodeOfPc()...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("Record\n------\n%s", record_text)
}

// Dump - 内部状態を全部出力しようぜ（＾～＾）？
func (pPos *Position) Dump() string {
	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buffer bytes.Buffer

	buffer.WriteString("Board:")
	for i := 0; i < BOARD_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.Board[i]))
	}
	buffer.WriteString("\n")

	king1 := pPos.PieceLocations[PCLOC_K1]
	king2 := pPos.PieceLocations[PCLOC_K2]
	buffer.WriteString(fmt.Sprintf("KingLocations:%d,%d,\n", king1, king2))

	buffer.WriteString("BishopLocations:")
	for i := 0; i < 2; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.PieceLocations[PCLOC_B1:PCLOC_B2][i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("LanceLocations:")
	for i := 0; i < 2; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.PieceLocations[PCLOC_L1:PCLOC_L4][i]))
	}
	buffer.WriteString("\n")

	for phase := 0; phase < 2; phase += 1 {
		// 利きボード
		for layer := 0; layer < CONTROL_LAYER_ALL_SIZE; layer += 1 {
			buffer.WriteString(pPos.SprintControl(l06.Phase(phase+1), layer))
		}
	}

	buffer.WriteString("Hands:")
	for i := 0; i < 14; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.Hands[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString(fmt.Sprintf("Phase:%d,\n", pPos.GetPhase()))

	buffer.WriteString(fmt.Sprintf("StartMovesNum:%d,\n", pPos.StartMovesNum))

	buffer.WriteString(fmt.Sprintf("OffsetMovesIndex:%d,\n", pPos.OffsetMovesIndex))

	buffer.WriteString("Moves:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.Moves[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("CapturedList:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPos.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
