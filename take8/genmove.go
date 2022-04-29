package take8

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// GenMoveEnd - 利いているマスの一覧を返します。動けるマスではありません。
func GenMoveEnd(pPos *Position, from l03.Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	if from == l03.SQ_EMPTY {
		panic(fmt.Errorf("GenMoveEnd has empty square"))
	} else if OnBoard(from) {
		// 盤上の駒の利き
		piece := pPos.Board[from]

		// ２つ先のマスから斜めに長い利き
		switch piece {
		case l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc():
			if l03.File(from) < 8 && l03.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
				for to := from + 18; l03.File(to) != 0 && l03.Rank(to) != 0; to += 9 { // ２つ左上から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l03.File(from) > 2 && l03.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
				for to := from - 22; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 11 { // ２つ右上から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l03.File(from) < 8 && l03.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
				for to := from + 22; l03.File(to) != 0 && l03.Rank(to) != 0; to += 11 { // ２つ左下から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l03.File(from) > 2 && l03.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
				for to := from - 18; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 9 { // ２つ右下から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
		default:
			// Ignored
		}

		// ２つ先のマスから先手香車の長い利き
		switch piece {
		case l03.PIECE_L1.ToCodeOfPc(), l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc():
			if l03.Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
				for to := from - 2; l03.Rank(to) != 0; to -= 1 { // 上
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
		default:
			// Ignored
		}

		// ２つ先のマスから後手香車の長い利き
		switch piece {
		case l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_L2.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc():
			if l03.Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
				for to := from + 2; l03.Rank(to) != 0; to += 1 { // 下
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
		default:
			// Ignored
		}

		// ２つ横のマスから飛の長い利き
		switch piece {
		case l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc():
			if l03.File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
				for to := from + 20; l03.File(to) != 0; to += 10 { // 左
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l03.File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
				for to := from - 20; l03.File(to) != 0; to -= 10 { // 右
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
		default:
			// Ignored
		}

		// 先手桂の利き
		if piece == l03.PIECE_N1.ToCodeOfPc() {
			if to := from + 8; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 左上桂馬飛び
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 12; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 右上桂馬飛び
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 後手桂の利き
		if piece == l03.PIECE_N2.ToCodeOfPc() {
			if to := from + 12; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 左下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 8; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 右下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 先手歩の利き
		switch piece {
		case l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_S1.ToCodeOfPc(), l03.PIECE_L1.ToCodeOfPc(), l03.PIECE_P1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc(),
			l03.PIECE_PN1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc(), l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc(),
			l03.PIECE_PN2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc():
			if to := from - 1; l03.Rank(to) != 0 { // 上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手歩の利き
		switch piece {
		case l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_S2.ToCodeOfPc(), l03.PIECE_L2.ToCodeOfPc(), l03.PIECE_P2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc(),
			l03.PIECE_PN2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc(), l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc(),
			l03.PIECE_PN1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc():
			if to := from + 1; l03.Rank(to) != 0 { // 下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 先手斜め前の利き
		switch piece {
		case l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_S1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc(), l03.PIECE_PN1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc(),
			l03.PIECE_PP1.ToCodeOfPc(), l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_S2.ToCodeOfPc():
			if to := from + 9; l03.File(to) != 0 && l03.Rank(to) != 0 { // 左上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 11; l03.File(to) != 0 && l03.Rank(to) != 0 { // 右上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手斜め前の利き
		switch piece {
		case l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_B2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_S2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc(), l03.PIECE_PN2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc(),
			l03.PIECE_PP2.ToCodeOfPc(), l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_B1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_S1.ToCodeOfPc():
			if to := from + 11; l03.File(to) != 0 && l03.Rank(to) != 0 { // 左下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 9; l03.File(to) != 0 && l03.Rank(to) != 0 { // 右下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 横１マスの利き
		switch piece {
		case l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_R1.ToCodeOfPc(), l03.PIECE_PR1.ToCodeOfPc(), l03.PIECE_PB1.ToCodeOfPc(), l03.PIECE_G1.ToCodeOfPc(), l03.PIECE_PS1.ToCodeOfPc(), l03.PIECE_PN1.ToCodeOfPc(), l03.PIECE_PL1.ToCodeOfPc(), l03.PIECE_PP1.ToCodeOfPc(),
			l03.PIECE_K2.ToCodeOfPc(), l03.PIECE_R2.ToCodeOfPc(), l03.PIECE_PR2.ToCodeOfPc(), l03.PIECE_PB2.ToCodeOfPc(), l03.PIECE_G2.ToCodeOfPc(), l03.PIECE_PS2.ToCodeOfPc(), l03.PIECE_PN2.ToCodeOfPc(), l03.PIECE_PL2.ToCodeOfPc(), l03.PIECE_PP2.ToCodeOfPc():
			if to := from + 10; l03.File(to) != 0 { // 左
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 10; l03.File(to) != 0 { // 右
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}
	} else {
		// どこに打てるか
		var start_rank l03.Square
		var end_rank l03.Square

		switch l03.FromSqToHandSq(from) {
		case l03.HANDSQ_R1, l03.HANDSQ_B1, l03.HANDSQ_G1, l03.HANDSQ_S1, l03.HANDSQ_R2, l03.HANDSQ_B2, l03.HANDSQ_G2, l03.HANDSQ_S2: // 81マスに打てる
			start_rank = 1
			end_rank = 10
		case l03.HANDSQ_N1: // 3～9段目に打てる
			start_rank = 3
			end_rank = 10
		case l03.HANDSQ_L1, l03.HANDSQ_P1: // 2～9段目に打てる
			start_rank = 2
			end_rank = 10
		case l03.HANDSQ_N2: // 1～7段目に打てる
			start_rank = 1
			end_rank = 8
		case l03.HANDSQ_L2, l03.HANDSQ_P2: // 1～8段目に打てる
			start_rank = 1
			end_rank = 9
		default:
			panic(fmt.Errorf("unknown hand from=%d", from))
		}

		switch l03.FromSqToHandSq(from) {
		case l03.HANDSQ_P1:
			// TODO 打ち歩詰め禁止
			for rank := l03.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l03.Square(9); file > 0; file-- {
					if !NifuFirst(pPos, file) { // ２歩禁止
						var to = SquareFrom(file, rank)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		case l03.HANDSQ_P2:
			// TODO 打ち歩詰め禁止
			for rank := l03.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l03.Square(9); file > 0; file-- {
					if !NifuSecond(pPos, file) { // ２歩禁止
						var to = SquareFrom(file, rank)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		default:
			for rank := l03.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l03.Square(9); file > 0; file-- {
					var to = SquareFrom(file, rank)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
			}
		}
	}

	return moveEndList
}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file l03.Square) bool {
	for rank := l03.Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l03.PIECE_P1.ToCodeOfPc() {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file l03.Square) bool {
	for rank := l03.Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l03.PIECE_P2.ToCodeOfPc() {
			return true
		}
	}

	return false
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []Move {

	move_list := []Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friendKingSq := pPos.PieceLocations[PCLOC_K1:PCLOC_K2][pPos.Phase-1]
	opponent := FlipPhase(pPos.Phase)

	if pPos.ControlBoards[opponent-1][friendKingSq] > 0 {
		// 王手されています
		// TODO アタッカーがどの駒か調べたいが……。一手前に動かした駒か、空き王手のどちらかしかなくないか（＾～＾）？
		// 王手されているところが開始局面だと、一手前を調べることができないので、やっぱ調べるしか（＾～＾）
		// 空き王手を利用して、2箇所から 長い利きが飛んでくることはある（＾～＾）

		// 駒を動かしてみて、王手が解除されるか調べるか（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := l03.Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							// 敵の長い駒の利きは、玉が逃げても伸びてくる方向があるので、
							// いったん玉を動かしてから 再チェックするぜ（＾～＾）
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMove(from, to, pro)
								pPos.DoMove(move)

								if pPos.ControlBoards[opponent-1][to] == 0 {
									// よっしゃ利きから逃げ切った（＾～＾）
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPos.UndoMove()
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMove(from, to, pro)
								pPos.DoMove(move)

								if pPos.ControlBoards[opponent-1][friendKingSq] == 0 {
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPos.UndoMove()
							}
						}
					}
				}
			}
		}
	} else {
		// 王手されていないぜ（＾～＾）
		// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := l03.Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) && pPos.ControlBoards[opponent-1][to] == 0 { // 自駒の上、敵の利きには移動できません
								move_list = append(move_list, NewMove(from, to, pro))
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move_list = append(move_list, NewMove(from, to, pro))
							}
						}
					}
				}
			}
		}

		// 駒台もスキャンしよ（＾～＾）
		phase_index := l03.Square(pPos.Phase - 1)
		for hand := l03.Square(phase_index * l03.HANDSQ_TYPE_SIZE.ToSq()); hand < (phase_index+1)*l03.HANDSQ_TYPE_SIZE.ToSq(); hand += 1 {
			if pPos.Hands[hand] > 0 {
				hand_sq := hand + l03.HANDSQ_ORIGIN.ToSq()
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, pro := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move_list = append(move_list, NewMove(hand_sq, to, pro))
					}
				}
			}
		}
	}

	// TODO 打もやりたい（＾～＾）

	return move_list
}
