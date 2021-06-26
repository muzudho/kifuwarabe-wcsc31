package take16position

import "fmt"

// 1:先手 2:後手
type Phase byte

// マス番号 00～99,100～113
type Square uint32

// 先後付きの駒
type Piece uint8

// 評価値
type Value int32

// 00～99
const BOARD_SIZE = 100

// Piece location
const (
	PCLOC_K1 = iota
	PCLOC_K2
	PCLOC_R1
	PCLOC_R2
	PCLOC_B1
	PCLOC_B2
	PCLOC_L1
	PCLOC_L2
	PCLOC_L3
	PCLOC_L4
	PCLOC_START = 0
	PCLOC_END   = 10 //この数を含まない
	PCLOC_SIZE  = 10
)

const (
	// 持ち駒を打つ 0～15 (Index)
	HAND_K1_IDX    = 0
	HAND_R1_IDX    = 1 // 先手飛打
	HAND_B1_IDX    = 2
	HAND_G1_IDX    = 3
	HAND_S1_IDX    = 4
	HAND_N1_IDX    = 5
	HAND_L1_IDX    = 6
	HAND_P1_IDX    = 7
	HAND_K2_IDX    = 8
	HAND_R2_IDX    = 9
	HAND_B2_IDX    = 10
	HAND_G2_IDX    = 11
	HAND_S2_IDX    = 12
	HAND_N2_IDX    = 13
	HAND_L2_IDX    = 14
	HAND_P2_IDX    = 15
	HAND_SIZE      = 16
	HAND_TYPE_SIZE = 8
	HAND_IDX_START = HAND_K1_IDX
	HAND_IDX_END   = HAND_SIZE // この数を含まない
)

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

// 駒
const (
	PIECE_EMPTY = iota
	PIECE_K1
	PIECE_R1
	PIECE_B1
	PIECE_G1
	PIECE_S1
	PIECE_N1
	PIECE_L1
	PIECE_P1
	PIECE_PR1
	PIECE_PB1
	PIECE_PS1
	PIECE_PN1
	PIECE_PL1
	PIECE_PP1
	PIECE_K2
	PIECE_R2
	PIECE_B2
	PIECE_G2
	PIECE_S2
	PIECE_N2
	PIECE_L2
	PIECE_P2
	PIECE_PR2
	PIECE_PB2
	PIECE_PS2
	PIECE_PN2
	PIECE_PL2
	PIECE_PP2
)

const (
	// 空マス
	ZEROTH = Phase(0)
	// 先手
	FIRST = Phase(1)
	// 後手
	SECOND = Phase(2)
)

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}

// Who - 駒が先手か後手か空升かを返します
func Who(piece Piece) Phase {
	switch piece {
	case PIECE_EMPTY: // 空きマス
		return ZEROTH
	case PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1, PIECE_PR1, PIECE_PB1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1:
		return FIRST
	case PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2, PIECE_PR2, PIECE_PB2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		return SECOND
	default:
		panic(fmt.Errorf("Error: Unknown piece=[%d]", piece))
	}
}

// ToCode - 文字列
func (pc Piece) ToCode() string {
	switch pc {
	case PIECE_EMPTY:
		return ""
	case PIECE_K1:
		return "K"
	case PIECE_R1:
		return "R"
	case PIECE_B1:
		return "B"
	case PIECE_G1:
		return "G"
	case PIECE_S1:
		return "S"
	case PIECE_N1:
		return "N"
	case PIECE_L1:
		return "L"
	case PIECE_P1:
		return "P"
	case PIECE_PR1:
		return "+R"
	case PIECE_PB1:
		return "+B"
	case PIECE_PS1:
		return "+S"
	case PIECE_PN1:
		return "+N"
	case PIECE_PL1:
		return "+L"
	case PIECE_PP1:
		return "+P"
	case PIECE_K2:
		return "k"
	case PIECE_R2:
		return "r"
	case PIECE_B2:
		return "b"
	case PIECE_G2:
		return "g"
	case PIECE_S2:
		return "s"
	case PIECE_N2:
		return "n"
	case PIECE_L2:
		return "l"
	case PIECE_P2:
		return "p"
	case PIECE_PR2:
		return "+r"
	case PIECE_PB2:
		return "+b"
	case PIECE_PS2:
		return "+s"
	case PIECE_PN2:
		return "+n"
	case PIECE_PL2:
		return "+l"
	case PIECE_PP2:
		return "+p"
	default:
		panic(fmt.Errorf("Unknown piece=%d", pc))
	}
}

// Position - 局面
// TODO 利きボードも含めたい
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [BOARD_SIZE]Piece
	// 駒の場所
	// [0]先手玉 [1]後手玉 [2:3]飛 [4:5]角 [6:9]香
	PieceLocations [PCLOC_SIZE]Square
	// 持ち駒の数だぜ（＾～＾）玉もある（＾～＾） K, R, B, G, S, N, L, P, k, r, b, g, s, n, l, p
	Hands1 [HAND_SIZE]int

	// 現局面の手番から見た駒得評価値
	MaterialValue Value
}

func NewPosition() *Position {
	var pPos = new(Position)

	pPos.Board = [BOARD_SIZE]Piece{
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
	}

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return pPos
}

// SetToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) SetToStartpos() {
	// 初期局面にします
	pPos.Board = [BOARD_SIZE]Piece{
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
	pPos.PieceLocations = [PCLOC_SIZE]Square{59, 51, 28, 82, 22, 88, 11, 19, 91, 99}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (pPos *Position) GetPieceLocation(index int) Square {
	return pPos.PieceLocations[index]
}

// ClearBoard - 駒を置いていな状態でリセットします
func (pPos *Position) ClearBoard() {
	pPos.Board = [BOARD_SIZE]Piece{
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
	}

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
