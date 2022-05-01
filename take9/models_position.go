package take9

import (
	"fmt"
	"strconv"
	"strings"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// position sfen の盤のスペース数に使われますN
var OneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// FlipPhase - 先後を反転します
func FlipPhase(phase l03.Phase) l03.Phase {
	return phase%2 + 1
}

// From - 筋と段からマス番号を作成します
func SquareFrom(file l03.Square, rank l03.Square) l03.Square {
	return l03.Square(file*10 + rank)
}

// 盤上なら真。ラベルのとこを指定しても場所によっては真だがめんどくさいんでOKで（＾～＾）
func OnBoard(sq l03.Square) bool {
	return 10 < sq && sq < 100
}

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [l03.BOARD_SIZE]l03.Piece
	// 玉と長い利きの駒の場所。長い利きを消すのに使う
	// [0]先手玉 [1]後手玉 [2:3]飛 [4:5]角 [6:9]香
	PieceLocations [PCLOC_SIZE]l03.Square
	// 利きテーブル [0]先手 [1]後手
	// マスへの利き数が入っています
	ControlBoards [2][l03.BOARD_SIZE]int8
	// マスへの利き数の差分が入っています。デバッグ目的で無駄に分けてるんだけどな（＾～＾）
	// プレイヤー１つにつき、５レイヤーあるぜ（＾～＾）
	ControlBoardsDiff [2][5][l03.BOARD_SIZE]int8

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
	CapturedList [l04.MOVES_SIZE]l03.Piece
}

func NewPosition() *Position {
	var ins = new(Position)
	ins.resetToZero()
	return ins
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPos *Position) resetToZero() {
	// 筋、段のラベルだけ入れとくぜ（＾～＾）
	pPos.Board = [l03.BOARD_SIZE]l03.Piece{
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
	}
	pPos.ControlBoards = [2][l03.BOARD_SIZE]int8{{
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
	pPos.ControlBoardsDiff = [2][5][l03.BOARD_SIZE]int8{{
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
	pPos.PieceLocations = [PCLOC_SIZE]l03.Square{
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
		l03.SQ_EMPTY,
	}

	// 持ち駒の数
	pPos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pPos.Phase = l03.FIRST
	// 何手目か
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0
	// 指し手のリスト
	pPos.Moves = [l04.MOVES_SIZE]l03.Move{}
	// 取った駒のリスト
	pPos.CapturedList = [l04.MOVES_SIZE]l03.Piece{}
}

// setToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) setToStartpos() {
	// 初期局面にします
	pPos.Board = [l03.BOARD_SIZE]l03.Piece{
		l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY,
		l03.PIECE_EMPTY, l03.PIECE_L2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_L1,
		l03.PIECE_EMPTY, l03.PIECE_N2, l03.PIECE_B2, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_R1, l03.PIECE_N1,
		l03.PIECE_EMPTY, l03.PIECE_S2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_S1,
		l03.PIECE_EMPTY, l03.PIECE_G2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_G1,
		l03.PIECE_EMPTY, l03.PIECE_K2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_K1,
		l03.PIECE_EMPTY, l03.PIECE_G2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_G1,
		l03.PIECE_EMPTY, l03.PIECE_S2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_S1,
		l03.PIECE_EMPTY, l03.PIECE_N2, l03.PIECE_R2, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_B1, l03.PIECE_N1,
		l03.PIECE_EMPTY, l03.PIECE_L2, l03.PIECE_EMPTY, l03.PIECE_P2, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_EMPTY, l03.PIECE_P1, l03.PIECE_EMPTY, l03.PIECE_L1,
	}
	pPos.PieceLocations = [PCLOC_SIZE]l03.Square{
		l03.Square(59),
		l03.Square(51),
		l03.Square(82),
		l03.Square(28),
		l03.Square(88),
		l03.Square(22),
		l03.Square(91),
		l03.Square(99),
		l03.Square(11),
		l03.Square(19),
	}
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
				pPos.Board[file*10+rank] = l03.FromCodeToPiece(string(pc))
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = l03.PIECE_EMPTY
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
					pPos.Board[file*10+rank] = l03.FromCodeToPiece("+" + string(pc2))
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.PieceLocations[PCLOC_K1:PCLOC_K2][0] = l03.Square((file+1)*10 + rank)
			case 'k':
				pPos.PieceLocations[PCLOC_K1:PCLOC_K2][1] = l03.Square((file+1)*10 + rank)
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i, sq := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[PCLOC_R1:PCLOC_R2][i] = l03.Square((file+1)*10 + rank)
						break
					}
				}
			case 'B', 'b':
				for i, sq := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[PCLOC_B1:PCLOC_B2][i] = l03.Square((file+1)*10 + rank)
						break
					}
				}
			case 'L', 'l':
				for i, sq := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[PCLOC_L1:PCLOC_L4][i] = l03.Square((file+1)*10 + rank)
						break
					}
				}
			}
		}

		// 手番
		switch command[i] {
		case 'b':
			pPos.Phase = l03.FIRST
			i += 1
		case 'w':
			pPos.Phase = l03.SECOND
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

				pPos.Hands[handSq.ToSq()-l03.HANDSQ_ORIGIN.ToSq()] = number

				// 長い利きの駒は位置を覚えておくぜ（＾～＾）
				switch handSq {
				case l03.HANDSQ_R1, l03.HANDSQ_R2:
					for i, sq := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
						if sq == l03.SQ_EMPTY {
							pPos.PieceLocations[PCLOC_R1:PCLOC_R2][i] = handSq.ToSq()
							break
						}
					}
				case l03.HANDSQ_B1, l03.HANDSQ_B2:
					for i, sq := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
						if sq == l03.SQ_EMPTY {
							pPos.PieceLocations[PCLOC_B1:PCLOC_B2][i] = handSq.ToSq()
							break
						}
					}
				case l03.HANDSQ_L1, l03.HANDSQ_L2:
					for i, sq := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
						if sq == l03.SQ_EMPTY {
							pPos.PieceLocations[PCLOC_L1:PCLOC_L4][i] = handSq.ToSq()
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
		fmt.Printf("unknown command=[%s]", command)
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
				fmt.Println(SprintBoard(pPos))
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
	for sq := l03.Square(11); sq < 100; sq += 1 {
		if l03.File(sq) != 0 && l03.Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				//fmt.Printf("Debug: sq=%d\n", sq)
				pPos.AddControlDiff(0, sq, 1)
			}
		}
	}
	//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
	pPos.MergeControlDiff()

	// 読込んだ l03.Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPos.OffsetMovesIndex = 0
	pPos.Phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPos.DoMove(pPos.Moves[i])
	}
}

// DoMove - 一手指すぜ（＾～＾）
func (pPos *Position) DoMove(move l03.Move) {
	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := l03.PIECE_TYPE_EMPTY
	cap_piece_type := l03.PIECE_TYPE_EMPTY

	from, to, pro := move.Destructure()

	if pPos.IsEmptySq(from) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", from)
	}

	var cap_src_sq l03.Square
	var cap_dst_sq = l03.SQ_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	pPos.AddControlDiffAllSlidingPiece(0, -1, from)

	// まず、打かどうかで処理を分けます
	hand := from
	var piece l03.Piece
	switch from {
	case l03.HANDSQ_R1.ToSq():
		piece = l03.PIECE_R1
	case l03.HANDSQ_B1.ToSq():
		piece = l03.PIECE_B1
	case l03.HANDSQ_G1.ToSq():
		piece = l03.PIECE_G1
	case l03.HANDSQ_S1.ToSq():
		piece = l03.PIECE_S1
	case l03.HANDSQ_N1.ToSq():
		piece = l03.PIECE_N1
	case l03.HANDSQ_L1.ToSq():
		piece = l03.PIECE_L1
	case l03.HANDSQ_P1.ToSq():
		piece = l03.PIECE_P1
	case l03.HANDSQ_R2.ToSq():
		piece = l03.PIECE_R2
	case l03.HANDSQ_B2.ToSq():
		piece = l03.PIECE_B2
	case l03.HANDSQ_G2.ToSq():
		piece = l03.PIECE_G2
	case l03.HANDSQ_S2.ToSq():
		piece = l03.PIECE_S2
	case l03.HANDSQ_N2.ToSq():
		piece = l03.PIECE_N2
	case l03.HANDSQ_L2.ToSq():
		piece = l03.PIECE_L2
	case l03.HANDSQ_P2.ToSq():
		piece = l03.PIECE_P2
	default:
		// Not hand
		hand = l03.Square(0)
	}

	if hand != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands[hand-l03.HANDSQ_ORIGIN.ToSq()] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		pPos.AddControlDiff(1, to, 1)
		mov_piece_type = l03.What(piece)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[to]
		if captured != l03.PIECE_EMPTY {
			pieceType := l03.What(captured)
			switch pieceType {
			case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR, l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB, l03.PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				pPos.AddControlDiff(1, to, -1)
			}
			cap_piece_type = l03.What(captured)
			cap_src_sq = to
		}

		// 元位置の駒を除去
		pPos.AddControlDiff(2, from, -1)

		// 行き先の駒の上書き
		if pro {
			// 駒を成りに変換します
			pPos.Board[to] = Promote(pPos.Board[from])
		} else {
			pPos.Board[to] = pPos.Board[from]
		}
		// 元位置の駒の削除pos
		mov_piece_type = l03.What(pPos.Board[to])
		pPos.Board[from] = l03.PIECE_EMPTY
		pPos.AddControlDiff(3, to, 1)

		switch captured {
		case l03.PIECE_EMPTY: // Ignored
		case l03.PIECE_K1: // Second player win
			// Lost l03.FIRST king
		case l03.PIECE_R1, l03.PIECE_PR1:
			cap_dst_sq = l03.HANDSQ_R2.ToSq()
		case l03.PIECE_B1, l03.PIECE_PB1:
			cap_dst_sq = l03.HANDSQ_B2.ToSq()
		case l03.PIECE_G1:
			cap_dst_sq = l03.HANDSQ_G2.ToSq()
		case l03.PIECE_S1, l03.PIECE_PS1:
			cap_dst_sq = l03.HANDSQ_S2.ToSq()
		case l03.PIECE_N1, l03.PIECE_PN1:
			cap_dst_sq = l03.HANDSQ_N2.ToSq()
		case l03.PIECE_L1, l03.PIECE_PL1:
			cap_dst_sq = l03.HANDSQ_L2.ToSq()
		case l03.PIECE_P1, l03.PIECE_PP1:
			cap_dst_sq = l03.HANDSQ_P2.ToSq()
		case l03.PIECE_K2: // l03.FIRST player win
			// Lost second king
		case l03.PIECE_R2, l03.PIECE_PR2:
			cap_dst_sq = l03.HANDSQ_R1.ToSq()
		case l03.PIECE_B2, l03.PIECE_PB2:
			cap_dst_sq = l03.HANDSQ_B1.ToSq()
		case l03.PIECE_G2:
			cap_dst_sq = l03.HANDSQ_G1.ToSq()
		case l03.PIECE_S2, l03.PIECE_PS2:
			cap_dst_sq = l03.HANDSQ_S1.ToSq()
		case l03.PIECE_N2, l03.PIECE_PN2:
			cap_dst_sq = l03.HANDSQ_N1.ToSq()
		case l03.PIECE_L2, l03.PIECE_PL2:
			cap_dst_sq = l03.HANDSQ_L1.ToSq()
		case l03.PIECE_P2, l03.PIECE_PP2:
			cap_dst_sq = l03.HANDSQ_P1.ToSq()
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		if cap_dst_sq != l03.SQ_EMPTY {
			pPos.CapturedList[pPos.OffsetMovesIndex] = captured
			pPos.Hands[cap_dst_sq-l03.HANDSQ_ORIGIN.ToSq()] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pPos.CapturedList[pPos.OffsetMovesIndex] = l03.PIECE_EMPTY
		}
	}

	pPos.Moves[pPos.OffsetMovesIndex] = move
	pPos.OffsetMovesIndex += 1
	prev_phase := pPos.Phase
	pPos.Phase = FlipPhase(pPos.Phase)

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []l03.PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []l03.Square{from, cap_src_sq}
	dst_sq_list := []l03.Square{to, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case l03.PIECE_TYPE_K:
			switch prev_phase {
			case l03.FIRST:
				pPos.PieceLocations[PCLOC_K1:PCLOC_K2][prev_phase-1] = dst_sq_list[j]
			case l03.SECOND:
				pPos.PieceLocations[PCLOC_K1:PCLOC_K2][prev_phase-1] = dst_sq_list[j]
			default:
				panic(App.LogNotEcho.Fatal("unknown prev_phase=%d", prev_phase))
			}
		case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR:
			for i, sq := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
				if sq == src_sq_list[j] {
					pPos.PieceLocations[PCLOC_R1:PCLOC_R2][i] = dst_sq_list[j]
				}
			}
		case l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB:
			for i, sq := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
				if sq == src_sq_list[j] {
					pPos.PieceLocations[PCLOC_B1:PCLOC_B2][i] = dst_sq_list[j]
				}
			}
		case l03.PIECE_TYPE_L, l03.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i, sq := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
				if sq == src_sq_list[j] {
					pPos.PieceLocations[PCLOC_L1:PCLOC_L4][i] = dst_sq_list[j]
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし動かした駒を除きます
	pPos.AddControlDiffAllSlidingPiece(4, 1, to)

	pPos.MergeControlDiff()
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pPos *Position) UndoMove() {

	// App.Log.Trace(pPos.Sprint())

	if pPos.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := l03.PIECE_TYPE_EMPTY
	cap_piece_type := l03.PIECE_TYPE_EMPTY

	prev_phase := pPos.Phase
	pPos.Phase = FlipPhase(pPos.Phase)

	pPos.OffsetMovesIndex -= 1
	move := pPos.Moves[pPos.OffsetMovesIndex]
	captured := pPos.CapturedList[pPos.OffsetMovesIndex]

	from, to, pro := move.Destructure()

	var cap_dst_sq l03.Square
	var cap_src_sq = l03.SQ_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	pPos.AddControlDiffAllSlidingPiece(0, -1, to)

	// 打かどうかで分けます
	switch from {
	case l03.HANDSQ_R1.ToSq(), l03.HANDSQ_B1.ToSq(), l03.HANDSQ_G1.ToSq(), l03.HANDSQ_S1.ToSq(), l03.HANDSQ_N1.ToSq(), l03.HANDSQ_L1.ToSq(), l03.HANDSQ_P1.ToSq(), l03.HANDSQ_R2.ToSq(), l03.HANDSQ_B2.ToSq(), l03.HANDSQ_G2.ToSq(), l03.HANDSQ_S2.ToSq(), l03.HANDSQ_N2.ToSq(), l03.HANDSQ_L2.ToSq(), l03.HANDSQ_P2.ToSq():
		// 打なら
		hand := from
		// 盤上から駒を除去します
		mov_piece_type = l03.What(pPos.Board[to])
		pPos.Board[to] = l03.PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands[hand-l03.HANDSQ_ORIGIN.ToSq()] += 1
		cap_dst_sq = 0
	default:
		// 打でないなら

		// 行き先の駒の除去
		mov_piece_type = l03.What(pPos.Board[to])
		pPos.AddControlDiff(1, to, -1)

		// 移動元への駒の配置
		if pro {
			// 成りを元に戻します
			pPos.Board[from] = Demote(pPos.Board[to])
		} else {
			pPos.Board[from] = pPos.Board[to]
		}

		// あれば、取った駒は駒台から下ろします
		switch captured {
		case l03.PIECE_EMPTY: // Ignored
		case l03.PIECE_K1: // Second player win
			// Lost l03.FIRST king
		case l03.PIECE_R1, l03.PIECE_PR1:
			cap_src_sq = l03.HANDSQ_R2.ToSq()
		case l03.PIECE_B1, l03.PIECE_PB1:
			cap_src_sq = l03.HANDSQ_B2.ToSq()
		case l03.PIECE_G1:
			cap_src_sq = l03.HANDSQ_G2.ToSq()
		case l03.PIECE_S1, l03.PIECE_PS1:
			cap_src_sq = l03.HANDSQ_S2.ToSq()
		case l03.PIECE_N1, l03.PIECE_PN1:
			cap_src_sq = l03.HANDSQ_N2.ToSq()
		case l03.PIECE_L1, l03.PIECE_PL1:
			cap_src_sq = l03.HANDSQ_L2.ToSq()
		case l03.PIECE_P1, l03.PIECE_PP1:
			cap_src_sq = l03.HANDSQ_P2.ToSq()
		case l03.PIECE_K2: // l03.FIRST player win
			// Lost second king
		case l03.PIECE_R2, l03.PIECE_PR2:
			cap_src_sq = l03.HANDSQ_R1.ToSq()
		case l03.PIECE_B2, l03.PIECE_PB2:
			cap_src_sq = l03.HANDSQ_B1.ToSq()
		case l03.PIECE_G2:
			cap_src_sq = l03.HANDSQ_G1.ToSq()
		case l03.PIECE_S2, l03.PIECE_PS2:
			cap_src_sq = l03.HANDSQ_S1.ToSq()
		case l03.PIECE_N2, l03.PIECE_PN2:
			cap_src_sq = l03.HANDSQ_N1.ToSq()
		case l03.PIECE_L2, l03.PIECE_PL2:
			cap_src_sq = l03.HANDSQ_L1.ToSq()
		case l03.PIECE_P2, l03.PIECE_PP2:
			cap_src_sq = l03.HANDSQ_P1.ToSq()
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		if cap_src_sq != l03.SQ_EMPTY {
			cap_dst_sq = cap_src_sq
			pPos.Hands[cap_src_sq-l03.HANDSQ_ORIGIN.ToSq()] -= 1

			// 取った駒を行き先に戻します
			cap_piece_type = l03.What(captured)
			pPos.Board[to] = captured
			pPos.AddControlDiff(2, from, 1)

			// pieceType := l03.What(captured)
			// switch pieceType {
			// case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR, l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB, l03.PIECE_TYPE_L:
			// 	// Ignored: 長い利きの駒は あとで追加するので、ここでは無視します
			// default:
			// 取った駒は盤上になかったので、ここで利きを復元させます
			pPos.AddControlDiff(3, to, 1)
			// }

		} else {
			pPos.Board[to] = l03.PIECE_EMPTY
		}
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []l03.PieceType{mov_piece_type, cap_piece_type}
	dst_sq_list := []l03.Square{to, cap_dst_sq}
	src_sq_list := []l03.Square{from, cap_src_sq}
	for j, moving_piece_type := range piece_type_list {
		switch moving_piece_type {
		case l03.PIECE_TYPE_K:
			switch prev_phase {
			case l03.FIRST:
				pPos.PieceLocations[PCLOC_K1:PCLOC_K2][prev_phase-1] = src_sq_list[j]
			case l03.SECOND:
				pPos.PieceLocations[PCLOC_K1:PCLOC_K2][prev_phase-1] = src_sq_list[j]
			default:
				panic(App.LogNotEcho.Fatal("unknown prev_phase=%d", prev_phase))
			}
		case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR:
			for i, sq := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
				if sq == dst_sq_list[j] {
					pPos.PieceLocations[PCLOC_R1:PCLOC_R2][i] = src_sq_list[j]
				}
			}
		case l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB:
			for i, sq := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
				if sq == dst_sq_list[j] {
					pPos.PieceLocations[PCLOC_B1:PCLOC_B2][i] = src_sq_list[j]
				}
			}
		case l03.PIECE_TYPE_L, l03.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i, sq := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
				if sq == dst_sq_list[j] {
					pPos.PieceLocations[PCLOC_L1:PCLOC_L4][i] = src_sq_list[j]
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	pPos.AddControlDiffAllSlidingPiece(4, 1, from)

	pPos.MergeControlDiff()
}

// AddControlDiffAllSlidingPiece - すべての長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlDiffAllSlidingPiece(layer int, sign int8, excludeFrom l03.Square) {
	for _, from := range pPos.PieceLocations[PCLOC_R1:PCLOC_R2] {
		if OnBoard(from) && from != excludeFrom {
			pPos.AddControlDiff(layer, from, sign)
		}
	}
	for _, from := range pPos.PieceLocations[PCLOC_B1:PCLOC_B2] {
		if OnBoard(from) && from != excludeFrom {
			pPos.AddControlDiff(layer, from, sign)
		}
	}
	for _, from := range pPos.PieceLocations[PCLOC_L1:PCLOC_L4] {
		if OnBoard(from) && from != excludeFrom && l03.PIECE_TYPE_PL != l03.What(pPos.Board[from]) { // 杏は除外
			pPos.AddControlDiff(layer, from, sign)
		}
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPos *Position) AddControlDiff(layer int, from l03.Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[from]
	if piece == l03.PIECE_EMPTY {
		panic(App.LogNotEcho.Fatal("LogicalError: Piece from empty square. It has no control. from=%d", from))
	}

	ph := int(l03.Who(piece)) - 1
	// fmt.Printf("Debug: ph=%d\n", ph)

	sq_list := GenMoveEnd(pPos, from)

	for _, to := range sq_list {
		// fmt.Printf("Debug: to=%d\n", to)
		// 差分の方のテーブルを更新（＾～＾）
		pPos.ControlBoardsDiff[ph][layer][to] += sign * 1
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pPos *Position) ClearControlDiff() {
	for sq := l03.Square(11); sq < 100; sq += 1 {
		if l03.File(sq) != 0 && l03.Rank(sq) != 0 {
			for layer := 0; layer < 5; layer += 1 {
				pPos.ControlBoardsDiff[0][layer][sq] = 0
				pPos.ControlBoardsDiff[1][layer][sq] = 0
			}
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pPos *Position) MergeControlDiff() {
	for sq := l03.Square(11); sq < 100; sq += 1 {
		if l03.File(sq) != 0 && l03.Rank(sq) != 0 {
			for layer := 0; layer < 5; layer += 1 {
				pPos.ControlBoards[0][sq] += pPos.ControlBoardsDiff[0][layer][sq]
				pPos.ControlBoards[1][sq] += pPos.ControlBoardsDiff[1][layer][sq]
			}
		}
	}
}

// Homo - 移動元と移動先の駒を持つプレイヤーが等しければ真。移動先が空なら偽
// 持ち駒は指定してはいけません。
func (pPos *Position) Homo(from l03.Square, to l03.Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return l03.Who(pPos.Board[from]) == l03.Who(pPos.Board[to])
}

// Hetero - 移動元と移動先の駒を持つプレイヤーが異なれば真。移動先が空マスでも真
// 持ち駒は指定してはいけません。
// Homo の逆だぜ（＾～＾）片方ありゃいいんだけど（＾～＾）
func (pPos *Position) Hetero(from l03.Square, to l03.Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return l03.Who(pPos.Board[from]) != l03.Who(pPos.Board[to])
}

// IsEmptySq - 空きマスなら真。持ち駒は偽
func (pPos *Position) IsEmptySq(sq l03.Square) bool {
	if sq > 99 {
		return false
	}
	return pPos.Board[sq] == l03.PIECE_EMPTY
}
