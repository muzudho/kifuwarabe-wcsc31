package take12

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

// ControlBoard - 利きボード
type ControlBoard struct {
	// 表示用の名前
	Title string
	// マスへの利き数、または差分が入っています
	Board [BOARD_SIZE]int8
}

// NewControlBoard - 利きボード生成
func NewControlBoard(title string) *ControlBoard {
	c := new(ControlBoard)
	c.Title = title
	c.Board = [BOARD_SIZE]int8{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	return c
}

// Clear - 利きボードのクリアー
func (pCB *ControlBoard) Clear() {
	for sq := l04.Square(11); sq < 100; sq += 1 {
		if l04.File(sq) != 0 && l04.Rank(sq) != 0 {
			pCB.Board[sq] = 0
		}
	}
}
