package take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Print - 局面出力（＾ｑ＾）
func Sprint(pPos *Position) string {
	var phase_str = "?"
	if pPos.Phase == l06.FIRST {
		phase_str = "First"
	} else if pPos.Phase == l06.SECOND {
		phase_str = "Second"
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
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pPos.GetPieceAtIndex(90).ToCodeOfPc(), pPos.GetPieceAtIndex(80).ToCodeOfPc(), pPos.GetPieceAtIndex(70).ToCodeOfPc(), pPos.GetPieceAtIndex(60).ToCodeOfPc(), pPos.GetPieceAtIndex(50).ToCodeOfPc(), pPos.GetPieceAtIndex(40).ToCodeOfPc(), pPos.GetPieceAtIndex(30).ToCodeOfPc(), pPos.GetPieceAtIndex(20).ToCodeOfPc(), pPos.GetPieceAtIndex(10).ToCodeOfPc(), pPos.GetPieceAtIndex(0).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(91).ToCodeOfPc(), pPos.GetPieceAtIndex(81).ToCodeOfPc(), pPos.GetPieceAtIndex(71).ToCodeOfPc(), pPos.GetPieceAtIndex(61).ToCodeOfPc(), pPos.GetPieceAtIndex(51).ToCodeOfPc(), pPos.GetPieceAtIndex(41).ToCodeOfPc(), pPos.GetPieceAtIndex(31).ToCodeOfPc(), pPos.GetPieceAtIndex(21).ToCodeOfPc(), pPos.GetPieceAtIndex(11).ToCodeOfPc(), pPos.GetPieceAtIndex(1).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(92).ToCodeOfPc(), pPos.GetPieceAtIndex(82).ToCodeOfPc(), pPos.GetPieceAtIndex(72).ToCodeOfPc(), pPos.GetPieceAtIndex(62).ToCodeOfPc(), pPos.GetPieceAtIndex(52).ToCodeOfPc(), pPos.GetPieceAtIndex(42).ToCodeOfPc(), pPos.GetPieceAtIndex(32).ToCodeOfPc(), pPos.GetPieceAtIndex(22).ToCodeOfPc(), pPos.GetPieceAtIndex(12).ToCodeOfPc(), pPos.GetPieceAtIndex(2).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(93).ToCodeOfPc(), pPos.GetPieceAtIndex(83).ToCodeOfPc(), pPos.GetPieceAtIndex(73).ToCodeOfPc(), pPos.GetPieceAtIndex(63).ToCodeOfPc(), pPos.GetPieceAtIndex(53).ToCodeOfPc(), pPos.GetPieceAtIndex(43).ToCodeOfPc(), pPos.GetPieceAtIndex(33).ToCodeOfPc(), pPos.GetPieceAtIndex(23).ToCodeOfPc(), pPos.GetPieceAtIndex(13).ToCodeOfPc(), pPos.GetPieceAtIndex(3).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(94).ToCodeOfPc(), pPos.GetPieceAtIndex(84).ToCodeOfPc(), pPos.GetPieceAtIndex(74).ToCodeOfPc(), pPos.GetPieceAtIndex(64).ToCodeOfPc(), pPos.GetPieceAtIndex(54).ToCodeOfPc(), pPos.GetPieceAtIndex(44).ToCodeOfPc(), pPos.GetPieceAtIndex(34).ToCodeOfPc(), pPos.GetPieceAtIndex(24).ToCodeOfPc(), pPos.GetPieceAtIndex(14).ToCodeOfPc(), pPos.GetPieceAtIndex(4).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(95).ToCodeOfPc(), pPos.GetPieceAtIndex(85).ToCodeOfPc(), pPos.GetPieceAtIndex(75).ToCodeOfPc(), pPos.GetPieceAtIndex(65).ToCodeOfPc(), pPos.GetPieceAtIndex(55).ToCodeOfPc(), pPos.GetPieceAtIndex(45).ToCodeOfPc(), pPos.GetPieceAtIndex(35).ToCodeOfPc(), pPos.GetPieceAtIndex(25).ToCodeOfPc(), pPos.GetPieceAtIndex(15).ToCodeOfPc(), pPos.GetPieceAtIndex(5).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(96).ToCodeOfPc(), pPos.GetPieceAtIndex(86).ToCodeOfPc(), pPos.GetPieceAtIndex(76).ToCodeOfPc(), pPos.GetPieceAtIndex(66).ToCodeOfPc(), pPos.GetPieceAtIndex(56).ToCodeOfPc(), pPos.GetPieceAtIndex(46).ToCodeOfPc(), pPos.GetPieceAtIndex(36).ToCodeOfPc(), pPos.GetPieceAtIndex(26).ToCodeOfPc(), pPos.GetPieceAtIndex(16).ToCodeOfPc(), pPos.GetPieceAtIndex(6).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(97).ToCodeOfPc(), pPos.GetPieceAtIndex(87).ToCodeOfPc(), pPos.GetPieceAtIndex(77).ToCodeOfPc(), pPos.GetPieceAtIndex(67).ToCodeOfPc(), pPos.GetPieceAtIndex(57).ToCodeOfPc(), pPos.GetPieceAtIndex(47).ToCodeOfPc(), pPos.GetPieceAtIndex(37).ToCodeOfPc(), pPos.GetPieceAtIndex(27).ToCodeOfPc(), pPos.GetPieceAtIndex(17).ToCodeOfPc(), pPos.GetPieceAtIndex(7).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(98).ToCodeOfPc(), pPos.GetPieceAtIndex(88).ToCodeOfPc(), pPos.GetPieceAtIndex(78).ToCodeOfPc(), pPos.GetPieceAtIndex(68).ToCodeOfPc(), pPos.GetPieceAtIndex(58).ToCodeOfPc(), pPos.GetPieceAtIndex(48).ToCodeOfPc(), pPos.GetPieceAtIndex(38).ToCodeOfPc(), pPos.GetPieceAtIndex(28).ToCodeOfPc(), pPos.GetPieceAtIndex(18).ToCodeOfPc(), pPos.GetPieceAtIndex(8).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(99).ToCodeOfPc(), pPos.GetPieceAtIndex(89).ToCodeOfPc(), pPos.GetPieceAtIndex(79).ToCodeOfPc(), pPos.GetPieceAtIndex(69).ToCodeOfPc(), pPos.GetPieceAtIndex(59).ToCodeOfPc(), pPos.GetPieceAtIndex(49).ToCodeOfPc(), pPos.GetPieceAtIndex(39).ToCodeOfPc(), pPos.GetPieceAtIndex(29).ToCodeOfPc(), pPos.GetPieceAtIndex(19).ToCodeOfPc(), pPos.GetPieceAtIndex(9).ToCodeOfPc()) +
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
// * `flag` - 0: 利き数ボード, 1-5:利き数の差分ボードのレイヤー[0]～[4]
func (pPos *Position) SprintControl(phase l06.Phase, flag int) string {
	var board [l03.BOARD_SIZE]int8
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
	if ph < 2 { // 0 <= ph &&
		if flag == 0 {
			title = "Control"
			board = pPos.ControlBoards[ph]
		} else {
			// 利き数の差分
			var layer = flag - 1
			title = fmt.Sprintf("ControlDiff%d", layer)
			board = pPos.ControlBoardsDiff[ph][layer]
		}
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
			var piece = pPos.GetPieceAtSq(SquareFrom(file, rank))
			var pieceCode = piece.ToCodeOfPc()

			length := len(pieceCode)

			if length > 0 && spaces > 0 {
				buf = append(buf, OneDigitNumbers[spaces])
				spaces = 0
			}

			switch length {
			case 2:
				buf = append(buf, pieceCode[0])
				buf = append(buf, pieceCode[1])
			case 1:
				buf = append(buf, pieceCode[0])
			default:
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
	switch pPos.Phase {
	case l06.FIRST:
		phaseStr = "b"
	case l06.SECOND:
		phaseStr = "w"
	default:
		panic(fmt.Errorf("LogicalError: Unknows phase=[%d]", pPos.Phase))
	}

	// 持ち駒
	hands := ""
	num := pPos.Hands[0]
	if num == 1 {
		hands += "R"
	} else if num > 1 {
		hands += fmt.Sprintf("%dR", num)
	}

	num = pPos.Hands[1]
	if num == 1 {
		hands += "B"
	} else if num > 1 {
		hands += fmt.Sprintf("%dB", num)
	}

	num = pPos.Hands[2]
	if num == 1 {
		hands += "G"
	} else if num > 1 {
		hands += fmt.Sprintf("%dG", num)
	}

	num = pPos.Hands[3]
	if num == 1 {
		hands += "S"
	} else if num > 1 {
		hands += fmt.Sprintf("%dS", num)
	}

	num = pPos.Hands[4]
	if num == 1 {
		hands += "N"
	} else if num > 1 {
		hands += fmt.Sprintf("%dN", num)
	}

	num = pPos.Hands[5]
	if num == 1 {
		hands += "L"
	} else if num > 1 {
		hands += fmt.Sprintf("%dL", num)
	}

	num = pPos.Hands[6]
	if num == 1 {
		hands += "P"
	} else if num > 1 {
		hands += fmt.Sprintf("%dP", num)
	}

	num = pPos.Hands[7]
	if num == 1 {
		hands += "r"
	} else if num > 1 {
		hands += fmt.Sprintf("%dr", num)
	}

	num = pPos.Hands[8]
	if num == 1 {
		hands += "b"
	} else if num > 1 {
		hands += fmt.Sprintf("%db", num)
	}

	num = pPos.Hands[9]
	if num == 1 {
		hands += "g"
	} else if num > 1 {
		hands += fmt.Sprintf("%dg", num)
	}

	num = pPos.Hands[10]
	if num == 1 {
		hands += "s"
	} else if num > 1 {
		hands += fmt.Sprintf("%ds", num)
	}

	num = pPos.Hands[11]
	if num == 1 {
		hands += "n"
	} else if num > 1 {
		hands += fmt.Sprintf("%dn", num)
	}

	num = pPos.Hands[12]
	if num == 1 {
		hands += "l"
	} else if num > 1 {
		hands += fmt.Sprintf("%dl", num)
	}

	num = pPos.Hands[13]
	if num == 1 {
		hands += "p"
	} else if num > 1 {
		hands += fmt.Sprintf("%dp", num)
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
