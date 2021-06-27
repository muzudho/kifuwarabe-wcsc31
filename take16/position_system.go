package take16

import (
	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// 盤レイヤー・インデックス型
type PosLayerT int

const (
	POS_LAYER_MAIN  = PosLayerT(0)
	POS_LAYER_COPY  = PosLayerT(1) // テスト用
	POS_LAYER_DIFF1 = PosLayerT(2) // テスト用
	POS_LAYER_DIFF2 = PosLayerT(3) // テスト用
	POS_LAYER_SIZE  = 4
)

// FlipPhase - 先後を反転します
func FlipPhase(phase p.Phase) p.Phase {
	return phase%2 + 1
}

// [0], [1]
const PHASE_ARRAY_SIZE = 2

// PieceFromPhPt - 駒作成。空マスは作れません
func PieceFromPhPt(phase p.Phase, pieceType PieceType) p.Piece {
	switch phase {
	case p.FIRST:
		switch pieceType {
		case PIECE_TYPE_K:
			return p.PIECE_K1
		case PIECE_TYPE_R:
			return p.PIECE_R1
		case PIECE_TYPE_B:
			return p.PIECE_B1
		case PIECE_TYPE_G:
			return p.PIECE_G1
		case PIECE_TYPE_S:
			return p.PIECE_S1
		case PIECE_TYPE_N:
			return p.PIECE_N1
		case PIECE_TYPE_L:
			return p.PIECE_L1
		case PIECE_TYPE_P:
			return p.PIECE_P1
		case PIECE_TYPE_PR:
			return p.PIECE_PR1
		case PIECE_TYPE_PB:
			return p.PIECE_PB1
		case PIECE_TYPE_PS:
			return p.PIECE_PS1
		case PIECE_TYPE_PN:
			return p.PIECE_PN1
		case PIECE_TYPE_PL:
			return p.PIECE_PL1
		case PIECE_TYPE_PP:
			return p.PIECE_PP1
		default:
			panic(G.Log.Fatal("Unknown pieceType=%d", pieceType))
		}
	case p.SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return p.PIECE_K2
		case PIECE_TYPE_R:
			return p.PIECE_R2
		case PIECE_TYPE_B:
			return p.PIECE_B2
		case PIECE_TYPE_G:
			return p.PIECE_G2
		case PIECE_TYPE_S:
			return p.PIECE_S2
		case PIECE_TYPE_N:
			return p.PIECE_N2
		case PIECE_TYPE_L:
			return p.PIECE_L2
		case PIECE_TYPE_P:
			return p.PIECE_P2
		case PIECE_TYPE_PR:
			return p.PIECE_PR2
		case PIECE_TYPE_PB:
			return p.PIECE_PB2
		case PIECE_TYPE_PS:
			return p.PIECE_PS2
		case PIECE_TYPE_PN:
			return p.PIECE_PN2
		case PIECE_TYPE_PL:
			return p.PIECE_PL2
		case PIECE_TYPE_PP:
			return p.PIECE_PP2
		default:
			panic(G.Log.Fatal("Unknown pieceType=%d", pieceType))
		}
	default:
		panic(G.Log.Fatal("Unknown phase=%d", phase))
	}
}

// PositionSystem - 局面にいろいろな機能を付けたもの
type PositionSystem struct {
	// 局面
	PPosition [POS_LAYER_SIZE]*p.Position

	// 先手が1、後手が2（＾～＾）
	phase p.Phase
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int

	// 差分での連続局面記録。つまり、ふつうの棋譜（＾～＾）
	PRecord *DifferenceRecord
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)

	pPosSys.PRecord = NewDifferenceRecord()
	pPosSys.PPosition = [POS_LAYER_SIZE]*p.Position{p.NewPosition(), p.NewPosition(), p.NewPosition(), p.NewPosition()}

	pPosSys.resetPosition()
	return pPosSys
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	// 先手の局面
	pPosSys.phase = p.FIRST
	// 何手目か
	pPosSys.PRecord.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.PRecord.Moves = [MOVES_SIZE]p.Move{}
	// 取った駒のリスト
	pPosSys.PRecord.CapturedList = [MOVES_SIZE]p.Piece{}
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() p.Phase {
	return pPosSys.phase
}
