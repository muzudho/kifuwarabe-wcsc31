package lesson03

import (
	"fmt"
	"strconv"
)

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase Phase) (Move, error) {
	var len = len(command)
	var move = RESIGN_MOVE

	var sqOfHand Square

	// file
	switch ch := command[*i]; ch {
	case 'K':
		sqOfHand = SQ_K1
	case 'R':
		sqOfHand = SQ_R1
	case 'B':
		sqOfHand = SQ_B1
	case 'G':
		sqOfHand = SQ_G1
	case 'S':
		sqOfHand = SQ_S1
	case 'N':
		sqOfHand = SQ_N1
	case 'L':
		sqOfHand = SQ_L1
	case 'P':
		sqOfHand = SQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if sqOfHand != SQ_EMPTY {
		switch phase {
		case FIRST:
			//AssertSqOfHand(sqOfHand, command)
			move = move.ReplaceSource(sqOfHand)
		case SECOND:
			fmt.Printf("sqOfHand=%d, HAND_TYPE_SIZE_SQ=%d, 2p command [%d]='%c'\n", sqOfHand, HAND_TYPE_SIZE_SQ, *i, command[*i])
			//AssertSqOfHand(sqOfHand, command)
			sqOfHand += HAND_TYPE_SIZE_SQ
			//AssertSqOfHand(sqOfHand, command)
			move = move.ReplaceSource(sqOfHand)
		default:
			return *new(Move), fmt.Errorf("fatal: unknown phase=%d", phase)
		}

		*i += 1
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

	//AssertMove(move)
	return move, nil
}

// ToCodeOfM - SFEN の moves の後に続く指し手に使える文字列を返します
func (move Move) ToCodeOfM() string {

	// 投了（＾～＾）
	if uint32(move) == 0 {
		return "resign"
	}

	str := make([]byte, 0, 5)
	count := 0

	// 移動元マス、移動先マス、成りの有無
	from, to, pro := move.Destructure()

	// 移動元マス(Source square)
	switch from {
	case SQ_R1, SQ_R2:
		str = append(str, 'R')
		count = 1
	case SQ_B1, SQ_B2:
		str = append(str, 'B')
		count = 1
	case SQ_G1, SQ_G2:
		str = append(str, 'G')
		count = 1
	case SQ_S1, SQ_S2:
		str = append(str, 'S')
		count = 1
	case SQ_N1, SQ_N2:
		str = append(str, 'N')
		count = 1
	case SQ_L1, SQ_L2:
		str = append(str, 'L')
		count = 1
	case SQ_P1, SQ_P2:
		str = append(str, 'P')
		count = 1
	default:
		// Ignored
	}

	if count == 1 {
		// 打
		str = append(str, '*')
	}

	for count < 2 {
		var sq Square // マス番号
		if count == 0 {
			// 移動元
			sq = from
		} else if count == 1 {
			// 移動先
			sq = to
		} else {
			panic(App.LogNotEcho.Fatal("LogicError: count=%d", count))
		}
		// 正常時は必ず２桁（＾～＾）
		file := byte(sq / 10)
		rank := byte(sq % 10)
		// ASCII Code
		// '0'=48, '9'=57, 'a'=97, 'i'=105
		str = append(str, file+48)
		str = append(str, rank+96)
		// fmt.Printf("Debug: move=%d sq=%d count=%d file=%d rank=%d\n", uint32(move), sq, count, file, rank)
		count += 1
	}

	if pro {
		str = append(str, '+')
	}

	return string(str)
}
