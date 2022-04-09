package take15

import (
	"fmt"

	l13 "github.com/muzudho/kifuwarabe-wcsc31/take13"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// Who - 駒が先手か後手か空升かを返します
func Who(piece l09.Piece) Phase {
	switch piece {
	case l13.PIECE_EMPTY: // 空きマス
		return ZEROTH
	case l13.PIECE_K1, l13.PIECE_R1, l13.PIECE_B1, l13.PIECE_G1, l13.PIECE_S1, l13.PIECE_N1, l13.PIECE_L1, l13.PIECE_P1, l13.PIECE_PR1, l13.PIECE_PB1, l13.PIECE_PS1, l13.PIECE_PN1, l13.PIECE_PL1, l13.PIECE_PP1:
		return FIRST
	case l13.PIECE_K2, l13.PIECE_R2, l13.PIECE_B2, l13.PIECE_G2, l13.PIECE_S2, l13.PIECE_N2, l13.PIECE_L2, l13.PIECE_P2, l13.PIECE_PR2, l13.PIECE_PB2, l13.PIECE_PS2, l13.PIECE_PN2, l13.PIECE_PL2, l13.PIECE_PP2:
		return SECOND
	default:
		panic(fmt.Errorf("unknown piece=[%d]", piece))
	}
}
