package take9 // same take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// 先後のない駒種類
type PieceType byte

const (
	PIECE_TYPE_EMPTY PieceType = 0 + iota // 空マス
	PIECE_TYPE_K
	PIECE_TYPE_R
	PIECE_TYPE_B
	PIECE_TYPE_G
	PIECE_TYPE_S
	PIECE_TYPE_N
	PIECE_TYPE_L
	PIECE_TYPE_P
	PIECE_TYPE_PR
	PIECE_TYPE_PB
	PIECE_TYPE_PS
	PIECE_TYPE_PN
	PIECE_TYPE_PL
	PIECE_TYPE_PP
)

// What - 先後のない駒種類を返します。
func What(piece l03.Piece) PieceType {
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
		panic(fmt.Errorf("unknown piece=[%s]", piece.ToCodeOfPc()))
	}
}

// WhatHand - 盤外のマス番号を、駒台の持ち駒の種類と識別し、先後なしの駒種類を返します
func WhatHand(sq l03.Square) PieceType {
	var handSq = l03.FromSqToHandSq(sq)
	switch handSq {
	case l03.HANDSQ_R1, l03.HANDSQ_R2:
		return PIECE_TYPE_R
	case l03.HANDSQ_B1, l03.HANDSQ_B2:
		return PIECE_TYPE_B
	case l03.HANDSQ_G1, l03.HANDSQ_G2:
		return PIECE_TYPE_G
	case l03.HANDSQ_S1, l03.HANDSQ_S2:
		return PIECE_TYPE_S
	case l03.HANDSQ_N1, l03.HANDSQ_N2:
		return PIECE_TYPE_N
	case l03.HANDSQ_L1, l03.HANDSQ_L2:
		return PIECE_TYPE_L
	case l03.HANDSQ_P1, l03.HANDSQ_P2:
		return PIECE_TYPE_P
	default:
		panic(fmt.Errorf("unknown hand sq=[%d]", handSq))
	}
}
