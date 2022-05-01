package take16

import (
	"bytes"
	"fmt"

	l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"
	l08 "github.com/muzudho/kifuwarabe-wcsc31/take8"
)

// Dump - 内部状態を全部出力しようぜ（＾～＾）？
func (pNerve *Nerve) Dump() string {
	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buffer bytes.Buffer

	// Each position
	for b := PosLayerT(0); b < 2; b += 1 {
		pPos := pNerve.PPosSys.PPosition[b]
		buffer.WriteString(fmt.Sprintf("Position[%d]:\n", b))

		// Board, Hands
		buffer.WriteString(pPos.SprintBoard())
		buffer.WriteString("\n")

		// PieceLocation
		l08.SprintLocation(pPos)
	}

	// 利きボード全部
	for c := ControlLayerT(0); c < CONTROL_LAYER_ALL_SIZE; c += 1 {
		buffer.WriteString(pNerve.PCtrlBrdSys.SprintControl(c))
	}

	buffer.WriteString(fmt.Sprintf("Phase:%d,\n", pNerve.PPosSys.GetPhase()))

	buffer.WriteString(fmt.Sprintf("StartMovesNum:%d,\n", pNerve.PRecord.StartMovesNum))

	buffer.WriteString(fmt.Sprintf("OffsetMovesIndex:%d,\n", pNerve.PRecord.OffsetMovesIndex))

	// moves
	buffer.WriteString(pNerve.SprintBoardFooter())

	buffer.WriteString("CapturedList:")
	for i := 0; i < l02.MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pNerve.PRecord.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
