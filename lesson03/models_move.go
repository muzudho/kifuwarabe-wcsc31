package lesson03

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
func NewMove(from Square, to Square, promotion bool) Move {
	// if 116 <= from {
	// 	panic(App.Log.Fatal(fmt.Sprintf("new.move 1: abnormal from=%d\n", from)))
	// }
	// if App.IsDebug {
	// 	var fromHandSq = FromSqToHandSq(from)
	// 	if 116 <= from || fromHandSq < HANDSQ_BEGIN || HANDSQ_END <= fromHandSq {
	// 		panic(App.Log.Fatal(fmt.Sprintf("new move: abnormal from=%d, fromHandSq=%d, HANDSQ_BEGIN=%d, HANDSQ_END=%d\n", from, fromHandSq, HANDSQ_BEGIN, HANDSQ_END)))
	// 	}
	// }

	move := RESIGN_MOVE

	// Replace source square bits
	move = move.ReplaceSource(from)

	// Replace destination square bits
	move = move.ReplaceDestination(to)

	// Replace promotion bit
	move = move.ReplacePromotion(promotion)

	//AssertMove(move)
	return move
}

// ReplaceSource - Replace 7 source square bits
// 1111 1111 1000 0000 (Clear) 0xff80
// .pdd dddd dsss ssss
func (move Move) ReplaceSource(sq Square) Move {
	return Move(uint16(move)&0xff80 | uint16(sq))
}

// ReplaceDestination - Replace 7 destination square bits
// 1100 0000 0111 1111 (Clear) 0xc07f
// .pdd dddd dsss ssss
func (move Move) ReplaceDestination(sq Square) Move {
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
func (move Move) Destructure() (Square, Square, bool) {
	//AssertMove(move)

	var from = Square(uint16(move) & 0x007f)
	var to = Square((uint16(move) & 0x3f80) >> 7)
	var pro = uint16(move)&0x4000 != 0

	// if 116 <= from {
	// 	panic(App.Log.Fatal(fmt.Sprintf("move.destructure 1: abnormal from=%d\n", from)))
	// }
	// if App.IsDebug {
	// 	var fromHandSq = FromSqToHandSq(from)
	// 	if 116 <= from || fromHandSq < HANDSQ_BEGIN || HANDSQ_END <= fromHandSq {
	// 		panic(App.Log.Fatal(fmt.Sprintf("move.destructure: abnormal from=%d, fromHandSq=%d, HANDSQ_BEGIN=%d, HANDSQ_END=%d\n", from, fromHandSq, HANDSQ_BEGIN, HANDSQ_END)))
	// 	}
	// }

	return from, to, pro
}

// func AssertMove(move Move) {
// 	var from = Square(uint16(move) & 0x007f)
// 	if 116 <= from {
// 		panic(App.Log.Fatal(fmt.Sprintf("assert move 1: abnormal from=%d\n", from)))
// 	}
// }
