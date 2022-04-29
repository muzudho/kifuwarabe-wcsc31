package take11 // same take12

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
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
		panic(App.LogNotEcho.Fatal("unknown piece=[%d]", piece))
	}
}

// WhatHand - 持ち駒のマス番号から、先後なしの駒種類を返します
func WhatHand(hand l03.Square) PieceType {
	switch hand {
	case l03.SQ_R1, l03.SQ_R2:
		return PIECE_TYPE_R
	case l03.SQ_B1, l03.SQ_B2:
		return PIECE_TYPE_B
	case l03.SQ_G1, l03.SQ_G2:
		return PIECE_TYPE_G
	case l03.SQ_S1, l03.SQ_S2:
		return PIECE_TYPE_S
	case l03.SQ_N1, l03.SQ_N2:
		return PIECE_TYPE_N
	case l03.SQ_L1, l03.SQ_L2:
		return PIECE_TYPE_L
	case l03.SQ_P1, l03.SQ_P2:
		return PIECE_TYPE_P
	default:
		panic(App.LogNotEcho.Fatal("unknown hand=[%d]", hand))
	}
}

// PieceFromPhPt - 駒作成。空マスは作れません
func PieceFromPhPt(phase l06.Phase, pieceType PieceType) l03.Piece {
	switch phase {
	case l06.FIRST:
		switch pieceType {
		case PIECE_TYPE_K:
			return l03.PIECE_K1
		case PIECE_TYPE_R:
			return l03.PIECE_R1
		case PIECE_TYPE_B:
			return l03.PIECE_B1
		case PIECE_TYPE_G:
			return l03.PIECE_G1
		case PIECE_TYPE_S:
			return l03.PIECE_S1
		case PIECE_TYPE_N:
			return l03.PIECE_N1
		case PIECE_TYPE_L:
			return l03.PIECE_L1
		case PIECE_TYPE_P:
			return l03.PIECE_P1
		case PIECE_TYPE_PR:
			return l03.PIECE_PR1
		case PIECE_TYPE_PB:
			return l03.PIECE_PB1
		case PIECE_TYPE_PS:
			return l03.PIECE_PS1
		case PIECE_TYPE_PN:
			return l03.PIECE_PN1
		case PIECE_TYPE_PL:
			return l03.PIECE_PL1
		case PIECE_TYPE_PP:
			return l03.PIECE_PP1
		default:
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	case l06.SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return l03.PIECE_K2
		case PIECE_TYPE_R:
			return l03.PIECE_R2
		case PIECE_TYPE_B:
			return l03.PIECE_B2
		case PIECE_TYPE_G:
			return l03.PIECE_G2
		case PIECE_TYPE_S:
			return l03.PIECE_S2
		case PIECE_TYPE_N:
			return l03.PIECE_N2
		case PIECE_TYPE_L:
			return l03.PIECE_L2
		case PIECE_TYPE_P:
			return l03.PIECE_P2
		case PIECE_TYPE_PR:
			return l03.PIECE_PR2
		case PIECE_TYPE_PB:
			return l03.PIECE_PB2
		case PIECE_TYPE_PS:
			return l03.PIECE_PS2
		case PIECE_TYPE_PN:
			return l03.PIECE_PN2
		case PIECE_TYPE_PL:
			return l03.PIECE_PL2
		case PIECE_TYPE_PP:
			return l03.PIECE_PP2
		default:
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
	}
}
