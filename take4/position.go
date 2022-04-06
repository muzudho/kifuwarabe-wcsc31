package take4

import (
	"fmt"
	"strconv"
	"strings"
)

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
	// 何手目か（＾～＾）
	MovesNum int
	// 指し手のリスト（＾～＾）
	Moves []Move
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
	pos.MovesNum = 1
	// 指し手のリスト
	pos.Moves = []Move{}
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
	pos.MovesNum = 0
MovesNumLoop:
	for i < len {
		switch figure := command[i]; figure {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num, err := strconv.Atoi(string(figure))
			if err != nil {
				panic(err)
			}
			i += 1
			pos.MovesNum *= 10
			pos.MovesNum += num
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
			fmt.Println(Sprint(pos))
			panic(err)
		}
		pos.Moves = append(pos.Moves, move)
	}

	// 読込んだ Move を全て実行します
	for _, move := range pos.Moves {
		pos.DoMove(move)
	}
}

// ParseMove
func ParseMove(command string, i *int, phase int) (Move, error) {
	var len = len(command)

	var from Square
	var to Square
	var pro = false

	// 0=移動元 1=移動先
	var count = 0

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_R1)
		case SECOND:
			from = Square(DROP_R2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'B':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_B1)
		case SECOND:
			from = Square(DROP_B2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'G':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_G1)
		case SECOND:
			from = Square(DROP_G2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'S':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_S1)
		case SECOND:
			from = Square(DROP_S2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'N':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_N1)
		case SECOND:
			from = Square(DROP_N2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'L':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_L1)
		case SECOND:
			from = Square(DROP_L2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'P':
		*i += 1
		switch phase {
		case FIRST:
			from = Square(DROP_P1)
		case SECOND:
			from = Square(DROP_P2)
		default:
			return *new(Move), fmt.Errorf("Fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	default:
		// Ignored
	}

	if count == 1 {
		if command[*i] != '+' {
			return *new(Move), fmt.Errorf("Fatal: +じゃなかった（＾～＾）")
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
				return *new(Move), fmt.Errorf("Fatal: なんか分かんないfileかrank（＾～＾） ch2='%c'", ch2)
			}
			*i += 1

			sq := Square(file*10 + rank)
			if count == 0 {
				from = sq
			} else if count == 1 {
				to = sq
			} else {
				return *new(Move), fmt.Errorf("Fatal: なんか分かんないcount（＾～＾） count='%c'", count)
			}
		default:
			return *new(Move), fmt.Errorf("Fatal: なんか分かんないmove（＾～＾） ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pro = true
	}

	return NewMove(from, to, pro), nil
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
}
