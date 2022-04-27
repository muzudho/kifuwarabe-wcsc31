package take9

import (
	"fmt"

	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// GenMoveEnd - 利いているマスの一覧を返します。動けるマスではありません。
func GenMoveEnd(pPos *Position, from l04.Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	if from == l04.SQ_EMPTY {
		panic(fmt.Errorf("GenMoveEnd has empty square"))
	} else if OnBoard(from) {
		// 盤上の駒の利き
		piece := pPos.Board[from]

		// ２つ先のマスから斜めに長い利き
		switch piece {
		case PIECE_B1, PIECE_PB1, PIECE_B2, PIECE_PB2:
			if l04.File(from) < 8 && l04.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
				for to := from + 18; l04.File(to) != 0 && l04.Rank(to) != 0; to += 9 { // ２つ左上から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) > 2 && l04.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
				for to := from - 22; l04.File(to) != 0 && l04.Rank(to) != 0; to -= 11 { // ２つ右上から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) < 8 && l04.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
				for to := from + 22; l04.File(to) != 0 && l04.Rank(to) != 0; to += 11 { // ２つ左下から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) > 2 && l04.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
				for to := from - 18; l04.File(to) != 0 && l04.Rank(to) != 0; to -= 9 { // ２つ右下から
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
		case PIECE_L1, PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
			if l04.Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
				for to := from - 2; l04.Rank(to) != 0; to -= 1 { // 上
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
		case PIECE_R1, PIECE_PR1, PIECE_L2, PIECE_R2, PIECE_PR2:
			if l04.Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
				for to := from + 2; l04.Rank(to) != 0; to += 1 { // 下
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
		case PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2:
			if l04.File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
				for to := from + 20; l04.File(to) != 0; to += 10 { // 左
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
				for to := from - 20; l04.File(to) != 0; to -= 10 { // 右
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
		if piece == PIECE_N1 {
			if to := from + 8; l04.File(to) != 0 && l04.Rank(to) != 0 && l04.Rank(to) != 9 { // 左上桂馬飛び
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 12; l04.File(to) != 0 && l04.Rank(to) != 0 && l04.Rank(to) != 9 { // 右上桂馬飛び
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 後手桂の利き
		if piece == PIECE_N2 {
			if to := from + 12; l04.File(to) != 0 && l04.Rank(to) != 0 && l04.Rank(to) != 9 { // 左下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 8; l04.File(to) != 0 && l04.Rank(to) != 0 && l04.Rank(to) != 9 { // 右下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 先手歩の利き
		switch piece {
		case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_L1, PIECE_P1, PIECE_PS1,
			PIECE_PN1, PIECE_PL1, PIECE_PP1, PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2,
			PIECE_PN2, PIECE_PL2, PIECE_PP2:
			if to := from - 1; l04.Rank(to) != 0 { // 上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手歩の利き
		switch piece {
		case PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_L2, PIECE_P2, PIECE_PS2,
			PIECE_PN2, PIECE_PL2, PIECE_PP2, PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1,
			PIECE_PN1, PIECE_PL1, PIECE_PP1:
			if to := from + 1; l04.Rank(to) != 0 { // 下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 先手斜め前の利き
		switch piece {
		case PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_G1, PIECE_S1, PIECE_PS1, PIECE_PN1, PIECE_PL1,
			PIECE_PP1, PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_S2:
			if to := from + 9; l04.File(to) != 0 && l04.Rank(to) != 0 { // 左上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 11; l04.File(to) != 0 && l04.Rank(to) != 0 { // 右上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手斜め前の利き
		switch piece {
		case PIECE_K2, PIECE_PR2, PIECE_B2, PIECE_PB2, PIECE_G2, PIECE_S2, PIECE_PS2, PIECE_PN2, PIECE_PL2,
			PIECE_PP2, PIECE_K1, PIECE_PR1, PIECE_B1, PIECE_PB1, PIECE_S1:
			if to := from + 11; l04.File(to) != 0 && l04.Rank(to) != 0 { // 左下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 9; l04.File(to) != 0 && l04.Rank(to) != 0 { // 右下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 横１マスの利き
		switch piece {
		case PIECE_K1, PIECE_R1, PIECE_PR1, PIECE_PB1, PIECE_G1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1,
			PIECE_K2, PIECE_R2, PIECE_PR2, PIECE_PB2, PIECE_G2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
			if to := from + 10; l04.File(to) != 0 { // 左
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 10; l04.File(to) != 0 { // 右
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}
	} else {
		// どこに打てるか
		var start_rank l04.Square
		var end_rank l04.Square

		switch from {
		case HAND_R1, HAND_B1, HAND_G1, HAND_S1, HAND_R2, HAND_B2, HAND_G2, HAND_S2: // 81マスに打てる
			start_rank = 1
			end_rank = 10
		case HAND_N1: // 3～9段目に打てる
			start_rank = 3
			end_rank = 10
		case HAND_L1, HAND_P1: // 2～9段目に打てる
			start_rank = 2
			end_rank = 10
		case HAND_N2: // 1～7段目に打てる
			start_rank = 1
			end_rank = 8
		case HAND_L2, HAND_P2: // 1～8段目に打てる
			start_rank = 1
			end_rank = 9
		default:
			panic(fmt.Errorf("unknown hand from=%d", from))
		}

		switch from {
		case HAND_P1:
			// TODO 打ち歩詰め禁止
			for rank := l04.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l04.Square(9); file > 0; file-- {
					if !NifuFirst(pPos, file) { // ２歩禁止
						var to = SquareFrom(file, rank)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		case HAND_P2:
			// TODO 打ち歩詰め禁止
			for rank := l04.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l04.Square(9); file > 0; file-- {
					if !NifuSecond(pPos, file) { // ２歩禁止
						var to = SquareFrom(file, rank)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		default:
			for rank := l04.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l04.Square(9); file > 0; file-- {
					var to = SquareFrom(file, rank)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
			}
		}
	}

	return moveEndList
}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file l04.Square) bool {
	for rank := l04.Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file l04.Square) bool {
	for rank := l04.Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == PIECE_P2 {
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
				from := l04.Square(file*10 + rank)
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
				from := l04.Square(file*10 + rank)
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
		phase_index := l04.Square(pPos.Phase - 1)
		for hand := l04.Square(phase_index * HAND_TYPE_SIZE); hand < (phase_index+1)*HAND_TYPE_SIZE; hand += 1 {
			if pPos.Hands[hand] > 0 {
				hand_sq := hand + HAND_ORIGIN
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
