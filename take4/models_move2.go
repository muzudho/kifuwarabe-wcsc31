package take4

import (
	"fmt"
	"strconv"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// ParseMove
func ParseMove(command string, i *int, phase int) (l03.Move, error) {
	var len = len(command)

	var from l03.Square
	var to l03.Square
	var pro = false

	// 0=移動元 1=移動先
	var count = 0

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_R1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_R2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'B':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_B1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_B2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'G':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_G1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_G2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'S':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_S1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_S2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'N':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_N1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_N2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'L':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_L1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_L2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'P':
		*i += 1
		switch phase {
		case FIRST:
			from = l03.Square(l03.HANDSQ_P1)
		case SECOND:
			from = l03.Square(l03.HANDSQ_P2)
		default:
			return *new(l03.Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	default:
		// Ignored
	}

	if count == 1 {
		if command[*i] != '+' {
			return *new(l03.Move), fmt.Errorf("fatal: +じゃなかった（＾～＾）")
		}
		*i += 1
	}

	// file, rank
	for count < 2 {
		switch ch := command[*i]; ch {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			*i += 1
			file_int, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			file := byte(file_int)

			var rank byte
			switch ch2 := command[*i]; ch2 {
			case 'a':
				rank = 1
			case 'b':
				rank = 2
			case 'c':
				rank = 3
			case 'd':
				rank = 4
			case 'e':
				rank = 5
			case 'f':
				rank = 6
			case 'g':
				rank = 7
			case 'h':
				rank = 8
			case 'i':
				rank = 9
			default:
				return *new(l03.Move), fmt.Errorf("fatal: なんか分かんないfileかrank（＾～＾） ch2='%c'", ch2)
			}
			*i += 1

			sq := l03.Square(file*10 + rank)
			if count == 0 {
				from = sq
			} else if count == 1 {
				to = sq
			} else {
				return *new(l03.Move), fmt.Errorf("fatal: なんか分かんないcount（＾～＾） count='%c'", count)
			}
		default:
			return *new(l03.Move), fmt.Errorf("fatal: なんか分かんないmove（＾～＾） ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pro = true
	}

	return l03.NewMove(from, to, pro), nil
}
