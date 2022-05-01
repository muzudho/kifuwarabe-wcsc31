package take10

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// Move - 指し手
//
// 15bit で表せるはず（＾～＾）
// .pdd dddd dsss ssss
//
// 1～7bit: 移動元(0～127)
// 8～14bit: 移動先(0～127)
// 15bit: 成(0～1)
type Move uint16

// 0 は 投了ということにするぜ（＾～＾）
const RESIGN_MOVE = Move(0)

// NewMove - 初期値として 移動元マス、移動先マス、成りの有無 を指定してください
func NewMove(from l03.Square, to l03.Square, promotion bool) Move {
	move := RESIGN_MOVE

	// Replace source square bits
	move = move.ReplaceSource(from)

	// Replace destination square bits
	move = move.ReplaceDestination(to)

	// Replace promotion bit
	return move.ReplacePromotion(promotion)
}

// ToCodeOfM - SFEN の moves の後に続く指し手に使える文字列を返します
func (move Move) ToCodeOfM() string {

	// 投了（＾～＾）
	if uint32(move) == 0 {
		return "resign"
	}

	str := make([]byte, 0, 5)
	count := 0

	// 移動元マス、移動先マス、成りの有無
	from, to, pro := move.Destructure()

	// 移動元マス(Source square)
	switch from {
	case l03.SQ_R1, l03.SQ_R2:
		str = append(str, 'R')
		count = 1
	case l03.SQ_B1, l03.SQ_B2:
		str = append(str, 'B')
		count = 1
	case l03.SQ_G1, l03.SQ_G2:
		str = append(str, 'G')
		count = 1
	case l03.SQ_S1, l03.SQ_S2:
		str = append(str, 'S')
		count = 1
	case l03.SQ_N1, l03.SQ_N2:
		str = append(str, 'N')
		count = 1
	case l03.SQ_L1, l03.SQ_L2:
		str = append(str, 'L')
		count = 1
	case l03.SQ_P1, l03.SQ_P2:
		str = append(str, 'P')
		count = 1
	default:
		// Ignored
	}

	if count == 1 {
		// 打
		str = append(str, '*')
	}

	for count < 2 {
		var sq l03.Square // マス番号
		if count == 0 {
			// 移動元
			sq = from
		} else if count == 1 {
			// 移動先
			sq = to
		} else {
			panic(App.LogNotEcho.Fatal("LogicError: count=%d", count))
		}
		// 正常時は必ず２桁（＾～＾）
		file := byte(sq / 10)
		rank := byte(sq % 10)
		// ASCII Code
		// '0'=48, '9'=57, 'a'=97, 'i'=105
		str = append(str, file+48)
		str = append(str, rank+96)
		// fmt.Printf("Debug: move=%d sq=%d count=%d file=%d rank=%d\n", uint32(move), sq, count, file, rank)
		count += 1
	}

	if pro {
		str = append(str, '+')
	}

	return string(str)
}

// ReplaceSource - Replace 7 source square bits
// 1111 1111 1000 0000 (Clear) 0xff80
// .pdd dddd dsss ssss
func (move Move) ReplaceSource(sq l03.Square) Move {
	return Move(uint16(move)&0xff80 | uint16(sq))
}

// ReplaceDestination - Replace 7 destination square bits
// 1100 0000 0111 1111 (Clear) 0xc07f
// .pdd dddd dsss ssss
func (move Move) ReplaceDestination(sq l03.Square) Move {
	return Move(uint16(move)&0xc07f | (uint16(sq) << 7))
}

// ReplacePromotion - Replace 1 promotion bit
// 0100 0000 0000 0000 (Stand) 0x4000
// 1011 1111 1111 1111 (Clear) 0xbfff
// .pdd dddd dsss ssss
func (move Move) ReplacePromotion(promotion bool) Move {
	if promotion {
		return Move(uint16(move) | 0x4000)
	}

	return Move(uint16(move) & 0xbfff)
}

// Destructure - 移動元マス、移動先マス、成りの有無
//
// 移動元マス
// 0000 0000 0111 1111 (Mask) 0x007f
// .pdd dddd dsss ssss
//
// 移動先マス
// 0011 1111 1000 0000 (Mask) 0x3f80
// .pdd dddd dsss ssss
//
// 成
// 0100 0000 0000 0000 (Mask) 0x4000
// .pdd dddd dsss ssss
func (move Move) Destructure() (l03.Square, l03.Square, bool) {
	var from = l03.Square(uint16(move) & 0x007f)
	var to = l03.Square((uint16(move) & 0x3f80) >> 7)
	var pro = uint16(move)&0x4000 != 0
	return from, to, pro
}
