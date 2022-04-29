package take12

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// マス番号が正常値でなければ強制終了させます
func ValidateSq(sq l04.Square) {
	if !OnBoard(sq) && !OnHands(sq) {
		panic(fmt.Errorf("TestSq: sq=%d", sq))
	}
}

// GenMoveEnd - 利いているマスの一覧を返します。動けるマスではありません。
func GenMoveEnd(pPos *Position, from l04.Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	if from == l04.SQ_EMPTY {
		panic(fmt.Errorf("GenMoveEnd has empty square"))
	} else if OnHands(from) {
		// どこに打てるか
		var start_rank l04.Square
		var end_rank l04.Square

		switch from {
		case l04.SQ_R1, l04.SQ_B1, l04.SQ_G1, l04.SQ_S1, l04.SQ_R2, l04.SQ_B2, l04.SQ_G2, l04.SQ_S2: // 81マスに打てる
			start_rank = 1
			end_rank = 10
		case l04.SQ_N1: // 3～9段目に打てる
			start_rank = 3
			end_rank = 10
		case l04.SQ_L1, l04.SQ_P1: // 2～9段目に打てる
			start_rank = 2
			end_rank = 10
		case l04.SQ_N2: // 1～7段目に打てる
			start_rank = 1
			end_rank = 8
		case l04.SQ_L2, l04.SQ_P2: // 1～8段目に打てる
			start_rank = 1
			end_rank = 9
		default:
			panic(fmt.Errorf("unknown hand from=%d", from))
		}

		switch from {
		case l04.SQ_P1:
			// TODO 打ち歩詰め禁止
			for rank := l04.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l04.Square(9); file > 0; file-- {
					if !NifuFirst(pPos, file) { // ２歩禁止
						sq := SquareFrom(file, rank)
						ValidateSq(sq)
						moveEndList = append(moveEndList, NewMoveEnd(sq, false))
					}
				}
			}
		case l04.SQ_P2:
			// TODO 打ち歩詰め禁止
			for rank := l04.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l04.Square(9); file > 0; file-- {
					if !NifuSecond(pPos, file) { // ２歩禁止
						sq := SquareFrom(file, rank)
						ValidateSq(sq)
						moveEndList = append(moveEndList, NewMoveEnd(sq, false))
					}
				}
			}
		default:
			for rank := l04.Square(start_rank); rank < end_rank; rank += 1 {
				for file := l04.Square(9); file > 0; file-- {
					sq := SquareFrom(file, rank)
					ValidateSq(sq)
					moveEndList = append(moveEndList, NewMoveEnd(sq, false))
				}
			}
		}
	} else {
		// 盤上の駒の利き
		piece := pPos.Board[from]

		// ２つ先のマスから斜めに長い利き
		switch piece {
		case l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_B2, l03.PIECE_PB2:
			if l04.File(from) < 8 && l04.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
				for to := from + 18; l04.File(to) != 0 && l04.Rank(to) != 0; to += 9 { // ２つ左上から
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) > 2 && l04.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
				for to := from - 22; l04.File(to) != 0 && l04.Rank(to) != 0; to -= 11 { // ２つ右上から
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) < 8 && l04.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
				for to := from + 22; l04.File(to) != 0 && l04.Rank(to) != 0; to += 11 { // ２つ左下から
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) > 2 && l04.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
				for to := from - 18; l04.File(to) != 0 && l04.Rank(to) != 0; to -= 9 { // ２つ右下から
					ValidateSq(to)
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
		case l03.PIECE_L1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_R2, l03.PIECE_PR2:
			if l04.Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
				for to := from - 2; l04.Rank(to) != 0; to -= 1 { // 上
					ValidateSq(to)
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
		case l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_L2, l03.PIECE_R2, l03.PIECE_PR2:
			if l04.Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
				for to := from + 2; l04.Rank(to) != 0; to += 1 { // 下
					ValidateSq(to)
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
		case l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_R2, l03.PIECE_PR2:
			if l04.File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
				for to := from + 20; l04.File(to) != 0; to += 10 { // 左
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if l04.File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
				for to := from - 20; l04.File(to) != 0; to -= 10 { // 右
					ValidateSq(to)
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
		if piece == l03.PIECE_N1 && 2 < l04.Rank(from) && l04.Rank(from) < 10 {
			if 0 < l04.File(from) && l04.File(from) < 9 { // 左上桂馬飛び
				to := from + 8
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if 1 < l04.File(from) && l04.File(from) < 10 { // 右上桂馬飛び
				to := from - 12
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 後手桂の利き
		if piece == l03.PIECE_N2 {
			if to := from + 12; l04.File(to) != 0 && l04.Rank(to) != 0 && l04.Rank(to) != 9 { // 左下
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 8; l04.File(to) != 0 && l04.Rank(to) != 0 && l04.Rank(to) != 9 { // 右下
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		}

		// 先手歩の利き
		switch piece {
		case l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_S1, l03.PIECE_L1, l03.PIECE_P1, l03.PIECE_PS1,
			l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1, l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_PS2,
			l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2:
			if to := from - 1; l04.Rank(to) != 0 { // 上
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手歩の利き
		switch piece {
		case l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_S2, l03.PIECE_L2, l03.PIECE_P2, l03.PIECE_PS2,
			l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2, l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_PS1,
			l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1:
			if to := from + 1; l04.Rank(to) != 0 { // 下
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 先手斜め前の利き
		switch piece {
		case l03.PIECE_K1, l03.PIECE_PR1, l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_S1, l03.PIECE_PS1, l03.PIECE_PN1, l03.PIECE_PL1,
			l03.PIECE_PP1, l03.PIECE_K2, l03.PIECE_PR2, l03.PIECE_B2, l03.PIECE_PB2, l03.PIECE_S2:
			if to := from + 9; l04.File(to) != 0 && l04.Rank(to) != 0 { // 左上
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 11; l04.File(to) != 0 && l04.Rank(to) != 0 { // 右上
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 後手斜め前の利き
		switch piece {
		case l03.PIECE_K2, l03.PIECE_PR2, l03.PIECE_B2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_S2, l03.PIECE_PS2, l03.PIECE_PN2, l03.PIECE_PL2,
			l03.PIECE_PP2, l03.PIECE_K1, l03.PIECE_PR1, l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_S1:
			if to := from + 11; l04.File(to) != 0 && l04.Rank(to) != 0 { // 左下
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 9; l04.File(to) != 0 && l04.Rank(to) != 0 { // 右下
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}

		// 横１マスの利き
		switch piece {
		case l03.PIECE_K1, l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_PB1, l03.PIECE_G1, l03.PIECE_PS1, l03.PIECE_PN1, l03.PIECE_PL1, l03.PIECE_PP1,
			l03.PIECE_K2, l03.PIECE_R2, l03.PIECE_PR2, l03.PIECE_PB2, l03.PIECE_G2, l03.PIECE_PS2, l03.PIECE_PN2, l03.PIECE_PL2, l03.PIECE_PP2:
			if to := from + 10; l04.File(to) != 0 { // 左
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if to := from - 10; l04.File(to) != 0 { // 右
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
		default:
			// Ignored
		}
	}

	return moveEndList
}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file l04.Square) bool {
	for rank := l04.Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l03.PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file l04.Square) bool {
	for rank := l04.Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l03.PIECE_P2 {
			return true
		}
	}

	return false
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPosSys *PositionSystem, pPos *Position) []Move {

	move_list := []Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friend := pPosSys.GetPhase()
	var friendKingSq l04.Square
	var hand_start l11.HandIdx
	var hand_end l11.HandIdx
	// var opponentKingSq l04.Square
	var pOpponentSumCB *ControlBoard
	if friend == l06.FIRST {
		friendKingSq = pPos.GetPieceLocation(l11.PCLOC_K1)
		hand_start = l11.HAND_IDX_START
		pOpponentSumCB = pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_SUM2]
	} else if friend == l06.SECOND {
		friendKingSq = pPos.GetPieceLocation(l11.PCLOC_K2)
		hand_start = l11.HAND_IDX_START + l11.HAND_TYPE_SIZE
		pOpponentSumCB = pPosSys.PControlBoardSystem.Boards[CONTROL_LAYER_SUM1]
	} else {
		panic(fmt.Errorf("unknown phase=%d", friend))
	}
	hand_end = hand_start + l11.HAND_TYPE_SIZE

	// 相手の利きテーブルの自玉のマスに利きがあるか
	if pOpponentSumCB.Board[friendKingSq] > 0 {
		// 王手されています
		// fmt.Printf("Debug: Checked friendKingSq=%d opponentKingSq=%d friend=%d opponent=%d\n", friendKingSq, opponentKingSq, friend, opponent)
		// TODO アタッカーがどの駒か調べたいが……。一手前に動かした駒か、空き王手のどちらかしかなくないか（＾～＾）？
		// 王手されているところが開始局面だと、一手前を調べることができないので、やっぱ調べるしか（＾～＾）
		// 空き王手を利用して、2箇所から 長い利きが飛んでくることはある（＾～＾）

		// 盤上の駒を動かしてみて、王手が解除されるか調べるか（＾～＾）
		for rank := 1; rank < 10; rank += 1 {
			for file := 1; file < 10; file += 1 {
				from := l04.Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := l11.What(piece)

					if pieceType == l11.PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, _ := moveEnd.Destructure()
							// 敵の長い駒の利きは、玉が逃げても伸びてくる方向があるので、
							// いったん玉を動かしてから 再チェックするぜ（＾～＾）
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMove2(from, to)
								pPosSys.DoMove(pPos, move)

								if pOpponentSumCB.Board[to] == 0 {
									// よっしゃ利きから逃げ切った（＾～＾）
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPosSys.UndoMove(pPos)
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, _ := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMove2(from, to)
								pPosSys.DoMove(pPos, move)

								if pOpponentSumCB.Board[friendKingSq] == 0 {
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPosSys.UndoMove(pPos)
							}
						}
					}
				}
			}
		}

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands1[hand_index] > 0 {
				hand_sq := l04.Square(hand_index) + l04.SQ_HAND_START
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, _ := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move := NewMove2(hand_sq, to)
						pPosSys.DoMove(pPos, move)

						if pOpponentSumCB.Board[friendKingSq] == 0 {
							// 王手が解除されてるから採用（＾～＾）
							move_list = append(move_list, move)
						}

						pPosSys.UndoMove(pPos)

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
				from := l04.Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := l11.What(piece)

					if pieceType == l11.PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, _ := moveEnd.Destructure()
							if pPos.Hetero(from, to) && pOpponentSumCB.Board[to] == 0 { // 自駒の上、敵の利きには移動できません
								move_list = append(move_list, NewMove2(from, to))
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, _ := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move_list = append(move_list, NewMove2(from, to))
							}
						}
					}
				}
			}
		}

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands1[hand_index] > 0 {
				hand_sq := l04.Square(hand_index) + l04.SQ_HAND_START
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, _ := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move_list = append(move_list, NewMove2(hand_sq, to))
					}
				}
			}
		}
	}

	return move_list
}
