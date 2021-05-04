package take12

import (
	"fmt"
	"math"
)

// SprintControl - 利き数ボード出力（＾ｑ＾）
//
// Parameters
// ----------
// * `c` - 利き数ボードのレイヤー番号（＾～＾）
func (pPosSys *PositionSystem) SprintControl(c ControlLayerT) string {
	title := fmt.Sprintf("Control(%d)%s", c, pPosSys.PControlBoardSystem.Boards[c].Title)
	board := pPosSys.PControlBoardSystem.Boards[c].Board

	// 表示桁数を調べます
	max_num := int8(math.MinInt8)
	min_num := int8(math.MaxInt8)
	for _, val := range board {
		if max_num < val {
			max_num = val
		}

		if min_num > val {
			min_num = val
		}
	}

	if -10 < min_num && max_num < 100 {
		// 表示桁数２桁
		return "\n" +
			//
			fmt.Sprintf("[%s]\n", title) +
			//
			"\n" +
			//
			"  9  8  7  6  5  4  3  2  1\n" +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| a\n", board[91], board[81], board[71], board[61], board[51], board[41], board[31], board[21], board[11]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| b\n", board[92], board[82], board[72], board[62], board[52], board[42], board[32], board[22], board[12]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| c\n", board[93], board[83], board[73], board[63], board[53], board[43], board[33], board[23], board[13]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| d\n", board[94], board[84], board[74], board[64], board[54], board[44], board[34], board[24], board[14]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| e\n", board[95], board[85], board[75], board[65], board[55], board[45], board[35], board[25], board[15]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| f\n", board[96], board[86], board[76], board[66], board[56], board[46], board[36], board[26], board[16]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| g\n", board[97], board[87], board[77], board[67], board[57], board[47], board[37], board[27], board[17]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| h\n", board[98], board[88], board[78], board[68], board[58], board[48], board[38], board[28], board[18]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d|%2d| i\n", board[99], board[89], board[79], board[69], board[59], board[49], board[39], board[29], board[19]) +
			//
			"+--+--+--+--+--+--+--+--+--+\n" +
			//
			"\n"
	} else {
		// 表示桁数４桁 Example: '-100'
		return "\n" +
			//
			fmt.Sprintf("[%s]\n", title) +
			//
			"\n" +
			//
			"    9    8    7    6    5    4    3    2    1\n" +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| a\n", board[91], board[81], board[71], board[61], board[51], board[41], board[31], board[21], board[11]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| b\n", board[92], board[82], board[72], board[62], board[52], board[42], board[32], board[22], board[12]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| c\n", board[93], board[83], board[73], board[63], board[53], board[43], board[33], board[23], board[13]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| d\n", board[94], board[84], board[74], board[64], board[54], board[44], board[34], board[24], board[14]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| e\n", board[95], board[85], board[75], board[65], board[55], board[45], board[35], board[25], board[15]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| f\n", board[96], board[86], board[76], board[66], board[56], board[46], board[36], board[26], board[16]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| g\n", board[97], board[87], board[77], board[67], board[57], board[47], board[37], board[27], board[17]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| h\n", board[98], board[88], board[78], board[68], board[58], board[48], board[38], board[28], board[18]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			fmt.Sprintf("|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d|%4d| i\n", board[99], board[89], board[79], board[69], board[59], board[49], board[39], board[29], board[19]) +
			//
			"+----+----+----+----+----+----+----+----+----+\n" +
			//
			"\n"
	}
}
