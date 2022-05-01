package lesson03

// WhatHand - 盤外のマス番号を、駒台の持ち駒の種類と識別し、先後なしの駒種類を返します
func WhatHand(sq Square) PieceType {
	if sq < 100 || 114 <= sq {
		App.Log.Debug("unknown hand sq=[%d]", sq)
		panic(App.LogNotEcho.Fatal("unknown hand sq=[%d]", sq))
	}
	var handSq = FromSqToHandSq(sq)
	return WhatHandSq(handSq)
}

func WhatHandSq(handSq HandSq) PieceType {
	switch handSq {
	case HANDSQ_K1, HANDSQ_K2:
		return PIECE_TYPE_K
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
		App.Log.Debug("unknown hand sq=[%d]", handSq)
		panic(App.LogNotEcho.Fatal("unknown hand sq=[%d]", handSq))
	}
}
