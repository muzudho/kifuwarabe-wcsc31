package take13

import "fmt"

const (
	// 持ち駒を打つ 100～115
	// 先手飛打
	SQ_K1         = Square(100)
	SQ_R1         = Square(101)
	SQ_B1         = Square(102)
	SQ_G1         = Square(103)
	SQ_S1         = Square(104)
	SQ_N1         = Square(105)
	SQ_L1         = Square(106)
	SQ_P1         = Square(107)
	SQ_K2         = Square(108)
	SQ_R2         = Square(109)
	SQ_B2         = Square(110)
	SQ_G2         = Square(111)
	SQ_S2         = Square(112)
	SQ_N2         = Square(113)
	SQ_L2         = Square(114)
	SQ_P2         = Square(115)
	SQ_HAND_START = SQ_K1
	SQ_HAND_END   = SQ_P2 + 1 // この数を含まない
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

func NewMoveValue() Move {
	return Move(0)
}

// NewMoveValue2 - 初期値として 移動元マス、移動先マス、成りの有無 を指定してください
func NewMoveValue2(from Square, to Square, promote bool) Move {
	move := Move(0)
	move = move.ReplaceSource(from)
	move = move.ReplacePromotion(promote)
	return move.ReplaceDestination(to)
}

// ToCode - SFEN の moves の後に続く指し手に使える文字列を返します
func (move Move) ToCode() string {

	// 投了（＾～＾）
	if uint32(move) == 0 {
		return "resign"
	}

	str := make([]byte, 0, 5)
	count := 0

	// 移動元マス(Source square)
	source_sq := Square(move.GetSource())
	switch source_sq {
	case SQ_R1, SQ_R2:
		str = append(str, 'R')
		count = 1
	case SQ_B1, SQ_B2:
		str = append(str, 'B')
		count = 1
	case SQ_G1, SQ_G2:
		str = append(str, 'G')
		count = 1
	case SQ_S1, SQ_S2:
		str = append(str, 'S')
		count = 1
	case SQ_N1, SQ_N2:
		str = append(str, 'N')
		count = 1
	case SQ_L1, SQ_L2:
		str = append(str, 'L')
		count = 1
	case SQ_P1, SQ_P2:
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
		var sq Square // マス番号
		if count == 0 {
			// 移動元
			sq = source_sq
		} else if count == 1 {
			// 移動先
			sq = Square(move.GetDestination())
		} else {
			panic(fmt.Errorf("LogicError: count=%d", count))
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

	if move.IsPromotion() {
		str = append(str, '+')
	}

	return string(str)
}

// ReplaceSource - 移動元マス
// 1111 1111 1000 0000 (Clear) 0xff80
// .pdd dddd dsss ssss
func (move Move) ReplaceSource(sq Square) Move {
	return Move(uint16(move)&0xff80 | uint16(sq))
}

// ReplaceDestination - 移動先マス
// 1100 0000 0111 1111 (Clear) 0xc07f
// .pdd dddd dsss ssss
func (move Move) ReplaceDestination(sq Square) Move {
	return Move(uint16(move)&0xc07f | (uint16(sq) << 7))
}

// ReplacePromotion - 成
// 0100 0000 0000 0000 (Stand) 0x4000
// 1011 1111 1111 1111 (Clear) 0xbfff
// .pdd dddd dsss ssss
func (move Move) ReplacePromotion(promotion bool) Move {
	if promotion {
		return Move(uint16(move) | 0x4000)
	}

	return Move(uint16(move) & 0xbfff)
}

// GetSource - 移動元マス
// 0000 0000 0111 1111 (Mask) 0x007f
// .pdd dddd dsss ssss
func (move Move) GetSource() Square {
	return Square(uint16(move) & 0x007f)
}

// GetDestination - 移動元マス
// 0011 1111 1000 0000 (Mask) 0x3f80
// .pdd dddd dsss ssss
func (move Move) GetDestination() Square {
	return Square((uint16(move) & 0x3f80) >> 7)
}

// IsPromotion - 成
// 0100 0000 0000 0000 (Mask) 0x4000
// .pdd dddd dsss ssss
func (move Move) IsPromotion() bool {
	return uint16(move)&0x4000 != 0
}
