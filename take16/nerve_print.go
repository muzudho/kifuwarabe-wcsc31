package take16

import (
	"bytes"
	"fmt"

	r "github.com/muzudho/kifuwarabe-wcsc31/take16record"
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
		pPos.SprintLocation()
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
	for i := 0; i < r.MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pNerve.PRecord.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}

// SprintBoardFooter - 局面出力（＾ｑ＾）
func (pNerve *Nerve) SprintBoardFooter() string {
	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return "moves" + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return "moves" + createMovesText(pNerve.PRecord) + "\n"
}
