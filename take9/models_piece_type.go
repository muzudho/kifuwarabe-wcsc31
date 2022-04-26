package take9

import (
	"fmt"

	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
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
func What(piece Piece) PieceType {
	switch piece {
	case PIECE_EMPTY: // 空きマス
		return PIECE_TYPE_EMPTY
	case PIECE_K1, PIECE_K2:
		return PIECE_TYPE_K
	case PIECE_R1, PIECE_R2:
		return PIECE_TYPE_R
	case PIECE_B1, PIECE_B2:
		return PIECE_TYPE_B
	case PIECE_G1, PIECE_G2:
		return PIECE_TYPE_G
	case PIECE_S1, PIECE_S2:
		return PIECE_TYPE_S
	case PIECE_N1, PIECE_N2:
		return PIECE_TYPE_N
	case PIECE_L1, PIECE_L2:
		return PIECE_TYPE_L
	case PIECE_P1, PIECE_P2:
		return PIECE_TYPE_P
	case PIECE_PR1, PIECE_PR2:
		return PIECE_TYPE_PR
	case PIECE_PB1, PIECE_PB2:
		return PIECE_TYPE_PB
	case PIECE_PS1, PIECE_PS2:
		return PIECE_TYPE_PS
	case PIECE_PN1, PIECE_PN2:
		return PIECE_TYPE_PN
	case PIECE_PL1, PIECE_PL2:
		return PIECE_TYPE_PL
	case PIECE_PP1, PIECE_PP2:
		return PIECE_TYPE_PP
	default:
		panic(fmt.Errorf("unknown piece=[%d]", piece))
	}
}

// WhatHand - 持ち駒のマス番号から、先後なしの駒種類を返します
func WhatHand(hand l04.Square) PieceType {
	switch hand {
	case HAND_R1, HAND_R2:
		return PIECE_TYPE_R
	case HAND_B1, HAND_B2:
		return PIECE_TYPE_B
	case HAND_G1, HAND_G2:
		return PIECE_TYPE_G
	case HAND_S1, HAND_S2:
		return PIECE_TYPE_S
	case HAND_N1, HAND_N2:
		return PIECE_TYPE_N
	case HAND_L1, HAND_L2:
		return PIECE_TYPE_L
	case HAND_P1, HAND_P2:
		return PIECE_TYPE_P
	default:
		panic(fmt.Errorf("unknown hand=[%d]", hand))
	}
}
