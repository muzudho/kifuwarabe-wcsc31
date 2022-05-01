package take5

import (
	"fmt"
	"strconv"
	"strings"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board []string
	// 持ち駒の数だぜ（＾～＾） R, B, G, S, N, L, P, r, b, g, s, n, l, p
	Hands []int
	// 先手が1、後手が2（＾～＾）
	Phase l03.Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [l04.MOVES_SIZE]l03.Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [l04.MOVES_SIZE]string
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
	pos.Moves = [l04.MOVES_SIZE]l03.Move{}
	// 取った駒のリスト
	pos.CapturedList = [l04.MOVES_SIZE]string{}
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
		panic("fatal: 手番わかんない（＾～＾）")
	}

	if command[i] != ' ' {
		panic("fatal: 手番の後ろにスペースがない（＾～＾）")
	}
	i += 1

	// 持ち駒
	if command[i] == '-' {
		i += 1
		if command[i] != ' ' {
			panic("fatal: 持ち駒 - の後ろにスペースがない（＾～＾）")
		}
		i += 1
	} else {
	HandLoop:
		for {
			var handSq l03.HandSq
			var piece = command[i]

			var isBreak = false
			var convertAlternativeValue = func(code byte) l03.HandSq {
				if code == ' ' {
					i += 1
					isBreak = true
					return l03.HANDSQ_SIZE // この値は使いません
				} else {
					panic(App.LogNotEcho.Fatal("fatal: unknown piece=%c", piece))
				}
			}

			handSq = l03.FromCodeToHandSq(byte(piece), &convertAlternativeValue)

			if isBreak {
				// ループを抜けます
				break HandLoop
			}
			i += 1

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

			pos.Hands[handSq-l03.HANDSQ_ORIGIN] = number
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
		var move, err = l03.ParseMove(command, &i, pos.Phase)
		if err != nil {
			fmt.Println(SprintBoard(pos))
			panic(err)
		}
		pos.Moves[pos.OffsetMovesIndex] = move
		pos.OffsetMovesIndex += 1
	}

	// 読込んだ l03.Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pos.OffsetMovesIndex = 0
	for i = 0; i < moves_size; i += 1 {
		pos.DoMove(pos.Moves[i])
	}
}

// DoMove - 一手指すぜ（＾～＾）
func (pos *Position) DoMove(move l03.Move) {
	from, to, _ := move.Destructure()
	switch l03.FromSqToHandSq(from) {
	case l03.HANDSQ_R1:
		pos.Hands[l03.HANDSQ_R1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_R1.ToCodeOfPc()
	case l03.HANDSQ_B1:
		pos.Hands[l03.HANDSQ_B1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_B1.ToCodeOfPc()
	case l03.HANDSQ_G1:
		pos.Hands[l03.HANDSQ_G1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_G1.ToCodeOfPc()
	case l03.HANDSQ_S1:
		pos.Hands[l03.HANDSQ_S1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_S1.ToCodeOfPc()
	case l03.HANDSQ_N1:
		pos.Hands[l03.HANDSQ_N1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_N1.ToCodeOfPc()
	case l03.HANDSQ_L1:
		pos.Hands[l03.HANDSQ_L1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_L1.ToCodeOfPc()
	case l03.HANDSQ_P1:
		pos.Hands[l03.HANDSQ_P1-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_P1.ToCodeOfPc()
	case l03.HANDSQ_R2:
		pos.Hands[l03.HANDSQ_R2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_R2.ToCodeOfPc()
	case l03.HANDSQ_B2:
		pos.Hands[l03.HANDSQ_B2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_B2.ToCodeOfPc()
	case l03.HANDSQ_G2:
		pos.Hands[l03.HANDSQ_G2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_G2.ToCodeOfPc()
	case l03.HANDSQ_S2:
		pos.Hands[l03.HANDSQ_S2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_S2.ToCodeOfPc()
	case l03.HANDSQ_N2:
		pos.Hands[l03.HANDSQ_N2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_N2.ToCodeOfPc()
	case l03.HANDSQ_L2:
		pos.Hands[l03.HANDSQ_L2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_L2.ToCodeOfPc()
	case l03.HANDSQ_P2:
		pos.Hands[l03.HANDSQ_P2-l03.HANDSQ_ORIGIN] -= 1
		pos.Board[to] = l03.PIECE_P2.ToCodeOfPc()
	default:
		// あれば、取った駒
		captured := pos.Board[to]
		pos.Board[to] = pos.Board[from]
		pos.Board[from] = l03.PIECE_EMPTY.ToCodeOfPc()
		switch captured {
		case l03.PIECE_EMPTY.ToCodeOfPc(): // Ignored
		case l03.PIECE_K1.ToCodeOfPc(): // Second player win
			// Lost first king
		case l03.PIECE_R1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_B1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_G1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_G2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_S1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_N1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_L1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_P1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PR1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PB1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PS1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PN1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PL1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PP1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P2-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_K2.ToCodeOfPc(): // First player win
			// Lost second king
		case l03.PIECE_R2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_B2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_G2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_G1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_S2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_N2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_L2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_P2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PR2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PB2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PS2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PN2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PL2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L1-l03.HANDSQ_ORIGIN] += 1
		case l03.PIECE_PP2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P1-l03.HANDSQ_ORIGIN] += 1
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

	switch l03.FromSqToHandSq(from) {
	case l03.HANDSQ_R1:
		pos.Hands[l03.HANDSQ_R1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_B1:
		pos.Hands[l03.HANDSQ_B1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_G1:
		pos.Hands[l03.HANDSQ_G1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_S1:
		pos.Hands[l03.HANDSQ_S1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_N1:
		pos.Hands[l03.HANDSQ_N1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_L1:
		pos.Hands[l03.HANDSQ_L1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_P1:
		pos.Hands[l03.HANDSQ_P1-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_R2:
		pos.Hands[l03.HANDSQ_R2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_B2:
		pos.Hands[l03.HANDSQ_B2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_G2:
		pos.Hands[l03.HANDSQ_G2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_S2:
		pos.Hands[l03.HANDSQ_S2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_N2:
		pos.Hands[l03.HANDSQ_N2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_L2:
		pos.Hands[l03.HANDSQ_L2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	case l03.HANDSQ_P2:
		pos.Hands[l03.HANDSQ_P2-l03.HANDSQ_ORIGIN] += 1
		pos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()
	default:
		pos.Board[from] = pos.Board[to]
		// あれば、取った駒
		pos.Board[to] = captured
		switch captured {
		case l03.PIECE_EMPTY.ToCodeOfPc(): // Ignored
		case l03.PIECE_K1.ToCodeOfPc(): // Second player win
			// Lost first king
		case l03.PIECE_R1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_B1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_G1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_G2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_S1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_N1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_L1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_P1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PR1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PB1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PS1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PN1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PL1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PP1.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P2-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_K2.ToCodeOfPc(): // First player win
			// Lost second king
		case l03.PIECE_R2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_B2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_G2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_G1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_S2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_N2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_L2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_P2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PR2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_R1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PB2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_B1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PS2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_S1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PN2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_N1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PL2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_L1-l03.HANDSQ_ORIGIN] -= 1
		case l03.PIECE_PP2.ToCodeOfPc():
			pos.Hands[l03.HANDSQ_P1-l03.HANDSQ_ORIGIN] -= 1
		default:
			fmt.Printf("Error: 知らん駒を取ったぜ（＾～＾） captured=[%s]", captured)
		}
	}

}
