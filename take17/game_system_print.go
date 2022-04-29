package take17

import (
	"bytes"
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// Print - ２局面の比較用画面出力（＾ｑ＾）
func sprintPositionDiff(pPosSys *PositionSystem, b1 PosLayerT, b2 PosLayerT, pRecord *DifferenceRecord) string {
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
	lines = append(lines, fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pRecord.StartMovesNum, (pRecord.StartMovesNum+pRecord.OffsetMovesIndex), phase_str))
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
	for i := l03.HAND_IDX_START; i < l03.HAND_TYPE_SIZE; i++ {
		buf.WriteString(fmt.Sprintf("|%2d%2d", pPosSys.PPosition[b1].Hands1[i], pPosSys.PPosition[b2].Hands1[i]))
	}
	buf.WriteString("|\n")
	lines = append(lines, buf.String())

	lines = append(lines, " +----+----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "moves")

	lines = append(lines, createMovesText(pRecord))
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("KomawariValue: %d %d\n", pPosSys.PPosition[b1].MaterialValue, pPosSys.PPosition[b2].MaterialValue))

	buf.Reset()
	for _, line := range lines {
		buf.WriteString(line)
	}
	return buf.String()
}

// SprintSfen - SFEN文字列返せよ（＾～＾）投了図を返すぜ（＾～＾）棋譜の部分を捨てるぜ（＾～＾）
func sprintSfenResignation(pPosSys *PositionSystem, pPos *l15.Position, pRecord *DifferenceRecord) string {
	// 9x9=81 + 8slash = 89 文字 なんだが成り駒で増えるし めんどくさ（＾～＾）多めに取っとくか（＾～＾）
	// 成り駒２文字なんで、byte型だとめんどくさ（＾～＾）
	buf := make([]byte, 0, 200)

	spaces := 0
	for rank := l03.Square(1); rank < 10; rank += 1 {
		for file := l03.Square(9); file > 0; file -= 1 {
			piece := pPos.Board[l15.SquareFrom(file, rank)]

			if piece != l03.PIECE_EMPTY {
				if spaces > 0 {
					buf = append(buf, oneDigitNumbers[spaces])
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
			buf = append(buf, oneDigitNumbers[spaces])
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
	// num := pPos.Hands1[l03.HAND_K1]
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
	// num := pPos.Hands1[l03.HAND_K2]
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
	movesNum := pRecord.StartMovesNum + pRecord.OffsetMovesIndex

	// 指し手
	// moves_text := pPosSys.createMovesText()

	// return fmt.Sprintf("position sfen %s %s %s %d moves%s\n", buf, phaseStr, hands, movesNum, moves_text)
	return fmt.Sprintf("position sfen %s %s %s %d\n", buf, phaseStr, hands, movesNum)
}

// sprintRecord - 棋譜表示（＾～＾）
func sprintRecord(pRecord *DifferenceRecord) string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, l04.MOVES_SIZE*9)
	for i := 0; i < pRecord.OffsetMovesIndex; i += 1 {
		record_text = append(record_text, pRecord.Moves[i].ToCodeOfM()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, pRecord.CapturedList[i].ToCodeOfPc()...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("Record\n------\n%s", record_text)
}
