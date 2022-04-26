// 利きボード
package take11

import (
	"fmt"

	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// 利きテーブル・インデックス型
type ControlLayerT int

const (
	CONTROL_LAYER_SUM                = ControlLayerT(0)
	CONTROL_LAYER_DIFF_ROOK_OFF      = ControlLayerT(1)
	CONTROL_LAYER_DIFF_BISHOP_OFF    = ControlLayerT(2)
	CONTROL_LAYER_DIFF_LANCE_OFF     = ControlLayerT(3)
	CONTROL_LAYER_DIFF_PUT           = ControlLayerT(4) // 打とか指すとか
	CONTROL_LAYER_DIFF_REMOVE        = ControlLayerT(5)
	CONTROL_LAYER_DIFF_CAPTURED      = ControlLayerT(6)
	CONTROL_LAYER_DIFF_LANCE_ON      = ControlLayerT(7)
	CONTROL_LAYER_DIFF_BISHOP_ON     = ControlLayerT(8)
	CONTROL_LAYER_DIFF_ROOK_ON       = ControlLayerT(9)
	CONTROL_LAYER_TEST_COPY          = ControlLayerT(10) // テスト用
	CONTROL_LAYER_TEST_ERROR         = ControlLayerT(11) // テスト用
	CONTROL_LAYER_TEST_RECALCULATION = ControlLayerT(12) // テスト用 再計算
	CONTROL_LAYER_DIFF_START         = ControlLayerT(1)
	CONTROL_LAYER_DIFF_END           = ControlLayerT(10) // この数を含まない。テスト用も含まない
	CONTROL_LAYER_ALL_SIZE           = 13                // この数を含まない
)

// GetControlLayerName - 利きボードのレイヤーの名前
func GetControlLayerName(c ControlLayerT) string {
	switch c {
	case CONTROL_LAYER_SUM:
		return "Sum"
	case CONTROL_LAYER_DIFF_ROOK_OFF:
		return "RookOff"
	case CONTROL_LAYER_DIFF_BISHOP_OFF:
		return "BishopOff"
	case CONTROL_LAYER_DIFF_LANCE_OFF:
		return "LanceOff"
	case CONTROL_LAYER_DIFF_PUT:
		return "Put"
	case CONTROL_LAYER_DIFF_REMOVE:
		return "Remove"
	case CONTROL_LAYER_DIFF_CAPTURED:
		return "Captured"
	case CONTROL_LAYER_DIFF_LANCE_ON:
		return "LanceOn"
	case CONTROL_LAYER_DIFF_BISHOP_ON:
		return "BishopOn"
	case CONTROL_LAYER_DIFF_ROOK_ON:
		return "RookOn"
	case CONTROL_LAYER_TEST_COPY:
		return "TestCopy"
	case CONTROL_LAYER_TEST_ERROR:
		return "TestError"
	case CONTROL_LAYER_TEST_RECALCULATION:
		return "TestRecalc"
	default:
		panic(fmt.Errorf("unknown control layer=%d", c))
	}
}

// AddControlRook - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlRook(pPos *Position, c ControlLayerT, sign int8, excludeFrom l04.Square) {
	for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 飛落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPosSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlBishop - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlBishop(pPos *Position, c ControlLayerT, sign int8, excludeFrom l04.Square) {
	for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 角落ちも考えて 空マスは除外
			from != excludeFrom { // 除外マスは除外
			pPosSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlLance - 長い利きの駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlLance(pPos *Position, c ControlLayerT, sign int8, excludeFrom l04.Square) {
	for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
		from := pPos.PieceLocations[i]
		if !OnHands(from) && // 持ち駒は除外
			!pPos.IsEmptySq(from) && // 香落ちも考えて 空マスは除外
			from != excludeFrom && // 除外マスは除外
			PIECE_TYPE_PL != What(pPos.Board[from]) { // 杏は除外
			pPosSys.AddControlDiff(pPos, c, from, sign)
		}
	}
}

// AddControlDiff - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pPosSys *PositionSystem) AddControlDiff(pPos *Position, c ControlLayerT, from l04.Square, sign int8) {
	if from > 99 {
		// 持ち駒は無視します
		return
	}

	piece := pPos.Board[from]
	if piece == l09.PIECE_EMPTY {
		panic(fmt.Errorf("LogicalError: Piece from empty square. It has no control. from=%d", from))
	}

	ph := int(Who(piece)) - 1
	// fmt.Printf("Debug: ph=%d\n", ph)

	sq_list := GenMoveEnd(pPos, from)

	for _, to := range sq_list {
		// fmt.Printf("Debug: ph=%d c=%d to=%d\n", ph, c, to)
		// 差分の方のテーブルを更新（＾～＾）
		pPosSys.ControlBoards[ph][c][to] += sign * 1
	}
}

// ClearControlDiff - 利きの差分テーブルをクリアーするぜ（＾～＾）
func (pPosSys *PositionSystem) ClearControlDiff() {
	// c=0 を除く
	for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
		pPosSys.ClearControlLayer(c)
	}
}

func (pPosSys *PositionSystem) ClearControlLayer(c ControlLayerT) {
	for sq := l04.Square(11); sq < 100; sq += 1 {
		if l04.File(sq) != 0 && l04.Rank(sq) != 0 {
			pPosSys.ControlBoards[0][c][sq] = 0
			pPosSys.ControlBoards[1][c][sq] = 0
		}
	}
}

// MergeControlDiff - 利きの差分を解消するぜ（＾～＾）
func (pPosSys *PositionSystem) MergeControlDiff() {
	for sq := l04.Square(11); sq < BOARD_SIZE; sq += 1 {
		if l04.File(sq) != 0 && l04.Rank(sq) != 0 {
			// c=0 を除く
			for c := CONTROL_LAYER_DIFF_START; c < CONTROL_LAYER_DIFF_END; c += 1 {
				pPosSys.ControlBoards[0][CONTROL_LAYER_SUM][sq] += pPosSys.ControlBoards[0][c][sq]
				pPosSys.ControlBoards[1][CONTROL_LAYER_SUM][sq] += pPosSys.ControlBoards[1][c][sq]
			}
		}
	}
}

// RecalculateControl - 利きの再計算
func (pPosSys *PositionSystem) RecalculateControl(pPos *Position, c1 ControlLayerT) {

	pPosSys.ClearControlLayer(c1)

	for from := l04.Square(11); from < BOARD_SIZE; from += 1 {
		if l04.File(from) != 0 && l04.Rank(from) != 0 && !pPos.IsEmptySq(from) {
			piece := pPos.Board[from]
			phase := Who(piece)
			sq_list := GenMoveEnd(pPos, from)

			for _, to := range sq_list {
				pPosSys.ControlBoards[phase-1][c1][to] += 1
			}

		}
	}
}

// DiffControl - 利きテーブルの差分計算
func (pPosSys *PositionSystem) DiffControl(c1 ControlLayerT, c2 ControlLayerT, c3 ControlLayerT) {

	pPosSys.ClearControlLayer(c3)

	for phase := 0; phase < 2; phase += 1 {
		for from := l04.Square(11); from < BOARD_SIZE; from += 1 {
			if l04.File(from) != 0 && l04.Rank(from) != 0 {

				pPosSys.ControlBoards[phase][c3][from] = pPosSys.ControlBoards[phase][c1][from] - pPosSys.ControlBoards[phase][c2][from]

			}
		}
	}
}
