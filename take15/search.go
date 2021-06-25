package take15

import (
	"fmt"
	"math/rand"
)

const RESIGN_VALUE = Value(-2_147_483_647)     // Value(-32767)
const ANTI_RESIGN_VALUE = Value(2_147_483_647) // Value(32767)

var nodesNum int

// 0 にすると 1手読み（＾～＾）
var depthEnd int = 0 // 3 はまだ遅い。 2 だと駒を取り返さない。

type CuttingType int

const (
	CuttingNone        = CuttingType(0)
	CuttingKingCapture = CuttingType(1)
)

// Search - 探索部
func Search(pBrain *Brain) Move {

	nodesNum = 0
	curDepth := 0
	//fmt.Printf("Search: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	bestMove, bestVal := search2(pBrain, curDepth)

	// 評価値出力（＾～＾）
	G.Chat.Print("info depth %d nodes %d score cp %d currmove %s pv %s\n",
		curDepth, nodesNum, bestVal, bestMove.ToCode(), bestMove.ToCode())

	// ゲーム向けの軽い乱数
	return bestMove
}

// search2 - 探索部
func search2(pBrain *Brain, curDepth int) (Move, Value) {
	//fmt.Printf("Search2: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	someMoves := GenMoveList(pBrain, pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
	lenOfMoves := len(someMoves)
	//fmt.Printf("%d/%d lenOfMoves=%d\n", curDepth, depthEnd, lenOfMoves)

	if lenOfMoves == 0 {
		// ステイルメートされたら負け（＾～＾）
		return RESIGN_MOVE, RESIGN_VALUE
	}

	// 同じ価値のベストムーブがいっぱいあるかも（＾～＾）
	var someBestMoves []Move

	// 次の相手の手の評価値（自分は これを最小にしたい）
	var opponentWorstVal Value = ANTI_RESIGN_VALUE
	var younger_sibling_move = RESIGN_MOVE
	// 探索終了
	var cutting = CuttingNone

	// その手を指してみるぜ（＾～＾）
	for i, move := range someMoves {
		// G.Chat.Debug("move=%s\n", move.ToCode())
		from, _, _ := move.Destructure()

		// デバッグに使うために、盤をコピーしておきます
		pPosCopy := NewPosition()
		copyBoard(pBrain.PPosSys.PPosition[0], pPosCopy)

		// DoMove と UndoMove を繰り返していると、ずれてくる（＾～＾）
		if pBrain.PPosSys.PPosition[POS_LAYER_MAIN].IsEmptySq(from) {
			// 強制終了した局面（＾～＾）
			G.Chat.Debug(pBrain.PPosSys.PPosition[POS_LAYER_MAIN].Sprint(
				pBrain.PPosSys.phase,
				pBrain.PPosSys.StartMovesNum,
				pBrain.PPosSys.OffsetMovesIndex,
				pBrain.PPosSys.createMovesText()))
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pBrain.PPosSys.PPosition[POS_LAYER_MAIN].SprintLocation())
			panic(fmt.Errorf("Move.Source(%d) has empty square. i=%d/%d. younger_sibling_move=%s",
				from, i, lenOfMoves, younger_sibling_move.ToCode()))
		}

		pBrain.DoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex-1]

		if pBrain.IsCheckmate(FlipPhase(pBrain.PPosSys.phase)) {
			// ここで指した方の玉に王手がかかるようなら、被空き王手（＾～＾）
			// この手は見なかったことにするぜ（＾～＾）
		} else if What(captured) == PIECE_TYPE_K {
			// 玉を取るのは最善手
			someBestMoves = nil
			someBestMoves = append(someBestMoves, move)
			opponentWorstVal = RESIGN_VALUE
			// bestVal = ANTI_RESIGN_VALUE // pBrain.PPosSys.PPosition[POS_LAYER_MAIN].MaterialValue
			cutting = CuttingKingCapture
		} else {
			if curDepth < depthEnd {

				// 再帰
				_, opponentVal := search2(pBrain, curDepth+1)
				// 再帰直後（＾～＾）
				// G.Chat.Debug(pBrain.PPosSys.Sprint(POS_LAYER_MAIN))

				if opponentVal < opponentWorstVal {
					// より低い価値が見つかったら更新
					someBestMoves = nil
					someBestMoves = append(someBestMoves, move)
					opponentWorstVal = opponentVal
				} else if opponentVal == opponentWorstVal {
					// 最低値が並んだら配列の要素として追加
					someBestMoves = append(someBestMoves, move)
				}

			} else {
				// 駒割り評価値は、相手の手番のものになっています。
				materialVal := pBrain.PPosSys.PPosition[POS_LAYER_MAIN].MaterialValue
				//fmt.Printf("move=%s leafVal=%6d materialVal=%6d(%s) control_val=%6d\n", move.ToCode(), leafVal, materialVal, captured.ToCode(), control_val)

				if materialVal < opponentWorstVal {
					// より低い価値が見つかったら更新
					someBestMoves = nil
					someBestMoves = append(someBestMoves, move)
					opponentWorstVal = materialVal
				} else if materialVal == opponentWorstVal {
					// 最低値が並んだら配列の要素として追加
					someBestMoves = append(someBestMoves, move)
				}
			}
		}

		pBrain.UndoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN])

		// 盤と、コピー盤を比較します
		diffBoard(pBrain.PPosSys.PPosition[0], pPosCopy, pBrain.PPosSys.PPosition[2], pBrain.PPosSys.PPosition[3])
		// 異なる箇所を数えます
		errorNum := errorBoard(pBrain.PPosSys.PPosition[0], pPosCopy, pBrain.PPosSys.PPosition[2], pBrain.PPosSys.PPosition[3])
		if errorNum != 0 {
			// 違いのあった局面（＾～＾）
			G.Chat.Debug(pBrain.PPosSys.SprintDiff(0, 1))
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pBrain.PPosSys.PPosition[0].SprintLocation())
			G.Chat.Debug(pPosCopy.SprintLocation())
			panic(fmt.Errorf("Error: count=%d younger_sibling_move=%s move=%s", errorNum, younger_sibling_move.ToCode(), move.ToCode()))
		}

		younger_sibling_move = move

		if cutting != CuttingNone {
			break
		}

		/*
			// Debug ここから
			var debugBestMove = RESIGN_MOVE
			bestmoveListLen := len(someBestMoves)
			if bestmoveListLen > 0 {
				debugBestMove = someBestMoves[rand.Intn(bestmoveListLen)]
			}
			G.Chat.Debug("info string Debug: depth=%d nodes=%d value=%d move.best=%s.%s\n", curDepth, nodesNum, -opponentWorstVal, move.ToCode(), debugBestMove.ToCode())
			// Debug ここまで
		*/
	}

	var bestMove = RESIGN_MOVE
	bestmoveListLen := len(someBestMoves)
	//fmt.Printf("%d/%d bestmoveListLen=%d\n", curDepth, depthEnd, bestmoveListLen)
	if bestmoveListLen > 0 {
		// 0件を避ける（＾～＾）
		bestMove = someBestMoves[rand.Intn(bestmoveListLen)]
	}
	// 評価値出力（＾～＾）
	// G.Chat.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, bestMove.ToCode(), bestMove.ToCode())

	// 相手の評価値の逆が、自分の評価値
	return bestMove, -opponentWorstVal
}
