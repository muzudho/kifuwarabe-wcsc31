package take16

import (
	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
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
func What(piece p.Piece) PieceType {
	switch piece {
	case p.PIECE_EMPTY: // 空きマス
		return PIECE_TYPE_EMPTY
	case p.PIECE_K1, p.PIECE_K2:
		return PIECE_TYPE_K
	case p.PIECE_R1, p.PIECE_R2:
		return PIECE_TYPE_R
	case p.PIECE_B1, p.PIECE_B2:
		return PIECE_TYPE_B
	case p.PIECE_G1, p.PIECE_G2:
		return PIECE_TYPE_G
	case p.PIECE_S1, p.PIECE_S2:
		return PIECE_TYPE_S
	case p.PIECE_N1, p.PIECE_N2:
		return PIECE_TYPE_N
	case p.PIECE_L1, p.PIECE_L2:
		return PIECE_TYPE_L
	case p.PIECE_P1, p.PIECE_P2:
		return PIECE_TYPE_P
	case p.PIECE_PR1, p.PIECE_PR2:
		return PIECE_TYPE_PR
	case p.PIECE_PB1, p.PIECE_PB2:
		return PIECE_TYPE_PB
	case p.PIECE_PS1, p.PIECE_PS2:
		return PIECE_TYPE_PS
	case p.PIECE_PN1, p.PIECE_PN2:
		return PIECE_TYPE_PN
	case p.PIECE_PL1, p.PIECE_PL2:
		return PIECE_TYPE_PL
	case p.PIECE_PP1, p.PIECE_PP2:
		return PIECE_TYPE_PP
	default:
		panic(G.Log.Fatal("Error: Unknown piece=[%d]", piece))
	}
}

// WhatHand - 持ち駒のマス番号から、先後なしの駒種類を返します
func WhatHand(hand p.Square) PieceType {
	switch hand {
	case p.SQ_R1, p.SQ_R2:
		return PIECE_TYPE_R
	case p.SQ_B1, p.SQ_B2:
		return PIECE_TYPE_B
	case p.SQ_G1, p.SQ_G2:
		return PIECE_TYPE_G
	case p.SQ_S1, p.SQ_S2:
		return PIECE_TYPE_S
	case p.SQ_N1, p.SQ_N2:
		return PIECE_TYPE_N
	case p.SQ_L1, p.SQ_L2:
		return PIECE_TYPE_L
	case p.SQ_P1, p.SQ_P2:
		return PIECE_TYPE_P
	default:
		panic(G.Log.Fatal("Error: Unknown hand=[%d]", hand))
	}
}
