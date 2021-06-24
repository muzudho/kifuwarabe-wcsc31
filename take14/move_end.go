// 移動先と成り

package take14

// MoveEnd - 移動先と成り
//
// pddd dddd
//
// 1～7bit: 移動先(0～127)
// 8bit: 成(0～1)
type MoveEnd uint8

// 0 は 投了ということにするぜ（＾～＾）
const RESIGN_MOVE_END = MoveEnd(0)

// NewMoveEnd - 移動先マス、成りの有無 を指定してください
func NewMoveEnd(to Square, promote bool) MoveEnd {
	moveEnd := RESIGN_MOVE_END
	moveEnd = moveEnd.ReplaceDestination(to)
	return moveEnd.ReplacePromotion(promote)
}

// ReplaceDestination - 移動先マス
// 1000 0000 (Clear) 0x80
// pddd dddd
func (moveEnd MoveEnd) ReplaceDestination(sq Square) MoveEnd {
	return MoveEnd(uint8(moveEnd)&0x80 | uint8(sq))
}

// ReplacePromotion - 成
// 1000 0000 (Stand) 0x80
// 0111 1111 (Clear) 0x7f
// pddd dddd
func (moveEnd MoveEnd) ReplacePromotion(promotion bool) MoveEnd {
	if promotion {
		return MoveEnd(uint8(moveEnd) | 0x80)
	}

	return MoveEnd(uint8(moveEnd) & 0x7f)
}

// Destructure
//
// 移動先マス
// 0111 1111 (Mask) 0x7f
// pddd dddd
//
// 成
// 1000 0000 (Mask) 0x80
// pddd dddd
func (moveEnd MoveEnd) Destructure() (Square, bool) {
	var to = Square(uint8(moveEnd) & 0x7f)
	var pro = uint8(moveEnd)&0x80 != 0
	return to, pro
}

// ToString - 確認用の文字列
func (moveEnd MoveEnd) ToString() string {

	// 投了（＾～＾）
	if uint8(moveEnd) == 0 {
		return "resign"
	}

	str := make([]byte, 0, 3)

	to, _ := moveEnd.Destructure()
	// 正常時は必ず２桁（＾～＾）
	file := byte(to / 10)
	rank := byte(to % 10)
	// ASCII Code
	// '0'=48, '9'=57, 'a'=97, 'i'=105
	str = append(str, file+48)
	str = append(str, rank+96)

	return string(str)
}
