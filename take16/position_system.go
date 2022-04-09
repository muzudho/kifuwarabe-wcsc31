package take16

import (
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// 盤レイヤー・インデックス型
type PosLayerT int

const (
	POS_LAYER_MAIN  = PosLayerT(0)
	POS_LAYER_COPY  = PosLayerT(1) // テスト用
	POS_LAYER_DIFF1 = PosLayerT(2) // テスト用
	POS_LAYER_DIFF2 = PosLayerT(3) // テスト用
	POS_LAYER_SIZE  = 4
)

// position sfen の盤のスペース数に使われますN
var oneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// FlipPhase - 先後を反転します
func FlipPhase(phase Phase) Phase {
	return phase%2 + 1
}

// [0], [1]
const PHASE_ARRAY_SIZE = 2

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
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
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
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
	}
}

// PositionSystem - 局面にいろいろな機能を付けたもの
type PositionSystem struct {
	// 局面
	PPosition [POS_LAYER_SIZE]*Position

	// 先手が1、後手が2（＾～＾）
	phase Phase
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)
	pPosSys.PPosition = [POS_LAYER_SIZE]*Position{NewPosition(), NewPosition(), NewPosition(), NewPosition()}
	pPosSys.ResetPosition()
	return pPosSys
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) ResetPosition() {
	// 先手の局面
	pPosSys.phase = FIRST
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() Phase {
	return pPosSys.phase
}
