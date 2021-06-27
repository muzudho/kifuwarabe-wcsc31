package take15

import (
	b "github.com/muzudho/kifuwarabe-wcsc31/take16base"
	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

const (
	// 持ち駒を打つ 100～115
	// 先手飛打
	SQ_K1         = p.Square(100)
	SQ_R1         = p.Square(101)
	SQ_B1         = p.Square(102)
	SQ_G1         = p.Square(103)
	SQ_S1         = p.Square(104)
	SQ_N1         = p.Square(105)
	SQ_L1         = p.Square(106)
	SQ_P1         = p.Square(107)
	SQ_K2         = p.Square(108)
	SQ_R2         = p.Square(109)
	SQ_B2         = p.Square(110)
	SQ_G2         = p.Square(111)
	SQ_S2         = p.Square(112)
	SQ_N2         = p.Square(113)
	SQ_L2         = p.Square(114)
	SQ_P2         = p.Square(115)
	SQ_HAND_START = SQ_K1
	SQ_HAND_END   = SQ_P2 + 1 // この数を含まない
)

// 0 は 投了ということにするぜ（＾～＾）
const RESIGN_MOVE = b.Move(0)

// NewMove - 初期値として 移動元マス、移動先マス、成りの有無 を指定してください
func NewMove(from p.Square, to p.Square, promotion bool) b.Move {
	move := RESIGN_MOVE

	// replaceSource - 移動元マス
	// 1111 1111 1000 0000 (Clear) 0xff80
	// .pdd dddd dsss ssss
	move = b.Move(uint16(move)&0xff80 | uint16(from))

	// replaceDestination - 移動先マス
	// 1100 0000 0111 1111 (Clear) 0xc07f
	// .pdd dddd dsss ssss
	move = b.Move(uint16(move)&0xc07f | (uint16(to) << 7))

	// replacePromotion - 成
	// 0100 0000 0000 0000 (Stand) 0x4000
	// 1011 1111 1111 1111 (Clear) 0xbfff
	// .pdd dddd dsss ssss
	if promotion {
		return b.Move(uint16(move) | 0x4000)
	}

	return b.Move(uint16(move) & 0xbfff)
}
