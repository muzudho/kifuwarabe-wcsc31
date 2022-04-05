package lesson02

import (
	"fmt"
	"strconv"
	"strings"
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
	Moves []string
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
	pos.Moves = []string{}
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pos *Position) ReadPosition(command string) {
	G.Log.Trace("command=%s\n", command)

	if strings.HasPrefix(command, "position startpos") {
		pos.ResetToStartpos()
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
			// fmt.Printf("(%d,%d) [%s]\n", file, rank, string(pc))
			pos.Board[file*10+rank] = string(pc)
			file -= 1
			i += 1
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// fmt.Printf("(%d,%d) [%s]\n", file, rank, string(pc))
			var spaces, _ = strconv.Atoi(string(pc))
			// fmt.Printf("[%s]=%d spaces\n", string(pc), spaces)
			for sp := 0; sp < spaces; sp += 1 {
				pos.Board[file*10+rank] = ""
				file -= 1
			}
			i += 1
		case '+':
			// fmt.Printf("(%d,%d) [%s]\n", file, rank, string(pc))
			i += 1
			switch pc2 := command[i]; pc2 {
			case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
				fmt.Printf("(%d,%d) [%s]\n", file, rank, string(pc))
				pos.Board[file*10+rank] = "+" + string(pc2)
				file -= 1
				i += 1
			default:
				panic("Undefined sfen board+")
			}
		case '/':
			// fmt.Printf("(%d,%d) [%s]\n", file, rank, string(pc))
			file = 9
			rank += 1
			i += 1
		case ' ':
			// fmt.Printf("(%d,%d) [%s]\n", file, rank, string(pc))
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
	MoveLoop:
		for {
			var piece_type = command[i]
			switch piece_type {
			case 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'r', 'b', 'g', 's', 'n', 'l', 'p':

			case ' ':
				i += 1
				break MoveLoop
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
					break MoveLoop
				default:
					break NumberLoop
				}
			}
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

	for i < len {
		var move = make([]byte, 0, 5)
		if command[i] != ' ' {
			break
		}
		i += 1

		var count = 0

		// file
		switch ch := command[i]; ch {
		case 'R', 'B', 'G', 'S', 'N', 'L', 'P':
			i += 1
			move = append(move, ch)

			if command[i] != '+' {
				panic("Fatal: +じゃなかった（＾～＾）")
			}

			i += 1
			move = append(move, '+')
			count = 1
		default:
			// Ignored
		}

		// file, rank
		for count < 2 {
			switch ch := command[i]; ch {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				i += 1
				move = append(move, ch)

				switch ch2 := command[i]; ch2 {
				case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i':
					i += 1
					move = append(move, ch2)
				default:
					panic(fmt.Errorf("Fatal: なんか分かんないfileかrank（＾～＾） ch2='%c'", ch2))
				}

			default:
				fmt.Println(pos.Sprint())
				panic(fmt.Errorf("Fatal: なんか分かんないmove（＾～＾） ch='%c' move=%s", ch, string(move)))
			}

			count += 1
		}

		if i < len && command[i] == '+' {
			i += 1
			move = append(move, '+')
		}

		pos.Moves = append(pos.Moves, string(move))
	}
}

// Print - 局面出力（＾ｑ＾）
func (pos *Position) Sprint() string {
	var phase_str = "First"
	if pos.Phase == 2 {
		phase_str = "Second"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d moves / %s / ? repeats]\n", pos.MovesNum, phase_str) +
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

	moves_list := make([]byte, 0, 512*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for _, move := range pos.Moves {
		moves_list = append(moves_list, ' ')
		moves_list = append(moves_list, move...)
	}

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	// return s1 + *(*string)(unsafe.Pointer(&moves_list)) + "\n"
	return s1 + string(moves_list) + "\n"
}
