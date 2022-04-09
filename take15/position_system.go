package take15

import (
	l12 "github.com/muzudho/kifuwarabe-wcsc31/take12"
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
var OneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

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
			return l12.PIECE_K1
		case PIECE_TYPE_R:
			return l12.PIECE_R1
		case PIECE_TYPE_B:
			return l12.PIECE_B1
		case PIECE_TYPE_G:
			return l12.PIECE_G1
		case PIECE_TYPE_S:
			return l12.PIECE_S1
		case PIECE_TYPE_N:
			return l12.PIECE_N1
		case PIECE_TYPE_L:
			return l12.PIECE_L1
		case PIECE_TYPE_P:
			return l12.PIECE_P1
		case PIECE_TYPE_PR:
			return l12.PIECE_PR1
		case PIECE_TYPE_PB:
			return l12.PIECE_PB1
		case PIECE_TYPE_PS:
			return l12.PIECE_PS1
		case PIECE_TYPE_PN:
			return l12.PIECE_PN1
		case PIECE_TYPE_PL:
			return l12.PIECE_PL1
		case PIECE_TYPE_PP:
			return l12.PIECE_PP1
		default:
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	case SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return l12.PIECE_K2
		case PIECE_TYPE_R:
			return l12.PIECE_R2
		case PIECE_TYPE_B:
			return l12.PIECE_B2
		case PIECE_TYPE_G:
			return l12.PIECE_G2
		case PIECE_TYPE_S:
			return l12.PIECE_S2
		case PIECE_TYPE_N:
			return l12.PIECE_N2
		case PIECE_TYPE_L:
			return l12.PIECE_L2
		case PIECE_TYPE_P:
			return l12.PIECE_P2
		case PIECE_TYPE_PR:
			return l12.PIECE_PR2
		case PIECE_TYPE_PB:
			return l12.PIECE_PB2
		case PIECE_TYPE_PS:
			return l12.PIECE_PS2
		case PIECE_TYPE_PN:
			return l12.PIECE_PN2
		case PIECE_TYPE_PL:
			return l12.PIECE_PL2
		case PIECE_TYPE_PP:
			return l12.PIECE_PP2
		default:
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
	}
}

var HandPieceMap1 = [HAND_SIZE]l09.Piece{
	l12.PIECE_K1, l12.PIECE_R1, l12.PIECE_B1, l12.PIECE_G1, l12.PIECE_S1, l12.PIECE_N1, l12.PIECE_L1, l12.PIECE_P1,
	l12.PIECE_K2, l12.PIECE_R2, l12.PIECE_B2, l12.PIECE_G2, l12.PIECE_S2, l12.PIECE_N2, l12.PIECE_L2, l12.PIECE_P2}

// 開発 or リリース モード
type BuildT int

const (
	BUILD_DEV     = BuildT(0)
	BUILD_RELEASE = BuildT(1)
)

// PositionSystem - 局面にいろいろな機能を付けたもの
type PositionSystem struct {
	// 開発モードフラグ。デフォルト値：真。 'usi' コマンドで解除
	BuildType BuildT
	// 局面
	PPosition [POS_LAYER_SIZE]*Position

	// 先手が1、後手が2（＾～＾）
	phase Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [l09.MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [l09.MOVES_SIZE]l09.Piece
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)
	pPosSys.BuildType = BUILD_DEV

	pPosSys.PPosition = [POS_LAYER_SIZE]*Position{NewPosition(), NewPosition(), NewPosition(), NewPosition()}

	pPosSys.resetPosition()
	return pPosSys
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() Phase {
	return pPosSys.phase
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	// 先手の局面
	pPosSys.phase = FIRST
	// 何手目か
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPosSys.CapturedList = [MOVES_SIZE]l09.Piece{}
}
