package take16

import (
	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// ControlBoard - 利きボード
type ControlBoard struct {
	// 表示用の名前
	Title string
	// マスへの利き数、または差分、さらには評価値が入っています
	Board1 [p.BOARD_SIZE]int16
}

// NewControlBoard - 利きボード生成
func NewControlBoard(title string) *ControlBoard {
	c := new(ControlBoard)
	c.Title = title
	c.Board1 = [p.BOARD_SIZE]int16{
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
	for sq := p.Square(11); sq < 100; sq += 1 {
		if p.File(sq) != 0 && p.Rank(sq) != 0 {
			pCB.Board1[sq] = 0
		}
	}
}

// AddControl - 盤上のマスを指定することで、そこにある駒の利きを調べて、利きの差分テーブルの値を増減させます
func (pCB *ControlBoard) AddControl(sq_list []p.Square, from p.Square, sign int16) {

	// if from > 99 {
	// 	// 持ち駒は無視します
	// 	return
	// }

	//sq_list := GenMoveEnd(pPos, from)
	for _, to := range sq_list {
		// fmt.Printf("Debug: ph=%d c=%d to=%d\n", ph, c, to)
		// 差分の方のテーブルを更新（＾～＾）
		pCB.Board1[to] += sign * 1
	}
}
