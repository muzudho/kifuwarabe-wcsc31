package take7

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

func GenMoveEnd(pPos *Position, from l03.Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	piece := pPos.Board[from]

	// ２つ先のマスから斜めに長い利き
	switch piece {
	case l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_B2, l03.PIECE_PB2:
		if l03.File(from) < 8 && l03.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
			for to := from + 18; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to); to += 9 { // ２つ左上から
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
		if l03.File(from) > 2 && l03.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
			for to := from - 22; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to); to -= 11 { // ２つ右上から
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
		if l03.File(from) < 8 && l03.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
			for to := from + 22; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to); to += 11 { // ２つ左下から
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
		if l03.File(from) > 2 && l03.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
			for to := from - 18; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to); to -= 9 { // ２つ右下から
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
	default:
		// Ignored
	}

	// ２つ先のマスから先手香車の長い利き
	switch piece {
	case l03.PIECE_L1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_R2, l03.PIECE_PR2:
		if l03.Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
			for to := from - 2; l03.Rank(to) != 0 && pPos.Hetero(to); to -= 1 { // 上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
	default:
		// Ignored
	}

	// ２つ先のマスから後手香車の長い利き
	switch piece {
	case l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_L2, l03.PIECE_R2, l03.PIECE_PR2:
		if l03.Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
			for to := from + 2; l03.Rank(to) != 0 && pPos.Hetero(to); to += 1 { // 下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
	default:
		// Ignored
	}

	// ２つ横のマスから飛の長い利き
	switch piece {
	case l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_R2, l03.PIECE_PR2:
		if l03.File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
			for to := from + 20; l03.File(to) != 0 && pPos.Hetero(to); to += 10 { // 左
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
		if l03.File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
			for to := from - 20; l03.File(to) != 0 && pPos.Hetero(to); to -= 10 { // 右
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}
	default:
		// Ignored
	}

	// 先手桂の動き
	if piece == l03.PIECE_N1 {
		if to := from + 8; l03.File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 左上桂馬飛び
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if to := from - 12; l03.File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 右上桂馬飛び
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	}

	// 後手桂の動き
	if piece == l03.PIECE_N2 {
		if to := from + 12; l03.File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 左下
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if to := from - 8; l03.File(to) != 0 && to%10 != 0 && to%10 != 9 && pPos.Hetero(to) { // 右下
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	}

	// 先手歩の動き
	switch piece {
	case l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_S1, l03.PIECE_L1, l03.PIECE_P1, l03.PIECE_PS1,
		l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1, l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_PS2,
		l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2:
		if to := from - 1; l03.Rank(to) != 0 && pPos.Hetero(to) { // 上
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	default:
		// Ignored
	}

	// 後手歩の動き
	switch piece {
	case l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_S2, l03.PIECE_L2, l03.PIECE_P2, l03.PIECE_PS2,
		l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2, l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_PS1,
		l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1:
		if to := from + 1; l03.Rank(to) != 0 && pPos.Hetero(to) { // 下
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	default:
		// Ignored
	}

	// 先手斜め前の動き
	switch piece {
	case l03.PIECE_K1, l03.PIECE_PR1, l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_S1, l03.PIECE_PS1, l03.PIECE_PN1, l03.PIECE_PL1,
		l03.PIECE_PP1, l03.PIECE_K2, l03.PIECE_PR2, l03.PIECE_B2, l03.PIECE_PB2, l03.PIECE_S2:
		if to := from + 9; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to) { // 左上
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if to := from - 11; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to) { // 右上
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	default:
		// Ignored
	}

	// 後手斜め前の動き
	switch piece {
	case l03.PIECE_K2, l03.PIECE_PR2, l03.PIECE_B2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_S2, l03.PIECE_PS2, l03.PIECE_PN2, l03.PIECE_PL2,
		l03.PIECE_PP2, l03.PIECE_K1, l03.PIECE_PR1, l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_S1:
		if to := from + 11; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to) { // 左下
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if to := from - 9; l03.File(to) != 0 && l03.Rank(to) != 0 && pPos.Hetero(to) { // 右下
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	default:
		// Ignored
	}

	// 横１マスの動き
	switch piece {
	case l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_PS1, l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1,
		l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_PS2, l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2:
		if to := from + 10; l03.File(to) != 0 && pPos.Hetero(to) { // 左
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if to := from - 10; l03.File(to) != 0 && pPos.Hetero(to) { // 右
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	default:
		// Ignored
	}

	return moveEndList
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []l03.Move {

	move_list := []l03.Move{}

	// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
	for rank := 1; rank < 10; rank += 1 {
		for file := 1; file < 10; file += 1 {
			from := l03.Square(file*10 + rank)
			if pPos.Homo(from) {
				moveEndList := GenMoveEnd(pPos, from)

				for _, moveEnd := range moveEndList {
					to, pro := moveEnd.Destructure()
					move_list = append(move_list, l03.NewMove(from, to, pro))
				}
			}
		}
	}

	return move_list
}
