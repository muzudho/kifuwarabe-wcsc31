package take15

import l13 "github.com/muzudho/kifuwarabe-wcsc31/take13"

// GenMoveEnd - 利いているマスの一覧を返します。動けるマスではありません。
// 成らないと移動できないが、成れば移動できるマスがあるので、移動先と成りの２つセットで返します。
func GenMoveEnd(pPos *Position, from Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	var rank_from = Rank(from)

	/*
		// 盤上の駒、駒台の駒に対して、37個のルールを実装すればいいはず（＾～＾）
		k1 := []int{2}
		r1 := []int{13}
		b1 := []int{14}
		g1 := []int{3}
		s1 := []int{19, 21}
		n1 := []int{15}
		l1 := []int{17, 22}
		p1 := []int{17}
		pr1 := []int{5}
		pb1 := []int{6}
		ps1 := []int{3}
		pn1 := []int{3}
		pl1 := []int{3}
		pp1 := []int{3}
		k2 := []int{2}
		r2 := []int{13}
		b2 := []int{14}
		g2 := []int{4}
		s2 := []int{20, 21}
		n2 := []int{16}
		l2 := []int{18, 23}
		p2 := []int{18}
		pr2 := []int{5}
		pb2 := []int{6}
		ps2 := []int{4}
		pn2 := []int{4}
		pl2 := []int{4}
		pp2 := []int{4}

		dr1 := []int{31}
		db1 := []int{31}
		dg1 := []int{31}
		ds1 := []int{31}
		dn1 := []int{32}
		dl1 := []int{34}
		dp1 := []int{36}
		dr2 := []int{31}
		db2 := []int{31}
		dg2 := []int{31}
		ds2 := []int{31}
		dn2 := []int{33}
		dl2 := []int{35}
		dp2 := []int{37}
	*/

	if from == SQUARE_EMPTY {
		panic(App.LogNotEcho.Fatal("GenMoveEnd has empty square"))
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
			panic(App.LogNotEcho.Fatal("unknown hand from=%d", from))
		}

		switch from {
		case SQ_P1: // 先手Pawn
			// TODO 打ち歩詰め禁止
			for rank := Square(start_rank); rank < end_rank; rank += 1 {
				for file := Square(9); file > 0; file-- {
					if !NifuFirst(pPos, file) { // ２歩禁止
						to := SquareFrom(file, rank)
						ValidateSq(to)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		case SQ_P2:
			// TODO 打ち歩詰め禁止
			for rank := Square(start_rank); rank < end_rank; rank += 1 {
				for file := Square(9); file > 0; file-- {
					if !NifuSecond(pPos, file) { // ２歩禁止
						to := SquareFrom(file, rank)
						ValidateSq(to)
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
				}
			}
		default:
			for rank := Square(start_rank); rank < end_rank; rank += 1 {
				for file := Square(9); file > 0; file-- {
					to := SquareFrom(file, rank)
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
			}
		}
	} else {
		// 盤上の駒の利き
		piece := pPos.Board[from]

		// ２つ先のマスから斜めに長い利き
		switch piece {
		case l13.PIECE_B1, l13.PIECE_PB1, l13.PIECE_B2, l13.PIECE_PB2:
			if File(from) < 8 && rank_from > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
				for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) > 2 && rank_from > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
				for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) < 8 && rank_from < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
				for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) > 2 && rank_from < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
				for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
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

		// ２つ先のマスから先手香車の長い利き（不成の動き）
		// 4～9段目にある香で、１つ上が空マスなら
		if piece == l13.PIECE_L1 && 4 <= rank_from && pPos.IsEmptySq(from-1) {
			for to := from - 2; 2 <= Rank(to); to -= 1 { // fromの２段上～2段目
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}

		// ２つ先のマスから後手香車の長い利き（不成の動き）
		// 1～6段目にある香で、１つ下が空マスなら
		if piece == l13.PIECE_L2 && rank_from <= 6 && pPos.IsEmptySq(from+1) {
			for to := from + 2; Rank(to) <= 8; to += 1 { // fromの2段下～8段目
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
				if !pPos.IsEmptySq(to) {
					break
				}
			}
		}

		// ２つ先のマスから先手飛・竜の長い利き（不成の動き）
		switch piece {
		case l13.PIECE_L1, l13.PIECE_R1, l13.PIECE_PR1, l13.PIECE_R2, l13.PIECE_PR2:
			if rank_from > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
				for to := from - 2; Rank(to) != 0; to -= 1 { // 上
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

		// ２つ先のマスから後手飛・竜の長い利き（不成の動き）
		switch piece {
		case l13.PIECE_R1, l13.PIECE_PR1, l13.PIECE_L2, l13.PIECE_R2, l13.PIECE_PR2:
			if rank_from < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
				for to := from + 2; Rank(to) != 0; to += 1 { // 下
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
		case l13.PIECE_R1, l13.PIECE_PR1, l13.PIECE_R2, l13.PIECE_PR2:
			if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
				for to := from + 20; File(to) != 0; to += 10 { // 左
					ValidateSq(to)
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
					if !pPos.IsEmptySq(to) {
						break
					}
				}
			}
			if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
				for to := from - 20; File(to) != 0; to -= 10 { // 右
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
		if piece == l13.PIECE_N1 {
			// 成らず駒の 成らず の動きを作るか（＾～＾）？
			var keepGoing bool
			if 5 <= rank_from {
				keepGoing = true
			} else {
				keepGoing = false
			}
			// 成り の動きを作るか（＾～＾）？
			var promote bool
			if 3 <= rank_from && rank_from <= 5 {
				promote = true
			} else {
				keepGoing = true
			}

			if 2 < rank_from && rank_from < 10 {
				if 0 < File(from) && File(from) < 9 { // 左上桂馬飛び
					to := from + 8
					ValidateSq(to)
					if keepGoing {
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
					if promote {
						moveEndList = append(moveEndList, NewMoveEnd(to, true))
					}
				}
				if 1 < File(from) && File(from) < 10 { // 右上桂馬飛び
					to := from - 12
					ValidateSq(to)
					if keepGoing {
						moveEndList = append(moveEndList, NewMoveEnd(to, false))
					}
					if promote {
						moveEndList = append(moveEndList, NewMoveEnd(to, true))
					}
				}
			}
		}

		// 後手桂の利き
		if piece == l13.PIECE_N2 {
			// 成らず の動きを作るか（＾～＾）？
			var keepGoing bool
			if rank_from <= 5 {
				keepGoing = true
			} else {
				keepGoing = false
			}
			// 成り の動きを作るか（＾～＾）？
			var promote bool
			if 5 <= rank_from && rank_from <= 7 {
				promote = true
			} else {
				promote = false
			}

			if to := from + 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左下
				ValidateSq(to)
				if keepGoing {
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
				if promote {
					moveEndList = append(moveEndList, NewMoveEnd(to, true))
				}
			}
			if to := from - 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右下
				ValidateSq(to)
				if keepGoing {
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
				if promote {
					moveEndList = append(moveEndList, NewMoveEnd(to, true))
				}
			}
		}

		// 先手歩の利き
		switch piece {
		case l13.PIECE_P1:
			if 2 <= rank_from {
				// 成らず駒の 成らず の動きを作るか（＾～＾）？
				var keepGoing bool
				if 3 <= rank_from {
					keepGoing = true
				} else {
					keepGoing = false
				}
				// 成り の動きを作るか（＾～＾）？
				var promote bool
				if 2 <= rank_from && rank_from <= 4 {
					promote = true
				} else {
					promote = false
				}

				to := from - 1 // 上
				ValidateSq(to)
				if keepGoing {
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
				if promote {
					moveEndList = append(moveEndList, NewMoveEnd(to, true))
				}
			}

		case l13.PIECE_L1:
			// 成り の動きを作るか（＾～＾）？
			var promote bool
			if 2 <= rank_from && rank_from <= 4 {
				promote = true
			} else {
				promote = false
			}

			if to := from - 1; Rank(to) != 0 { // 上
				ValidateSq(to)

				// 成らず駒の 成らず の動きを作るか（＾～＾）？
				var keepGoing bool
				if 2 <= Rank(to) {
					keepGoing = true
				} else {
					keepGoing = false
				}

				if keepGoing {
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
				if promote {
					moveEndList = append(moveEndList, NewMoveEnd(to, true))
				}
			}
		case l13.PIECE_K1, l13.PIECE_R1, l13.PIECE_PR1, l13.PIECE_PB1, l13.PIECE_G1, l13.PIECE_S1, l13.PIECE_PS1,
			l13.PIECE_PN1, l13.PIECE_PL1, l13.PIECE_PP1, l13.PIECE_K2, l13.PIECE_R2, l13.PIECE_PR2, l13.PIECE_PB2, l13.PIECE_G2, l13.PIECE_PS2,
			l13.PIECE_PN2, l13.PIECE_PL2, l13.PIECE_PP2:
			if to := from - 1; Rank(to) != 0 { // 上
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
				// moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
		default:
			// Ignored
		}

		// 後手歩の利き
		switch piece {
		case l13.PIECE_P2:
			if rank_from <= 8 {
				to := from + 1 // 下
				// 成らず駒の 成らず の動きを作るか（＾～＾）？
				var keepGoing bool
				if rank_from <= 7 {
					keepGoing = true
				} else {
					keepGoing = false
				}
				// 成り の動きを作るか（＾～＾）？
				var promote bool
				if 6 <= rank_from && rank_from <= 8 {
					promote = true
				} else {
					promote = false
				}

				ValidateSq(to)
				if keepGoing {
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
				if promote {
					moveEndList = append(moveEndList, NewMoveEnd(to, true))
				}
			}
		case l13.PIECE_L2:
			// 成り の動きを作るか（＾～＾）？
			var promote bool
			if 6 <= rank_from && rank_from <= 8 {
				promote = true
			} else {
				promote = false
			}

			if to := from + 1; Rank(to) != 0 { // 下
				ValidateSq(to)

				// 成らず駒の 成らず の動きを作るか（＾～＾）？
				var keepGoing bool
				if Rank(to) <= 8 {
					keepGoing = true
				} else {
					keepGoing = false
				}

				if keepGoing {
					moveEndList = append(moveEndList, NewMoveEnd(to, false))
				}
				if promote {
					moveEndList = append(moveEndList, NewMoveEnd(to, true))
				}
			}
		case l13.PIECE_K2, l13.PIECE_R2, l13.PIECE_PR2, l13.PIECE_PB2, l13.PIECE_G2, l13.PIECE_S2, l13.PIECE_PS2,
			l13.PIECE_PN2, l13.PIECE_PL2, l13.PIECE_PP2, l13.PIECE_K1, l13.PIECE_R1, l13.PIECE_PR1, l13.PIECE_PB1, l13.PIECE_G1, l13.PIECE_PS1,
			l13.PIECE_PN1, l13.PIECE_PL1, l13.PIECE_PP1:
			if to := from + 1; Rank(to) != 0 { // 下
				ValidateSq(to)
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
				// moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
		default:
			// Ignored
		}

		// 先手斜め前の利き
		switch piece {
		case l13.PIECE_K1, l13.PIECE_PR1, l13.PIECE_B1, l13.PIECE_PB1, l13.PIECE_G1, l13.PIECE_S1, l13.PIECE_PS1, l13.PIECE_PN1, l13.PIECE_PL1,
			l13.PIECE_PP1, l13.PIECE_K2, l13.PIECE_PR2, l13.PIECE_B2, l13.PIECE_PB2, l13.PIECE_S2:
			if to := from + 9; File(to) != 0 && Rank(to) != 0 { // 左上
				ValidateSq(to)
				moveEnd := NewMoveEnd(to, false)
				moveEndList = append(moveEndList, moveEnd)
			}
			if to := from - 11; File(to) != 0 && Rank(to) != 0 { // 右上
				ValidateSq(to)
				moveEnd := NewMoveEnd(to, false)
				moveEndList = append(moveEndList, moveEnd)
			}
		default:
			// Ignored
		}

		// 後手斜め前の利き
		switch piece {
		case l13.PIECE_K2, l13.PIECE_PR2, l13.PIECE_B2, l13.PIECE_PB2, l13.PIECE_G2, l13.PIECE_S2, l13.PIECE_PS2, l13.PIECE_PN2, l13.PIECE_PL2,
			l13.PIECE_PP2, l13.PIECE_K1, l13.PIECE_PR1, l13.PIECE_B1, l13.PIECE_PB1, l13.PIECE_S1:
			if to := from + 11; File(to) != 0 && Rank(to) != 0 { // 左下
				ValidateSq(to)
				moveEnd := NewMoveEnd(to, false)
				moveEndList = append(moveEndList, moveEnd)
			}
			if to := from - 9; File(to) != 0 && Rank(to) != 0 { // 右下
				ValidateSq(to)
				moveEnd := NewMoveEnd(to, false)
				moveEndList = append(moveEndList, moveEnd)
			}
		default:
			// Ignored
		}

		// 横１マスの利き
		switch piece {
		case l13.PIECE_K1, l13.PIECE_R1, l13.PIECE_PR1, l13.PIECE_PB1, l13.PIECE_G1, l13.PIECE_PS1, l13.PIECE_PN1, l13.PIECE_PL1, l13.PIECE_PP1,
			l13.PIECE_K2, l13.PIECE_R2, l13.PIECE_PR2, l13.PIECE_PB2, l13.PIECE_G2, l13.PIECE_PS2, l13.PIECE_PN2, l13.PIECE_PL2, l13.PIECE_PP2:
			if to := from + 10; File(to) != 0 { // 左
				ValidateSq(to)
				moveEnd := NewMoveEnd(to, false)
				moveEndList = append(moveEndList, moveEnd)
			}
			if to := from - 10; File(to) != 0 { // 右
				ValidateSq(to)
				moveEnd := NewMoveEnd(to, false)
				moveEndList = append(moveEndList, moveEnd)
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
		if pPos.Board[SquareFrom(file, rank)] == l13.PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file Square) bool {
	for rank := Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l13.PIECE_P2 {
			return true
		}
	}

	return false
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pBrain *Brain, pPos *Position) []Move {

	move_list := []Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friend := pBrain.PPosSys.GetPhase()
	var friendKingSq Square
	var hand_start int
	var hand_end int
	// var opponentKingSq Square
	var pOpponentSumCB *ControlBoard
	if friend == FIRST {
		friendKingSq = pPos.GetPieceLocation(PCLOC_K1)
		hand_start = HAND_IDX_START
		pOpponentSumCB = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2]
	} else if friend == SECOND {
		friendKingSq = pPos.GetPieceLocation(PCLOC_K2)
		hand_start = HAND_IDX_START + HAND_TYPE_SIZE
		pOpponentSumCB = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1]
	} else {
		panic(App.LogNotEcho.Fatal("unknown phase=%d", friend))
	}
	hand_end = hand_start + HAND_TYPE_SIZE

	if !OnBoard(friendKingSq) {
		// 自玉が盤上にない場合は、指し手を返しません

	} else if pOpponentSumCB.Board1[friendKingSq] > 0 {
		// 相手の利きテーブルの自玉のマスに利きがあるか
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
								pBrain.DoMove(pPos, move)

								if pOpponentSumCB.Board1[to] == 0 {
									// よっしゃ利きから逃げ切った（＾～＾）
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pBrain.UndoMove(pPos)
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMove(from, to, pro)
								pBrain.DoMove(pPos, move)

								if pOpponentSumCB.Board1[friendKingSq] == 0 {
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pBrain.UndoMove(pPos)
							}
						}
					}
				}
			}
		}

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands1[hand_index] > 0 {
				hand_sq := Square(hand_index) + SQ_HAND_START
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, pro := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move := NewMove(hand_sq, to, pro)
						pBrain.DoMove(pPos, move)

						if pOpponentSumCB.Board1[friendKingSq] == 0 {
							// 王手が解除されてるから採用（＾～＾）
							move_list = append(move_list, move)
						}

						pBrain.UndoMove(pPos)

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
							if pPos.Hetero(from, to) && pOpponentSumCB.Board1[to] == 0 { // 自駒の上、敵の利きには移動できません
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
			if pPos.Hands1[hand_index] > 0 {
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
