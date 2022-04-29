package take17

import (
	"math/rand"
	"strconv"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l13 "github.com/muzudho/kifuwarabe-wcsc31/take13"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
	l08 "github.com/muzudho/kifuwarabe-wcsc31/take8"
)

type SearchType uint8

// 探索への指定
const (
	// 特になし
	SEARCH_NONE = SearchType(0)
	// 駒の取り合い
	SEARCH_CAPTURE = SearchType(1)
)

// 最大限に使わなくても、十分に大きければ十分だが（＾～＾）
const VALUE_INFINITE_1 = 1_000_001

var nodesNum int

type CuttingType int

const (
	CuttingNone = CuttingType(0)
	// 玉を取った
	CuttingKingCapture = CuttingType(1)
	// 駒の取り合いが終わった
	//CuttingEndCapture = CuttingType(2)
)

// IterativeDeepeningSearch - 探索部（開始）
func IterativeDeepeningSearch(pNerve *Nerve, tokens []string) l13.Move {
	pNerve.ClearBySearchEntry()
	// # Example
	//
	// ```
	// go btime 60000 wtime 50000 byoyomi 10000
	// .  .     2     .     4     .       6
	//
	// go btime 40000 wtime 50000 binc 10000 winc 10000
	// .  .     2     .     4     .    6     .    8
	// ```

	// パース
	var btime int = 0
	var wtime int = 0
	var byoyomi int = 0
	var binc int = 0
	var winc int = 0
	if 8 <= len(tokens) && tokens[5] == "binc" {
		// フィッシャー・クロック・ルール
		btime, _ = strconv.Atoi(tokens[2])
		wtime, _ = strconv.Atoi(tokens[4])
		byoyomi = 0
		binc, _ = strconv.Atoi(tokens[6])
		winc, _ = strconv.Atoi(tokens[8])
	} else if 6 <= len(tokens) {
		// 秒読みルール
		btime, _ = strconv.Atoi(tokens[2])
		wtime, _ = strconv.Atoi(tokens[4])
		byoyomi, _ = strconv.Atoi(tokens[6])
		binc = 0
		winc = 0
	}

	var time_sec int = 0
	var inc_sec int = 0
	switch pNerve.PPosSys.phase {
	case 1:
		time_sec = (btime + byoyomi + binc) / 1000
		inc_sec = binc / 1000
	case 2:
		time_sec = (wtime + byoyomi + winc) / 1000
		inc_sec = winc / 1000
	}

	if pNerve.OneMoveSec == 0 {
		// 対局開始時に計算しておく
		// １手に割り当てる消費時間
		pNerve.OneMoveSec = time_sec / 130
	}

	// 2手目以降
	var think_sec = pNerve.OneMoveSec

	if 0 < inc_sec {
		// フィッシャー・クロック・ルール:
		// 最低でも （加算時間-1秒）は使おう
		if 1 < inc_sec && think_sec < inc_sec-1 {
			think_sec = (inc_sec - 1)
		}
	} else {
		// 最低でも 1秒は使おう
		if think_sec < 1 {
			think_sec = 1
		}
	}

	nodesNum = 0

	var alpha l15.Value = -VALUE_INFINITE_1
	var beta l15.Value = VALUE_INFINITE_1
	var bestValue l15.Value = -VALUE_INFINITE_1
	var bestMove l13.Move = l13.RESIGN_MOVE // 指し手が無いとき投了

	// Iterative Deepening
	for depth := 1; depth < pNerve.MaxDepth+1; depth += 1 {
		value, move := search(pNerve, alpha, beta, depth, SEARCH_NONE)
		if pNerve.IsStopSearch {
			// タイムアップしたときの探索結果は使わないぜ（＾～＾）
			// 評価値出力（＾～＾）
			App.Out.Print("# Time up\n")
			break
		} else {
			bestValue = value
			bestMove = move
		}

		// 評価値出力（＾～＾）
		App.Out.Print("info depth %d nodes %d score cp %d currmove %s pv %s\n",
			depth, nodesNum, bestValue, bestMove.ToCodeOfM(), bestMove.ToCodeOfM())
	}
	//fmt.Printf("Search: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	// ゲーム向けの軽い乱数
	return bestMove
}

// search - 探索部
func search(pNerve *Nerve, alpha l15.Value, beta l15.Value, depth int, search_type SearchType) (l15.Value, l13.Move) {
	//fmt.Printf("Search2: depth=%d/%d nodesNum=%d\n", curDepth, depthEnd, nodesNum)

	// TODO 葉ノード
	if depth <= 0 {
		// 葉ノード
		// 駒割り評価値は、相手の手番のものになっています。
		materialValue := pNerve.PPosSys.PPosition[POS_LAYER_MAIN].MaterialValue
		//fmt.Printf("move=%s leafVal=%6d materialVal=%6d(%s) control_val=%6d\n", move.ToCode(), leafVal, materialVal, captured.ToCode(), control_val)

		return materialValue, l13.RESIGN_MOVE // 葉では、指し手は使わないから、返さないぜ（＾～＾）
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	someMoves := GenMoveList(pNerve, pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
	lenOfMoves := len(someMoves)
	//fmt.Printf("%d/%d lenOfMoves=%d\n", curDepth, depthEnd, lenOfMoves)

	if lenOfMoves == 0 {
		return -VALUE_INFINITE_1, l13.RESIGN_MOVE // ステイルメート（指し手がない）されたら投了（＾～＾）
	}

	// 同じ価値のベストムーブがいっぱいあるかも（＾～＾）
	var someBestMoves []l13.Move

	// 前回のムーブ（デバッグ用）
	// var younger_sibling_move = l13.RESIGN_MOVE
	// 探索終了
	var cutting = CuttingNone

	// その手を指してみるぜ（＾～＾）
	for i, move := range someMoves {
		// TODO タイムアップ判定（＾～＾）
		sec := pNerve.PStopwatchSearch.ElapsedSeconds()
		if sec >= 20.0 {
			App.Out.Print("# Time up. sec=%d\n", sec)
			pNerve.IsStopSearch = true
			return -VALUE_INFINITE_1, l13.RESIGN_MOVE // タイムアップしたときの探索結果は使わないぜ（＾～＾）
		}

		// App.Out.Debug("move=%s\n", move.ToCode())
		from, _, _ := move.Destructure()

		var pPosCopy *l15.Position
		if App.IsDebug {
			// デバッグに使うために、盤をコピーしておきます
			pPosCopy = l15.NewPosition()
			copyBoard(pNerve.PPosSys.PPosition[0], pPosCopy)
		}

		// DoMove と UndoMove を繰り返していると、ずれてくる（＾～＾）
		if pNerve.PPosSys.PPosition[POS_LAYER_MAIN].IsEmptySq(from) {
			if App.IsDebug {
				// 強制終了した局面（＾～＾）
				App.Out.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoardHeader(
					pNerve.PPosSys.phase,
					pNerve.PRecord.StartMovesNum,
					pNerve.PRecord.OffsetMovesIndex))
				App.Out.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoard())
				App.Out.Debug(pNerve.SprintBoardFooter())
				// あの駒、どこにいんの（＾～＾）？
				App.Out.Debug(l08.SprintLocation(pNerve.PPosSys.PPosition[POS_LAYER_MAIN]))
			}

			panic(App.LogNotEcho.Fatal("Move.Source(%d) has empty square. i=%d/%d.",
				from, i, lenOfMoves))
			//  younger_sibling_move=%s
			//, ToMoveCode(younger_sibling_move)
		}

		pNerve.DoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], move)
		nodesNum += 1

		// 取った駒は棋譜の１手前に記録されています
		captured := pNerve.PRecord.CapturedList[pNerve.PRecord.OffsetMovesIndex-1]

		if pNerve.IsCheckmate(FlipPhase(pNerve.PPosSys.phase)) {
			// ここで指した方の玉に王手がかかるようなら、被空き王手（＾～＾）
			// この手は見なかったことにするぜ（＾～＾）
		} else if l11.What(captured) == l11.PIECE_TYPE_K {
			// 玉を取るのは最善手
			someBestMoves = nil
			someBestMoves = append(someBestMoves, move)
			cutting = CuttingKingCapture
			alpha = VALUE_INFINITE_1
		} else if search_type == SEARCH_CAPTURE && captured == l03.PIECE_EMPTY {
			// 駒の取り合いを探索中に、駒を取らなかったら
			// この手は見なかったことにするぜ（＾～＾）
		} else {
			// 駒を取っている場合は、探索を延長します
			// TODO  if captured != PIECE_EMPTY
			var search_type2 SearchType
			if captured != l03.PIECE_EMPTY {
				search_type2 = SEARCH_CAPTURE
			} else {
				search_type2 = search_type
			}

			// 再帰
			nodeValue, _ := search(pNerve, -beta, -alpha, depth-1, search_type2)
			var edgeValue = -nodeValue
			// 再帰直後（＾～＾）
			// App.Out.Debug(pNerve.PPosSys.Sprint(POS_LAYER_MAIN))

			// 説明変数：何か１つは指し手を選んでおかないと、投了してしまうから、最初の１手は候補に入れておけだぜ（＾～＾）
			var isAnyOneMove = len(someBestMoves) == 0
			// 説明変数：アルファー・アップデートしないが、同着なら配列の要素として追加
			var isSameAlpha = alpha == edgeValue

			if isAnyOneMove || isSameAlpha {
				someBestMoves = append(someBestMoves, move)
			} else if alpha < edgeValue {
				// アルファー・アップデート
				someBestMoves = nil
				someBestMoves = append(someBestMoves, move)
				alpha = edgeValue
			}
		}

		pNerve.UndoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN])

		if App.IsDebug {
			// 盤と、コピー盤を比較します
			diffBoard(pNerve.PPosSys.PPosition[0], pPosCopy, pNerve.PPosSys.PPosition[2], pNerve.PPosSys.PPosition[3])
			// 異なる箇所を数えます
			errorNum := errorBoard(pNerve.PPosSys.PPosition[0], pPosCopy, pNerve.PPosSys.PPosition[2], pNerve.PPosSys.PPosition[3])
			if errorNum != 0 {
				if App.IsDebug {
					// 違いのあった局面（＾～＾）
					App.Out.Debug(sprintPositionDiff(pNerve.PPosSys, 0, 1, pNerve.PRecord))
					// あの駒、どこにいんの（＾～＾）？
					App.Out.Debug(l08.SprintLocation(pNerve.PPosSys.PPosition[0]))
					App.Out.Debug(l08.SprintLocation(pPosCopy))
				}

				panic(App.LogNotEcho.Fatal("Error: count=%d move=%s", errorNum, move.ToCodeOfM()))
				// younger_sibling_move=%s
				//, ToMoveCode(younger_sibling_move)
			}
		}

		// ベーター・カット
		if beta < alpha {
			// betaより1でもalphaが大きければalphaは使われないから投了を返すぜ（＾～＾）
			return alpha, l13.RESIGN_MOVE
		}

		// younger_sibling_move = move

		if cutting != CuttingNone {
			break
		}

		/*
			// Debug ここから
			var debugBestMove = l13.RESIGN_MOVE
			bestmoveListLen := len(someBestMoves)
			if bestmoveListLen > 0 {
				debugBestMove = someBestMoves[rand.Intn(bestmoveListLen)]
			}
			App.Out.Debug("info string Debug: depth=%d nodes=%d value=%d move.best=%s.%s\n", curDepth, nodesNum, -opponentWorstVal, move.ToCode(), debugBestMove.ToCode())
			// Debug ここまで
		*/
	}

	bestmoveListLen := len(someBestMoves)
	// // Debug出力
	// for i := 0; i < bestmoveListLen; i += 1 {
	// 	App.Out.Debug("i=%d depth=%d bestmoveListLen=%d\n", i, depth, bestmoveListLen)
	// }

	if bestmoveListLen < 1 {
		return -VALUE_INFINITE_1, l13.RESIGN_MOVE // 指せる手無いから投了（＾～＾）
	}
	var bestMove = someBestMoves[rand.Intn(bestmoveListLen)]
	// 評価値出力（＾～＾）
	// App.Out.Print("info depth 0 nodes %d score cp %d currmove %s pv %s\n", nodesNum, bestVal, bestMove.ToCode(), bestMove.ToCode())

	return alpha, bestMove
}
