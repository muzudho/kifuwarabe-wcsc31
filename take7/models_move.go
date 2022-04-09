package take7

import "fmt"

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

// NewMove - 初期値として 移動元マス、移動先マスを指定してください
func NewMove(from Square, to Square, promotion bool) Move {
	move := RESIGN_MOVE

	// ReplaceSource - 移動元マス
	move = Move(uint32(move)&0xffffff00 | uint32(from))

	// ReplaceDestination - 移動先マス
	move = Move(uint32(move)&0xffff00ff | (uint32(to) << 8))

	// ReplacePromotion - 成
	if promotion {
		return Move(uint32(move) | 0x00010000)
	}

	return Move(uint32(move) & 0xfffeffff)
}

// ToMCode - SFEN の moves の後に続く指し手に使える文字列を返します
func (move Move) ToMCode() string {

	// 投了（＾～＾）
	if uint32(move) == 0 {
		return "resign"
	}

	str := make([]byte, 0, 5)
	count := 0

	from, to, pro := move.Destructure()

	// 移動元マス(Source square)
	switch from {
	case HAND_R1, HAND_R2:
		str = append(str, 'R')
		count = 1
	case HAND_B1, HAND_B2:
		str = append(str, 'B')
		count = 1
	case HAND_G1, HAND_G2:
		str = append(str, 'G')
		count = 1
	case HAND_S1, HAND_S2:
		str = append(str, 'S')
		count = 1
	case HAND_N1, HAND_N2:
		str = append(str, 'N')
		count = 1
	case HAND_L1, HAND_L2:
		str = append(str, 'L')
		count = 1
	case HAND_P1, HAND_P2:
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

// Destructure - 移動元マス、移動先マス、成りの有無
func (move Move) Destructure() (Square, Square, bool) {
	var from = Square(uint32(move) & 0x000000ff)
	var to = Square((uint32(move) >> 8) & 0x000000ff)
	var pro = (uint32(move)>>9)&0x00000001 == 1
	return from, to, pro
}
