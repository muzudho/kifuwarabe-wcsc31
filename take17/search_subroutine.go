package take17

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
	l08 "github.com/muzudho/kifuwarabe-wcsc31/take8"
)

func subCopyBoard(pNerve *Nerve) *l15.Position {
	// デバッグに使うために、盤をコピーしておきます
	var pPosCopy = l15.NewPosition()
	copyBoard(pNerve.PPosSys.PPosition[0], pPosCopy)

	return pPosCopy
}

func subErrorBoard(pNerve *Nerve) {
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

func subErrorBoardAfterUndoMove(pNerve *Nerve, pPosCopy *l15.Position, move l03.Move) {
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
