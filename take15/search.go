package take15

import (
	"math/rand"
)

/*
type SearchType uint8

// 探索への指定
const (
	// 特になし
	SEARCH_NONE = SearchType(0)
	// 駒の取り合い
	SEARCH_CAPTURE = SearchType(1)
)
*/

const RESIGN_VALUE = Value(-2_147_483_647)     // Value(-32767)
const ANTI_RESIGN_VALUE = Value(2_147_483_647) // Value(32767)

var nodesNum int

// 0 にすると 1手読み（＾～＾）
var depthEnd int = 1 // 3 はまだ遅い。 2 だと駒を取り返さない。

type CuttingType int

const (
	CuttingNone = CuttingType(0)
	// 玉を取った
	CuttingKingCapture = CuttingType(1)
	// 駒の取り合いが終わった
	//CuttingEndCapture = CuttingType(2)
)

// Search - 探索部
func Search(pBrain *Brain) Move {

	nodesNum = 0
	curDepth := 0
	//fmt.Printf("Search: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	bestMove, bestVal := search2(pBrain, curDepth) //, SEARCH_NONE

	// 評価値出力（＾～＾）
	App.Out.Print("info depth %d nodes %d score cp %d currmove %s pv %s\n",
		curDepth, nodesNum, bestVal, ToMCode(bestMove), ToMCode(bestMove))

	// ゲーム向けの軽い乱数
	return bestMove
}

// search2 - 探索部
func search2(pBrain *Brain, curDepth int) (Move, Value) { //, search_type SearchType
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
	// 前回のムーブ
	var younger_sibling_move = RESIGN_MOVE
	// 探索終了
	var cutting = CuttingNone

	// その手を指してみるぜ（＾～＾）
	for i, move := range someMoves {
		// App.Out.Debug("move=%s\n", move.ToCode())
		from, _, _ := Destructure(move)

		// デバッグに使うために、盤をコピーしておきます
		pPosCopy := NewPosition()
		copyBoard(pBrain.PPosSys.PPosition[0], pPosCopy)

		// DoMove と UndoMove を繰り返していると、ずれてくる（＾～＾）
		if pBrain.PPosSys.PPosition[POS_LAYER_MAIN].IsEmptySq(from) {
			// 強制終了した局面（＾～＾）
			App.Out.Debug(Sprint(
				pBrain.PPosSys.PPosition[POS_LAYER_MAIN],
				pBrain.PPosSys.phase,
				pBrain.PPosSys.StartMovesNum,
				pBrain.PPosSys.OffsetMovesIndex,
				pBrain.PPosSys.createMovesText()))
			// あの駒、どこにいんの（＾～＾）？
			App.Out.Debug(pBrain.PPosSys.PPosition[POS_LAYER_MAIN].SprintLocation())
			panic(App.LogNotEcho.Fatal("Move.Source(%d) has empty square. i=%d/%d. younger_sibling_move=%s",
				from, i, lenOfMoves, ToMCode(younger_sibling_move)))
		}

		pBrain.DoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex-1]

		var leaf = false

		if pBrain.IsCheckmate(FlipPhase(pBrain.PPosSys.phase)) {
			// ここで指した方の玉に王手がかかるようなら、被空き王手（＾～＾）
			// この手は見なかったことにするぜ（＾～＾）
		} else if What(captured) == PIECE_TYPE_K {
			// 玉を取るのは最善手
			someBestMoves = nil
			someBestMoves = append(someBestMoves, move)
			opponentWorstVal = RESIGN_VALUE
			cutting = CuttingKingCapture
			/*
				} else if search_type == SEARCH_CAPTURE && captured == PIECE_EMPTY {
					// 駒の取り合いを探索中に、駒を取らなかったら
					// ただの葉
					leaf = true
			*/
		} else {
			// 駒を取っている場合は、探索を延長します
			if curDepth < depthEnd { //  || captured != PIECE_EMPTY
				/*
					var search_type2 SearchType
					if captured != PIECE_EMPTY {
						search_type2 = SEARCH_CAPTURE
					} else {
						search_type2 = search_type
					}
				*/

				// 再帰
				_, opponentVal := search2(pBrain, curDepth+1) //search_type2
				// 再帰直後（＾～＾）
				// App.Out.Debug(pBrain.PPosSys.Sprint(POS_LAYER_MAIN))

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
				// 葉ノード
				leaf = true
			}
		}

		if leaf {
			// 葉ノード
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

		pBrain.UndoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN])

		// 盤と、コピー盤を比較します
		diffBoard(pBrain.PPosSys.PPosition[0], pPosCopy, pBrain.PPosSys.PPosition[2], pBrain.PPosSys.PPosition[3])
		// 異なる箇所を数えます
		errorNum := errorBoard(pBrain.PPosSys.PPosition[0], pPosCopy, pBrain.PPosSys.PPosition[2], pBrain.PPosSys.PPosition[3])
		if errorNum != 0 {
			// 違いのあった局面（＾～＾）
			App.Out.Debug(pBrain.PPosSys.SprintDiff(0, 1))
			// あの駒、どこにいんの（＾～＾）？
			App.Out.Debug(pBrain.PPosSys.PPosition[0].SprintLocation())
			App.Out.Debug(pPosCopy.SprintLocation())
			panic(App.LogNotEcho.Fatal("Error: count=%d younger_sibling_move=%s move=%s", errorNum, ToMCode(younger_sibling_move), ToMCode(move)))
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
			App.Out.Debug("info string Debug: depth=%d nodes=%d value=%d move.best=%s.%s\n", curDepth, nodesNum, -opponentWorstVal, move.ToCode(), debugBestMove.ToCode())
			// Debug ここまで
		*/
	}

	// bestMoveは、１手目しか使わないけど（＾～＾）
	var bestMove = RESIGN_MOVE

	bestmoveListLen := len(someBestMoves)
	//fmt.Printf("%d/%d bestmoveListLen=%d\n", curDepth, depthEnd, bestmoveListLen)
	if bestmoveListLen < 1 {
		// 指せる手なし
		return RESIGN_MOVE, RESIGN_VALUE
	}
	bestMove = someBestMoves[rand.Intn(bestmoveListLen)]
	// 評価値出力（＾～＾）
	// App.Out.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, bestMove.ToCode(), bestMove.ToCode())

	// 相手の評価値の逆が、自分の評価値
	return bestMove, -opponentWorstVal
}
