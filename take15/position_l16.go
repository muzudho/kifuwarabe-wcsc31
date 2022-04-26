package take15

import (
	"fmt"
	"strconv"

	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// 評価値
type Value int32

// 00～99
const BOARD_SIZE = 100

// 持ち駒
var HandIndexToPiece = [l11.HAND_SIZE]l09.Piece{
	l09.PIECE_K1, l09.PIECE_R1, l09.PIECE_B1, l09.PIECE_G1, l09.PIECE_S1, l09.PIECE_N1, l09.PIECE_L1, l09.PIECE_P1,
	l09.PIECE_K2, l09.PIECE_R2, l09.PIECE_B2, l09.PIECE_G2, l09.PIECE_S2, l09.PIECE_N2, l09.PIECE_L2, l09.PIECE_P2}

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [BOARD_SIZE]l09.Piece
	// 駒の場所
	// [0]先手玉 [1]後手玉 [2:3]飛 [4:5]角 [6:9]香
	PieceLocations [PCLOC_SIZE]Square
	// 持ち駒の数だぜ（＾～＾）玉もある（＾～＾） K, R, B, G, S, N, L, P, k, r, b, g, s, n, l, p
	Hands1 [l11.HAND_SIZE]int

	// 現局面の手番から見た駒得評価値
	MaterialValue Value
}

func NewPosition() *Position {
	var pPos = new(Position)

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

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [l11.HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return pPos
}

// SetToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) SetToStartpos() {
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
	pPos.PieceLocations = [PCLOC_SIZE]Square{59, 51, 28, 82, 22, 88, 11, 19, 91, 99}

	// 持ち駒の数
	pPos.Hands1 = [l11.HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (pPos *Position) GetPieceLocation(index int) Square {
	return pPos.PieceLocations[index]
}

// ClearBoard - 駒を置いていな状態でリセットします
func (pPos *Position) ClearBoard() {
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

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [l11.HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
			from = hand_sq + l11.HAND_TYPE_SIZE
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
				return *new(Move), fmt.Errorf("fatal: Unknown count='%c'", count)
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

// From - 筋と段からマス番号を作成します
func SquareFrom(file Square, rank Square) Square {
	return Square(file*10 + rank)
}

// OnHands - 持ち駒なら真
func OnHands(sq Square) bool {
	return SQ_HAND_START <= sq && sq < SQ_HAND_END
}

// OnBoard - 盤上なら真
func OnBoard(sq Square) bool {
	return 10 < sq && sq < 100 && File(sq) != 0 && Rank(sq) != 0
}
