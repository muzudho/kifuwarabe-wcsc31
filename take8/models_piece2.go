package take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// Promote - 成ります
func Promote(piece string) string {
	switch piece {
	case l03.PIECE_EMPTY.ToCodeOfPc(), l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc(), l03.PIECE_PN1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc(),
		l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc(), l03.PIECE_PN2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc(): // 成らずにそのまま返す駒
		return piece
	case l03.PIECE_R1.ToCodeOfPc():
		return l03.PIECE_PR1.ToCodeOfPc()
	case l03.PIECE_B1.ToCodeOfPc():
		return l03.PIECE_PB1.ToCodeOfPc()
	case l03.PIECE_S1.ToCodeOfPc():
		return l03.PIECE_PS1.ToCodeOfPc()
	case l03.PIECE_N1.ToCodeOfPc():
		return l03.PIECE_PN1.ToCodeOfPc()
	case l03.PIECE_L1.ToCodeOfPc():
		return l03.PIECE_PL1.ToCodeOfPc()
	case l03.PIECE_P1.ToCodeOfPc():
		return l03.PIECE_PP1.ToCodeOfPc()
	case l03.PIECE_R2.ToCodeOfPc():
		return l03.PIECE_PR2.ToCodeOfPc()
	case l03.PIECE_B2.ToCodeOfPc():
		return l03.PIECE_PB2.ToCodeOfPc()
	case l03.PIECE_S2.ToCodeOfPc():
		return l03.PIECE_PS2.ToCodeOfPc()
	case l03.PIECE_N2.ToCodeOfPc():
		return l03.PIECE_PN2.ToCodeOfPc()
	case l03.PIECE_L2.ToCodeOfPc():
		return l03.PIECE_PL2.ToCodeOfPc()
	case l03.PIECE_P2.ToCodeOfPc():
		return l03.PIECE_PP2.ToCodeOfPc()
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}

// Demote - 成っている駒を、成っていない駒に戻します
func Demote(piece string) string {
	switch piece {
	case l03.PIECE_EMPTY.ToCodeOfPc(), l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_S1.ToCodeOfPc(), l03.PIECE_N1.ToCodeOfPc(), l03.PIECE_L1.ToCodeOfPc(), l03.PIECE_P1.ToCodeOfPc(),
		l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_S2.ToCodeOfPc(), l03.PIECE_N2.ToCodeOfPc(), l03.PIECE_L2.ToCodeOfPc(), l03.PIECE_P2.ToCodeOfPc(): // 裏返さずにそのまま返す駒
		return piece
	case l03.PIECE_PR1.ToCodeOfPc():
		return l03.PIECE_P1.ToCodeOfPc()
	case l03.PIECE_PB1.ToCodeOfPc():
		return l03.PIECE_B1.ToCodeOfPc()
	case l03.PIECE_PS1.ToCodeOfPc():
		return l03.PIECE_S1.ToCodeOfPc()
	case l03.PIECE_PN1.ToCodeOfPc():
		return l03.PIECE_N1.ToCodeOfPc()
	case l03.PIECE_PL1.ToCodeOfPc():
		return l03.PIECE_L1.ToCodeOfPc()
	case l03.PIECE_PP1.ToCodeOfPc():
		return l03.PIECE_P1.ToCodeOfPc()
	case l03.PIECE_PR2.ToCodeOfPc():
		return l03.PIECE_R2.ToCodeOfPc()
	case l03.PIECE_PB2.ToCodeOfPc():
		return l03.PIECE_B2.ToCodeOfPc()
	case l03.PIECE_PS2.ToCodeOfPc():
		return l03.PIECE_S2.ToCodeOfPc()
	case l03.PIECE_PN2.ToCodeOfPc():
		return l03.PIECE_N2.ToCodeOfPc()
	case l03.PIECE_PL2.ToCodeOfPc():
		return l03.PIECE_L2.ToCodeOfPc()
	case l03.PIECE_PP2.ToCodeOfPc():
		return l03.PIECE_P2.ToCodeOfPc()
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}
