// 評価関数
package take15

// 評価値の型。
// int16 に収まるように設計できてないので（＾～＾）
type Value int32

// EvalControlVal - 葉局面での利きの評価
func EvalControlVal(pPosSys *PositionSystem) Value {
	var control_val Value = 0

	// 何もしない方がマシかも（＾～＾）
	/*
		switch pPosSys.phase {
		case FIRST:
			WaterColor(
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL1],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL2],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL3])
			my_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[PCLOC_K1]
			oppo_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[PCLOC_K2]
			control_val = pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL3].Board1[my_king_sq] +
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL3].Board1[oppo_king_sq]
		case SECOND:
			WaterColor(
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL1],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL2],
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL3])
			my_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[PCLOC_K2]
			oppo_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[PCLOC_K1]
			control_val = pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL3].Board1[my_king_sq] +
				pPosSys.PCtrlBrdSys.PBoards[CONTROL_LAYER_EVAL3].Board1[oppo_king_sq]
		default:
			panic(fmt.Errorf("Unknown phase=%d", pPosSys.phase))
		}

		// 利き評価が強すぎると 指し手がバラけません
		control_val /= 50

		// 乱数を使って 確率的にします。
		if control_val != 0 {
			var sign int16
			if control_val < 0 {
				sign = -1
			}
			control_val = sign * int16(rand.Intn(int(math.Abs(float64(control_val)))))
		}
	*/

	return control_val
}
