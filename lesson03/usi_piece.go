package lesson03

// FromCodeToPiece - 文字列からPieceを作成
func FromCodeToPiece(piece string) Piece {
	switch piece {
	case "":
		return PIECE_EMPTY
	case "K":
		return PIECE_K1
	case "R":
		return PIECE_R1
	case "B":
		return PIECE_B1
	case "G":
		return PIECE_G1
	case "S":
		return PIECE_S1
	case "N":
		return PIECE_N1
	case "L":
		return PIECE_L1
	case "P":
		return PIECE_P1
	case "+R":
		return PIECE_PR1
	case "+B":
		return PIECE_PB1
	case "+S":
		return PIECE_PS1
	case "+N":
		return PIECE_PN1
	case "+L":
		return PIECE_PL1
	case "+P":
		return PIECE_PP1
	case "k":
		return PIECE_K2
	case "r":
		return PIECE_R2
	case "b":
		return PIECE_B2
	case "g":
		return PIECE_G2
	case "s":
		return PIECE_S2
	case "n":
		return PIECE_N2
	case "l":
		return PIECE_L2
	case "p":
		return PIECE_P2
	case "+r":
		return PIECE_PR2
	case "+b":
		return PIECE_PB2
	case "+s":
		return PIECE_PS2
	case "+n":
		return PIECE_PN2
	case "+l":
		return PIECE_PL2
	case "+p":
		return PIECE_PP2
	default:
		panic(App.LogNotEcho.Fatal("unknown piece=[%s]", piece))
	}
}
