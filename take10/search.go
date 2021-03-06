package take10

import (
	"math/rand"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// Search - 探索部
func Search(pPos *Position) l03.Move {

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPos)
	size := len(move_list)

	if size == 0 {
		return l03.RESIGN_MOVE
	}

	nodesNum := 0

	// 同じ価値のベストムーブがいっぱいあるかも（＾～＾）
	var bestMoveList []l03.Move
	// 最初に最低値を入れておけば、更新されるだろ（＾～＾）
	var bestVal int16 = -32768

	// その手を指してみるぜ（＾～＾）
	for _, move := range move_list {
		pPos.DoMove(move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pPos.CapturedList[pPos.OffsetMovesIndex-1]
		materialVal := EvalMaterial(captured)

		if bestVal < materialVal {
			// より高い価値が見つかったら更新
			bestMoveList = nil
			bestMoveList = append(bestMoveList, move)
			bestVal = materialVal
		} else if bestVal == materialVal {
			// 最高値が並んだら配列の要素として追加
			bestMoveList = append(bestMoveList, move)
		}

		pPos.UndoMove()
	}

	// 0件にはならないはず（＾～＾）
	currMove := bestMoveList[rand.Intn(len(bestMoveList))]

	// 評価値出力（＾～＾）
	App.Out.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, currMove.ToCodeOfM(), currMove.ToCodeOfM())

	// ゲーム向けの軽い乱数
	return currMove
}
