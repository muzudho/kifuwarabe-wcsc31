package take7

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}

func GenControl(pPos *Position, from Square) []Square {
	sq_list := []Square{}

	piece := pPos.Board[from]

	// ２つ先のマスから斜めに長い利き
	switch piece {
	case PIECE_B1, PIECE_PB1, PIECE_B2, PIECE_PB2:
		if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
			for to := from + 18; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to); to += 9 { // ２つ左上から
				sq_list = append(sq_list, to)
			}
		}
		if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
			for to := from - 22; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to); to -= 11 { // ２つ右上から
				sq_list = append(sq_list, to)
			}
		}
		if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
			for to := from + 22; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to); to += 11 { // ２つ左下から
				sq_list = append(sq_list, to)
			}
		}
		if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
			for to := from - 18; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to); to -= 9 { // ２つ右下から
				sq_list = append(sq_list, to)
			}
		}
	default:
		// Ignored
	}

	// ２つ先のマスから先手香車の長い利き
	switch piece {
	case PIECE_L1, PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
		if Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
			for to := from - 2; Rank(to) != 0 && pPos.Hetero(to); to -= 1 { // 上
				sq_list = append(sq_list, to)
			}
		}
	default:
		// Ignored
	}

	// ２つ先のマスから後手香車の長い利き
	switch piece {
	case PIECE_R1, PIECE_PR1, PIECE_L2, PIECE_R2, PIECE_PR2:
		if Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
			for to := from + 2; Rank(to) != 0 && pPos.Hetero(to); to += 1 { // 下
				sq_list = append(sq_list, to)
			}
		}
	default:
		// Ignored
	}

	// ２つ横のマスから飛の長い利き
	switch piece {
	case PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
		if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
			for to := from + 20; File(to) != 0 && pPos.Hetero(to); to += 10 { // 左
				sq_list = append(sq_list, to)
			}
		}
		if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
			for to := from - 20; File(to) != 0 && pPos.Hetero(to); to -= 10 { // 右
				sq_list = append(sq_list, to)
			}
		}
	default:
		// Ignored
	}

	// 先手桂の動き
	if piece == PIECE_N1 {
		if to := from + 8; File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 左上桂馬飛び
			sq_list = append(sq_list, to)
		}
		if to := from - 12; File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 右上桂馬飛び
			sq_list = append(sq_list, to)
		}
	}

	// 後手桂の動き
	if piece == PIECE_N2 {
		if to := from + 12; File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 左下
			sq_list = append(sq_list, to)
		}
		if to := from - 8; File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 右下
			sq_list = append(sq_list, to)
		}
	}

	// 先手歩の動き
	switch piece {
	case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_L1, PIECE_P1, PIECE_PS1,
		PIECE_PN1, PIECE_PL1, PIECE_PP1, PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2,
		PIECE_PN2, PIECE_PL2, PIECE_PP2:
		if to := from - 1; Rank(to) != 0 && pPos.Hetero(to) { // 上
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 後手歩の動き
	switch piece {
	case PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_L2, PIECE_P2, PIECE_PS2,
		PIECE_PN2, PIECE_PL2, PIECE_PP2, PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1,
		PIECE_PN1, PIECE_PL1, PIECE_PP1:
		if to := from + 1; Rank(to) != 0 && pPos.Hetero(to) { // 下
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 先手斜め前の動き
	switch piece {
	case PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_PS1, PIECE_PN1, PIECE_PL1,
		PIECE_PP1, PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_S2:
		if to := from + 9; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to) { // 左上
			sq_list = append(sq_list, to)
		}
		if to := from - 11; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to) { // 右上
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 後手斜め前の動き
	switch piece {
	case PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_PS2, PIECE_PN2, PIECE_PL2,
		PIECE_PP2, PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_S1:
		if to := from + 11; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to) { // 左下
			sq_list = append(sq_list, to)
		}
		if to := from - 9; File(to) != 0 && Rank(to) != 0 && pPos.Hetero(to) { // 右下
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	// 横１マスの動き
	switch piece {
	case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1,
		PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		if to := from + 10; File(to) != 0 && pPos.Hetero(to) { // 左
			sq_list = append(sq_list, to)
		}
		if to := from - 10; File(to) != 0 && pPos.Hetero(to) { // 右
			sq_list = append(sq_list, to)
		}
	default:
		// Ignored
	}

	return sq_list
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []Move {

	move_list := []Move{}

	// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
	for rank := 1; rank < 10; rank += 1 {
		for file := 1; file < 10; file += 1 {
			from := Square(file*10 + rank)
			if pPos.Homo(from) {
				sq_list := GenControl(pPos, from)

				for _, to := range sq_list {
					move_list = append(move_list, NewMove2(from, to))
				}
			}
		}
	}

	return move_list
}
