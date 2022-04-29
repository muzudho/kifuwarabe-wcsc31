package take17

import (
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
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
var oneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// FlipPhase - 先後を反転します
func FlipPhase(phase l06.Phase) l06.Phase {
	return phase%2 + 1
}

// [0], [1]
const PHASE_ARRAY_SIZE = 2

// PositionSystem - 局面にいろいろな機能を付けたもの
type PositionSystem struct {
	// 局面
	PPosition [POS_LAYER_SIZE]*l15.Position

	// 先手が1、後手が2（＾～＾）
	phase l06.Phase
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)
	pPosSys.PPosition = [POS_LAYER_SIZE]*l15.Position{l15.NewPosition(), l15.NewPosition(), l15.NewPosition(), l15.NewPosition()}
	pPosSys.ResetPosition()
	return pPosSys
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) ResetPosition() {
	// 先手の局面
	pPosSys.phase = l06.FIRST
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() l06.Phase {
	return pPosSys.phase
}
