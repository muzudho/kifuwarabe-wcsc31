package lesson03

import (
	"fmt"
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
		panic(App.LogNotEcho.Fatal("unknown piece=[%s]", piece.ToCodeOfPc()))
	}
}

// WhatHand - 盤外のマス番号を、駒台の持ち駒の種類と識別し、先後なしの駒種類を返します
func WhatHand(sq Square) PieceType {
	var handSq = FromSqToHandSq(sq)
	switch handSq {
	case HANDSQ_R1, HANDSQ_R2:
		return PIECE_TYPE_R
	case HANDSQ_B1, HANDSQ_B2:
		return PIECE_TYPE_B
	case HANDSQ_G1, HANDSQ_G2:
		return PIECE_TYPE_G
	case HANDSQ_S1, HANDSQ_S2:
		return PIECE_TYPE_S
	case HANDSQ_N1, HANDSQ_N2:
		return PIECE_TYPE_N
	case HANDSQ_L1, HANDSQ_L2:
		return PIECE_TYPE_L
	case HANDSQ_P1, HANDSQ_P2:
		return PIECE_TYPE_P
	default:
		panic(fmt.Errorf("unknown hand sq=[%d]", handSq))
	}
}

// PieceFromPhPt - 駒作成。空マスは作れません
func PieceFromPhPt(phase Phase, pieceType PieceType) Piece {
	switch phase {
	case FIRST:
		switch pieceType {
		case PIECE_TYPE_K:
			return PIECE_K1
		case PIECE_TYPE_R:
			return PIECE_R1
		case PIECE_TYPE_B:
			return PIECE_B1
		case PIECE_TYPE_G:
			return PIECE_G1
		case PIECE_TYPE_S:
			return PIECE_S1
		case PIECE_TYPE_N:
			return PIECE_N1
		case PIECE_TYPE_L:
			return PIECE_L1
		case PIECE_TYPE_P:
			return PIECE_P1
		case PIECE_TYPE_PR:
			return PIECE_PR1
		case PIECE_TYPE_PB:
			return PIECE_PB1
		case PIECE_TYPE_PS:
			return PIECE_PS1
		case PIECE_TYPE_PN:
			return PIECE_PN1
		case PIECE_TYPE_PL:
			return PIECE_PL1
		case PIECE_TYPE_PP:
			return PIECE_PP1
		default:
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	case SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return PIECE_K2
		case PIECE_TYPE_R:
			return PIECE_R2
		case PIECE_TYPE_B:
			return PIECE_B2
		case PIECE_TYPE_G:
			return PIECE_G2
		case PIECE_TYPE_S:
			return PIECE_S2
		case PIECE_TYPE_N:
			return PIECE_N2
		case PIECE_TYPE_L:
			return PIECE_L2
		case PIECE_TYPE_P:
			return PIECE_P2
		case PIECE_TYPE_PR:
			return PIECE_PR2
		case PIECE_TYPE_PB:
			return PIECE_PB2
		case PIECE_TYPE_PS:
			return PIECE_PS2
		case PIECE_TYPE_PN:
			return PIECE_PN2
		case PIECE_TYPE_PL:
			return PIECE_PL2
		case PIECE_TYPE_PP:
			return PIECE_PP2
		default:
			panic(App.LogNotEcho.Fatal("unknown piece type=%d", pieceType))
		}
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
	}
}

// Who - 駒が先手か後手か空升かを返します
func Who(piece Piece) Phase {
	switch piece {
	case PIECE_EMPTY: // 空きマス
		return ZEROTH
	case PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1, PIECE_PR1, PIECE_PB1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1:
		return FIRST
	case PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2, PIECE_PR2, PIECE_PB2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		return SECOND
	default:
		panic(fmt.Errorf("error: 知らん駒（＾～＾） piece=[%s]", piece.ToCodeOfPc()))
	}
}
