package take7

import (
	"fmt"
	"strconv"
	"strings"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// 00～99
const BOARD_SIZE = 100

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [BOARD_SIZE]string

	// 玉と長い利きの駒の場所。長い利きを消すのに使う
	// [0]先手玉 [1]後手玉 [2:3]飛 [4:5]角 [6:9]香
	PieceLocations [PCLOC_SIZE]l03.Square

	// 利きテーブル [0]先手 [1]後手
	// マスへの利き数が入っています
	ControlBoards [2][BOARD_SIZE]int8

	// 持ち駒の数だぜ（＾～＾） R, B, G, S, N, L, P, r, b, g, s, n, l, p
	Hands []int
	// 先手が1、後手が2（＾～＾）
	Phase l06.Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]l04.Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [MOVES_SIZE]string
}

func NewPosition() *Position {
	var ins = new(Position)
	ins.ResetToStartpos()
	return ins
}

// ResetToStartpos - 初期局面にします。
func (pPos *Position) ResetToStartpos() {
	// 初期局面にします
	pPos.Board = [BOARD_SIZE]string{
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
	pPos.ControlBoards = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 2, 2, 1,
		0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
		0, 0, 0, 0, 0, 0, 1, 1, 3, 1,
		0, 0, 0, 0, 0, 0, 1, 0, 4, 2,
		0, 0, 0, 0, 0, 0, 1, 0, 4, 1,
		0, 0, 0, 0, 0, 0, 1, 0, 4, 2,
		0, 0, 0, 0, 0, 0, 1, 2, 3, 1,
		0, 0, 0, 0, 0, 0, 1, 0, 2, 0,
		0, 0, 0, 0, 0, 0, 1, 3, 1, 1,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 1, 3, 1, 0, 0, 0, 0, 0,
		0, 0, 2, 0, 1, 0, 0, 0, 0, 0,
		0, 2, 3, 2, 1, 0, 0, 0, 0, 0,
		0, 1, 4, 0, 1, 0, 0, 0, 0, 0,
		0, 2, 4, 0, 1, 0, 0, 0, 0, 0,
		0, 1, 4, 0, 1, 0, 0, 0, 0, 0,
		0, 1, 3, 1, 1, 0, 0, 0, 0, 0,
		0, 1, 1, 1, 1, 0, 0, 0, 0, 0,
		0, 0, 2, 2, 1, 0, 0, 0, 0, 0,
	}}

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]l03.Square{
		51,
		59,
		28,
		82,
		22,
		88,
		11,
		19,
		91,
		99,
	}

	// 持ち駒の数
	pPos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pPos.Phase = l06.FIRST
	// 何手目か
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0
	// 指し手のリスト
	pPos.Moves = [MOVES_SIZE]l04.Move{}
	// 取った駒のリスト
	pPos.CapturedList = [MOVES_SIZE]string{}
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pPos *Position) ReadPosition(command string) {
	// めんどくさいんで、初期化の代わりに 平手初期局面をセットするぜ（＾～＾） 盤面は あとで上書きされるから大丈夫（＾～＾）
	pPos.ResetToStartpos()

	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面が指定されたら、さっき初期化したんで、そのまま終了だぜ（＾～＾）
		i = 17

		if i >= len || command[i] != ' ' {
			return
		}
		i += 1
		// moves へ続くぜ（＾～＾）

	} else {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		i = 14
		var rank = 1
		var file = 9

	BoardLoop:
		for {
			switch pc := command[i]; pc {
			case 'K', 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'k', 'r', 'b', 'g', 's', 'n', 'l', 'p':
				pPos.Board[file*10+rank] = string(pc)
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = ""
					file -= 1
				}
				i += 1
			case '+':
				i += 1
				switch pc2 := command[i]; pc2 {
				case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
					pPos.Board[file*10+rank] = "+" + string(pc2)
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
			pPos.Phase = l06.FIRST
			i += 1
		case 'w':
			pPos.Phase = l06.SECOND
			i += 1
		default:
			panic("fatal: unknown phase")
		}

		if command[i] != ' ' {
			// 手番の後ろにスペースがない（＾～＾）
			panic("fatal: Nothing space")
		}
		i += 1

		// 持ち駒
		if command[i] == '-' {
			i += 1
			if command[i] != ' ' {
				// 持ち駒 - の後ろにスペースがない（＾～＾）
				panic("fatal: Nothing space after -")
			}
			i += 1
		} else {
		HandLoop:
			for {
				var handSq l03.HandSq
				var piece = command[i]
				switch piece {
				case 'R':
					handSq = l03.HANDSQ_R1
				case 'B':
					handSq = l03.HANDSQ_B1
				case 'G':
					handSq = l03.HANDSQ_G1
				case 'S':
					handSq = l03.HANDSQ_S1
				case 'N':
					handSq = l03.HANDSQ_N1
				case 'L':
					handSq = l03.HANDSQ_L1
				case 'P':
					handSq = l03.HANDSQ_P1
				case 'r':
					handSq = l03.HANDSQ_R2
				case 'b':
					handSq = l03.HANDSQ_B2
				case 'g':
					handSq = l03.HANDSQ_G2
				case 's':
					handSq = l03.HANDSQ_S2
				case 'n':
					handSq = l03.HANDSQ_N2
				case 'l':
					handSq = l03.HANDSQ_L2
				case 'p':
					handSq = l03.HANDSQ_P2
				case ' ':
					i += 1
					break HandLoop
				default:
					panic(fmt.Errorf("fatal: unknown piece=%c", piece))
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

				pPos.Hands[handSq] = number
			}
		}

		// 手数
		pPos.StartMovesNum = 0
	MovesNumLoop:
		for i < len {
			switch figure := command[i]; figure {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num, err := strconv.Atoi(string(figure))
				if err != nil {
					panic(err)
				}
				i += 1
				pPos.StartMovesNum *= 10
				pPos.StartMovesNum += num
			case ' ':
				i += 1
				break MovesNumLoop
			default:
				break MovesNumLoop
			}
		}

	}

	// fmt.Printf("command[i:]=[%s]\n", command[i:])

	if strings.HasPrefix(command[i:], "moves") {
		i += 5
	} else {
		return
	}

	// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
	start_phase := pPos.Phase
	for i < len {
		if command[i] != ' ' {
			break
		}
		i += 1

		// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
		var move, err = ParseMove(command, &i, pPos.Phase)
		if err != nil {
			fmt.Println(err)
			fmt.Println(Sprint(pPos))
			panic(err)
		}
		pPos.Moves[pPos.OffsetMovesIndex] = move
		pPos.OffsetMovesIndex += 1
		pPos.Phase = pPos.Phase%2 + 1
	}

	// 読込んだ l04.Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPos.OffsetMovesIndex = 0
	pPos.Phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPos.DoMove(pPos.Moves[i])
	}
}

// ParseMove
func ParseMove(command string, i *int, phase l06.Phase) (l04.Move, error) {
	var len = len(command)
	var handSq1 = l03.HandSq(0)

	var from l03.Square
	var to l03.Square
	var pro = false

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		handSq1 = l03.HANDSQ_R1
	case 'B':
		*i += 1
		handSq1 = l03.HANDSQ_B1
	case 'G':
		*i += 1
		handSq1 = l03.HANDSQ_G1
	case 'S':
		*i += 1
		handSq1 = l03.HANDSQ_S1
	case 'N':
		*i += 1
		handSq1 = l03.HANDSQ_N1
	case 'L':
		*i += 1
		handSq1 = l03.HANDSQ_L1
	case 'P':
		*i += 1
		handSq1 = l03.HANDSQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if handSq1 != 0 {
		switch phase {
		case l06.FIRST:
			from = handSq1.ToSq()
		case l06.SECOND:
			from = handSq1.ToSq() + l03.HANDSQ_TYPE_SIZE_SQ
		default:
			return *new(l04.Move), fmt.Errorf("fatal: unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(l04.Move), fmt.Errorf("fatal: not *")
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
				return *new(l04.Move), fmt.Errorf("fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := l03.Square(file*10 + rank)
			if count == 0 {
				from = sq
			} else if count == 1 {
				to = sq
			} else {
				return *new(l04.Move), fmt.Errorf("fatal: Unknown count='%c'", count)
			}
		default:
			return *new(l04.Move), fmt.Errorf("fatal: Unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pro = true
	}

	return l04.NewMove(from, to, pro), nil
}

// Print - 利き数ボード出力（＾ｑ＾）
func (pPos *Position) SprintControl(phase l06.Phase) string {
	var board [BOARD_SIZE]int8
	var phase_str string
	switch phase {
	case l06.FIRST:
		phase_str = "First"
		board = pPos.ControlBoards[0]
	case l06.SECOND:
		phase_str = "Second"
		board = pPos.ControlBoards[1]
	default:
		return "\n"
	}

	return "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPos.StartMovesNum, (pPos.StartMovesNum+pPos.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pPos.Board[90], pPos.Board[80], pPos.Board[70], pPos.Board[60], pPos.Board[50], pPos.Board[40], pPos.Board[30], pPos.Board[20], pPos.Board[10], pPos.Board[0]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[91], board[81], board[71], board[61], board[51], board[41], board[31], board[21], board[11], board[1]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[92], board[82], board[72], board[62], board[52], board[42], board[32], board[22], board[12], board[2]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[93], board[83], board[73], board[63], board[53], board[43], board[33], board[23], board[13], board[3]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[94], board[84], board[74], board[64], board[54], board[44], board[34], board[24], board[14], board[4]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[95], board[85], board[75], board[65], board[55], board[45], board[35], board[25], board[15], board[5]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[96], board[86], board[76], board[66], board[56], board[46], board[36], board[26], board[16], board[6]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[97], board[87], board[77], board[67], board[57], board[47], board[37], board[27], board[17], board[7]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[98], board[88], board[78], board[68], board[58], board[48], board[38], board[28], board[18], board[8]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d\n", board[99], board[89], board[79], board[69], board[59], board[49], board[39], board[29], board[19], board[9]) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n"
}

// DoMove - 一手指すぜ（＾～＾）
func (pPos *Position) DoMove(move l04.Move) {
	// 作業前に、長い利きの駒の利きを -1 します
	pPos.AddControlAllSlidingPiece(-1)

	from, to, _ := move.Destructure()

	// [0]movingPieceType [1]capturedPieceType
	moving_piece_types := []PieceType{PIECE_TYPE_EMPTY, PIECE_TYPE_EMPTY}

	// まず、打かどうかで処理を分けます
	hand := l03.FromSqToHandSq(from)
	var piece string
	switch l03.FromSqToHandSq(from) {
	case l03.HANDSQ_R1:
		piece = l03.PIECE_R1.ToCodeOfPc()
	case l03.HANDSQ_B1:
		piece = l03.PIECE_B1.ToCodeOfPc()
	case l03.HANDSQ_G1:
		piece = l03.PIECE_G1.ToCodeOfPc()
	case l03.HANDSQ_S1:
		piece = l03.PIECE_S1.ToCodeOfPc()
	case l03.HANDSQ_N1:
		piece = l03.PIECE_N1.ToCodeOfPc()
	case l03.HANDSQ_L1:
		piece = l03.PIECE_L1.ToCodeOfPc()
	case l03.HANDSQ_P1:
		piece = l03.PIECE_P1.ToCodeOfPc()
	case l03.HANDSQ_R2:
		piece = l03.PIECE_R2.ToCodeOfPc()
	case l03.HANDSQ_B2:
		piece = l03.PIECE_B2.ToCodeOfPc()
	case l03.HANDSQ_G2:
		piece = l03.PIECE_G2.ToCodeOfPc()
	case l03.HANDSQ_S2:
		piece = l03.PIECE_S2.ToCodeOfPc()
	case l03.HANDSQ_N2:
		piece = l03.PIECE_N2.ToCodeOfPc()
	case l03.HANDSQ_L2:
		piece = l03.PIECE_L2.ToCodeOfPc()
	case l03.HANDSQ_P2:
		hand = l03.FromSqToHandSq(from)
		piece = l03.PIECE_P2.ToCodeOfPc()
	default:
		// Not hand
		hand = l03.FromSqToHandSq(0)
	}

	if hand != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands[hand-l03.HANDSQ_ORIGIN] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		pPos.AddControl(to, 1)
		moving_piece_types[0] = What(piece)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します
		captured := pPos.Board[to]
		if captured != l03.PIECE_EMPTY.ToCodeOfPc() {
			pPos.AddControl(to, -1)
			moving_piece_types[1] = What(captured)
		}

		// 元位置の駒を除去
		pPos.AddControl(from, -1)

		// 行き先の駒の配置
		pPos.Board[to] = pPos.Board[from]
		moving_piece_types[0] = What(pPos.Board[to])
		pPos.Board[from] = l03.PIECE_EMPTY.ToCodeOfPc()
		pPos.AddControl(to, 1)

		hand := l03.HandSq(0)
		switch captured {
		case l03.PIECE_EMPTY.ToCodeOfPc(): // Ignored
		case l03.PIECE_K1.ToCodeOfPc(): // Second player win
			// Lost l06.FIRST king
		case l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc():
			hand = l03.HANDSQ_R2
		case l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc():
			hand = l03.HANDSQ_B2
		case l03.PIECE_G1.ToCodeOfPc():
			hand = l03.HANDSQ_G2
		case l03.PIECE_S1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc():
			hand = l03.HANDSQ_S2
		case l03.PIECE_N1.ToCodeOfPc(), l03.PIECE_PN1.ToCodeOfPc():
			hand = l03.HANDSQ_N2
		case l03.PIECE_L1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc():
			hand = l03.HANDSQ_L2
		case l03.PIECE_P1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc():
			hand = l03.HANDSQ_P2
		case l03.PIECE_K2.ToCodeOfPc(): // l06.FIRST player win
			// Lost second king
		case l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc():
			hand = l03.HANDSQ_R1
		case l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc():
			hand = l03.HANDSQ_B1
		case l03.PIECE_G2.ToCodeOfPc():
			hand = l03.HANDSQ_G1
		case l03.PIECE_S2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc():
			hand = l03.HANDSQ_S1
		case l03.PIECE_N2.ToCodeOfPc(), l03.PIECE_PN2.ToCodeOfPc():
			hand = l03.HANDSQ_N1
		case l03.PIECE_L2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc():
			hand = l03.HANDSQ_L1
		case l03.PIECE_P2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc():
			hand = l03.HANDSQ_P1
		default:
			fmt.Printf("unknown captured=[%s]", captured)
		}

		if hand != 0 {
			pPos.Hands[hand-l03.HANDSQ_ORIGIN] += 1
		}
	}

	pPos.Moves[pPos.OffsetMovesIndex] = move
	pPos.OffsetMovesIndex += 1
	pPos.Phase = pPos.Phase%2 + 1

	// 長い利きの駒が動いたときは、位置情報更新
	for _, moving_piece_type := range moving_piece_types {
		switch moving_piece_type {
		case PIECE_TYPE_R:
			for i, sq := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
				if sq == from {
					pPos.PieceLocations[i] = to
				}
			}
		case PIECE_TYPE_B:
			for i, sq := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
				if sq == from {
					pPos.PieceLocations[i] = to
				}
			}
		case PIECE_TYPE_L:
			for i, sq := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
				if sq == from {
					pPos.PieceLocations[i] = to
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します
	pPos.AddControlAllSlidingPiece(1)
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pPos *Position) UndoMove() {
	if pPos.OffsetMovesIndex < 1 {
		return
	}

	// [0]movingPieceType [1]capturedPieceType
	moving_piece_types := []PieceType{PIECE_TYPE_EMPTY, PIECE_TYPE_EMPTY}

	// 作業前に、長い利きの駒の利きを -1 します
	pPos.AddControlAllSlidingPiece(-1)

	pPos.OffsetMovesIndex -= 1
	pPos.Phase = pPos.Phase%2 + 1
	move := pPos.Moves[pPos.OffsetMovesIndex]
	captured := pPos.CapturedList[pPos.OffsetMovesIndex]

	from, to, _ := move.Destructure()

	// 打かどうかで分けます
	switch l03.FromSqToHandSq(from) {
	case l03.HANDSQ_R1, l03.HANDSQ_B1, l03.HANDSQ_G1, l03.HANDSQ_S1, l03.HANDSQ_N1, l03.HANDSQ_L1, l03.HANDSQ_P1, l03.HANDSQ_R2, l03.HANDSQ_B2, l03.HANDSQ_G2, l03.HANDSQ_S2, l03.HANDSQ_N2, l03.HANDSQ_L2, l03.HANDSQ_P2:
		// 打なら
		hand := l03.FromSqToHandSq(from)
		// 盤上から駒を除去します
		moving_piece_types[0] = What(pPos.Board[to])
		pPos.Board[to] = l03.PIECE_EMPTY.ToCodeOfPc()

		// 駒台に駒を戻します
		pPos.Hands[hand-l03.HANDSQ_ORIGIN] += 1
	default:
		// 打でないなら

		// 行き先の駒の除去
		moving_piece_types[0] = What(pPos.Board[to])
		pPos.AddControl(to, -1)
		// 移動元への駒の配置
		pPos.Board[from] = pPos.Board[to]

		// あれば、取った駒は駒台から下ろします
		cap := l03.HandSq(0)
		switch captured {
		case l03.PIECE_EMPTY.ToCodeOfPc(): // Ignored
		case l03.PIECE_K1.ToCodeOfPc(): // Second player win
			// Lost l06.FIRST king
		case l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc():
			cap = l03.HANDSQ_R2
		case l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc():
			cap = l03.HANDSQ_B2
		case l03.PIECE_G1.ToCodeOfPc():
			cap = l03.HANDSQ_G2
		case l03.PIECE_S1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc():
			cap = l03.HANDSQ_S2
		case l03.PIECE_N1.ToCodeOfPc(), l03.PIECE_PN1.ToCodeOfPc():
			cap = l03.HANDSQ_N2
		case l03.PIECE_L1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc():
			cap = l03.HANDSQ_L2
		case l03.PIECE_P1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc():
			cap = l03.HANDSQ_P2
		case l03.PIECE_K2.ToCodeOfPc(): // l06.FIRST player win
			// Lost second king
		case l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc():
			cap = l03.HANDSQ_R1
		case l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc():
			cap = l03.HANDSQ_B1
		case l03.PIECE_G2.ToCodeOfPc():
			cap = l03.HANDSQ_G1
		case l03.PIECE_S2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc():
			cap = l03.HANDSQ_S1
		case l03.PIECE_N2.ToCodeOfPc(), l03.PIECE_PN2.ToCodeOfPc():
			cap = l03.HANDSQ_N1
		case l03.PIECE_L2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc():
			cap = l03.HANDSQ_L1
		case l03.PIECE_P2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc():
			cap = l03.HANDSQ_P1
		default:
			fmt.Printf("unknown captured=[%s]", captured)
		}

		if cap != 0 {
			pPos.Hands[cap-l03.HANDSQ_ORIGIN] -= 1

			// 取った駒を行き先に戻します
			moving_piece_types[1] = What(captured)
			pPos.Board[to] = captured
			pPos.AddControl(from, 1)
			pPos.AddControl(to, 1)
		}
	}

	// 長い利きの駒が動いたときは、位置情報更新
	for _, moving_piece_type := range moving_piece_types {
		switch moving_piece_type {
		case PIECE_TYPE_R:
			for i, sq := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
				if sq == from {
					pPos.PieceLocations[PCLOC_R1:PCLOC_R2][i] = to
				}
			}
		case PIECE_TYPE_B:
			for i, sq := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
				if sq == from {
					pPos.PieceLocations[PCLOC_B1:PCLOC_B2][i] = to
				}
			}
		case PIECE_TYPE_L:
			for i, sq := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
				if sq == from {
					pPos.PieceLocations[PCLOC_L1:PCLOC_L4][i] = to
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します
	pPos.AddControlAllSlidingPiece(1)
}

// AddControlAllSlidingPiece - すべての長い利きの駒の利きを増減させます
func (pPos *Position) AddControlAllSlidingPiece(sign int8) {
	for _, from := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
		pPos.AddControl(from, sign)
	}
	for _, from := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
		pPos.AddControl(from, sign)
	}
	for _, from := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
		pPos.AddControl(from, sign)
	}
}

// AddControl - 盤上のマスを指定することで、そこにある駒の利きを増減させます
func (pPos *Position) AddControl(from l03.Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[from]
	if piece == l03.PIECE_EMPTY.ToCodeOfPc() {
		panic(fmt.Errorf("LogicalError: Empty square has no control"))
	}

	ph := int(Who(piece)) - 1

	moveEndList := GenMoveEnd(pPos, from)

	for _, moveEnd := range moveEndList {
		to, _ := moveEnd.Destructure()
		pPos.ControlBoards[ph][to] += sign * 1
	}
}

// Homo - 手番と移動元の駒を持つプレイヤーが等しければ真。移動先が空なら偽
func (pPos *Position) Homo(to l03.Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return pPos.Phase == Who(pPos.Board[to])
}

// Hetero - 手番と移動先の駒を持つプレイヤーが異なれば真。移動先が空マスでも真
// Homo の逆だぜ（＾～＾）片方ありゃいいんだけど（＾～＾）
func (pPos *Position) Hetero(to l03.Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return pPos.Phase != Who(pPos.Board[to])
}

func (pPos *Position) IsEmptySq(sq l03.Square) bool {
	return pPos.Board[sq] == l03.PIECE_EMPTY.ToCodeOfPc()
}
