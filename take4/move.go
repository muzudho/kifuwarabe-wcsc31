package take4

import "fmt"

const (
	// 持ち駒を打つ 100～113
	// 先手飛打
	DROP_R1 = iota + 100
	DROP_B1
	DROP_G1
	DROP_S1
	DROP_N1
	DROP_L1
	DROP_P1
	DROP_R2
	DROP_B2
	DROP_G2
	DROP_S2
	DROP_N2
	DROP_L2
	DROP_P2
	DROP_ORIGIN = DROP_R1
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

func NewMoveValue() Move {
	return Move(0)
}

// ToCode - SFEN の moves の後に続く指し手に使える文字列を返します
func (move Move) ToCode() string {
	str := make([]byte, 0, 5)
	count := 0

	// 移動元マス(Source square)
	source_sq := move.GetSource()
	switch source_sq {
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
		var sq byte // マス番号
		if count == 0 {
			// 移動元
			sq = source_sq
		} else if count == 1 {
			// 移動先
			sq = move.GetDestination()
		} else {
			panic(fmt.Errorf("LogicError: count=%d", count))
		}
		// 正常時は必ず２桁（＾～＾）
		file := sq / 10
		rank := sq % 10
		// ASCII Code
		// '0'=48, '9'=57, 'a'=97, 'i'=105
		str = append(str, file+48)
		str = append(str, rank+96)
		fmt.Printf("Debug: file=%d rank=%d\n", file, rank)
		count += 1
	}

	// if move.IsPromotion() {
	// 	str = append(str, '+')
	// }

	return string(str)
}

// ReplaceSource - 移動元マス
func (move Move) ReplaceSource(sq uint32) Move {
	return Move(uint32(move)&0xffffff00 | sq)
}

// ReplaceDestination - 移動先マス
func (move Move) ReplaceDestination(sq uint32) Move {
	return Move(uint32(move)&0xffff00ff | (sq << 8))
}

// ReplacePromotion - 成
func (move Move) ReplacePromotion(promotion bool) Move {
	if promotion {
		return Move(uint32(move) | 0x00010000)
	}

	return Move(uint32(move) & 0xfffeffff)
}

// GetSource - 移動元マス
func (move Move) GetSource() byte {
	return byte(uint32(move) & 0x000000ff)
}

// GetDestination - 移動元マス
func (move Move) GetDestination() byte {
	return byte((uint32(move) >> 8) & 0x000000ff)
}

// IsPromotion - 成
func (move Move) IsPromotion() byte {
	return byte((uint32(move) >> 9) & 0x00000001)
}
