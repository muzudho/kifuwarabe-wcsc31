package take16

import (
	"bytes"
	"fmt"

	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// Print - ２局面の比較用画面出力（＾ｑ＾）
func (pPosSys *PositionSystem) SprintDiff(pRecord *DifferenceRecord, b1 PosLayerT, b2 PosLayerT) string {
	var phase_str string
	switch pPosSys.GetPhase() {
	case p.FIRST:
		phase_str = "First"
	case p.SECOND:
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
		if !pPosSys.PPosition[b1].IsEmptySq(p.Square(file*10)) || !pPosSys.PPosition[b2].IsEmptySq(p.Square(file*10)) {
			zeroRanks[10-file] = fmt.Sprintf("%2s%2s", pPosSys.PPosition[b1].Board[file*10].ToCode(), pPosSys.PPosition[b2].Board[file*10].ToCode())
		}
	}

	// 0筋目
	for rank := p.Square(1); rank < 10; rank += 1 {
		if !pPosSys.PPosition[b1].IsEmptySq(rank) || !pPosSys.PPosition[b2].IsEmptySq(rank) {
			zeroFiles[rank-1] = fmt.Sprintf("%2s%2s", pPosSys.PPosition[b1].Board[rank].ToCode(), pPosSys.PPosition[b2].Board[rank].ToCode())
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
	for i := p.HAND_TYPE_SIZE; i < p.HAND_IDX_END; i++ {
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
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 2
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 3
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 4
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 5
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 6
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 7
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 8
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	rank = 9
	for file := 9; file > 0; file-- {
		buf.WriteString(fmt.Sprintf("|%2s%2s", pPosSys.PPosition[b1].Board[file*10+rank].ToCode(), pPosSys.PPosition[b2].Board[file*10+rank].ToCode()))
	}
	buf.WriteString(fmt.Sprintf("|%s\n", zeroFiles[rank-1]))
	lines = append(lines, buf.String())

	lines = append(lines, "+----+----+----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "     K    R    B    G    S    N    L    P\n")
	lines = append(lines, " +----+----+----+----+----+----+----+----+\n")

	buf.Reset()
	buf.WriteString(" ")
	for i := p.HAND_IDX_START; i < p.HAND_TYPE_SIZE; i++ {
		buf.WriteString(fmt.Sprintf("|%2d%2d", pPosSys.PPosition[b1].Hands1[i], pPosSys.PPosition[b2].Hands1[i]))
	}
	buf.WriteString("|\n")
	lines = append(lines, buf.String())

	lines = append(lines, " +----+----+----+----+----+----+----+----+\n")
	lines = append(lines, "\n")
	lines = append(lines, "moves")

	lines = append(lines, pPosSys.createMovesText(pRecord))
	lines = append(lines, "\n")
	lines = append(lines, fmt.Sprintf("KomawariValue: %d %d\n", pPosSys.PPosition[b1].MaterialValue, pPosSys.PPosition[b2].MaterialValue))

	buf.Reset()
	for _, line := range lines {
		buf.WriteString(line)
	}
	return buf.String()
}

// CreateMovesList - " 7g7f 3c3d" みたいな部分を返します。最初は半角スペースです
func (pPosSys *PositionSystem) createMovesText(pRecord *DifferenceRecord) string {
	moves_text := make([]byte, 0, pRecord.OffsetMovesIndex*6) // スペース含めて１手最大6文字（＾～＾）
	for i := 0; i < pRecord.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pRecord.Moves[i].ToCode()...)
	}
	return string(moves_text)
}

// position sfen の盤のスペース数に使われますN
var oneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// SprintSfen - SFEN文字列返せよ（＾～＾）投了図を返すぜ（＾～＾）棋譜の部分を捨てるぜ（＾～＾）
func (pPosSys *PositionSystem) SprintSfenResignation(pRecord *DifferenceRecord, pPos *p.Position) string {
	// 9x9=81 + 8slash = 89 文字 なんだが成り駒で増えるし めんどくさ（＾～＾）多めに取っとくか（＾～＾）
	// 成り駒２文字なんで、byte型だとめんどくさ（＾～＾）
	buf := make([]byte, 0, 200)

	spaces := 0
	for rank := p.Square(1); rank < 10; rank += 1 {
		for file := p.Square(9); file > 0; file -= 1 {
			piece := pPos.Board[p.SquareFrom(file, rank)]

			if piece != p.PIECE_EMPTY {
				if spaces > 0 {
					buf = append(buf, oneDigitNumbers[spaces])
					spaces = 0
				}

				pieceString := piece.ToCode()
				length := len(pieceString)
				switch length {
				case 2:
					buf = append(buf, pieceString[0])
					buf = append(buf, pieceString[1])
				case 1:
					buf = append(buf, pieceString[0])
				default:
					panic(G.Log.Fatal("LogicError: length=%d", length))
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
	case p.FIRST:
		phaseStr = "b"
	case p.SECOND:
		phaseStr = "w"
	default:
		panic(G.Log.Fatal("LogicalError: Unknows phase=[%d]", pPosSys.GetPhase()))
	}

	// 持ち駒
	hands := ""

	// 玉は出力できません
	// num := pPos.Hands1[HAND_K1_IDX]
	// if num == 1 {
	// 	hands += "K"
	// } else if num > 1 {
	// 	hands += fmt.Sprintf("K%d", num)
	// }

	num := pPos.Hands1[p.HAND_R1_IDX]
	if num == 1 {
		hands += "R"
	} else if num > 1 {
		hands += fmt.Sprintf("R%d", num)
	}

	num = pPos.Hands1[p.HAND_B1_IDX]
	if num == 1 {
		hands += "B"
	} else if num > 1 {
		hands += fmt.Sprintf("B%d", num)
	}

	num = pPos.Hands1[p.HAND_G1_IDX]
	if num == 1 {
		hands += "G"
	} else if num > 1 {
		hands += fmt.Sprintf("G%d", num)
	}

	num = pPos.Hands1[p.HAND_S1_IDX]
	if num == 1 {
		hands += "S"
	} else if num > 1 {
		hands += fmt.Sprintf("S%d", num)
	}

	num = pPos.Hands1[p.HAND_N1_IDX]
	if num == 1 {
		hands += "N"
	} else if num > 1 {
		hands += fmt.Sprintf("N%d", num)
	}

	num = pPos.Hands1[p.HAND_L1_IDX]
	if num == 1 {
		hands += "L"
	} else if num > 1 {
		hands += fmt.Sprintf("L%d", num)
	}

	num = pPos.Hands1[p.HAND_P1_IDX]
	if num == 1 {
		hands += "P"
	} else if num > 1 {
		hands += fmt.Sprintf("P%d", num)
	}

	// 玉は出力できません
	// num := pPos.Hands1[HAND_K2_IDX]
	// if num == 1 {
	// 	hands += "k"
	// } else if num > 1 {
	// 	hands += fmt.Sprintf("k%d", num)
	// }

	num = pPos.Hands1[p.HAND_R2_IDX]
	if num == 1 {
		hands += "r"
	} else if num > 1 {
		hands += fmt.Sprintf("r%d", num)
	}

	num = pPos.Hands1[p.HAND_B2_IDX]
	if num == 1 {
		hands += "b"
	} else if num > 1 {
		hands += fmt.Sprintf("b%d", num)
	}

	num = pPos.Hands1[p.HAND_G2_IDX]
	if num == 1 {
		hands += "g"
	} else if num > 1 {
		hands += fmt.Sprintf("g%d", num)
	}

	num = pPos.Hands1[p.HAND_S2_IDX]
	if num == 1 {
		hands += "s"
	} else if num > 1 {
		hands += fmt.Sprintf("s%d", num)
	}

	num = pPos.Hands1[p.HAND_N2_IDX]
	if num == 1 {
		hands += "n"
	} else if num > 1 {
		hands += fmt.Sprintf("n%d", num)
	}

	num = pPos.Hands1[p.HAND_L2_IDX]
	if num == 1 {
		hands += "l"
	} else if num > 1 {
		hands += fmt.Sprintf("l%d", num)
	}

	num = pPos.Hands1[p.HAND_P2_IDX]
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

// SprintRecord - 棋譜表示（＾～＾）
func (pPosSys *PositionSystem) SprintRecord(pRecord *DifferenceRecord) string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, MOVES_SIZE*9)
	for i := 0; i < pRecord.OffsetMovesIndex; i += 1 {
		record_text = append(record_text, pRecord.Moves[i].ToCode()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, pRecord.CapturedList[i].ToCode()...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("Record\n------\n%s", record_text)
}
