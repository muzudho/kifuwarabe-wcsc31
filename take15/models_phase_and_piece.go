package take15

import (
	"fmt"

	l10 "github.com/muzudho/kifuwarabe-wcsc31/take10"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// Who - 駒が先手か後手か空升かを返します
func Who(piece l09.Piece) Phase {
	switch piece {
	case l10.PIECE_EMPTY: // 空きマス
		return ZEROTH
	case l10.PIECE_K1, l10.PIECE_R1, l10.PIECE_B1, l10.PIECE_G1, l10.PIECE_S1, l10.PIECE_N1, l10.PIECE_L1, l10.PIECE_P1, l10.PIECE_PR1, l10.PIECE_PB1, l10.PIECE_PS1, l10.PIECE_PN1, l10.PIECE_PL1, l10.PIECE_PP1:
		return FIRST
	case l10.PIECE_K2, l10.PIECE_R2, l10.PIECE_B2, l10.PIECE_G2, l10.PIECE_S2, l10.PIECE_N2, l10.PIECE_L2, l10.PIECE_P2, l10.PIECE_PR2, l10.PIECE_PB2, l10.PIECE_PS2, l10.PIECE_PN2, l10.PIECE_PL2, l10.PIECE_PP2:
		return SECOND
	default:
		panic(fmt.Errorf("unknown piece=[%d]", piece))
	}
}
