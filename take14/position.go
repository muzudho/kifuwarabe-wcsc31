package take14

import (
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// Position - 局面
// TODO 利きボードも含めたい
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [l11.BOARD_SIZE]l09.Piece
	// 駒の場所
	// [0]先手玉 [1]後手玉 [2:3]飛 [4:5]角 [6:9]香
	PieceLocations [l11.PCLOC_SIZE]Square
	// 持ち駒の数だぜ（＾～＾）玉もある（＾～＾） K, R, B, G, S, N, L, P, k, r, b, g, s, n, l, p
	Hands1 [HAND_SIZE]int

	// 先手から見た駒得評価値
	MaterialValue int16
}

func NewPosition() *Position {
	var pPos = new(Position)

	pPos.Board = [BOARD_SIZE]l09.Piece{
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
	}

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return pPos
}

// setToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) setToStartpos() {
	// 初期局面にします
	pPos.Board = [BOARD_SIZE]l09.Piece{
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_L2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_L1,
		l11.PIECE_EMPTY, l11.PIECE_N2, l11.PIECE_B2, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_R1, l11.PIECE_N1,
		l11.PIECE_EMPTY, l11.PIECE_S2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_S1,
		l11.PIECE_EMPTY, l11.PIECE_G2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_G1,
		l11.PIECE_EMPTY, l11.PIECE_K2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_K1,
		l11.PIECE_EMPTY, l11.PIECE_G2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_G1,
		l11.PIECE_EMPTY, l11.PIECE_S2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_S1,
		l11.PIECE_EMPTY, l11.PIECE_N2, l11.PIECE_R2, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_B1, l11.PIECE_N1,
		l11.PIECE_EMPTY, l11.PIECE_L2, l11.PIECE_EMPTY, l11.PIECE_P2, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_P1, l11.PIECE_EMPTY, l11.PIECE_L1,
	}
	pPos.PieceLocations = [PCLOC_SIZE]Square{59, 51, 28, 82, 22, 88, 11, 19, 91, 99}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (pPos *Position) GetPieceLocation(index int) Square {
	return pPos.PieceLocations[index]
}

// clearBoard - 駒を置いていな状態でリセットします
func (pPos *Position) clearBoard() {
	pPos.Board = [BOARD_SIZE]l09.Piece{
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
		l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY, l11.PIECE_EMPTY,
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
	return pPos.Board[sq] == l11.PIECE_EMPTY
}
