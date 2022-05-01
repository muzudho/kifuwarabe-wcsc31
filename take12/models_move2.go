package take12 // not same take13

import (
	"fmt"
	"strconv"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase l03.Phase) (l03.Move, error) {
	var len = len(command)
	var move = l03.RESIGN_MOVE

	var hand_sq = l03.SQ_EMPTY

	// file
	switch ch := command[*i]; ch {
	case 'R':
		hand_sq = l03.SQ_R1
	case 'B':
		hand_sq = l03.SQ_B1
	case 'G':
		hand_sq = l03.SQ_G1
	case 'S':
		hand_sq = l03.SQ_S1
	case 'N':
		hand_sq = l03.SQ_N1
	case 'L':
		hand_sq = l03.SQ_L1
	case 'P':
		hand_sq = l03.SQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand_sq != l03.SQ_EMPTY {
		*i += 1
		switch phase {
		case l03.FIRST:
			move = move.ReplaceSource(hand_sq)
		case l03.SECOND:
			move = move.ReplaceSource(hand_sq + l03.HAND_TYPE_SIZE_SQ)
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
				return *new(l03.Move), fmt.Errorf("fatal: unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := l03.Square(file*10 + rank)
			if count == 0 {
				move = move.ReplaceSource(sq)
			} else if count == 1 {
				move = move.ReplaceDestination(sq)
			} else {
				return *new(l03.Move), fmt.Errorf("fatal: unknown count='%c'", count)
			}
		default:
			return *new(l03.Move), fmt.Errorf("fatal: unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		move = move.ReplacePromotion(true)
	}

	return move, nil
}
