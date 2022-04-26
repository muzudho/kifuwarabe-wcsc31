package take10

import (
	"fmt"

	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// GenMoveEnd - 利いているマスの一覧を返します。動けるマスではありません。
func GenMoveEnd(pPos *Position, from Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	if from == SQUARE_EMPTY {
		panic(fmt.Errorf("GenMoveEnd has empty square"))
	} else if OnHands(from) {
		// どこに打てるか
		var start_rank Square
		var end_rank Square

		switch from {
		case SQ_R1, SQ_B1, SQ_G1, SQ_S1, SQ_R2, SQ_B2, SQ_G2, SQ_S2: // 81マスに打てる
			start_rank = 1
			end_rank = 10
		case SQ_N1: // 3～9段目に打てる
			start_rank = 3
			end_rank = 10
		case SQ_L1, SQ_P1: // 2～9段目に打てる
			start_rank = 2
			end_rank = 10
		case SQ_N2: // 1～7段目に打てる
			start_rank = 1
			end_rank = 8
		case SQ_L2, SQ_P2: // 1～8段目に打てる
			start_rank = 1
			end_rank = 9
		default:
			panic(fmt.Errorf("unknown hand from=%d", from))
		}

		switch from {
		case SQ_P1:
			// TODO 打ち歩詰め禁止
			for rank := Square(start_rank); rank < end_rank; rank += 1 {
				for file := Square(9); file > 0; file-- {
					if !NifuFirst(pPos, file) { // ２歩禁止
						var to = SquareFrom(file, rank)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		case SQ_P2:
			// TODO 打ち歩詰め禁止
			for rank := Square(start_rank); rank < end_rank; rank += 1 {
				for file := Square(9); file > 0; file-- {
					if !NifuSecond(pPos, file) { // ２歩禁止
						var to = SquareFrom(file, rank)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		default:
			for rank := Square(start_rank); rank < end_rank; rank += 1 {
				for file := Square(9); file > 0; file-- {
					var to = SquareFrom(file, rank)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
			}
		}
	} else {
		// 盤上の駒の利き
		piece := pPos.Board[from]

		// ２つ先のマスから斜めに長い利き
		switch piece {
		case l09.PIECE_B1, l09.PIECE_PB1, l09.PIECE_B2, l09.PIECE_PB2:
			if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
				for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
				for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
				for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
				for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
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
		case l09.PIECE_L1, l09.PIECE_R1, l09.PIECE_PR1, l09.PIECE_R2, l09.PIECE_PR2:
			if Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
				for to := from - 2; Rank(to) != 0; to -= 1 { // 上
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
		case l09.PIECE_R1, l09.PIECE_PR1, l09.PIECE_L2, l09.PIECE_R2, l09.PIECE_PR2:
			if Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
				for to := from + 2; Rank(to) != 0; to += 1 { // 下
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
		case l09.PIECE_R1, l09.PIECE_PR1, l09.PIECE_R2, l09.PIECE_PR2:
			if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
				for to := from + 20; File(to) != 0; to += 10 { // 左
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
				for to := from - 20; File(to) != 0; to -= 10 { // 右
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
		if piece == l09.PIECE_N1 {
			if to := from + 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左上桂馬飛び
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右上桂馬飛び
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 後手桂の利き
		if piece == l09.PIECE_N2 {
			if to := from + 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 先手歩の利き
		switch piece {
		case l09.PIECE_K1, l09.PIECE_R1, l09.PIECE_PR1, l09.PIECE_PB1, l09.PIECE_G1, l09.PIECE_S1, l09.PIECE_L1, l09.PIECE_P1, l09.PIECE_PS1,
			l09.PIECE_PN1, l09.PIECE_PL1, l09.PIECE_PP1, l09.PIECE_K2, l09.PIECE_R2, l09.PIECE_PR2, l09.PIECE_PB2, l09.PIECE_G2, l09.PIECE_PS2,
			l09.PIECE_PN2, l09.PIECE_PL2, l09.PIECE_PP2:
			if to := from - 1; Rank(to) != 0 { // 上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手歩の利き
		switch piece {
		case l09.PIECE_K2, l09.PIECE_R2, l09.PIECE_PR2, l09.PIECE_PB2, l09.PIECE_G2, l09.PIECE_S2, l09.PIECE_L2, l09.PIECE_P2, l09.PIECE_PS2,
			l09.PIECE_PN2, l09.PIECE_PL2, l09.PIECE_PP2, l09.PIECE_K1, l09.PIECE_R1, l09.PIECE_PR1, l09.PIECE_PB1, l09.PIECE_G1, l09.PIECE_PS1,
			l09.PIECE_PN1, l09.PIECE_PL1, l09.PIECE_PP1:
			if to := from + 1; Rank(to) != 0 { // 下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 先手斜め前の利き
		switch piece {
		case l09.PIECE_K1, l09.PIECE_PR1, l09.PIECE_B1, l09.PIECE_PB1, l09.PIECE_G1, l09.PIECE_S1, l09.PIECE_PS1, l09.PIECE_PN1, l09.PIECE_PL1,
			l09.PIECE_PP1, l09.PIECE_K2, l09.PIECE_PR2, l09.PIECE_B2, l09.PIECE_PB2, l09.PIECE_S2:
			if to := from + 9; File(to) != 0 && Rank(to) != 0 { // 左上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 11; File(to) != 0 && Rank(to) != 0 { // 右上
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手斜め前の利き
		switch piece {
		case l09.PIECE_K2, l09.PIECE_PR2, l09.PIECE_B2, l09.PIECE_PB2, l09.PIECE_G2, l09.PIECE_S2, l09.PIECE_PS2, l09.PIECE_PN2, l09.PIECE_PL2,
			l09.PIECE_PP2, l09.PIECE_K1, l09.PIECE_PR1, l09.PIECE_B1, l09.PIECE_PB1, l09.PIECE_S1:
			if to := from + 11; File(to) != 0 && Rank(to) != 0 { // 左下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 9; File(to) != 0 && Rank(to) != 0 { // 右下
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 横１マスの利き
		switch piece {
		case l09.PIECE_K1, l09.PIECE_R1, l09.PIECE_PR1, l09.PIECE_PB1, l09.PIECE_G1, l09.PIECE_PS1, l09.PIECE_PN1, l09.PIECE_PL1, l09.PIECE_PP1,
			l09.PIECE_K2, l09.PIECE_R2, l09.PIECE_PR2, l09.PIECE_PB2, l09.PIECE_G2, l09.PIECE_PS2, l09.PIECE_PN2, l09.PIECE_PL2, l09.PIECE_PP2:
			if to := from + 10; File(to) != 0 { // 左
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 10; File(to) != 0 { // 右
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}
	}

	return moveEndList
}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file Square) bool {
	for rank := Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l09.PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file Square) bool {
	for rank := Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l09.PIECE_P2 {
			return true
		}
	}

	return false
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []Move {

	move_list := []Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friend := pPos.GetPhase()
	opponent := FlipPhase(pPos.GetPhase())
	var friendKingSq Square
	var hand_start int
	var hand_end int
	// var opponentKingSq Square
	if friend == l06.FIRST {
		friendKingSq, _ = pPos.GetKingLocations()
		hand_start = HAND_IDX_START
	} else if friend == l06.SECOND {
		_, friendKingSq = pPos.GetKingLocations()
		hand_start = HAND_IDX_START + HAND_TYPE_SIZE
	} else {
		panic(fmt.Errorf("unknown phase=%d", friend))
	}
	hand_end = hand_start + HAND_TYPE_SIZE

	// 相手の利きテーブルの自玉のマスに利きがあるか
	if pPos.ControlBoards[opponent-1][CONTROL_LAYER_SUM][friendKingSq] > 0 {
		// 王手されています
		// fmt.Printf("Debug: Checked friendKingSq=%d opponentKingSq=%d friend=%d opponent=%d\n", friendKingSq, opponentKingSq, friend, opponent)
		// TODO アタッカーがどの駒か調べたいが……。一手前に動かした駒か、空き王手のどちらかしかなくないか（＾～＾）？
		// 王手されているところが開始局面だと、一手前を調べることができないので、やっぱ調べるしか（＾～＾）
		// 空き王手を利用して、2箇所から 長い利きが飛んでくることはある（＾～＾）

		// 盤上の駒を動かしてみて、王手が解除されるか調べるか（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := Square(file*10 + rank)
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

								if pPos.ControlBoards[opponent-1][CONTROL_LAYER_SUM][to] == 0 {
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

								if pPos.ControlBoards[opponent-1][CONTROL_LAYER_SUM][friendKingSq] == 0 {
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

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands[hand_index] > 0 {
				hand_sq := Square(hand_index) + SQ_HAND_START
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, pro := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move := NewMove(hand_sq, to, pro)
						pPos.DoMove(move)

						if pPos.ControlBoards[opponent-1][CONTROL_LAYER_SUM][friendKingSq] == 0 {
							// 王手が解除されてるから採用（＾～＾）
							move_list = append(move_list, move)
						}

						pPos.UndoMove()

					}
				}
			}
		}

	} else {
		// 王手されていないぜ（＾～＾）
		// fmt.Printf("Debug: Not checked\n")

		// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) && pPos.ControlBoards[opponent-1][CONTROL_LAYER_SUM][to] == 0 { // 自駒の上、敵の利きには移動できません
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

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands[hand_index] > 0 {
				hand_sq := Square(hand_index) + SQ_HAND_START
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

	return move_list
}
