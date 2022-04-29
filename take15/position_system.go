package take15

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l13 "github.com/muzudho/kifuwarabe-wcsc31/take13"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
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
func FlipPhase(phase l06.Phase) l06.Phase {
	return phase%2 + 1
}

// [0], [1]
const PHASE_ARRAY_SIZE = 2

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
	phase l06.Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [l04.MOVES_SIZE]l13.Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [l04.MOVES_SIZE]l03.Piece
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
func (pPosSys *PositionSystem) GetPhase() l06.Phase {
	return pPosSys.phase
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	// 先手の局面
	pPosSys.phase = l06.FIRST
	// 何手目か
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.Moves = [l04.MOVES_SIZE]l13.Move{}
	// 取った駒のリスト
	pPosSys.CapturedList = [l04.MOVES_SIZE]l03.Piece{}
}
