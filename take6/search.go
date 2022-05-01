package take6

import (
	"math/rand"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// Search - 探索部
func Search(pPos *Position) l03.Move {

	// 指し手生成
	legal_move_list := GenMoveList(pPos)
	size := len(legal_move_list)

	if size == 0 {
		return l03.RESIGN_MOVE
	}

	// ゲーム向けの軽い乱数
	return legal_move_list[rand.Intn(size)]
}
