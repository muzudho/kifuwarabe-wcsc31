package take11

import (
	"fmt"
	"math/rand"
)

const RESIGN_VALUE = -32768
const MAX_VALUE = 32767

var nodesNum int
var depthEnd int = 1

// Search - 探索部
func Search(pPosSys *PositionSystem) Move {

	nodesNum = 0
	curDepth := 0
	//fmt.Printf("Search: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	bestmove, bestVal := search2(pPosSys, curDepth)

	// 評価値出力（＾～＾）
	G.Chat.Print("info depth %d nodes %d score cp %d currmove %s pv %s\n",
		curDepth, nodesNum, bestVal, bestmove.ToCode(), bestmove.ToCode())

	// ゲーム向けの軽い乱数
	return bestmove
}

// search2 - 探索部
func search2(pPosSys *PositionSystem, curDepth int) (Move, int16) {
	//fmt.Printf("Search2: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPosSys, pPosSys.PPosition[POS_LAYER_MAIN])
	move_length := len(move_list)
	//fmt.Printf("%d/%d move_length=%d\n", curDepth, depthEnd, move_length)

	if move_length == 0 {
		return RESIGN_MOVE, RESIGN_VALUE
	}

	// 同じ価値のベストムーブがいっぱいあるかも（＾～＾）
	var bestMoveList []Move
	// 最初に最低値を入れておけば、更新されるだろ（＾～＾）
	var bestVal int16 = RESIGN_VALUE

	// 相手の評価値
	var opponentWorstVal int16 = MAX_VALUE
	var younger_sibling_move = RESIGN_MOVE

	// その手を指してみるぜ（＾～＾）
	for i, move := range move_list {
		// G.Chat.Debug("move=%s\n", move.ToCode())

		from, _, _ := move.Destructure()

		// 盤をコピーしておきます
		pPosCopy := NewPosition()
		copyBoard(pPosSys.PPosition[0], pPosCopy)

		// DoMove と UndoMove を繰り返していると、ずれてくる（＾～＾）
		if pPosSys.PPosition[POS_LAYER_MAIN].IsEmptySq(from) {
			// 強制終了した局面（＾～＾）
			G.Chat.Debug(Sprint(
				pPosSys.PPosition[POS_LAYER_MAIN],
				pPosSys.phase,
				pPosSys.StartMovesNum,
				pPosSys.OffsetMovesIndex,
				pPosSys.createMovesText()))
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pPosSys.PPosition[POS_LAYER_MAIN].SprintLocation())
			panic(fmt.Errorf("Move.Source(%d) has empty square. i=%d/%d. younger_sibling_move=%s",
				from, i, move_length, younger_sibling_move.ToCode()))
		}

		pPosSys.DoMove(pPosSys.PPosition[POS_LAYER_MAIN], move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pPosSys.CapturedList[pPosSys.OffsetMovesIndex-1]
		materialVal := EvalMaterial(captured)

		if curDepth < depthEnd {
			// 再帰
			_, opponentVal := search2(pPosSys, curDepth+1)
			// 再帰直後（＾～＾）
			// G.Chat.Debug(pPosSys.Sprint(POS_LAYER_MAIN))

			if opponentVal < opponentWorstVal {
				// より低い価値が見つかったら更新
				bestMoveList = nil
				bestMoveList = append(bestMoveList, move)
				opponentWorstVal = opponentVal
			} else if opponentVal == opponentWorstVal {
				// 最低値が並んだら配列の要素として追加
				bestMoveList = append(bestMoveList, move)
			}

		} else {
			// 葉ノードでは、相手の手ではなく、自分の局面に点数を付けます

			if bestVal < materialVal {
				// より高い価値が見つかったら更新
				bestMoveList = nil
				bestMoveList = append(bestMoveList, move)
				bestVal = materialVal
			} else if bestVal == materialVal {
				// 最高値が並んだら配列の要素として追加
				bestMoveList = append(bestMoveList, move)
			}
		}

		pPosSys.UndoMove(pPosSys.PPosition[POS_LAYER_MAIN])

		// 盤と、コピー盤を比較します
		diffBoard(pPosSys.PPosition[0], pPosCopy, pPosSys.PPosition[2], pPosSys.PPosition[3])
		// 異なる箇所を数えます
		errorNum := errorBoard(pPosSys.PPosition[0], pPosCopy, pPosSys.PPosition[2], pPosSys.PPosition[3])
		if errorNum != 0 {
			// 違いのあった局面（＾～＾）
			G.Chat.Debug(pPosSys.SprintDiff(0, 1))
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pPosSys.PPosition[0].SprintLocation())
			G.Chat.Debug(pPosCopy.SprintLocation())
			panic(fmt.Errorf("Error: count=%d younger_sibling_move=%s move=%s", errorNum, younger_sibling_move.ToCode(), move.ToCode()))
		}

		younger_sibling_move = move
	}

	if curDepth < depthEnd {
		// 葉以外のノードでは、相手の評価値の逆が、自分の評価値
		bestVal = -opponentWorstVal
	}

	var bestmove = RESIGN_MOVE
	bestmove_length := len(bestMoveList)
	//fmt.Printf("%d/%d bestmove_length=%d\n", curDepth, depthEnd, bestmove_length)
	if bestmove_length > 0 {
		bestmove = bestMoveList[rand.Intn(bestmove_length)]
	}

	// 評価値出力（＾～＾）
	// G.Chat.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, bestmove.ToCode(), bestmove.ToCode())

	// 0件にはならないはず（＾～＾）
	return bestmove, bestVal
}
