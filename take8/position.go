package take8

import (
	"fmt"
	"strconv"
	"strings"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// 00～99
const BOARD_SIZE = 100

// position sfen の盤のスペース数に使われますN
var OneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// 1:先手 2:後手
type Phase byte

// FlipPhase - 先後を反転します
func FlipPhase(phase Phase) Phase {
	return phase%2 + 1
}

// マス番号 00～99,100～113
type Square uint32

// From - 筋と段からマス番号を作成します
func SquareFrom(file Square, rank Square) Square {
	return Square(file*10 + rank)
}

// 盤上なら真。ラベルのとこを指定しても場所によっては真だがめんどくさいんでOKで（＾～＾）
func OnBoard(sq Square) bool {
	return 10 < sq && sq < 100
}

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

const (
	// 空マス
	ZEROTH = Phase(0)
	// 先手
	FIRST = Phase(1)
	// 後手
	SECOND = Phase(2)
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
	PIECE_PS2   = "+s"
	PIECE_PN2   = "+n"
	PIECE_PL2   = "+l"
	PIECE_PP2   = "+p"
)

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [BOARD_SIZE]string
	// [0]先手 [1]後手
	KingLocations [2]Square
	// 飛車の場所。長い利きを消すために必要（＾～＾）
	RookLocations [2]Square
	// 角の場所。長い利きを消すために必要（＾～＾）
	BishopLocations [2]Square
	// 香の場所。長い利きを消すために必要（＾～＾）
	LanceLocations [4]Square
	// 利きテーブル [0]先手 [1]後手
	// マスへの利き数が入っています
	ControlBoards [2][BOARD_SIZE]int8
	// マスへの利き数の差分が入っています。デバッグ目的で無駄に分けてるんだけどな（＾～＾）
	// プレイヤー１つにつき、５レイヤーあるぜ（＾～＾）
	ControlBoardsDiff [2][5][BOARD_SIZE]int8

	// 持ち駒の数だぜ（＾～＾） R, B, G, S, N, L, P, r, b, g, s, n, l, p
	Hands []int
	// 先手が1、後手が2（＾～＾）
	Phase Phase
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
	ins.resetToZero()
	return ins
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPos *Position) resetToZero() {
	// 筋、段のラベルだけ入れとくぜ（＾～＾）
	pPos.Board = [BOARD_SIZE]string{
		"", "a", "b", "c", "d", "e", "f", "g", "h", "i",
		"1", "", "", "", "", "", "", "", "", "",
		"2", "", "", "", "", "", "", "", "", "",
		"3", "", "", "", "", "", "", "", "", "",
		"4", "", "", "", "", "", "", "", "", "",
		"5", "", "", "", "", "", "", "", "", "",
		"6", "", "", "", "", "", "", "", "", "",
		"7", "", "", "", "", "", "", "", "", "",
		"8", "", "", "", "", "", "", "", "", "",
		"9", "", "", "", "", "", "", "", "", "",
	}
	pPos.ControlBoards = [2][BOARD_SIZE]int8{{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}}
	pPos.ControlBoardsDiff = [2][5][BOARD_SIZE]int8{{
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}, {
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}}
	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.KingLocations = [2]Square{SQUARE_EMPTY, SQUARE_EMPTY}
	pPos.RookLocations = [2]Square{SQUARE_EMPTY, SQUARE_EMPTY}
	pPos.BishopLocations = [2]Square{SQUARE_EMPTY, SQUARE_EMPTY}
	pPos.LanceLocations = [4]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pPos.Phase = FIRST
	// 何手目か
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0
	// 指し手のリスト
	pPos.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPos.CapturedList = [MOVES_SIZE]string{}
}

// setToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) setToStartpos() {
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
	pPos.KingLocations = [2]Square{Square(59), Square(51)}
	pPos.RookLocations = [2]Square{28, 82}
	pPos.BishopLocations = [2]Square{22, 88}
	pPos.LanceLocations = [4]Square{11, 19, 91, 99}
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pPos *Position) ReadPosition(command string) {
	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面をセット（＾～＾）
		pPos.resetToZero()
		pPos.setToStartpos()
		i = 17

		if i < len && command[i] == ' ' {
			i += 1
		}
		// moves へ続くぜ（＾～＾）

	} else if strings.HasPrefix(command, "position sfen ") {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		pPos.resetToZero()
		i = 14
		var rank = 1
		var file = 9

	BoardLoop:
		for {
			promoted := false
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
				promoted = true
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

			if promoted {
				switch pc2 := command[i]; pc2 {
				case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
					pPos.Board[file*10+rank] = "+" + string(pc2)
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.KingLocations[0] = Square((file+1)*10 + rank)
			case 'k':
				pPos.KingLocations[1] = Square((file+1)*10 + rank)
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i, sq := range pPos.RookLocations {
					if sq == SQUARE_EMPTY {
						pPos.RookLocations[i] = Square((file+1)*10 + rank)
						break
					}
				}
			case 'B', 'b':
				for i, sq := range pPos.BishopLocations {
					if sq == SQUARE_EMPTY {
						pPos.BishopLocations[i] = Square((file+1)*10 + rank)
						break
					}
				}
			case 'L', 'l':
				for i, sq := range pPos.LanceLocations {
					if sq == SQUARE_EMPTY {
						pPos.LanceLocations[i] = Square((file+1)*10 + rank)
						break
					}
				}
			}
		}

		// 手番
		switch command[i] {
		case 'b':
			pPos.Phase = FIRST
			i += 1
		case 'w':
			pPos.Phase = SECOND
			i += 1
		default:
			panic("Fatal: Unknown phase")
		}

		if command[i] != ' ' {
			// 手番の後ろにスペースがない（＾～＾）
			panic("Fatal: Nothing space")
		}
		i += 1

		// 持ち駒
		if command[i] == '-' {
			i += 1
			if command[i] != ' ' {
				// 持ち駒 - の後ろにスペースがない（＾～＾）
				panic("Fatal: Nothing space after -")
			}
			i += 1
		} else {
		HandLoop:
			for {
				var drop_index Square
				var piece = command[i]
				switch piece {
				case 'R':
					drop_index = HAND_R1
				case 'B':
					drop_index = HAND_B1
				case 'G':
					drop_index = HAND_G1
				case 'S':
					drop_index = HAND_S1
				case 'N':
					drop_index = HAND_N1
				case 'L':
					drop_index = HAND_L1
				case 'P':
					drop_index = HAND_P1
				case 'r':
					drop_index = HAND_R2
				case 'b':
					drop_index = HAND_B2
				case 'g':
					drop_index = HAND_G2
				case 's':
					drop_index = HAND_S2
				case 'n':
					drop_index = HAND_N2
				case 'l':
					drop_index = HAND_L2
				case 'p':
					drop_index = HAND_P2
				case ' ':
					i += 1
					break HandLoop
				default:
					panic(fmt.Errorf("Fatal: Unknown piece=%c", piece))
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

				pPos.Hands[drop_index-HAND_ORIGIN] = number

				// 長い利きの駒は位置を覚えておくぜ（＾～＾）
				switch drop_index {
				case HAND_R1, HAND_R2:
					for i, sq := range pPos.RookLocations {
						if sq == SQUARE_EMPTY {
							pPos.RookLocations[i] = drop_index
							break
						}
					}
				case HAND_B1, HAND_B2:
					for i, sq := range pPos.BishopLocations {
						if sq == SQUARE_EMPTY {
							pPos.BishopLocations[i] = drop_index
							break
						}
					}
				case HAND_L1, HAND_L2:
					for i, sq := range pPos.LanceLocations {
						if sq == SQUARE_EMPTY {
							pPos.LanceLocations[i] = drop_index
							break
						}
					}
				}

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

	} else {
		fmt.Printf("Error: Unknown command=[%s]", command)
	}

	// fmt.Printf("command[i:]=[%s]\n", command[i:])

	start_phase := pPos.Phase
	if strings.HasPrefix(command[i:], "moves") {
		i += 5

		// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
		for i < len {
			if command[i] != ' ' {
				break
			}
			i += 1

			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			var move, err = ParseMove(command, &i, pPos.Phase)
			if err != nil {
				fmt.Println(err)
				fmt.Println(pPos.Sprint())
				panic(err)
			}
			pPos.Moves[pPos.OffsetMovesIndex] = move
			pPos.OffsetMovesIndex += 1
			pPos.Phase = FlipPhase(pPos.Phase)
		}
	}

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 開始局面の利きを計算（＾～＾）
	//fmt.Printf("Debug: 開始局面の利きを計算（＾～＾）\n")
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				//fmt.Printf("Debug: sq=%d\n", sq)
				pPos.AddControlDiff(0, sq, 1)
			}
		}
	}
	//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
	pPos.MergeControlDiff()

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPos.OffsetMovesIndex = 0
	pPos.Phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPos.DoMove(pPos.Moves[i])
	}
}

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase Phase) (Move, error) {
	var len = len(command)
	var move = NewMoveValue()

	var hand1 = Square(0)

	// file
	switch ch := command[*i]; ch {
	case 'R':
		*i += 1
		hand1 = HAND_R1
	case 'B':
		*i += 1
		hand1 = HAND_B1
	case 'G':
		*i += 1
		hand1 = HAND_G1
	case 'S':
		*i += 1
		hand1 = HAND_S1
	case 'N':
		*i += 1
		hand1 = HAND_N1
	case 'L':
		*i += 1
		hand1 = HAND_L1
	case 'P':
		*i += 1
		hand1 = HAND_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand1 != 0 {
		switch phase {
		case FIRST:
			move = move.ReplaceSource(hand1)
		case SECOND:
			move = move.ReplaceSource(hand1 + HAND_TYPE_SIZE)
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(Move), fmt.Errorf("Fatal: not *")
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
				return *new(Move), fmt.Errorf("Fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := Square(file*10 + rank)
			if count == 0 {
				move = move.ReplaceSource(sq)
			} else if count == 1 {
				move = move.ReplaceDestination(sq)
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

// DoMove - 一手指すぜ（＾～＾）
func (pPos *Position) DoMove(move Move) {
	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	mov_src_sq := move.GetSource()
	if pPos.IsEmptySq(mov_src_sq) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", mov_src_sq)
	}
	mov_dst_sq := move.GetDestination()
	var cap_src_sq Square
	var cap_dst_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	pPos.AddControlDiffAllSlidingPiece(0, -1, mov_src_sq)

	// まず、打かどうかで処理を分けます
	drop := mov_src_sq
	var piece string
	switch mov_src_sq {
	case HAND_R1:
		piece = PIECE_R1
	case HAND_B1:
		piece = PIECE_B1
	case HAND_G1:
		piece = PIECE_G1
	case HAND_S1:
		piece = PIECE_S1
	case HAND_N1:
		piece = PIECE_N1
	case HAND_L1:
		piece = PIECE_L1
	case HAND_P1:
		piece = PIECE_P1
	case HAND_R2:
		piece = PIECE_R2
	case HAND_B2:
		piece = PIECE_B2
	case HAND_G2:
		piece = PIECE_G2
	case HAND_S2:
		piece = PIECE_S2
	case HAND_N2:
		piece = PIECE_N2
	case HAND_L2:
		piece = PIECE_L2
	case HAND_P2:
		piece = PIECE_P2
	default:
		// Not drop
		drop = Square(0)
	}

	if drop != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands[drop-HAND_ORIGIN] -= 1

		// 行き先に駒を置きます
		pPos.Board[mov_dst_sq] = piece
		pPos.AddControlDiff(1, mov_dst_sq, 1)
		mov_piece_type = What(piece)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[mov_dst_sq]
		if captured != PIECE_EMPTY {
			pieceType := What(captured)
			switch pieceType {
			case PIECE_TYPE_R, PIECE_TYPE_PR, PIECE_TYPE_B, PIECE_TYPE_PB, PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				pPos.AddControlDiff(1, mov_dst_sq, -1)
			}
			cap_piece_type = What(captured)
			cap_src_sq = mov_dst_sq
		}

		// 元位置の駒を除去
		pPos.AddControlDiff(2, mov_src_sq, -1)

		// 行き先の駒の上書き
		if move.IsPromotion() {
			// 駒を成りに変換します
			pPos.Board[mov_dst_sq] = Promote(pPos.Board[mov_src_sq])
		} else {
			pPos.Board[mov_dst_sq] = pPos.Board[mov_src_sq]
		}
		// 元位置の駒の削除pos
		mov_piece_type = What(pPos.Board[mov_dst_sq])
		pPos.Board[mov_src_sq] = PIECE_EMPTY
		pPos.AddControlDiff(3, mov_dst_sq, 1)

		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1, PIECE_PR1:
			cap_dst_sq = HAND_R2
		case PIECE_B1, PIECE_PB1:
			cap_dst_sq = HAND_B2
		case PIECE_G1:
			cap_dst_sq = HAND_G2
		case PIECE_S1, PIECE_PS1:
			cap_dst_sq = HAND_S2
		case PIECE_N1, PIECE_PN1:
			cap_dst_sq = HAND_N2
		case PIECE_L1, PIECE_PL1:
			cap_dst_sq = HAND_L2
		case PIECE_P1, PIECE_PP1:
			cap_dst_sq = HAND_P2
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2, PIECE_PR2:
			cap_dst_sq = HAND_R1
		case PIECE_B2, PIECE_PB2:
			cap_dst_sq = HAND_B1
		case PIECE_G2:
			cap_dst_sq = HAND_G1
		case PIECE_S2, PIECE_PS2:
			cap_dst_sq = HAND_S1
		case PIECE_N2, PIECE_PN2:
			cap_dst_sq = HAND_N1
		case PIECE_L2, PIECE_PL2:
			cap_dst_sq = HAND_L1
		case PIECE_P2, PIECE_PP2:
			cap_dst_sq = HAND_P1
		default:
			fmt.Printf("Error: Unknown captured=[%s]", captured)
		}

		if cap_dst_sq != SQUARE_EMPTY {
			pPos.CapturedList[pPos.OffsetMovesIndex] = captured
			pPos.Hands[cap_dst_sq-HAND_ORIGIN] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pPos.CapturedList[pPos.OffsetMovesIndex] = PIECE_EMPTY
		}
	}

	pPos.Moves[pPos.OffsetMovesIndex] = move
	pPos.OffsetMovesIndex += 1
	prev_phase := pPos.Phase
	pPos.Phase = FlipPhase(pPos.Phase)

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []Square{mov_src_sq, cap_src_sq}
	dst_sq_list := []Square{mov_dst_sq, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case PIECE_TYPE_K:
			switch prev_phase {
			case FIRST:
				pPos.KingLocations[prev_phase-1] = dst_sq_list[j]
			case SECOND:
				pPos.KingLocations[prev_phase-1] = dst_sq_list[j]
			default:
				panic(fmt.Errorf("Unknown prev_phase=%d", prev_phase))
			}
		case PIECE_TYPE_R, PIECE_TYPE_PR:
			for i, sq := range pPos.RookLocations {
				if sq == src_sq_list[j] {
					pPos.RookLocations[i] = dst_sq_list[j]
				}
			}
		case PIECE_TYPE_B, PIECE_TYPE_PB:
			for i, sq := range pPos.BishopLocations {
				if sq == src_sq_list[j] {
					pPos.BishopLocations[i] = dst_sq_list[j]
				}
			}
		case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i, sq := range pPos.LanceLocations {
				if sq == src_sq_list[j] {
					pPos.LanceLocations[i] = dst_sq_list[j]
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし動かした駒を除きます
	pPos.AddControlDiffAllSlidingPiece(4, 1, mov_dst_sq)

	pPos.MergeControlDiff()
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pPos *Position) UndoMove() {

	// G.StderrChat.Trace(pPos.Sprint())

	if pPos.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	prev_phase := pPos.Phase
	pPos.Phase = FlipPhase(pPos.Phase)

	pPos.OffsetMovesIndex -= 1
	move := pPos.Moves[pPos.OffsetMovesIndex]
	captured := pPos.CapturedList[pPos.OffsetMovesIndex]

	mov_dst_sq := move.GetDestination()
	mov_src_sq := move.GetSource()
	var cap_dst_sq Square
	var cap_src_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	pPos.AddControlDiffAllSlidingPiece(0, -1, mov_dst_sq)

	// 打かどうかで分けます
	switch mov_src_sq {
	case HAND_R1, HAND_B1, HAND_G1, HAND_S1, HAND_N1, HAND_L1, HAND_P1, HAND_R2, HAND_B2, HAND_G2, HAND_S2, HAND_N2, HAND_L2, HAND_P2:
		// 打なら
		drop := mov_src_sq
		// 盤上から駒を除去します
		mov_piece_type = What(pPos.Board[mov_dst_sq])
		pPos.Board[mov_dst_sq] = PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands[drop-HAND_ORIGIN] += 1
		cap_dst_sq = 0
	default:
		// 打でないなら

		// 行き先の駒の除去
		mov_piece_type = What(pPos.Board[mov_dst_sq])
		pPos.AddControlDiff(1, mov_dst_sq, -1)

		// 移動元への駒の配置
		if move.IsPromotion() {
			// 成りを元に戻します
			pPos.Board[mov_src_sq] = Demote(pPos.Board[mov_dst_sq])
		} else {
			pPos.Board[mov_src_sq] = pPos.Board[mov_dst_sq]
		}

		// あれば、取った駒は駒台から下ろします
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			// Lost first king
		case PIECE_R1, PIECE_PR1:
			cap_src_sq = HAND_R2
		case PIECE_B1, PIECE_PB1:
			cap_src_sq = HAND_B2
		case PIECE_G1:
			cap_src_sq = HAND_G2
		case PIECE_S1, PIECE_PS1:
			cap_src_sq = HAND_S2
		case PIECE_N1, PIECE_PN1:
			cap_src_sq = HAND_N2
		case PIECE_L1, PIECE_PL1:
			cap_src_sq = HAND_L2
		case PIECE_P1, PIECE_PP1:
			cap_src_sq = HAND_P2
		case PIECE_K2: // First player win
			// Lost second king
		case PIECE_R2, PIECE_PR2:
			cap_src_sq = HAND_R1
		case PIECE_B2, PIECE_PB2:
			cap_src_sq = HAND_B1
		case PIECE_G2:
			cap_src_sq = HAND_G1
		case PIECE_S2, PIECE_PS2:
			cap_src_sq = HAND_S1
		case PIECE_N2, PIECE_PN2:
			cap_src_sq = HAND_N1
		case PIECE_L2, PIECE_PL2:
			cap_src_sq = HAND_L1
		case PIECE_P2, PIECE_PP2:
			cap_src_sq = HAND_P1
		default:
			fmt.Printf("Error: Unknown captured=[%s]", captured)
		}

		if cap_src_sq != SQUARE_EMPTY {
			cap_dst_sq = cap_src_sq
			pPos.Hands[cap_src_sq-HAND_ORIGIN] -= 1

			// 取った駒を行き先に戻します
			cap_piece_type = What(captured)
			pPos.Board[mov_dst_sq] = captured
			pPos.AddControlDiff(2, mov_src_sq, 1)

			// pieceType := What(captured)
			// switch pieceType {
			// case PIECE_TYPE_R, PIECE_TYPE_PR, PIECE_TYPE_B, PIECE_TYPE_PB, PIECE_TYPE_L:
			// 	// Ignored: 長い利きの駒は あとで追加するので、ここでは無視します
			// default:
			// 取った駒は盤上になかったので、ここで利きを復元させます
			pPos.AddControlDiff(3, mov_dst_sq, 1)
			// }

		} else {
			pPos.Board[mov_dst_sq] = PIECE_EMPTY
		}
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []PieceType{mov_piece_type, cap_piece_type}
	dst_sq_list := []Square{mov_dst_sq, cap_dst_sq}
	src_sq_list := []Square{mov_src_sq, cap_src_sq}
	for j, moving_piece_type := range piece_type_list {
		switch moving_piece_type {
		case PIECE_TYPE_K:
			switch prev_phase {
			case FIRST:
				pPos.KingLocations[prev_phase-1] = src_sq_list[j]
			case SECOND:
				pPos.KingLocations[prev_phase-1] = src_sq_list[j]
			default:
				panic(fmt.Errorf("Unknown prev_phase=%d", prev_phase))
			}
		case PIECE_TYPE_R, PIECE_TYPE_PR:
			for i, sq := range pPos.RookLocations {
				if sq == dst_sq_list[j] {
					pPos.RookLocations[i] = src_sq_list[j]
				}
			}
		case PIECE_TYPE_B, PIECE_TYPE_PB:
			for i, sq := range pPos.BishopLocations {
				if sq == dst_sq_list[j] {
					pPos.BishopLocations[i] = src_sq_list[j]
				}
			}
		case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i, sq := range pPos.LanceLocations {
				if sq == dst_sq_list[j] {
					pPos.LanceLocations[i] = src_sq_list[j]
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	pPos.AddControlDiffAllSlidingPiece(4, 1, mov_src_sq)

	pPos.MergeControlDiff()
}

// AddControlDiffAllSlidingPiece - すべての長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlDiffAllSlidingPiece(layer int, sign int8, excludeFrom Square) {
	for _, from := range pPos.RookLocations {
		if OnBoard(from) && from != excludeFrom {
			pPos.AddControlDiff(layer, from, sign)
		}
	}
	for _, from := range pPos.BishopLocations {
		if OnBoard(from) && from != excludeFrom {
			pPos.AddControlDiff(layer, from, sign)
		}
	}
	for _, from := range pPos.LanceLocations {
		if OnBoard(from) && from != excludeFrom && PIECE_TYPE_PL != What(pPos.Board[from]) { // 杏は除外
			pPos.AddControlDiff(layer, from, sign)
		}
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlDiff(layer int, from Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[from]
	if piece == PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: Piece from empty square. It has no control. from=%d", from))
	}

	ph := int(Who(piece)) - 1
	// fmt.Printf("Debug: ph=%d\n", ph)

	sq_list := GenControl(pPos, from)

	for _, to := range sq_list {
		// fmt.Printf("Debug: to=%d\n", to)
		// 差分の方のテーブルを更新（＾～＾）
		pPos.ControlBoardsDiff[ph][layer][to] += sign * 1
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pPos *Position) ClearControlDiff() {
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			for layer := 0; layer < 5; layer += 1 {
				pPos.ControlBoardsDiff[0][layer][sq] = 0
				pPos.ControlBoardsDiff[1][layer][sq] = 0
			}
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pPos *Position) MergeControlDiff() {
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			for layer := 0; layer < 5; layer += 1 {
				pPos.ControlBoards[0][sq] += pPos.ControlBoardsDiff[0][layer][sq]
				pPos.ControlBoards[1][sq] += pPos.ControlBoardsDiff[1][layer][sq]
			}
		}
	}
}

// Homo - 移動元と移動先の駒を持つプレイヤーが等しければ真。移動先が空なら偽
// 持ち駒は指定してはいけません。
func (pPos *Position) Homo(from Square, to Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return Who(pPos.Board[from]) == Who(pPos.Board[to])
}

// Hetero - 移動元と移動先の駒を持つプレイヤーが異なれば真。移動先が空マスでも真
// 持ち駒は指定してはいけません。
// Homo の逆だぜ（＾～＾）片方ありゃいいんだけど（＾～＾）
func (pPos *Position) Hetero(from Square, to Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return Who(pPos.Board[from]) != Who(pPos.Board[to])
}

// IsEmptySq - 空きマスなら真。持ち駒は偽
func (pPos *Position) IsEmptySq(sq Square) bool {
	if sq > 99 {
		return false
	}
	return pPos.Board[sq] == PIECE_EMPTY
}
