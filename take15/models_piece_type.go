package take15

import (
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
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
func What(piece l09.Piece) PieceType {
	switch piece {
	case l11.PIECE_EMPTY: // 空きマス
		return PIECE_TYPE_EMPTY
	case l11.PIECE_K1, l11.PIECE_K2:
		return PIECE_TYPE_K
	case l11.PIECE_R1, l11.PIECE_R2:
		return PIECE_TYPE_R
	case l11.PIECE_B1, l11.PIECE_B2:
		return PIECE_TYPE_B
	case l11.PIECE_G1, l11.PIECE_G2:
		return PIECE_TYPE_G
	case l11.PIECE_S1, l11.PIECE_S2:
		return PIECE_TYPE_S
	case l11.PIECE_N1, l11.PIECE_N2:
		return PIECE_TYPE_N
	case l11.PIECE_L1, l11.PIECE_L2:
		return PIECE_TYPE_L
	case l11.PIECE_P1, l11.PIECE_P2:
		return PIECE_TYPE_P
	case l11.PIECE_PR1, l11.PIECE_PR2:
		return PIECE_TYPE_PR
	case l11.PIECE_PB1, l11.PIECE_PB2:
		return PIECE_TYPE_PB
	case l11.PIECE_PS1, l11.PIECE_PS2:
		return PIECE_TYPE_PS
	case l11.PIECE_PN1, l11.PIECE_PN2:
		return PIECE_TYPE_PN
	case l11.PIECE_PL1, l11.PIECE_PL2:
		return PIECE_TYPE_PL
	case l11.PIECE_PP1, l11.PIECE_PP2:
		return PIECE_TYPE_PP
	default:
		panic(App.LogNotEcho.Fatal("unknown piece=[%d]", piece))
	}
}

// WhatHand - 持ち駒のマス番号から、先後なしの駒種類を返します
func WhatHand(hand Square) PieceType {
	switch hand {
	case SQ_R1, SQ_R2:
		return PIECE_TYPE_R
	case SQ_B1, SQ_B2:
		return PIECE_TYPE_B
	case SQ_G1, SQ_G2:
		return PIECE_TYPE_G
	case SQ_S1, SQ_S2:
		return PIECE_TYPE_S
	case SQ_N1, SQ_N2:
		return PIECE_TYPE_N
	case SQ_L1, SQ_L2:
		return PIECE_TYPE_L
	case SQ_P1, SQ_P2:
		return PIECE_TYPE_P
	default:
		panic(App.LogNotEcho.Fatal("unknown hand=[%d]", hand))
	}
}
