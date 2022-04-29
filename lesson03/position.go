package lesson03

import (
	"fmt"
	"strconv"
	"strings"
)

// Position - 局面
type Position struct {
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board []Piece
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
	pos.Board = []Piece{
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_L2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_L1,
		PIECE_EMPTY, PIECE_N2, PIECE_B2, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_R1, PIECE_N1,
		PIECE_EMPTY, PIECE_S2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_S1,
		PIECE_EMPTY, PIECE_G2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_G1,
		PIECE_EMPTY, PIECE_K2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_K1,
		PIECE_EMPTY, PIECE_G2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_G1,
		PIECE_EMPTY, PIECE_S2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_S1,
		PIECE_EMPTY, PIECE_N2, PIECE_R2, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_B1, PIECE_N1,
		PIECE_EMPTY, PIECE_L2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_L1,
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
			pos.Board[file*10+rank] = FromCodeToPiece(string(pc))
			file -= 1
			i += 1
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			var spaces, _ = strconv.Atoi(string(pc))
			for sp := 0; sp < spaces; sp += 1 {
				pos.Board[file*10+rank] = FromCodeToPiece("")
				file -= 1
			}
			i += 1
		case '+':
			i += 1
			switch pc2 := command[i]; pc2 {
			case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
				pos.Board[file*10+rank] = FromCodeToPiece("+" + string(pc2))
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
			var handSq HandSq
			var piece = command[i]
			switch piece {
			case 'R':
				handSq = HANDSQ_R1
			case 'B':
				handSq = HANDSQ_B1
			case 'G':
				handSq = HANDSQ_G1
			case 'S':
				handSq = HANDSQ_S1
			case 'N':
				handSq = HANDSQ_N1
			case 'L':
				handSq = HANDSQ_L1
			case 'P':
				handSq = HANDSQ_P1
			case 'r':
				handSq = HANDSQ_R2
			case 'b':
				handSq = HANDSQ_B2
			case 'g':
				handSq = HANDSQ_G2
			case 's':
				handSq = HANDSQ_S2
			case 'n':
				handSq = HANDSQ_N2
			case 'l':
				handSq = HANDSQ_L2
			case 'p':
				handSq = HANDSQ_P2
			case ' ':
				i += 1
				break HandLoop
			default:
				panic("fatal: 知らん持ち駒（＾～＾）")
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

			pos.Hands[handSq-HANDSQ_ORIGIN] = number
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
	var pMove = NewMove(0, 0, false)

	// 0=移動元 1=移動先
	var count = 0

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_R1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_R2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'B':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_B1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_B2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'G':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_G1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_G2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'S':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_S1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_S2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'N':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_N1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_N2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'L':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_L1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_L2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	case 'P':
		*i += 1
		switch phase {
		case FIRST:
			pMove.Squares[0] = HANDSQ_P1.ToSq()
		case SECOND:
			pMove.Squares[0] = HANDSQ_P2.ToSq()
		default:
			return *new(Move), fmt.Errorf("fatal: 分からんフェーズ（＾～＾） phase=%d", phase)
		}
	default:
		// Ignored
	}

	if count == 1 {
		if command[*i] != '+' {
			return *new(Move), fmt.Errorf("fatal: +じゃなかった（＾～＾）")
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
				return *new(Move), fmt.Errorf("fatal: なんか分かんないfileかrank（＾～＾） ch2='%c'", ch2)
			}
			*i += 1

			pMove.Squares[count] = Square(file*10 + rank)
		default:
			return *new(Move), fmt.Errorf("fatal: なんか分かんないmove（＾～＾） ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pMove.Promotion = true
	}

	return *pMove, nil
}

// DoMove - 一手指すぜ（＾～＾）
func (pos *Position) DoMove(move Move) {
	switch FromSqToHandSq(move.Squares[0]) {
	case HANDSQ_R1:
		pos.Hands[HANDSQ_R1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_R1
	case HANDSQ_B1:
		pos.Hands[HANDSQ_B1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_B1
	case HANDSQ_G1:
		pos.Hands[HANDSQ_G1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_G1
	case HANDSQ_S1:
		pos.Hands[HANDSQ_S1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_S1
	case HANDSQ_N1:
		pos.Hands[HANDSQ_N1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_N1
	case HANDSQ_L1:
		pos.Hands[HANDSQ_L1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_L1
	case HANDSQ_P1:
		pos.Hands[HANDSQ_P1-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_P1
	case HANDSQ_R2:
		pos.Hands[HANDSQ_R2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_R2
	case HANDSQ_B2:
		pos.Hands[HANDSQ_B2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_B2
	case HANDSQ_G2:
		pos.Hands[HANDSQ_G2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_G2
	case HANDSQ_S2:
		pos.Hands[HANDSQ_S2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_S2
	case HANDSQ_N2:
		pos.Hands[HANDSQ_N2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_N2
	case HANDSQ_L2:
		pos.Hands[HANDSQ_L2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_L2
	case HANDSQ_P2:
		pos.Hands[HANDSQ_P2-HANDSQ_ORIGIN] -= 1
		pos.Board[move.Squares[1]] = PIECE_P2
	default:
		// あれば、取った駒
		captured := pos.GetPieceAtSq(move.Squares[1])
		var piece2 = pos.GetPieceAtSq(move.Squares[0])

		pos.Board[move.Squares[1]] = piece2
		pos.Board[move.Squares[0]] = PIECE_EMPTY
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1:
			pos.Hands[HANDSQ_R2-HANDSQ_ORIGIN] += 1
		case PIECE_B1:
			pos.Hands[HANDSQ_B2-HANDSQ_ORIGIN] += 1
		case PIECE_G1:
			pos.Hands[HANDSQ_G2-HANDSQ_ORIGIN] += 1
		case PIECE_S1:
			pos.Hands[HANDSQ_S2-HANDSQ_ORIGIN] += 1
		case PIECE_N1:
			pos.Hands[HANDSQ_N2-HANDSQ_ORIGIN] += 1
		case PIECE_L1:
			pos.Hands[HANDSQ_L2-HANDSQ_ORIGIN] += 1
		case PIECE_P1:
			pos.Hands[HANDSQ_P2-HANDSQ_ORIGIN] += 1
		case PIECE_PR1:
			pos.Hands[HANDSQ_R2-HANDSQ_ORIGIN] += 1
		case PIECE_PB1:
			pos.Hands[HANDSQ_B2-HANDSQ_ORIGIN] += 1
		case PIECE_PS1:
			pos.Hands[HANDSQ_S2-HANDSQ_ORIGIN] += 1
		case PIECE_PN1:
			pos.Hands[HANDSQ_N2-HANDSQ_ORIGIN] += 1
		case PIECE_PL1:
			pos.Hands[HANDSQ_L2-HANDSQ_ORIGIN] += 1
		case PIECE_PP1:
			pos.Hands[HANDSQ_P2-HANDSQ_ORIGIN] += 1
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2:
			pos.Hands[HANDSQ_R1-HANDSQ_ORIGIN] += 1
		case PIECE_B2:
			pos.Hands[HANDSQ_B1-HANDSQ_ORIGIN] += 1
		case PIECE_G2:
			pos.Hands[HANDSQ_G1-HANDSQ_ORIGIN] += 1
		case PIECE_S2:
			pos.Hands[HANDSQ_S1-HANDSQ_ORIGIN] += 1
		case PIECE_N2:
			pos.Hands[HANDSQ_N1-HANDSQ_ORIGIN] += 1
		case PIECE_L2:
			pos.Hands[HANDSQ_L1-HANDSQ_ORIGIN] += 1
		case PIECE_P2:
			pos.Hands[HANDSQ_P1-HANDSQ_ORIGIN] += 1
		case PIECE_PR2:
			pos.Hands[HANDSQ_R1-HANDSQ_ORIGIN] += 1
		case PIECE_PB2:
			pos.Hands[HANDSQ_B1-HANDSQ_ORIGIN] += 1
		case PIECE_PS2:
			pos.Hands[HANDSQ_S1-HANDSQ_ORIGIN] += 1
		case PIECE_PN2:
			pos.Hands[HANDSQ_N1-HANDSQ_ORIGIN] += 1
		case PIECE_PL2:
			pos.Hands[HANDSQ_L1-HANDSQ_ORIGIN] += 1
		case PIECE_PP2:
			pos.Hands[HANDSQ_P1-HANDSQ_ORIGIN] += 1
		default:
			fmt.Printf("Error: 知らん駒を取ったぜ（＾～＾） captured=[%s]", captured.ToCodeOfPc())
		}
	}
}
