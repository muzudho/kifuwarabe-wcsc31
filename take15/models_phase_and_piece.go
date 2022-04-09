package take15

import (
	"fmt"

	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// Who - 駒が先手か後手か空升かを返します
func Who(piece l09.Piece) Phase {
	switch piece {
	case l11.PIECE_EMPTY: // 空きマス
		return ZEROTH
	case l11.PIECE_K1, l11.PIECE_R1, l11.PIECE_B1, l11.PIECE_G1, l11.PIECE_S1, l11.PIECE_N1, l11.PIECE_L1, l11.PIECE_P1, l11.PIECE_PR1, l11.PIECE_PB1, l11.PIECE_PS1, l11.PIECE_PN1, l11.PIECE_PL1, l11.PIECE_PP1:
		return FIRST
	case l11.PIECE_K2, l11.PIECE_R2, l11.PIECE_B2, l11.PIECE_G2, l11.PIECE_S2, l11.PIECE_N2, l11.PIECE_L2, l11.PIECE_P2, l11.PIECE_PR2, l11.PIECE_PB2, l11.PIECE_PS2, l11.PIECE_PN2, l11.PIECE_PL2, l11.PIECE_PP2:
		return SECOND
	default:
		panic(fmt.Errorf("unknown piece=[%d]", piece))
	}
}
