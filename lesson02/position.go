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
	My.Log.Trace("command=%s\n", command)

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
				fmt.Println(Sprint(pos))
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
