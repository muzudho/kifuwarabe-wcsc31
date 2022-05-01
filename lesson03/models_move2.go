package lesson03

import (
	"fmt"
	"strconv"
)

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase Phase) (Move, error) {
	var len = len(command)
	var move = RESIGN_MOVE

	var handSq = SQ_EMPTY

	// file
	switch ch := command[*i]; ch {
	case 'R':
		handSq = SQ_R1
	case 'B':
		handSq = SQ_B1
	case 'G':
		handSq = SQ_G1
	case 'S':
		handSq = SQ_S1
	case 'N':
		handSq = SQ_N1
	case 'L':
		handSq = SQ_L1
	case 'P':
		handSq = SQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if handSq != SQ_EMPTY {
		*i += 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(handSq)
		case SECOND:
			move = move.ReplaceSource(handSq + HAND_TYPE_SIZE_SQ)
		default:
			return *new(Move), fmt.Errorf("fatal: unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(Move), fmt.Errorf("fatal: not *")
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
				return *new(Move), fmt.Errorf("fatal: unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := Square(file*10 + rank)
			if count == 0 {
				move = move.ReplaceSource(sq)
			} else if count == 1 {
				move = move.ReplaceDestination(sq)
			} else {
				return *new(Move), fmt.Errorf("fatal: unknown count='%c'", count)
			}
		default:
			return *new(Move), fmt.Errorf("fatal: unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		move = move.ReplacePromotion(true)
	}

	return move, nil
}
