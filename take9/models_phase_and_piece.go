package take9

import (
	"fmt"

	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Who - 駒が先手か後手か空升かを返します
func Who(piece Piece) l06.Phase {
	switch piece {
	case PIECE_EMPTY: // 空きマス
		return l06.ZEROTH
	case PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1, PIECE_PR1, PIECE_PB1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1:
		return l06.FIRST
	case PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2, PIECE_PR2, PIECE_PB2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		return l06.SECOND
	default:
		panic(fmt.Errorf("unknown piece=[%d]", piece))
	}
}
