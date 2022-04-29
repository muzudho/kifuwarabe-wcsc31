package take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Who - 駒が先手か後手か空升かを返します
func Who(piece l03.Piece) l06.Phase {
	switch piece {
	case l03.PIECE_EMPTY: // 空きマス
		return l06.ZEROTH
	case l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_B1, l03.PIECE_G1, l03.PIECE_S1, l03.PIECE_N1, l03.PIECE_L1, l03.PIECE_P1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_PS1, l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1:
		return l06.FIRST
	case l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_B2, l03.PIECE_G2, l03.PIECE_S2, l03.PIECE_N2, l03.PIECE_L2, l03.PIECE_P2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_PS2, l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2:
		return l06.SECOND
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece.ToCodeOfPc()))
	}
}
