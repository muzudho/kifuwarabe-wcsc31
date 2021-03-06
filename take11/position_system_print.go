package take11

import (
	"bytes"
	"fmt"

	l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l07 "github.com/muzudho/kifuwarabe-wcsc31/take7"
)

// Print - ２局面の比較用画面出力（＾ｑ＾）
func (pPosSys *PositionSystem) SprintDiff(b1 PosLayerT, b2 PosLayerT) string {
	var phase_str string
	switch pPosSys.GetPhase() {
	case l03.FIRST:
		phase_str = "First"
	case l03.SECOND:
		phase_str = "Second"
	default:
		phase_str = "?"
	}

	// 0段目
	zeroRanks := [10]string{"    9", "    8", "    7", "    6", "    5", "    4", "    3", "    2", "    1", "     "}
	// 0筋目
	zeroFiles := [9]string{" a ", " b ", " c ", " d ", " e ", " f ", " g ", " h ", " i "}

	// 0段目、0筋目に駒置いてたらそれも表示（＾～＾）
	for file := 9; file > -1; file -= 1 {
		if !pPosSys.PPosition[b1].IsEmptySq(l03.Square(file*10)) || !pPosSys.PPosition[b2].IsEmptySq(l03.Square(file*10)) {
			zeroRanks[10-file] = fmt.Sprintf("%2s%2s", pPosSys.PPosition[b1].Board[file*10].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10].ToCodeOfPc())
		}
	}

	// 0筋目
	for rank := l03.Square(1); rank < 10; rank += 1 {
		if !pPosSys.PPosition[b1].IsEmptySq(rank) || !pPosSys.PPosition[b2].IsEmptySq(rank) {
			zeroFiles[rank-1] = fmt.Sprintf("%2s%2s", pPosSys.PPosition[b1].Board[rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[rank].ToCodeOfPc())
		}
	}

	lines := []string{}
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPosSys.StartMovesNum, (pPosSys.StartMovesNum+pPosSys.OffsetMovesIndex), phase_str))
	lines = append(lines, "\n")
	lines = append(lines, "    k    r    b    g    s    n    l    p\n")
	lines = append(lines, "+----+----+----+----+----+----+----+----+\n")

	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buf bytes.Buffer
	for i := l03.HAND_TYPE_SIZE; i < l03.HAND_IDX_END; i++ {
		buf.WriteString(fmt.Sprintf("|%2d%2d", pPosSys.PPosition[b1].Hands1[i], pPosSys.PPosition[b2].Hands1[i]))
	}
	buf.WriteString("|\n")
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	for i := 0; i < 10; i += 1 {
		buf.WriteString(zeroRanks[i])
	}
	buf.WriteString("\n")
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank := 1
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 2
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 3
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 4
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 5
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 6
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 7
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 8
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 9
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCodeOfPc(), pPosSys.PPosition[b2].Board[file*10+rank].ToCodeOfPc()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "     K    R    B    G    S    N    L    P\n")
	lines = append(lines, " +----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	buf.WriteString(" ")
	for i := l03.HAND_IDX_BEGIN; i < l03.HAND_TYPE_SIZE; i++ {
		buf.WriteString(fmt.Sprintf("|%2d%2d", pPosSys.PPosition[b1].Hands1[i], pPosSys.PPosition[b2].Hands1[i]))
	}
	buf.WriteString("|\n")
	lines = append(lines, buf.String())

	lines = append(lines, " +----+----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "moves")

	lines = append(lines, pPosSys.createMovesText())
	lines = append(lines, "\n")

	buf.Reset()
	for _, line := range lines {
		buf.WriteString(line)
	}
	return buf.String()
}

// SprintControl - 利き数ボード出力（＾ｑ＾）
//
// Parameters
// ----------
// * `c` - 利き数ボードのレイヤー番号（＾～＾）
func (pPosSys *PositionSystem) SprintControl(phase l03.Phase, c ControlLayerT) string {
	var board [l03.BOARD_SIZE]int8
	var phase_str string
	var title string

	switch phase {
	case l03.FIRST:
		phase_str = "First"
	case l03.SECOND:
		phase_str = "Second"
	default:
		return "\n"
	}

	var ph = phase - 1
	if ph < 2 { // 0 <= ph &&
		title = fmt.Sprintf("Control(%d)%s", c, GetControlLayerName(c))
		board = pPosSys.ControlBoards[ph][c]
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
func (pPosSys *PositionSystem) SprintSfen(pPos *Position) string {
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
					panic(App.LogNotEcho.Fatal("LogicError: length=%d", length))
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
	switch pPosSys.GetPhase() {
	case l03.FIRST:
		phaseStr = "b"
	case l03.SECOND:
		phaseStr = "w"
	default:
		panic(App.LogNotEcho.Fatal("LogicalError: Unknows phase=[%d]", pPosSys.GetPhase()))
	}

	// 持ち駒
	hands := ""

	// 玉は出力できません
	// num := pPos.Hands1[l03.HANDSQ_K1]
	// if num == 1 {
	// 	hands += "K"
	// } else if num > 1 {
	// 	hands += fmt.Sprintf("K%d", num)
	// }

	num := pPos.Hands1[l03.HAND_R1]
	if num == 1 {
		hands += "R"
	} else if num > 1 {
		hands += fmt.Sprintf("R%d", num)
	}

	num = pPos.Hands1[l03.HAND_B1]
	if num == 1 {
		hands += "B"
	} else if num > 1 {
		hands += fmt.Sprintf("B%d", num)
	}

	num = pPos.Hands1[l03.HAND_G1]
	if num == 1 {
		hands += "G"
	} else if num > 1 {
		hands += fmt.Sprintf("G%d", num)
	}

	num = pPos.Hands1[l03.HAND_S1]
	if num == 1 {
		hands += "S"
	} else if num > 1 {
		hands += fmt.Sprintf("S%d", num)
	}

	num = pPos.Hands1[l03.HAND_N1]
	if num == 1 {
		hands += "N"
	} else if num > 1 {
		hands += fmt.Sprintf("N%d", num)
	}

	num = pPos.Hands1[l03.HAND_L1]
	if num == 1 {
		hands += "L"
	} else if num > 1 {
		hands += fmt.Sprintf("L%d", num)
	}

	num = pPos.Hands1[l03.HAND_P1]
	if num == 1 {
		hands += "P"
	} else if num > 1 {
		hands += fmt.Sprintf("P%d", num)
	}

	// 玉は出力できません
	// num := pPos.Hands1[l03.HANDSQ_K2]
	// if num == 1 {
	// 	hands += "k"
	// } else if num > 1 {
	// 	hands += fmt.Sprintf("k%d", num)
	// }

	num = pPos.Hands1[l03.HAND_R2]
	if num == 1 {
		hands += "r"
	} else if num > 1 {
		hands += fmt.Sprintf("r%d", num)
	}

	num = pPos.Hands1[l03.HAND_B2]
	if num == 1 {
		hands += "b"
	} else if num > 1 {
		hands += fmt.Sprintf("b%d", num)
	}

	num = pPos.Hands1[l03.HAND_G2]
	if num == 1 {
		hands += "g"
	} else if num > 1 {
		hands += fmt.Sprintf("g%d", num)
	}

	num = pPos.Hands1[l03.HAND_S2]
	if num == 1 {
		hands += "s"
	} else if num > 1 {
		hands += fmt.Sprintf("s%d", num)
	}

	num = pPos.Hands1[l03.HAND_N2]
	if num == 1 {
		hands += "n"
	} else if num > 1 {
		hands += fmt.Sprintf("n%d", num)
	}

	num = pPos.Hands1[l03.HAND_L2]
	if num == 1 {
		hands += "l"
	} else if num > 1 {
		hands += fmt.Sprintf("l%d", num)
	}

	num = pPos.Hands1[l03.HAND_P2]
	if num == 1 {
		hands += "p"
	} else if num > 1 {
		hands += fmt.Sprintf("p%d", num)
	}

	if hands == "" {
		hands = "-"
	}

	// 手数
	movesNum := pPosSys.StartMovesNum + pPosSys.OffsetMovesIndex

	// 指し手
	moves_text := pPosSys.createMovesText()

	return fmt.Sprintf("position sfen %s %s %s %d moves%s\n", buf, phaseStr, hands, movesNum, moves_text)
}

// SprintRecord - 棋譜表示（＾～＾）
func (pPosSys *PositionSystem) SprintRecord() string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, l02.MOVES_SIZE*9)
	for i := 0; i < pPosSys.OffsetMovesIndex; i += 1 {
		record_text = append(record_text, pPosSys.Moves[i].ToCodeOfM()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, pPosSys.CapturedList[i].ToCodeOfPc()...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("Record\n------\n%s", record_text)
}

// Dump - 内部状態を全部出力しようぜ（＾～＾）？
func (pPosSys *PositionSystem) Dump() string {
	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buffer bytes.Buffer

	for b := PosLayerT(0); b < 2; b += 1 {
		pPos := pPosSys.PPosition[b]
		buffer.WriteString(fmt.Sprintf("Position[%d]:", b))
		for i := 0; i < POS_LAYER_SIZE; i += 1 {
			buffer.WriteString(fmt.Sprintf("%d,", pPosSys.PPosition[i].Board))
		}
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("KingLocations[%d]:%d,%d\n", b, pPos.PieceLocations[l07.PCLOC_K1], pPos.PieceLocations[l07.PCLOC_K2]))
		buffer.WriteString(fmt.Sprintf("RookLocations[%d]:%d,%d\n", b, pPos.PieceLocations[l07.PCLOC_R1], pPos.PieceLocations[l07.PCLOC_R2]))
		buffer.WriteString(fmt.Sprintf("BishopLocations[%d]:%d,%d\n", b, pPos.PieceLocations[l07.PCLOC_B1], pPos.PieceLocations[l07.PCLOC_B2]))
		buffer.WriteString(fmt.Sprintf("LanceLocations[%d]:%d,%d,%d,%d\n", b, pPos.PieceLocations[l07.PCLOC_L1], pPos.PieceLocations[l07.PCLOC_L2], pPos.PieceLocations[l07.PCLOC_L3], pPos.PieceLocations[l07.PCLOC_L4]))
	}

	for phase := 0; phase < 2; phase += 1 {
		// 利きボード
		for c := ControlLayerT(0); c < CONTROL_LAYER_ALL_SIZE; c += 1 {
			buffer.WriteString(pPosSys.SprintControl(l03.Phase(phase+1), c))
		}
	}

	for b := PosLayerT(0); b < 2; b += 1 {
		buffer.WriteString(fmt.Sprintf("Position[%d]:", b))
		buffer.WriteString("Hands:")
		for i := l03.HAND_IDX_BEGIN; i < l03.HAND_IDX_END; i += 1 {
			buffer.WriteString(fmt.Sprintf("%d,", pPosSys.PPosition[b].Hands1[i]))
		}
		buffer.WriteString("\n")
	}

	buffer.WriteString(fmt.Sprintf("Phase:%d,\n", pPosSys.GetPhase()))

	buffer.WriteString(fmt.Sprintf("StartMovesNum:%d,\n", pPosSys.StartMovesNum))

	buffer.WriteString(fmt.Sprintf("OffsetMovesIndex:%d,\n", pPosSys.OffsetMovesIndex))

	buffer.WriteString("Moves:")
	for i := 0; i < l02.MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPosSys.Moves[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("CapturedList:")
	for i := 0; i < l02.MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pPosSys.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
