package take10

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// 00～99
const BOARD_SIZE = 100

// position sfen の盤のスペース数に使われますN
var OneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// FlipPhase - 先後を反転します
func FlipPhase(phase Phase) Phase {
	return phase%2 + 1
}

// From - 筋と段からマス番号を作成します
func SquareFrom(file Square, rank Square) Square {
	return Square(file*10 + rank)
}

// OnHands - 持ち駒なら真
func OnHands(sq Square) bool {
	return 99 < sq && sq < 114
}

// OnBoard - 盤上なら真
func OnBoard(sq Square) bool {
	return 10 < sq && sq < 100 && File(sq) != 0 && Rank(sq) != 0
}

// PieceFromPhPt - 駒作成。空マスは作れません
func PieceFromPhPt(phase Phase, pieceType PieceType) l09.Piece {
	switch phase {
	case FIRST:
		switch pieceType {
		case PIECE_TYPE_K:
			return l09.PIECE_K1
		case PIECE_TYPE_R:
			return l09.PIECE_R1
		case PIECE_TYPE_B:
			return l09.PIECE_B1
		case PIECE_TYPE_G:
			return l09.PIECE_G1
		case PIECE_TYPE_S:
			return l09.PIECE_S1
		case PIECE_TYPE_N:
			return l09.PIECE_N1
		case PIECE_TYPE_L:
			return l09.PIECE_L1
		case PIECE_TYPE_P:
			return l09.PIECE_P1
		case PIECE_TYPE_PR:
			return l09.PIECE_PR1
		case PIECE_TYPE_PB:
			return l09.PIECE_PB1
		case PIECE_TYPE_PS:
			return l09.PIECE_PS1
		case PIECE_TYPE_PN:
			return l09.PIECE_PN1
		case PIECE_TYPE_PL:
			return l09.PIECE_PL1
		case PIECE_TYPE_PP:
			return l09.PIECE_PP1
		default:
			panic(fmt.Errorf("unknown piece type=%d", pieceType))
		}
	case SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return l09.PIECE_K2
		case PIECE_TYPE_R:
			return l09.PIECE_R2
		case PIECE_TYPE_B:
			return l09.PIECE_B2
		case PIECE_TYPE_G:
			return l09.PIECE_G2
		case PIECE_TYPE_S:
			return l09.PIECE_S2
		case PIECE_TYPE_N:
			return l09.PIECE_N2
		case PIECE_TYPE_L:
			return l09.PIECE_L2
		case PIECE_TYPE_P:
			return l09.PIECE_P2
		case PIECE_TYPE_PR:
			return l09.PIECE_PR2
		case PIECE_TYPE_PB:
			return l09.PIECE_PB2
		case PIECE_TYPE_PS:
			return l09.PIECE_PS2
		case PIECE_TYPE_PN:
			return l09.PIECE_PN2
		case PIECE_TYPE_PL:
			return l09.PIECE_PL2
		case PIECE_TYPE_PP:
			return l09.PIECE_PP2
		default:
			panic(fmt.Errorf("unknown piece type=%d", pieceType))
		}
	default:
		panic(fmt.Errorf("unknown phase=%d", phase))
	}
}

var HandPieceMap = [14]l09.Piece{
	l09.PIECE_R1, l09.PIECE_B1, l09.PIECE_G1, l09.PIECE_S1, l09.PIECE_N1, l09.PIECE_L1, l09.PIECE_P1,
	l09.PIECE_R2, l09.PIECE_B2, l09.PIECE_G2, l09.PIECE_S2, l09.PIECE_N2, l09.PIECE_L2, l09.PIECE_P2}

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [l09.BOARD_SIZE]l09.Piece
	// [0]先手 [1]後手
	kingLocations [2]Square
	// 飛車の場所。長い利きを消すために必要（＾～＾）
	RookLocations [2]Square
	// 角の場所。長い利きを消すために必要（＾～＾）
	BishopLocations [2]Square
	// 香の場所。長い利きを消すために必要（＾～＾）
	LanceLocations [4]Square
	// マスへの利き数、または差分が入っています。デバッグ目的で無駄に分けてるんだけどな（＾～＾）
	// 利きテーブル [0]先手 [1]後手
	// [0] 利き
	// [1] 飛の利き引く(差分)
	// [2] 角の利き引く(差分)
	// [3] 香の利き引く(差分)
	// [4] ムーブ用(差分)
	// [5] ムーブ用(差分)
	// [6] ムーブ用(差分)
	// [7] 香の利き戻す(差分)
	// [8] 角の利き戻す(差分)
	// [9] 飛の利き戻す(差分)
	// [10] テスト用
	// [11] テスト用
	// [12] テスト用(再計算)
	ControlBoards [2][CONTROL_LAYER_ALL_SIZE][BOARD_SIZE]int8

	// 持ち駒の数だぜ（＾～＾） R, B, G, S, N, L, P, r, b, g, s, n, l, p
	Hands []int
	// 先手が1、後手が2（＾～＾）
	phase Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [MOVES_SIZE]l09.Piece
}

func NewPosition() *Position {
	var ins = new(Position)
	ins.resetToZero()
	return ins
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPos *Position) FlipPhase() {
	pPos.phase = FlipPhase(pPos.phase)
}

// GetPhase - フェーズ
func (pPos *Position) GetPhase() Phase {
	return pPos.phase
}

func (pPos *Position) GetKingLocations() (Square, Square) {
	return pPos.kingLocations[0], pPos.kingLocations[1]
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPos *Position) resetToZero() {
	// 筋、段のラベルだけ入れとくぜ（＾～＾）
	pPos.Board = [BOARD_SIZE]l09.Piece{
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
	}
	pPos.ControlBoards = [2][CONTROL_LAYER_ALL_SIZE][BOARD_SIZE]int8{{
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
	pPos.kingLocations = [2]Square{SQUARE_EMPTY, SQUARE_EMPTY}
	pPos.RookLocations = [2]Square{SQUARE_EMPTY, SQUARE_EMPTY}
	pPos.BishopLocations = [2]Square{SQUARE_EMPTY, SQUARE_EMPTY}
	pPos.LanceLocations = [4]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands = []int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	// 先手の局面
	pPos.phase = FIRST
	// 何手目か
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0
	// 指し手のリスト
	pPos.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPos.CapturedList = [MOVES_SIZE]l09.Piece{}
}

// setToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) setToStartpos() {
	// 初期局面にします
	pPos.Board = [BOARD_SIZE]l09.Piece{
		l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY,
		l09.PIECE_EMPTY, l09.PIECE_L2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_L1,
		l09.PIECE_EMPTY, l09.PIECE_N2, l09.PIECE_B2, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_R1, l09.PIECE_N1,
		l09.PIECE_EMPTY, l09.PIECE_S2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_S1,
		l09.PIECE_EMPTY, l09.PIECE_G2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_G1,
		l09.PIECE_EMPTY, l09.PIECE_K2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_K1,
		l09.PIECE_EMPTY, l09.PIECE_G2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_G1,
		l09.PIECE_EMPTY, l09.PIECE_S2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_S1,
		l09.PIECE_EMPTY, l09.PIECE_N2, l09.PIECE_R2, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_B1, l09.PIECE_N1,
		l09.PIECE_EMPTY, l09.PIECE_L2, l09.PIECE_EMPTY, l09.PIECE_P2, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_EMPTY, l09.PIECE_P1, l09.PIECE_EMPTY, l09.PIECE_L1,
	}
	pPos.kingLocations = [2]Square{59, 51}
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
				pPos.Board[file*10+rank] = l09.FromStringToPiece(string(pc))
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = l09.PIECE_EMPTY
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
					pPos.Board[file*10+rank] = l09.FromStringToPiece("+" + string(pc2))
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.kingLocations[0] = Square((file+1)*10 + rank)
			case 'k':
				pPos.kingLocations[1] = Square((file+1)*10 + rank)
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
			pPos.phase = FIRST
			i += 1
		case 'w':
			pPos.phase = SECOND
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

			// R なら竜1枚
			// R2 なら竜2枚
			// P10 なら歩10枚。数が2桁になるのは歩だけ（＾～＾）
			// {アルファベット１文字}{数字1～2文字} になっている
			// アルファベットまたは半角スペースを見つけた時点で、以前の取り込み分が確定する
			var hand_index int = 999 //存在しない数
			var number = 0

		HandLoop:
			for {
				var piece = command[i]

				if unicode.IsLetter(rune(piece)) || piece == ' ' {

					if hand_index == 999 {
						// ループの１週目は無視します

					} else {
						// 数字が書いてなかったら１個
						if number == 0 {
							number = 1
						}

						pPos.Hands[hand_index] = number
						number = 0

						// 長い利きの駒は位置を覚えておくぜ（＾～＾）
						switch hand_index {
						case HAND_R1, HAND_R2:
							for i, sq := range pPos.RookLocations {
								if sq == SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.RookLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						case HAND_B1, HAND_B2:
							for i, sq := range pPos.BishopLocations {
								if sq == SQUARE_EMPTY {
									pPos.BishopLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						case HAND_L1, HAND_L2:
							for i, sq := range pPos.LanceLocations {
								if sq == SQUARE_EMPTY {
									pPos.LanceLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						}
					}
					i += 1

					switch piece {
					case 'R':
						hand_index = HAND_R1
					case 'B':
						hand_index = HAND_B1
					case 'G':
						hand_index = HAND_G1
					case 'S':
						hand_index = HAND_S1
					case 'N':
						hand_index = HAND_N1
					case 'L':
						hand_index = HAND_L1
					case 'P':
						hand_index = HAND_P1
					case 'r':
						hand_index = HAND_R2
					case 'b':
						hand_index = HAND_B2
					case 'g':
						hand_index = HAND_G2
					case 's':
						hand_index = HAND_S2
					case 'n':
						hand_index = HAND_N2
					case 'l':
						hand_index = HAND_L2
					case 'p':
						hand_index = HAND_P2
					case ' ':
						// ループを抜けます
						break HandLoop
					default:
						panic(fmt.Errorf("fatal: unknown piece=%c", piece))
					}
				} else if unicode.IsNumber(rune(piece)) {
					switch piece {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						num, err := strconv.Atoi(string(piece))
						if err != nil {
							panic(err)
						}
						i += 1
						number *= 10
						number += num
					default:
						panic(fmt.Errorf("fatal: unknown number character=%c", piece))
					}

				} else {
					panic(fmt.Errorf("fatal: unknown piece=%c", piece))
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

	start_phase := pPos.GetPhase()
	if strings.HasPrefix(command[i:], "moves") {
		i += 5

		// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
		for i < len {
			if command[i] != ' ' {
				break
			}
			i += 1

			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			var move, err = ParseMove(command, &i, pPos.GetPhase())
			if err != nil {
				fmt.Println(err)
				fmt.Println(Sprint(pPos))
				panic(err)
			}
			pPos.Moves[pPos.OffsetMovesIndex] = move
			pPos.OffsetMovesIndex += 1
			pPos.FlipPhase()
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
				// あとですぐクリアーするので、どのレイヤー使ってても関係ないんで、仮で PUTレイヤーを使っているぜ（＾～＾）
				pPos.AddControlDiff(CONTROL_LAYER_DIFF_PUT, sq, 1)
			}
		}
	}
	//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
	pPos.MergeControlDiff()

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPos.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPos.OffsetMovesIndex = 0
	pPos.phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPos.DoMove(pPos.Moves[i])
	}
}

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase Phase) (Move, error) {
	var len = len(command)
	var hand_sq = SQUARE_EMPTY

	var from Square
	var to Square
	var pro = false

	// file
	switch ch := command[*i]; ch {
	case 'R':
		hand_sq = SQ_R1
	case 'B':
		hand_sq = SQ_B1
	case 'G':
		hand_sq = SQ_G1
	case 'S':
		hand_sq = SQ_S1
	case 'N':
		hand_sq = SQ_N1
	case 'L':
		hand_sq = SQ_L1
	case 'P':
		hand_sq = SQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand_sq != SQUARE_EMPTY {
		*i += 1
		switch phase {
		case FIRST:
			from = hand_sq
		case SECOND:
			from = hand_sq + HAND_TYPE_SIZE
		default:
			return *new(Move), fmt.Errorf("fatal: unknown phase=%d", phase)
		}

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
				from = sq
			} else if count == 1 {
				to = sq
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
		pro = true
	}

	return NewMove(from, to, pro), nil
}

// DoMove - 一手指すぜ（＾～＾）
func (pPos *Position) DoMove(move Move) {
	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	from, to, pro := move.Destructure()

	if pPos.IsEmptySq(from) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", from)
	}

	var cap_src_sq Square
	var cap_dst_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	pPos.AddControlRook(CONTROL_LAYER_DIFF_ROOK_OFF, -1, from)
	pPos.AddControlBishop(CONTROL_LAYER_DIFF_BISHOP_OFF, -1, from)
	pPos.AddControlLance(CONTROL_LAYER_DIFF_LANCE_OFF, -1, from)

	// まず、打かどうかで処理を分けます
	sq_hand := from
	var piece l09.Piece
	switch from {
	case SQ_R1:
		piece = l09.PIECE_R1
	case SQ_B1:
		piece = l09.PIECE_B1
	case SQ_G1:
		piece = l09.PIECE_G1
	case SQ_S1:
		piece = l09.PIECE_S1
	case SQ_N1:
		piece = l09.PIECE_N1
	case SQ_L1:
		piece = l09.PIECE_L1
	case SQ_P1:
		piece = l09.PIECE_P1
	case SQ_R2:
		piece = l09.PIECE_R2
	case SQ_B2:
		piece = l09.PIECE_B2
	case SQ_G2:
		piece = l09.PIECE_G2
	case SQ_S2:
		piece = l09.PIECE_S2
	case SQ_N2:
		piece = l09.PIECE_N2
	case SQ_L2:
		piece = l09.PIECE_L2
	case SQ_P2:
		piece = l09.PIECE_P2
	default:
		// Not hand
		sq_hand = SQUARE_EMPTY
	}

	if sq_hand != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands[sq_hand-SQ_HAND_START] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		pPos.AddControlDiff(CONTROL_LAYER_DIFF_PUT, to, 1)
		mov_piece_type = What(piece)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[to]
		if captured != l09.PIECE_EMPTY {
			pieceType := What(captured)
			switch pieceType {
			case PIECE_TYPE_R, PIECE_TYPE_PR, PIECE_TYPE_B, PIECE_TYPE_PB, PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				pPos.AddControlDiff(CONTROL_LAYER_DIFF_CAPTURED, to, -1)
			}
			cap_piece_type = What(captured)
			cap_src_sq = to
		}

		// 元位置の駒の利きを除去
		pPos.AddControlDiff(CONTROL_LAYER_DIFF_REMOVE, from, -1)

		// 行き先の駒の上書き
		if pro {
			// 駒を成りに変換します
			pPos.Board[to] = l09.Promote(pPos.Board[from])
		} else {
			pPos.Board[to] = pPos.Board[from]
		}
		mov_piece_type = What(pPos.Board[to])
		// 元位置の駒を削除してから、移動先の駒の利きを追加
		pPos.Board[from] = l09.PIECE_EMPTY
		pPos.AddControlDiff(CONTROL_LAYER_DIFF_PUT, to, 1)

		switch captured {
		case l09.PIECE_EMPTY: // Ignored
		case l09.PIECE_K1: // Second player win
			// Lost first king
		case l09.PIECE_R1, l09.PIECE_PR1:
			cap_dst_sq = SQ_R2
		case l09.PIECE_B1, l09.PIECE_PB1:
			cap_dst_sq = SQ_B2
		case l09.PIECE_G1:
			cap_dst_sq = SQ_G2
		case l09.PIECE_S1, l09.PIECE_PS1:
			cap_dst_sq = SQ_S2
		case l09.PIECE_N1, l09.PIECE_PN1:
			cap_dst_sq = SQ_N2
		case l09.PIECE_L1, l09.PIECE_PL1:
			cap_dst_sq = SQ_L2
		case l09.PIECE_P1, l09.PIECE_PP1:
			cap_dst_sq = SQ_P2
		case l09.PIECE_K2: // First player win
			// Lost second king
		case l09.PIECE_R2, l09.PIECE_PR2:
			cap_dst_sq = SQ_R1
		case l09.PIECE_B2, l09.PIECE_PB2:
			cap_dst_sq = SQ_B1
		case l09.PIECE_G2:
			cap_dst_sq = SQ_G1
		case l09.PIECE_S2, l09.PIECE_PS2:
			cap_dst_sq = SQ_S1
		case l09.PIECE_N2, l09.PIECE_PN2:
			cap_dst_sq = SQ_N1
		case l09.PIECE_L2, l09.PIECE_PL2:
			cap_dst_sq = SQ_L1
		case l09.PIECE_P2, l09.PIECE_PP2:
			cap_dst_sq = SQ_P1
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		if cap_dst_sq != SQUARE_EMPTY {
			pPos.CapturedList[pPos.OffsetMovesIndex] = captured
			pPos.Hands[cap_dst_sq-SQ_HAND_START] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pPos.CapturedList[pPos.OffsetMovesIndex] = l09.PIECE_EMPTY
		}
	}

	// DoMoveでフェーズを１つ進めます
	pPos.Moves[pPos.OffsetMovesIndex] = move
	pPos.OffsetMovesIndex += 1
	prev_phase := pPos.GetPhase()
	pPos.FlipPhase()

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []Square{from, cap_src_sq}
	dst_sq_list := []Square{to, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case PIECE_TYPE_K:
			switch prev_phase {
			case FIRST:
				pPos.kingLocations[0] = dst_sq_list[j]
			case SECOND:
				pPos.kingLocations[1] = dst_sq_list[j]
			default:
				panic(fmt.Errorf("unknown prev_phase=%d", prev_phase))
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
	pPos.AddControlLance(CONTROL_LAYER_DIFF_LANCE_ON, 1, to)
	pPos.AddControlBishop(CONTROL_LAYER_DIFF_BISHOP_ON, 1, to)
	pPos.AddControlRook(CONTROL_LAYER_DIFF_ROOK_ON, 1, to)

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
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	// 先に 手目 を１つ戻すぜ（＾～＾）UndoMoveでフェーズもひっくり返すぜ（＾～＾）
	pPos.OffsetMovesIndex -= 1
	move := pPos.Moves[pPos.OffsetMovesIndex]
	// next_phase := pPos.GetPhase()
	pPos.FlipPhase()

	// 取った駒
	captured := pPos.CapturedList[pPos.OffsetMovesIndex]

	from, to, pro := move.Destructure()

	var cap_dst_sq Square
	var cap_src_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPos.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	pPos.AddControlRook(CONTROL_LAYER_DIFF_ROOK_ON, -1, to)
	pPos.AddControlBishop(CONTROL_LAYER_DIFF_BISHOP_ON, -1, to)
	pPos.AddControlLance(CONTROL_LAYER_DIFF_LANCE_ON, -1, to)

	// 打かどうかで分けます
	switch from {
	case SQ_R1, SQ_B1, SQ_G1, SQ_S1, SQ_N1, SQ_L1, SQ_P1, SQ_R2, SQ_B2, SQ_G2, SQ_S2, SQ_N2, SQ_L2, SQ_P2:
		// 打なら
		hand := from
		// 行き先から駒を除去します
		mov_piece_type = What(pPos.Board[to])
		pPos.AddControlDiff(CONTROL_LAYER_DIFF_PUT, to, -1)
		pPos.Board[to] = l09.PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands[hand-SQ_HAND_START] += 1
		cap_dst_sq = 0
	default:
		// 打でないなら

		// 行き先に進んでいた自駒の利きの除去
		mov_piece_type = What(pPos.Board[to])
		pPos.AddControlDiff(CONTROL_LAYER_DIFF_PUT, to, -1)

		// 自駒を移動元へ戻します
		if pro {
			// 成りを元に戻します
			pPos.Board[from] = l09.Demote(pPos.Board[to])
		} else {
			pPos.Board[from] = pPos.Board[to]
		}

		// あれば、取った駒は駒台から下ろします
		switch captured {
		case l09.PIECE_EMPTY: // Ignored
		case l09.PIECE_K1: // Second player win
			// Lost first king
		case l09.PIECE_R1, l09.PIECE_PR1:
			cap_src_sq = SQ_R2
		case l09.PIECE_B1, l09.PIECE_PB1:
			cap_src_sq = SQ_B2
		case l09.PIECE_G1:
			cap_src_sq = SQ_G2
		case l09.PIECE_S1, l09.PIECE_PS1:
			cap_src_sq = SQ_S2
		case l09.PIECE_N1, l09.PIECE_PN1:
			cap_src_sq = SQ_N2
		case l09.PIECE_L1, l09.PIECE_PL1:
			cap_src_sq = SQ_L2
		case l09.PIECE_P1, l09.PIECE_PP1:
			cap_src_sq = SQ_P2
		case l09.PIECE_K2: // First player win
			// Lost second king
		case l09.PIECE_R2, l09.PIECE_PR2:
			cap_src_sq = SQ_R1
		case l09.PIECE_B2, l09.PIECE_PB2:
			cap_src_sq = SQ_B1
		case l09.PIECE_G2:
			cap_src_sq = SQ_G1
		case l09.PIECE_S2, l09.PIECE_PS2:
			cap_src_sq = SQ_S1
		case l09.PIECE_N2, l09.PIECE_PN2:
			cap_src_sq = SQ_N1
		case l09.PIECE_L2, l09.PIECE_PL2:
			cap_src_sq = SQ_L1
		case l09.PIECE_P2, l09.PIECE_PP2:
			cap_src_sq = SQ_P1
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		if cap_src_sq != SQUARE_EMPTY {
			cap_dst_sq = cap_src_sq
			pPos.Hands[cap_src_sq-SQ_HAND_START] -= 1

			// 取っていた駒を行き先に戻します
			cap_piece_type = What(captured)
			pPos.Board[to] = captured

			// pieceType := What(captured)
			// switch pieceType {
			// case PIECE_TYPE_R, PIECE_TYPE_PR, PIECE_TYPE_B, PIECE_TYPE_PB, PIECE_TYPE_L:
			// 	// Ignored: 長い利きの駒は あとで追加するので、ここでは無視します
			// default:
			// 取った駒は盤上になかったので、ここで利きを復元させます
			// 行き先にある取られていた駒の利きの復元
			pPos.AddControlDiff(CONTROL_LAYER_DIFF_CAPTURED, to, 1)
			// }
		} else {
			pPos.Board[to] = l09.PIECE_EMPTY
		}

		// 元の場所に戻した自駒の利きを復元します
		pPos.AddControlDiff(CONTROL_LAYER_DIFF_REMOVE, from, 1)
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []PieceType{mov_piece_type, cap_piece_type}
	dst_sq_list := []Square{to, cap_dst_sq}
	src_sq_list := []Square{from, cap_src_sq}
	for j, moving_piece_type := range piece_type_list {
		switch moving_piece_type {
		case PIECE_TYPE_K:
			switch pPos.phase { // next_phase
			case FIRST:
				pPos.kingLocations[0] = src_sq_list[j]
			case SECOND:
				pPos.kingLocations[1] = src_sq_list[j]
			default:
				panic(fmt.Errorf("unknown pPos.phase=%d", pPos.phase))
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
	// アンドゥなので逆さになっているぜ（＾～＾）
	pPos.AddControlLance(CONTROL_LAYER_DIFF_LANCE_OFF, 1, from)
	pPos.AddControlBishop(CONTROL_LAYER_DIFF_BISHOP_OFF, 1, from)
	pPos.AddControlRook(CONTROL_LAYER_DIFF_ROOK_OFF, 1, from)

	pPos.MergeControlDiff()
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
	return pPos.Board[sq] == l09.PIECE_EMPTY
}
