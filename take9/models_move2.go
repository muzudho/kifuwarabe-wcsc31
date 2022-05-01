package take9

import (
	"fmt"
	"strconv"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase l03.Phase) (l03.Move, error) {
	var len = len(command)
	var hand1 = l03.Square(0)

	var from l03.Square
	var to l03.Square
	var pro = false

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		hand1 = l03.HANDSQ_R1.ToSq()
	case 'B':
		*i += 1
		hand1 = l03.HANDSQ_B1.ToSq()
	case 'G':
		*i += 1
		hand1 = l03.HANDSQ_G1.ToSq()
	case 'S':
		*i += 1
		hand1 = l03.HANDSQ_S1.ToSq()
	case 'N':
		*i += 1
		hand1 = l03.HANDSQ_N1.ToSq()
	case 'L':
		*i += 1
		hand1 = l03.HANDSQ_L1.ToSq()
	case 'P':
		*i += 1
		hand1 = l03.HANDSQ_P1.ToSq()
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand1 != 0 {
		switch phase {
		case l03.FIRST:
			from = hand1
		case l03.SECOND:
			from = hand1 + l03.HANDSQ_TYPE_SIZE.ToSq()
		default:
			return *new(l03.Move), fmt.Errorf("fatal: unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(l03.Move), fmt.Errorf("fatal: not *")
		}
		*i += 1
		count = 1
	}

	// file, rank
	for count < 2 {
		switch ch := command[*i]; ch {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			*i += 1
			file, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}

			var rank int
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
				return *new(l03.Move), fmt.Errorf("fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := l03.Square(file*10 + rank)
			if count == 0 {
				from = sq
			} else if count == 1 {
				to = sq
			} else {
				return *new(l03.Move), fmt.Errorf("fatal: Unknown count='%c'", count)
			}
		default:
			return *new(l03.Move), fmt.Errorf("fatal: Unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pro = true
	}

	return l03.NewMove(from, to, pro), nil
}
