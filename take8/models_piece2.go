package take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// Promote - 成ります
func Promote(piece string) string {
	switch piece {
	case l03.PIECE_EMPTY, l03.PIECE_K1, l03.PIECE_G1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_PS1, l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1,
		l03.PIECE_K2, l03.PIECE_G2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_PS2, l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2: // 成らずにそのまま返す駒
		return piece
	case l03.PIECE_R1:
		return l03.PIECE_PR1
	case l03.PIECE_B1:
		return l03.PIECE_PB1
	case l03.PIECE_S1:
		return l03.PIECE_PS1
	case l03.PIECE_N1:
		return l03.PIECE_PN1
	case l03.PIECE_L1:
		return l03.PIECE_PL1
	case l03.PIECE_P1:
		return l03.PIECE_PP1
	case l03.PIECE_R2:
		return l03.PIECE_PR2
	case l03.PIECE_B2:
		return l03.PIECE_PB2
	case l03.PIECE_S2:
		return l03.PIECE_PS2
	case l03.PIECE_N2:
		return l03.PIECE_PN2
	case l03.PIECE_L2:
		return l03.PIECE_PL2
	case l03.PIECE_P2:
		return l03.PIECE_PP2
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}

// Demote - 成っている駒を、成っていない駒に戻します
func Demote(piece string) string {
	switch piece {
	case l03.PIECE_EMPTY, l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_B1, l03.PIECE_G1, l03.PIECE_S1, l03.PIECE_N1, l03.PIECE_L1, l03.PIECE_P1,
		l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_B2, l03.PIECE_G2, l03.PIECE_S2, l03.PIECE_N2, l03.PIECE_L2, l03.PIECE_P2: // 裏返さずにそのまま返す駒
		return piece
	case l03.PIECE_PR1:
		return l03.PIECE_P1
	case l03.PIECE_PB1:
		return l03.PIECE_B1
	case l03.PIECE_PS1:
		return l03.PIECE_S1
	case l03.PIECE_PN1:
		return l03.PIECE_N1
	case l03.PIECE_PL1:
		return l03.PIECE_L1
	case l03.PIECE_PP1:
		return l03.PIECE_P1
	case l03.PIECE_PR2:
		return l03.PIECE_R2
	case l03.PIECE_PB2:
		return l03.PIECE_B2
	case l03.PIECE_PS2:
		return l03.PIECE_S2
	case l03.PIECE_PN2:
		return l03.PIECE_N2
	case l03.PIECE_PL2:
		return l03.PIECE_L2
	case l03.PIECE_PP2:
		return l03.PIECE_P2
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}
