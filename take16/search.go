package take16

import (
	"math/rand"

	b "github.com/muzudho/kifuwarabe-wcsc31/take16base"
	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// 0 にすると 1手読み（＾～＾）
// 1 の 2手読みにしておくと、玉を取りに行くぜ（＾～＾）
// 2 の 3手読みだと駒を取らない（＾～＾）駒のただ捨てをする（＾～＾）駒をとりかえさない（＾～＾）
// 3 の 4手読みは、まだ遅い（＾～＾）
var maxDepth int = 1

type SearchType uint8

// 探索への指定
const (
	// 特になし
	SEARCH_NONE = SearchType(0)
	// 駒の取り合い
	SEARCH_CAPTURE = SearchType(1)
)

// 最大限に使わなくても、十分に大きければ十分だが（＾～＾）
const VALUE_INFINITE = 2_147_483_647

var nodesNum int

type CuttingType int

const (
	CuttingNone = CuttingType(0)
	// 玉を取った
	CuttingKingCapture = CuttingType(1)
	// 駒の取り合いが終わった
	//CuttingEndCapture = CuttingType(2)
)

// SearchEntry - 探索部（開始）
func SearchEntry(pNerve *Nerve) b.Move {

	nodesNum = 0
	depth := maxDepth
	//fmt.Printf("Search: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	var alpha p.Value = -VALUE_INFINITE
	var beta p.Value = VALUE_INFINITE

	nodeValue, bestMove := search(pNerve, alpha, beta, depth, SEARCH_NONE)

	// 評価値出力（＾～＾）
	G.Chat.Print("info depth %d nodes %d score cp %d currmove %s pv %s\n",
		depth, nodesNum, nodeValue, p.ToMoveCode(bestMove), p.ToMoveCode(bestMove))

	// ゲーム向けの軽い乱数
	return bestMove
}

// search - 探索部
func search(pNerve *Nerve, alpha p.Value, beta p.Value, depth int, search_type SearchType) (p.Value, b.Move) {
	//fmt.Printf("Search2: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	// TODO 葉ノード
	if depth <= 0 {
		// 葉ノード
		// 駒割り評価値は、相手の手番のものになっています。
		materialValue := pNerve.PPosSys.PPosition[POS_LAYER_MAIN].MaterialValue
		//fmt.Printf("move=%s leafVal=%6d materialVal=%6d(%s) control_val=%6d\n", move.ToCode(), leafVal, materialVal, captured.ToCode(), control_val)

		// 指し手は返さないぜ（＾～＾）
		return materialValue, p.RESIGN_MOVE
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	someMoves := GenMoveList(pNerve, pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
	lenOfMoves := len(someMoves)
	//fmt.Printf("%d/%d lenOfMoves=%d\n", curDepth, depthEnd, lenOfMoves)

	if lenOfMoves == 0 {
		// ステイルメートされたら負け（＾～＾）
		return -VALUE_INFINITE, p.RESIGN_MOVE
	}

	// 同じ価値のベストムーブがいっぱいあるかも（＾～＾）
	var someBestMoves []b.Move

	// 前回のムーブ
	var younger_sibling_move = p.RESIGN_MOVE
	// 探索終了
	var cutting = CuttingNone

	// その手を指してみるぜ（＾～＾）
	for i, move := range someMoves {
		// G.Chat.Debug("move=%s\n", move.ToCode())
		from, _, _ := p.DestructureMove(move)

		// デバッグに使うために、盤をコピーしておきます
		pPosCopy := p.NewPosition()
		copyBoard(pNerve.PPosSys.PPosition[0], pPosCopy)

		// DoMove と UndoMove を繰り返していると、ずれてくる（＾～＾）
		if pNerve.PPosSys.PPosition[POS_LAYER_MAIN].IsEmptySq(from) {
			// 強制終了した局面（＾～＾）
			G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoardHeader(
				pNerve.PPosSys.phase,
				pNerve.PRecord.StartMovesNum,
				pNerve.PRecord.OffsetMovesIndex))
			G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoard())
			G.Chat.Debug(pNerve.SprintBoardFooter())
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintLocation())
			panic(G.Log.Fatal("Move.Source(%d) has empty square. i=%d/%d. younger_sibling_move=%s",
				from, i, lenOfMoves, p.ToMoveCode(younger_sibling_move)))
		}

		pNerve.DoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pNerve.PRecord.CapturedList[pNerve.PRecord.OffsetMovesIndex-1]

		var edgeValue p.Value = -VALUE_INFINITE

		if pNerve.IsCheckmate(FlipPhase(pNerve.PPosSys.phase)) {
			// ここで指した方の玉に王手がかかるようなら、被空き王手（＾～＾）
			// この手は見なかったことにするぜ（＾～＾）
		} else if What(captured) == PIECE_TYPE_K {
			// 玉を取るのは最善手
			someBestMoves = nil
			someBestMoves = append(someBestMoves, move)
			edgeValue = VALUE_INFINITE
			cutting = CuttingKingCapture
		} else if search_type == SEARCH_CAPTURE && captured == p.PIECE_EMPTY {
			// 駒の取り合いを探索中に、駒を取らなかったら
			// この手は見なかったことにするぜ（＾～＾）
		} else {
			// 駒を取っている場合は、探索を延長します
			// TODO  if captured != p.PIECE_EMPTY
			var search_type2 SearchType
			if captured != p.PIECE_EMPTY {
				search_type2 = SEARCH_CAPTURE
			} else {
				search_type2 = search_type
			}

			// 再帰
			nodeValue, _ := search(pNerve, -beta, -alpha, depth-1, search_type2)
			edgeValue = -nodeValue
			// 再帰直後（＾～＾）
			// G.Chat.Debug(pNerve.PPosSys.Sprint(POS_LAYER_MAIN))

			if alpha < edgeValue {
				// アルファー・アップデート
				someBestMoves = nil
				someBestMoves = append(someBestMoves, move)
				alpha = edgeValue
			} else if alpha <= edgeValue {
				// アルファー・アップデートしないが、同着なら配列の要素として追加
				someBestMoves = append(someBestMoves, move)
			}
		}

		// ベーター・カット
		if beta < alpha {
			return -VALUE_INFINITE, p.RESIGN_MOVE
		}

		pNerve.UndoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN])

		// 盤と、コピー盤を比較します
		diffBoard(pNerve.PPosSys.PPosition[0], pPosCopy, pNerve.PPosSys.PPosition[2], pNerve.PPosSys.PPosition[3])
		// 異なる箇所を数えます
		errorNum := errorBoard(pNerve.PPosSys.PPosition[0], pPosCopy, pNerve.PPosSys.PPosition[2], pNerve.PPosSys.PPosition[3])
		if errorNum != 0 {
			// 違いのあった局面（＾～＾）
			G.Chat.Debug(sprintPositionDiff(pNerve.PPosSys, 0, 1, pNerve.PRecord))
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pNerve.PPosSys.PPosition[0].SprintLocation())
			G.Chat.Debug(pPosCopy.SprintLocation())
			panic(G.Log.Fatal("Error: count=%d younger_sibling_move=%s move=%s", errorNum, p.ToMoveCode(younger_sibling_move), p.ToMoveCode(move)))
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

	// bestMoveは、１手目しか使わないけど（＾～＾）
	var bestMove = p.RESIGN_MOVE

	bestmoveListLen := len(someBestMoves)
	//fmt.Printf("%d/%d bestmoveListLen=%d\n", curDepth, depthEnd, bestmoveListLen)
	if bestmoveListLen < 1 {
		// 指せる手なし
		return -VALUE_INFINITE, p.RESIGN_MOVE
	}
	bestMove = someBestMoves[rand.Intn(bestmoveListLen)]
	// 評価値出力（＾～＾）
	// G.Chat.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, bestMove.ToCode(), bestMove.ToCode())

	return alpha, bestMove
}
