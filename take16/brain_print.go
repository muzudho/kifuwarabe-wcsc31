package take16

import (
	"bytes"
	"fmt"
)

// Dump - 内部状態を全部出力しようぜ（＾～＾）？
func (pBrain *Brain) Dump() string {
	// bytes.Bufferは、速くはないけど使いやすいぜ（＾～＾）
	var buffer bytes.Buffer

	for b := PosLayerT(0); b < 2; b += 1 {
		pPos := pBrain.PPosSys.PPosition[b]
		buffer.WriteString(fmt.Sprintf("Board[%d]:", b))
		for i := 0; i < POS_LAYER_SIZE; i += 1 {
			buffer.WriteString(fmt.Sprintf("%d,", pBrain.PPosSys.PPosition[i].Board))
		}
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("KingLocations[%d]:%d,%d\n", b, pPos.PieceLocations[PCLOC_K1], pPos.PieceLocations[PCLOC_K2]))
		buffer.WriteString(fmt.Sprintf("RookLocations[%d]:%d,%d\n", b, pPos.PieceLocations[PCLOC_R1], pPos.PieceLocations[PCLOC_R2]))
		buffer.WriteString(fmt.Sprintf("BishopLocations[%d]:%d,%d\n", b, pPos.PieceLocations[PCLOC_B1], pPos.PieceLocations[PCLOC_B2]))
		buffer.WriteString(fmt.Sprintf("LanceLocations[%d]:%d,%d,%d,%d\n", b, pPos.PieceLocations[PCLOC_L1], pPos.PieceLocations[PCLOC_L2], pPos.PieceLocations[PCLOC_L3], pPos.PieceLocations[PCLOC_L4]))
	}

	// 利きボード全部
	for c := ControlLayerT(0); c < CONTROL_LAYER_ALL_SIZE; c += 1 {
		buffer.WriteString(pBrain.PCtrlBrdSys.SprintControl(c))
	}

	for b := PosLayerT(0); b < 2; b += 1 {
		buffer.WriteString(fmt.Sprintf("Position[%d]:", b))
		buffer.WriteString("Hands:")
		for i := HAND_IDX_START; i < HAND_IDX_END; i += 1 {
			buffer.WriteString(fmt.Sprintf("%d,", pBrain.PPosSys.PPosition[b].Hands1[i]))
		}
		buffer.WriteString("\n")
	}

	buffer.WriteString(fmt.Sprintf("Phase:%d,\n", pBrain.PPosSys.GetPhase()))

	buffer.WriteString(fmt.Sprintf("StartMovesNum:%d,\n", pBrain.PPosSys.StartMovesNum))

	buffer.WriteString(fmt.Sprintf("OffsetMovesIndex:%d,\n", pBrain.PPosSys.OffsetMovesIndex))

	buffer.WriteString("Moves:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pBrain.PPosSys.Moves[i]))
	}
	buffer.WriteString("\n")

	buffer.WriteString("CapturedList:")
	for i := 0; i < MOVES_SIZE; i += 1 {
		buffer.WriteString(fmt.Sprintf("%d,", pBrain.PPosSys.CapturedList[i]))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
