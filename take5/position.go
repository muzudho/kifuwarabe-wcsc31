package take5

import (
	"fmt"
	"strconv"
	"strings"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// マス番号 00～99,100～113
type Square uint32

const (
	// 先手
	FIRST = iota + 1
	// 後手
	SECOND
)

// 駒
const (
	PIECE_EMPTY = ""
	PIECE_K1    = "K"
	PIECE_R1    = "R"
	PIECE_B1    = "B"
	PIECE_G1    = "G"
	PIECE_S1    = "S"
	PIECE_N1    = "N"
	PIECE_L1    = "L"
	PIECE_P1    = "P"
	PIECE_PR1   = "+R"
	PIECE_PB1   = "+B"
	PIECE_PG1   = "+G"
	PIECE_PS1   = "+S"
	PIECE_PN1   = "+N"
	PIECE_PL1   = "+L"
	PIECE_PP1   = "+P"
	PIECE_K2    = "k"
	PIECE_R2    = "r"
	PIECE_B2    = "b"
	PIECE_G2    = "g"
	PIECE_S2    = "s"
	PIECE_N2    = "n"
	PIECE_L2    = "l"
	PIECE_P2    = "p"
	PIECE_PR2   = "+r"
	PIECE_PB2   = "+b"
	PIECE_PG2   = "+g"
	PIECE_PS2   = "+s"
	PIECE_PN2   = "+n"
	PIECE_PL2   = "+l"
	PIECE_PP2   = "+p"
)

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board []string
	// 持ち駒の数だぜ（＾～＾） R, B, G, S, N, L, P, r, b, g, s, n, l, p
	Hands []int
	// 先手が1、後手が2（＾～＾）
	Phase int
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [MOVES_SIZE]string
}

func NewPosition() *Position {
	var ins = new(Position)
	ins.ResetToStartpos()
	return ins
}

// ResetToStartpos - 初期局面にします。
func (pos *Position) ResetToStartpos() {
	// 初期局面にします
	pos.Board = []string{
		"", "a", "b", "c", "d", "e", "f", "g", "h", "i",
		"1", "l", "", "p", "", "", "", "P", "", "L",
		"2", "n", "b", "p", "", "", "", "P", "R", "N",
		"3", "s", "", "p", "", "", "", "P", "", "S",
		"4", "g", "", "p", "", "", "", "P", "", "G",
		"5", "k", "", "p", "", "", "", "P", "", "K",
		"6", "g", "", "p", "", "", "", "P", "", "G",
		"7", "s", "", "p", "", "", "", "P", "", "S",
		"8", "n", "r", "p", "", "", "", "P", "B", "N",
		"9", "l", "", "p", "", "", "", "P", "", "L",
	}
	// 持ち駒の数
	pos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pos.Phase = 1
	// 何手目か
	pos.StartMovesNum = 1
	pos.OffsetMovesIndex = 0
	// 指し手のリスト
	pos.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pos.CapturedList = [MOVES_SIZE]string{}
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pos *Position) ReadPosition(command string) {
	// めんどくさいんで、初期化の代わりに 平手初期局面をセットするぜ（＾～＾） 盤面は あとで上書きされるから大丈夫（＾～＾）
	pos.ResetToStartpos()

	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面が指定されたら、さっき初期化したんで、そのまま終了だぜ（＾～＾）
		return
	}

	var len = len(command)
	// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
	var i = 14
	var rank = 1
	var file = 9

BoardLoop:
	for {
		switch pc := command[i]; pc {
		case 'K', 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'k', 'r', 'b', 'g', 's', 'n', 'l', 'p':
			pos.Board[file*10+rank] = string(pc)
			file -= 1
			i += 1
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var spaces, _ = strconv.Atoi(string(pc))
			for sp := 0; sp < spaces; sp += 1 {
				pos.Board[file*10+rank] = ""
				file -= 1
			}
			i += 1
		case '+':
			i += 1
			switch pc2 := command[i]; pc2 {
			case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
				pos.Board[file*10+rank] = "+" + string(pc2)
				file -= 1
				i += 1
			default:
				panic("Undefined sfen board+")
			}
		case '/':
			file = 9
			rank += 1
			i += 1
		case ' ':
			i += 1
			break BoardLoop
		default:
			panic("Undefined sfen board")
		}
	}

	// 手番
	switch command[i] {
	case 'b':
		pos.Phase = 1
		i += 1
	case 'w':
		pos.Phase = 2
		i += 1
	default:
		panic("Fatal: 手番わかんない（＾～＾）")
	}

	if command[i] != ' ' {
		panic("Fatal: 手番の後ろにスペースがない（＾～＾）")
	}
	i += 1

	// 持ち駒
	if command[i] == '-' {
		i += 1
		if command[i] != ' ' {
			panic("Fatal: 持ち駒 - の後ろにスペースがない（＾～＾）")
		}
		i += 1
	} else {
	HandLoop:
		for {
			var drop_index int
			var piece = command[i]
			switch piece {
			case 'R':
				drop_index = DROP_R1
			case 'B':
				drop_index = DROP_B1
			case 'G':
				drop_index = DROP_G1
			case 'S':
				drop_index = DROP_S1
			case 'N':
				drop_index = DROP_N1
			case 'L':
				drop_index = DROP_L1
			case 'P':
				drop_index = DROP_P1
			case 'r':
				drop_index = DROP_R2
			case 'b':
				drop_index = DROP_B2
			case 'g':
				drop_index = DROP_G2
			case 's':
				drop_index = DROP_S2
			case 'n':
				drop_index = DROP_N2
			case 'l':
				drop_index = DROP_L2
			case 'p':
				drop_index = DROP_P2
			case ' ':
				i += 1
				break HandLoop
			default:
				panic("Fatal: 知らん持ち駒（＾～＾）")
			}

			var number = 0
		NumberLoop:
			for {
				switch figure := command[i]; figure {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					num, err := strconv.Atoi(string(figure))
					if err != nil {
						panic(err)
					}
					i += 1
					number *= 10
					number += num
				case ' ':
					i += 1
					break HandLoop
				default:
					break NumberLoop
				}
			}

			pos.Hands[drop_index] = number
		}
	}

	// 手数
	pos.StartMovesNum = 0
MovesNumLoop:
	for i < len {
		switch figure := command[i]; figure {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num, err := strconv.Atoi(string(figure))
			if err != nil {
				panic(err)
			}
			i += 1
			pos.StartMovesNum *= 10
			pos.StartMovesNum += num
		case ' ':
			i += 1
			break MovesNumLoop
		default:
			break MovesNumLoop
		}
	}

	// fmt.Printf("command[i:]=[%s]\n", command[i:])

	if strings.HasPrefix(command[i:], "moves") {
		i += 5
	} else {
		return
	}

	// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
	for i < len {
		if command[i] != ' ' {
			break
		}
		i += 1

		// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
		var move, err = ParseMove(command, &i, pos.Phase)
		if err != nil {
			fmt.Println(pos.Sprint())
			panic(err)
		}
		pos.Moves[pos.OffsetMovesIndex] = move
		pos.OffsetMovesIndex += 1
	}

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pos.OffsetMovesIndex = 0
	for i = 0; i < moves_size; i += 1 {
		pos.DoMove(pos.Moves[i])
	}
}

// ParseMove
func ParseMove(command string, i *int, phase int) (Move, error) {
	var len = len(command)
	var move = RESIGN_MOVE

	// 0=移動元 1=移動先
	var count = 0

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_R1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_R2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	case 'B':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_B1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_B2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	case 'G':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_G1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_G2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	case 'S':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_S1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_S2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	case 'N':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_N1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_N2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	case 'L':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_L1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_L2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	case 'P':
		*i += 1
		count = 1
		switch phase {
		case FIRST:
			move = move.ReplaceSource(uint32(DROP_P1))
		case SECOND:
			move = move.ReplaceSource(uint32(DROP_P2))
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}
	default:
		// Ignored
	}

	if count == 1 {
		if command[*i] != '*' {
			return *new(Move), fmt.Errorf("Fatal: no *")
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
				return *new(Move), fmt.Errorf("Fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := file*10 + rank
			if count == 0 {
				move = move.ReplaceSource(uint32(sq))
			} else if count == 1 {
				move = move.ReplaceDestination(uint32(sq))
			} else {
				return *new(Move), fmt.Errorf("Fatal: Unknown count='%c'", count)
			}
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		move = move.ReplacePromotion(true)
	}

	return move, nil
}

// Print - 局面出力（＾ｑ＾）
func (pos *Position) Sprint() string {
	var phase_str = "First"
	if pos.Phase == 2 {
		phase_str = "Second"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pos.StartMovesNum, (pos.StartMovesNum+pos.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pos.Hands[7], pos.Hands[8], pos.Hands[9], pos.Hands[10], pos.Hands[11], pos.Hands[12], pos.Hands[13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pos.Board[90], pos.Board[80], pos.Board[70], pos.Board[60], pos.Board[50], pos.Board[40], pos.Board[30], pos.Board[20], pos.Board[10], pos.Board[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[91], pos.Board[81], pos.Board[71], pos.Board[61], pos.Board[51], pos.Board[41], pos.Board[31], pos.Board[21], pos.Board[11], pos.Board[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[92], pos.Board[82], pos.Board[72], pos.Board[62], pos.Board[52], pos.Board[42], pos.Board[32], pos.Board[22], pos.Board[12], pos.Board[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[93], pos.Board[83], pos.Board[73], pos.Board[63], pos.Board[53], pos.Board[43], pos.Board[33], pos.Board[23], pos.Board[13], pos.Board[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[94], pos.Board[84], pos.Board[74], pos.Board[64], pos.Board[54], pos.Board[44], pos.Board[34], pos.Board[24], pos.Board[14], pos.Board[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[95], pos.Board[85], pos.Board[75], pos.Board[65], pos.Board[55], pos.Board[45], pos.Board[35], pos.Board[25], pos.Board[15], pos.Board[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[96], pos.Board[86], pos.Board[76], pos.Board[66], pos.Board[56], pos.Board[46], pos.Board[36], pos.Board[26], pos.Board[16], pos.Board[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[97], pos.Board[87], pos.Board[77], pos.Board[67], pos.Board[57], pos.Board[47], pos.Board[37], pos.Board[27], pos.Board[17], pos.Board[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[98], pos.Board[88], pos.Board[78], pos.Board[68], pos.Board[58], pos.Board[48], pos.Board[38], pos.Board[28], pos.Board[18], pos.Board[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.Board[99], pos.Board[89], pos.Board[79], pos.Board[69], pos.Board[59], pos.Board[49], pos.Board[39], pos.Board[29], pos.Board[19], pos.Board[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pos.Hands[0], pos.Hands[1], pos.Hands[2], pos.Hands[3], pos.Hands[4], pos.Hands[5], pos.Hands[6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	moves_text := make([]byte, 0, MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for i := 0; i < pos.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pos.Moves[i].ToCode()...)
	}

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
}

// DoMove - 一手指すぜ（＾～＾）
func (pos *Position) DoMove(move Move) {
	from, to, _ := move.Destructure()
	switch from {
	case DROP_R1:
		pos.Hands[DROP_R1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_R1
	case DROP_B1:
		pos.Hands[DROP_B1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_B1
	case DROP_G1:
		pos.Hands[DROP_G1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_G1
	case DROP_S1:
		pos.Hands[DROP_S1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_S1
	case DROP_N1:
		pos.Hands[DROP_N1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_N1
	case DROP_L1:
		pos.Hands[DROP_L1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_L1
	case DROP_P1:
		pos.Hands[DROP_P1-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_P1
	case DROP_R2:
		pos.Hands[DROP_R2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_R2
	case DROP_B2:
		pos.Hands[DROP_B2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_B2
	case DROP_G2:
		pos.Hands[DROP_G2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_G2
	case DROP_S2:
		pos.Hands[DROP_S2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_S2
	case DROP_N2:
		pos.Hands[DROP_N2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_N2
	case DROP_L2:
		pos.Hands[DROP_L2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_L2
	case DROP_P2:
		pos.Hands[DROP_P2-DROP_ORIGIN] -= 1
		pos.Board[to] = PIECE_P2
	default:
		// あれば、取った駒
		captured := pos.Board[to]
		pos.Board[to] = pos.Board[from]
		pos.Board[from] = PIECE_EMPTY
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1:
			pos.Hands[DROP_R2-DROP_ORIGIN] += 1
		case PIECE_B1:
			pos.Hands[DROP_B2-DROP_ORIGIN] += 1
		case PIECE_G1:
			pos.Hands[DROP_G2-DROP_ORIGIN] += 1
		case PIECE_S1:
			pos.Hands[DROP_S2-DROP_ORIGIN] += 1
		case PIECE_N1:
			pos.Hands[DROP_N2-DROP_ORIGIN] += 1
		case PIECE_L1:
			pos.Hands[DROP_L2-DROP_ORIGIN] += 1
		case PIECE_P1:
			pos.Hands[DROP_P2-DROP_ORIGIN] += 1
		case PIECE_PR1:
			pos.Hands[DROP_R2-DROP_ORIGIN] += 1
		case PIECE_PB1:
			pos.Hands[DROP_B2-DROP_ORIGIN] += 1
		case PIECE_PG1:
			pos.Hands[DROP_G2-DROP_ORIGIN] += 1
		case PIECE_PS1:
			pos.Hands[DROP_S2-DROP_ORIGIN] += 1
		case PIECE_PN1:
			pos.Hands[DROP_N2-DROP_ORIGIN] += 1
		case PIECE_PL1:
			pos.Hands[DROP_L2-DROP_ORIGIN] += 1
		case PIECE_PP1:
			pos.Hands[DROP_P2-DROP_ORIGIN] += 1
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2:
			pos.Hands[DROP_R1-DROP_ORIGIN] += 1
		case PIECE_B2:
			pos.Hands[DROP_B1-DROP_ORIGIN] += 1
		case PIECE_G2:
			pos.Hands[DROP_G1-DROP_ORIGIN] += 1
		case PIECE_S2:
			pos.Hands[DROP_S1-DROP_ORIGIN] += 1
		case PIECE_N2:
			pos.Hands[DROP_N1-DROP_ORIGIN] += 1
		case PIECE_L2:
			pos.Hands[DROP_L1-DROP_ORIGIN] += 1
		case PIECE_P2:
			pos.Hands[DROP_P1-DROP_ORIGIN] += 1
		case PIECE_PR2:
			pos.Hands[DROP_R1-DROP_ORIGIN] += 1
		case PIECE_PB2:
			pos.Hands[DROP_B1-DROP_ORIGIN] += 1
		case PIECE_PG2:
			pos.Hands[DROP_G1-DROP_ORIGIN] += 1
		case PIECE_PS2:
			pos.Hands[DROP_S1-DROP_ORIGIN] += 1
		case PIECE_PN2:
			pos.Hands[DROP_N1-DROP_ORIGIN] += 1
		case PIECE_PL2:
			pos.Hands[DROP_L1-DROP_ORIGIN] += 1
		case PIECE_PP2:
			pos.Hands[DROP_P1-DROP_ORIGIN] += 1
		default:
			fmt.Printf("Error: 知らん駒を取ったぜ（＾～＾） captured=[%s]", captured)
		}
	}

	pos.Moves[pos.OffsetMovesIndex] = move
	pos.OffsetMovesIndex += 1
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pos *Position) UndoMove() {
	if pos.OffsetMovesIndex < 1 {
		return
	}

	pos.OffsetMovesIndex -= 1
	move := pos.Moves[pos.OffsetMovesIndex]
	captured := pos.CapturedList[pos.OffsetMovesIndex]

	from, to, _ := move.Destructure()

	switch from {
	case DROP_R1:
		pos.Hands[DROP_R1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_B1:
		pos.Hands[DROP_B1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_G1:
		pos.Hands[DROP_G1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_S1:
		pos.Hands[DROP_S1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_N1:
		pos.Hands[DROP_N1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_L1:
		pos.Hands[DROP_L1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_P1:
		pos.Hands[DROP_P1-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_R2:
		pos.Hands[DROP_R2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_B2:
		pos.Hands[DROP_B2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_G2:
		pos.Hands[DROP_G2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_S2:
		pos.Hands[DROP_S2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_N2:
		pos.Hands[DROP_N2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_L2:
		pos.Hands[DROP_L2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	case DROP_P2:
		pos.Hands[DROP_P2-DROP_ORIGIN] += 1
		pos.Board[to] = PIECE_EMPTY
	default:
		pos.Board[from] = pos.Board[to]
		// あれば、取った駒
		pos.Board[to] = captured
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1:
			pos.Hands[DROP_R2-DROP_ORIGIN] -= 1
		case PIECE_B1:
			pos.Hands[DROP_B2-DROP_ORIGIN] -= 1
		case PIECE_G1:
			pos.Hands[DROP_G2-DROP_ORIGIN] -= 1
		case PIECE_S1:
			pos.Hands[DROP_S2-DROP_ORIGIN] -= 1
		case PIECE_N1:
			pos.Hands[DROP_N2-DROP_ORIGIN] -= 1
		case PIECE_L1:
			pos.Hands[DROP_L2-DROP_ORIGIN] -= 1
		case PIECE_P1:
			pos.Hands[DROP_P2-DROP_ORIGIN] -= 1
		case PIECE_PR1:
			pos.Hands[DROP_R2-DROP_ORIGIN] -= 1
		case PIECE_PB1:
			pos.Hands[DROP_B2-DROP_ORIGIN] -= 1
		case PIECE_PG1:
			pos.Hands[DROP_G2-DROP_ORIGIN] -= 1
		case PIECE_PS1:
			pos.Hands[DROP_S2-DROP_ORIGIN] -= 1
		case PIECE_PN1:
			pos.Hands[DROP_N2-DROP_ORIGIN] -= 1
		case PIECE_PL1:
			pos.Hands[DROP_L2-DROP_ORIGIN] -= 1
		case PIECE_PP1:
			pos.Hands[DROP_P2-DROP_ORIGIN] -= 1
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2:
			pos.Hands[DROP_R1-DROP_ORIGIN] -= 1
		case PIECE_B2:
			pos.Hands[DROP_B1-DROP_ORIGIN] -= 1
		case PIECE_G2:
			pos.Hands[DROP_G1-DROP_ORIGIN] -= 1
		case PIECE_S2:
			pos.Hands[DROP_S1-DROP_ORIGIN] -= 1
		case PIECE_N2:
			pos.Hands[DROP_N1-DROP_ORIGIN] -= 1
		case PIECE_L2:
			pos.Hands[DROP_L1-DROP_ORIGIN] -= 1
		case PIECE_P2:
			pos.Hands[DROP_P1-DROP_ORIGIN] -= 1
		case PIECE_PR2:
			pos.Hands[DROP_R1-DROP_ORIGIN] -= 1
		case PIECE_PB2:
			pos.Hands[DROP_B1-DROP_ORIGIN] -= 1
		case PIECE_PG2:
			pos.Hands[DROP_G1-DROP_ORIGIN] -= 1
		case PIECE_PS2:
			pos.Hands[DROP_S1-DROP_ORIGIN] -= 1
		case PIECE_PN2:
			pos.Hands[DROP_N1-DROP_ORIGIN] -= 1
		case PIECE_PL2:
			pos.Hands[DROP_L1-DROP_ORIGIN] -= 1
		case PIECE_PP2:
			pos.Hands[DROP_P1-DROP_ORIGIN] -= 1
		default:
			fmt.Printf("Error: 知らん駒を取ったぜ（＾～＾） captured=[%s]", captured)
		}
	}

}
