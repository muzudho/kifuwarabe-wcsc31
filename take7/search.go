package take7

import (
	"math/rand"

	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// Search - 探索部
func Search(pPos *Position) l04.Move {

	// 指し手生成
	legal_move_list := GenMoveList(pPos)
	size := len(legal_move_list)

	if size == 0 {
		return l04.RESIGN_MOVE
	}

	// Debug表示
	//for i, move := range legal_move_list {
	//fmt.Printf("Debug: (%d) %s\n", i, move.ToCode())
	//}

	// ゲーム向けの軽い乱数
	return legal_move_list[rand.Intn(size)]
}
