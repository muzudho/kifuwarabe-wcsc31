package take7

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// 先後のない駒種類
type PieceType byte

const (
	PIECE_TYPE_EMPTY = PieceType(0) // 空マス
	PIECE_TYPE_K     = PieceType(1)
	PIECE_TYPE_R     = PieceType(2)
	PIECE_TYPE_B     = PieceType(3)
	PIECE_TYPE_G     = PieceType(4)
	PIECE_TYPE_S     = PieceType(5)
	PIECE_TYPE_N     = PieceType(6)
	PIECE_TYPE_L     = PieceType(7)
	PIECE_TYPE_P     = PieceType(8)
	PIECE_TYPE_PR    = PieceType(9)
	PIECE_TYPE_PB    = PieceType(10)
	PIECE_TYPE_PS    = PieceType(11)
	PIECE_TYPE_PN    = PieceType(12)
	PIECE_TYPE_PL    = PieceType(13)
	PIECE_TYPE_PP    = PieceType(14)
)

// What - 先後のない駒種類を返します。
func What(piece string) PieceType {
	switch piece {
	case l03.PIECE_EMPTY: // 空きマス
		return PIECE_TYPE_EMPTY
	case l03.PIECE_K1, l03.PIECE_K2:
		return PIECE_TYPE_K
	case l03.PIECE_R1, l03.PIECE_R2:
		return PIECE_TYPE_R
	case l03.PIECE_B1, l03.PIECE_B2:
		return PIECE_TYPE_B
	case l03.PIECE_G1, l03.PIECE_G2:
		return PIECE_TYPE_G
	case l03.PIECE_S1, l03.PIECE_S2:
		return PIECE_TYPE_S
	case l03.PIECE_N1, l03.PIECE_N2:
		return PIECE_TYPE_N
	case l03.PIECE_L1, l03.PIECE_L2:
		return PIECE_TYPE_L
	case l03.PIECE_P1, l03.PIECE_P2:
		return PIECE_TYPE_P
	case l03.PIECE_PR1, l03.PIECE_PR2:
		return PIECE_TYPE_PR
	case l03.PIECE_PB1, l03.PIECE_PB2:
		return PIECE_TYPE_PB
	case l03.PIECE_PS1, l03.PIECE_PS2:
		return PIECE_TYPE_PS
	case l03.PIECE_PN1, l03.PIECE_PN2:
		return PIECE_TYPE_PN
	case l03.PIECE_PL1, l03.PIECE_PL2:
		return PIECE_TYPE_PL
	case l03.PIECE_PP1, l03.PIECE_PP2:
		return PIECE_TYPE_PP
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}
