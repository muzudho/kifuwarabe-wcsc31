// 評価関数
package take13

// EvalControlVal - 葉局面での利きの評価
func EvalControlVal(pPosSys *PositionSystem) int16 {
	var control_val int16 = 0

	// 何もしない方がマシかも（＾～＾）
	/*
		switch pPosSys.phase {
		case l03.FIRST:
			WaterColor(
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL2],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL3])
			my_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[l11.PCLOC_K1]
			oppo_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[l11.PCLOC_K2]
			control_val = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL3].Board1[my_king_sq] +
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL3].Board1[oppo_king_sq]
		case l03.SECOND:
			WaterColor(
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL2],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL3])
			my_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[l11.PCLOC_K2]
			oppo_king_sq := pPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[l11.PCLOC_K1]
			control_val = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL3].Board1[my_king_sq] +
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_EVAL3].Board1[oppo_king_sq]
		default:
			panic(fmt.Errorf("unknown phase=%d", pPosSys.phase))
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
