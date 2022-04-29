package take7

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Who - 駒が先手か後手か空升かを返します
func Who(piece string) l06.Phase {
	switch piece {
	case l03.PIECE_EMPTY.ToCodeOfPc(): // 空きマス
		return l06.ZEROTH
	case l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_S1.ToCodeOfPc(), l03.PIECE_N1.ToCodeOfPc(), l03.PIECE_L1.ToCodeOfPc(), l03.PIECE_P1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc(), l03.PIECE_PN1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc():
		return l06.FIRST
	case l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_S2.ToCodeOfPc(), l03.PIECE_N2.ToCodeOfPc(), l03.PIECE_L2.ToCodeOfPc(), l03.PIECE_P2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc(), l03.PIECE_PN2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc():
		return l06.SECOND
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}
