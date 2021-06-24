package take7

import "fmt"

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	DROP_R1        = Square(100)
	DROP_B1        = Square(101)
	DROP_G1        = Square(102)
	DROP_S1        = Square(103)
	DROP_N1        = Square(104)
	DROP_L1        = Square(105)
	DROP_P1        = Square(106)
	DROP_R2        = Square(107)
	DROP_B2        = Square(108)
	DROP_G2        = Square(109)
	DROP_S2        = Square(110)
	DROP_N2        = Square(111)
	DROP_L2        = Square(112)
	DROP_P2        = Square(113)
	DROP_ORIGIN    = DROP_R1
	DROP_TYPE_SIZE = DROP_P1 - DROP_ORIGIN
)

// Move - 指し手
//
// 17bit で表せるはず（＾～＾）
// pddddddddssssssss
//
// 1～8bit: 移動元
// 9～16bit: 移動先
// 17bit: 成
type Move uint32

// 0 は 投了ということにするぜ（＾～＾）
const RESIGN_MOVE = Move(0)

// NewMoveValue2 - 初期値として 移動元マス、移動先マスを指定してください
func NewMoveValue2(src_sq Square, dst_sq Square) Move {
	move := RESIGN_MOVE
	move = move.ReplaceSource(src_sq)
	return move.ReplaceDestination(dst_sq)
}

// ToCode - SFEN の moves の後に続く指し手に使える文字列を返します
func (move Move) ToCode() string {

	// 投了（＾～＾）
	if uint32(move) == 0 {
		return "resign"
	}

	str := make([]byte, 0, 5)
	count := 0

	from, to, pro := move.Destructure()

	// 移動元マス(Source square)
	switch from {
	case DROP_R1, DROP_R2:
		str = append(str, 'R')
		count = 1
	case DROP_B1, DROP_B2:
		str = append(str, 'B')
		count = 1
	case DROP_G1, DROP_G2:
		str = append(str, 'G')
		count = 1
	case DROP_S1, DROP_S2:
		str = append(str, 'S')
		count = 1
	case DROP_N1, DROP_N2:
		str = append(str, 'N')
		count = 1
	case DROP_L1, DROP_L2:
		str = append(str, 'L')
		count = 1
	case DROP_P1, DROP_P2:
		str = append(str, 'P')
		count = 1
	default:
		// Ignored
	}

	if count == 1 {
		str = append(str, '+')
	}

	for count < 2 {
		var sq Square // マス番号
		if count == 0 {
			// 移動元
			sq = from
		} else if count == 1 {
			// 移動先
			sq = to
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

	if pro {
		str = append(str, '+')
	}

	return string(str)
}

// ReplaceSource - 移動元マス
func (move Move) ReplaceSource(sq Square) Move {
	return Move(uint32(move)&0xffffff00 | uint32(sq))
}

// ReplaceDestination - 移動先マス
func (move Move) ReplaceDestination(sq Square) Move {
	return Move(uint32(move)&0xffff00ff | (uint32(sq) << 8))
}

// ReplacePromotion - 成
func (move Move) ReplacePromotion(promotion bool) Move {
	if promotion {
		return Move(uint32(move) | 0x00010000)
	}

	return Move(uint32(move) & 0xfffeffff)
}

// Destructure - 移動元マス、移動先マス、成りの有無
func (move Move) Destructure() (Square, Square, bool) {
	var from = Square(uint32(move) & 0x000000ff)
	var to = Square((uint32(move) >> 8) & 0x000000ff)
	var pro = (uint32(move)>>9)&0x00000001 == 1
	return from, to, pro
}
