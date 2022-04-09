package take12

import (
	"fmt"

	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// 駒
const (
	PIECE_EMPTY = iota
	PIECE_K1
	PIECE_R1
	PIECE_B1
	PIECE_G1
	PIECE_S1
	PIECE_N1
	PIECE_L1
	PIECE_P1
	PIECE_PR1
	PIECE_PB1
	PIECE_PS1
	PIECE_PN1
	PIECE_PL1
	PIECE_PP1
	PIECE_K2
	PIECE_R2
	PIECE_B2
	PIECE_G2
	PIECE_S2
	PIECE_N2
	PIECE_L2
	PIECE_P2
	PIECE_PR2
	PIECE_PB2
	PIECE_PS2
	PIECE_PN2
	PIECE_PL2
	PIECE_PP2
)

// Promote - 成ります
func Promote(piece l09.Piece) l09.Piece {
	switch piece {
	case PIECE_EMPTY, PIECE_K1, PIECE_G1, PIECE_PR1, PIECE_PB1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1,
		PIECE_K2, PIECE_G2, PIECE_PR2, PIECE_PB2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2: // 成らずにそのまま返す駒
		return piece
	case PIECE_R1:
		return PIECE_PR1
	case PIECE_B1:
		return PIECE_PB1
	case PIECE_S1:
		return PIECE_PS1
	case PIECE_N1:
		return PIECE_PN1
	case PIECE_L1:
		return PIECE_PL1
	case PIECE_P1:
		return PIECE_PP1
	case PIECE_R2:
		return PIECE_PR2
	case PIECE_B2:
		return PIECE_PB2
	case PIECE_S2:
		return PIECE_PS2
	case PIECE_N2:
		return PIECE_PN2
	case PIECE_L2:
		return PIECE_PL2
	case PIECE_P2:
		return PIECE_PP2
	default:
		panic(fmt.Errorf("error: unknown piece=[%d]", piece))
	}
}

// Demote - 成っている駒を、成っていない駒に戻します
func Demote(piece l09.Piece) l09.Piece {
	switch piece {
	case PIECE_EMPTY, PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1,
		PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2: // 裏返さずにそのまま返す駒
		return piece
	case PIECE_PR1:
		return PIECE_R1
	case PIECE_PB1:
		return PIECE_B1
	case PIECE_PS1:
		return PIECE_S1
	case PIECE_PN1:
		return PIECE_N1
	case PIECE_PL1:
		return PIECE_L1
	case PIECE_PP1:
		return PIECE_P1
	case PIECE_PR2:
		return PIECE_R2
	case PIECE_PB2:
		return PIECE_B2
	case PIECE_PS2:
		return PIECE_S2
	case PIECE_PN2:
		return PIECE_N2
	case PIECE_PL2:
		return PIECE_L2
	case PIECE_PP2:
		return PIECE_P2
	default:
		panic(fmt.Errorf("error: unknown piece=[%d]", piece))
	}
}
