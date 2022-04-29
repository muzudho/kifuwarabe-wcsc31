package take14

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l13 "github.com/muzudho/kifuwarabe-wcsc31/take13"
)

// 条件
var cab = 20
var cac = cab + 20
var cb = cac + 20
var cc = cb + 20
var cd = cc + 20
var ce = cd + 20

// 盤上の駒、駒台の駒に対して、ルールを実装すればいいはず（＾～＾）
var genmv_k1 = []int{cb + 1, cc + 2, 3, 4, 5}
var genmv_k2 = genmv_k1
var genmv_r1 = []int{cb + 1, cc + 2, 5, cab + 6, cac + 7, 9, 10, 11, 12, 14, 16}
var genmv_r2 = genmv_r1
var genmv_pr1 = append(genmv_k1, []int{9, 10, 14}...)
var genmv_pr2 = genmv_pr1
var genmv_b1 = []int{3, 4, 8, 13, 15}
var genmv_b2 = genmv_b1
var genmv_pb1 = append(genmv_k1, []int{13}...)
var genmv_pb2 = genmv_pb1
var genmv_g1 = []int{cb + 1, cc + 2, 3, 5}
var genmv_ps1 = genmv_g1
var genmv_pn1 = genmv_g1
var genmv_pl1 = genmv_g1
var genmv_pp1 = genmv_g1
var genmv_g2 = []int{cb + 1, cc + 2, 4, 5}
var genmv_ps2 = genmv_g2
var genmv_pn2 = genmv_g2
var genmv_pl2 = genmv_g2
var genmv_pp2 = genmv_g2
var genmv_s1 = []int{cb + 1, 3, 4, cab + 6, 8}
var genmv_s2 = []int{cc + 2, 3, 4, cac + 7, 8}
var genmv_p1 = []int{cd + 1, 6}
var genmv_p2 = []int{ce + 2, 7}
var genmv_l1 = append(genmv_p1, []int{9, 11}...)
var genmv_l2 = append(genmv_p2, []int{10, 12}...)

// 打にはオフセット 30 足しとく
var genmv_dr1 = []int{31, 32, 33, 34, 35, 36, 37}
var genmv_db1 = genmv_dr1
var genmv_dg1 = genmv_dr1
var genmv_ds1 = genmv_dr1
var genmv_dr2 = genmv_dr1
var genmv_db2 = genmv_dr1
var genmv_dg2 = genmv_dr1
var genmv_ds2 = genmv_dr1
var genmv_dn1 = []int{33, 34, 35, 36, 37}
var genmv_dn2 = []int{31, 32, 33, 34, 35}
var genmv_dl1 = []int{32, 33, 34, 35, 36, 37}
var genmv_dl2 = []int{31, 32, 33, 34, 35, 36}
var genmv_dp1 = []int{32, 33, 34, 35, 36, 37} // と、二歩チェック
var genmv_dp2 = []int{31, 32, 33, 34, 35, 36} // と、二歩チェック

// 成りの手は生成しません
const NOT_PROMOTE = false

// 条件Aの1. 移動元が敵陣だ
func FromOpponent(phase l03.Phase, from l03.Square) bool {
	switch phase {
	case l03.FIRST:
		return l03.Rank(from) < 4
	case l03.SECOND:
		return l03.Rank(from) > 6
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
	}
}

// GenMoveEnd - 利いているマスの一覧を返します。動けるマスではありません。
// 成らないと移動できないが、成れば移動できるマスがあるので、移動先と成りの２つセットで返します。
// TODO 成る、成らないも入れたいぜ（＾～＾）
func GenMoveEnd(pPos *Position, from l03.Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	return moveEndList
	/*
		var genmv_list []int

		if from == l03.SQ_EMPTY {
			panic(App.LogNotEcho.Fatal("GenMoveEnd has empty square"))
		} else if OnHands(from) {
			// 打なら
			switch from {
			case l03.SQ_R1:
				genmv_list = genmv_dr1
			case l03.SQ_B1:
				genmv_list = genmv_db1
			case l03.SQ_G1:
				genmv_list = genmv_dg1
			case l03.SQ_S1:
				genmv_list = genmv_ds1
			case l03.SQ_N1:
				genmv_list = genmv_dn1
			case l03.SQ_L1:
				genmv_list = genmv_dl1
			case l03.SQ_P1:
				genmv_list = genmv_dp1
			case l03.SQ_R2:
				genmv_list = genmv_dr1
			case l03.SQ_B2:
				genmv_list = genmv_db1
			case l03.SQ_G2:
				genmv_list = genmv_dg1
			case l03.SQ_S2:
				genmv_list = genmv_ds1
			case l03.SQ_N2:
				genmv_list = genmv_dn1
			case l03.SQ_L2:
				genmv_list = genmv_dl2
			case l03.SQ_P2:
				genmv_list = genmv_dp2
			default:
				panic(App.LogNotEcho.Fatal("unknown hand from=%d", from))
			}

			// 打てる列（インデックスと列は等しい）。二歩チェックで使う
			var droppableFiles = [10]bool{false, true, true, true, true, true, true, true, true, true}

			for _, step := range genmv_list {
				switch step {
				case 31:
					makeDrop(pPos, droppableFiles, 1, moveEndList)
				case 32:
					makeDrop(pPos, droppableFiles, 2, moveEndList)
				case 33:
					makeDrop(pPos, droppableFiles, 3, moveEndList)
				case 34:
					makeDrop(pPos, droppableFiles, 4, moveEndList)
					makeDrop(pPos, droppableFiles, 5, moveEndList)
					makeDrop(pPos, droppableFiles, 6, moveEndList)
				case 35:
					makeDrop(pPos, droppableFiles, 7, moveEndList)
				case 36:
					makeDrop(pPos, droppableFiles, 8, moveEndList)
				case 37:
					makeDrop(pPos, droppableFiles, 9, moveEndList)
				// case 31:
				// 	genmv_list = append(genmv_list, 24)
				// 	genmv_list = append(genmv_list, 25)
				// 	genmv_list = append(genmv_list, 26)
				// 	genmv_list = append(genmv_list, 27)
				// 	genmv_list = append(genmv_list, 28)
				// 	genmv_list = append(genmv_list, 29)
				// 	genmv_list = append(genmv_list, 30)
				// case 32:
				// 	genmv_list = append(genmv_list, 26)
				// 	genmv_list = append(genmv_list, 27)
				// 	genmv_list = append(genmv_list, 28)
				// 	genmv_list = append(genmv_list, 29)
				// 	genmv_list = append(genmv_list, 30)
				// case 33:
				// 	genmv_list = append(genmv_list, 24)
				// 	genmv_list = append(genmv_list, 25)
				// 	genmv_list = append(genmv_list, 26)
				// 	genmv_list = append(genmv_list, 27)
				// 	genmv_list = append(genmv_list, 28)
				// case 34:
				// 	genmv_list = append(genmv_list, 25)
				// 	genmv_list = append(genmv_list, 26)
				// 	genmv_list = append(genmv_list, 27)
				// 	genmv_list = append(genmv_list, 28)
				// 	genmv_list = append(genmv_list, 29)
				// 	genmv_list = append(genmv_list, 30)
				// case 35:
				// 	genmv_list = append(genmv_list, 24)
				// 	genmv_list = append(genmv_list, 25)
				// 	genmv_list = append(genmv_list, 26)
				// 	genmv_list = append(genmv_list, 27)
				// 	genmv_list = append(genmv_list, 28)
				// 	genmv_list = append(genmv_list, 29)
				// case 36:
				// 	for file := l03.Square(9); file > 0; file -= 1 {
				// 		if NifuFirst(pPos, file) { // ２歩禁止
				// 			droppableFiles[file] = false
				// 		}
				// 	}

				// 	genmv_list = append(genmv_list, 34)
				// case 37:
				// 	for file := l03.Square(9); file > 0; file -= 1 {
				// 		if NifuSecond(pPos, file) { // ２歩禁止
				// 			droppableFiles[file] = false
				// 		}
				// 	}
				// 	genmv_list = append(genmv_list, 35)
				default:
					panic(App.LogNotEcho.Fatal("Unknown step=%d", step))
				}
			}

		} else {
			// 打でないなら
			piece := pPos.Board[from]
			phase := l03.Who(piece)

			switch piece {
			case l03.PIECE_EMPTY:
				panic(App.LogNotEcho.Fatal("Piece empty"))
			case l03.PIECE_K1:
				genmv_list = genmv_k1
			case l03.PIECE_R1:
				genmv_list = genmv_r1
			case l03.PIECE_B1:
				genmv_list = genmv_b1
			case l03.PIECE_G1:
				genmv_list = genmv_g1
			case l03.PIECE_S1:
				genmv_list = genmv_s1
			case l03.PIECE_N1:
				// 先手桂
				if FromOpponent(phase, from) {
					var promote = l03.File(from) < 6                         // 移動元または移動先が敵陣なら成れる
					makeFrontKnightPromotion(from, promote, moveEndList) // 先手桂の利き
				}
				if l03.Rank(from) != 3 { // 移動元が3段目でない
					var promote = l03.File(from) < 6                // 移動元または移動先が敵陣なら成れる
					makeFrontKnight(from, promote, moveEndList) // 先手桂の利き
				}
			case l03.PIECE_L1:
				genmv_list = genmv_l1
			case l03.PIECE_P1:
				genmv_list = genmv_p1
			case l03.PIECE_PR1:
				genmv_list = genmv_pr1
			case l03.PIECE_PB1:
				genmv_list = genmv_pb1
			case l03.PIECE_PS1:
				genmv_list = genmv_ps1
			case l03.PIECE_PN1:
				genmv_list = genmv_pn1
			case l03.PIECE_PL1:
				genmv_list = genmv_pl1
			case l03.PIECE_PP1:
				genmv_list = genmv_pp1
			case l03.PIECE_K2:
				genmv_list = genmv_k2
			case l03.PIECE_R2:
				genmv_list = genmv_r2
			case l03.PIECE_B2:
				genmv_list = genmv_b2
			case l03.PIECE_G2:
				genmv_list = genmv_g2
			case l03.PIECE_S2:
				genmv_list = genmv_s2
			case l03.PIECE_N2:
				// 後手桂
				if l03.Rank(from) != 7 { // 移動元が7段目でない
					var promote = l03.File(from) > 4               // 移動元または移動先が敵陣なら成れる
					makeBackKnight(from, promote, moveEndList) // 後手桂の利き
				}
				if FromOpponent(phase, from) {
					var promote = l03.File(from) < 6                        // 移動元または移動先が敵陣なら成れる
					makeBackKnightPromotion(from, promote, moveEndList) // 先手桂の利き
				}
			case l03.PIECE_L2:
				genmv_list = genmv_l2
			case l03.PIECE_P2:
				genmv_list = genmv_p2
			case l03.PIECE_PR2:
				genmv_list = genmv_pr2
			case l03.PIECE_PB2:
				genmv_list = genmv_pb2
			case l03.PIECE_PS2:
				genmv_list = genmv_ps2
			case l03.PIECE_PN2:
				genmv_list = genmv_pn2
			case l03.PIECE_PL2:
				genmv_list = genmv_pl2
			case l03.PIECE_PP2:
				genmv_list = genmv_pp2
			default:
				panic(App.LogNotEcho.Fatal("unknown piece=%d", piece))
			}

			for _, step := range genmv_list {
				switch step {
				case 1:
					// 先手から見て１つ上への利き
					makeFront(from, NOT_PROMOTE, moveEndList)
				case cb + 1:
					if l03.Rank(from) != 2 { // 移動元が2段目でない
						makeFront(from, NOT_PROMOTE, moveEndList)
					}
				case 2:
					// 先手から見て１つ後ろへの利き
					makeBack(from, NOT_PROMOTE, moveEndList)
				case cc + 2:
					if l03.Rank(from) != 8 { // 移動元が8段目でない
						makeBack(from, NOT_PROMOTE, moveEndList)
					}
				case 3:
					makeFrontDiagonal(from, NOT_PROMOTE, moveEndList) // 先手から見て斜め前の利き
				case 4:
					makeBackDiagonal(from, NOT_PROMOTE, moveEndList) // 先手から見て斜め後ろの利き
				case 5:
					makeSide(from, NOT_PROMOTE, moveEndList) // 先手から見て１つ横への利き
				case 6:
					if FromOpponent(phase, from) {
						makeFrontPromotion()
					}
				case cb + 6:
					if FromOpponent(phase, from) {
						makeFrontPromotion()
					}
				case 7:
					if FromOpponent(phase, from) {
						makeBackPromotion()
					}
				case cc + 7:
					if FromOpponent(phase, from) {
						makeBackPromotion()
					}
				case 8:
					if FromOpponent(phase, from) {
						makeDiagonalPromotion()
					}
				case 9:
					makeLongFront(pPos, from, NOT_PROMOTE, moveEndList) // ２つ先のマスからの上への長い利き
				case 10:
					makeLongBack(pPos, from, NOT_PROMOTE, moveEndList) // ２つ先のマスからの下への長い利き
				case 11:
					if FromOpponent(phase, from) {
						makeLongFrontPromotion(pPos, from, CAN_PROMOTE, moveEndList)
					}
				case 12:
					if FromOpponent(phase, from) {
						makeLongBackPromotion(pPos, from, CAN_PROMOTE, moveEndList)
					}
				case 13:
					makeLongDiagonal(pPos, from, moveEndList)
				case 14:
					makeLongSide(pPos, from, NOT_PROMOTE, moveEndList) // ２つ先のマスからの横への長い利き
				case 15:
					if FromOpponent(phase, from) {
						makeLongDiagonalPromotion(pPos, phase, from, moveEndList)
					}
				case 16:
					if FromOpponent(phase, from) {
						makeSide(from, promote, moveEndList)
					}
				default:
					panic(App.LogNotEcho.Fatal("Unknown step=%d", step))
				}
			}
		}

		return moveEndList
	*/
}

// 1 先手から見て１つ前への利き
func makeFront(from l03.Square, moveEndList []MoveEnd) {
	if to := from - 1; l03.Rank(to) != 0 { // 上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
	}
}

// 2 先手から見て１つ後ろへの利き
func makeBack(from l03.Square, moveEndList []MoveEnd) {
	// promote bool,
	// 移動元が８段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 8

	if to := from + 1; l03.Rank(to) != 0 { // 下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		// if promote {
		// 	moveEndList = append(moveEndList, NewMoveEnd(to, true))
		// }
	}
}

// 3 先手から見て斜め前の利き
func makeFrontDiagonal(from l03.Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 9; l03.File(to) != 0 && l03.Rank(to) != 0 { // 左上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
	if to := from - 11; l03.File(to) != 0 && l03.Rank(to) != 0 { // 右上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

// 4 先手から見て斜め後ろの利き
func makeBackDiagonal(from l03.Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 11; l03.File(to) != 0 && l03.Rank(to) != 0 { // 左下
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
	if to := from - 9; l03.File(to) != 0 && l03.Rank(to) != 0 { // 右下
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

// 5 先手から見て１つ横への利き
func makeSide(from l03.Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 10; l03.File(to) != 0 { // 左
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
	if to := from - 10; l03.File(to) != 0 { // 右
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEnd(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

// 6 先手から見て１つ上への利き（成りの動き、制約なし）
func makeFrontPromotion(from l03.Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が２段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 2

	if to := from - 1; l03.Rank(to) != 0 { // 上
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

// 7
func makeBackPromotion(from l03.Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が２段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 2

	if to := from + 1; l03.Rank(to) != 0 { // 下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

// 8 ２つ先のマスからの斜めへの長い利き（成り手のみの生成）
func makeDiagonalPromotion(pPos *Position, phase l03.Phase, from l03.Square, moveEndList []MoveEnd) {
	var src_pro bool
	if (phase == l03.FIRST && l03.Rank(from) < 4) || (phase == l03.SECOND && l03.Rank(from) > 6) {
		src_pro = true
	} else {
		src_pro = false
	}

	if l03.File(from) < 8 && l03.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; l03.File(to) != 0 && l03.Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) > 2 && l03.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) < 8 && l03.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; l03.File(to) != 0 && l03.Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) > 2 && l03.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 9 ２つ先のマスからの上への長い利き
func makeLongFront(pPos *Position, from l03.Square, promote bool, moveEndList []MoveEnd) {
	if l03.Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
		for to := from - 2; l03.Rank(to) != 0; to -= 1 { // 上
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 10 ２つ先のマスからの下への長い利き
func makeLongBack(pPos *Position, from l03.Square, promote bool, moveEndList []MoveEnd) {
	if l03.Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
		for to := from + 2; l03.Rank(to) != 0; to += 1 { // 下
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 11
func makeLongFrontPromotion(pPos *Position, from l03.Square, promote bool, moveEndList []MoveEnd) {
	if l03.Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
		for to := from - 2; l03.Rank(to) != 0; to -= 1 { // 上
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 12
func makeLongBackPromotion(pPos *Position, from l03.Square, promote bool, moveEndList []MoveEnd) {
	if l03.Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
		for to := from + 2; l03.Rank(to) != 0; to += 1 { // 下
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 13 ２つ先のマスからの斜めへの長い利き（成らず）
func makeLongDiagonal(pPos *Position, from l03.Square, moveEndList []MoveEnd) {
	if l03.File(from) < 8 && l03.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; l03.File(to) != 0 && l03.Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) > 2 && l03.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) < 8 && l03.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; l03.File(to) != 0 && l03.Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) > 2 && l03.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

}

// 14 ２つ先のマスからの横への長い利き
func makeLongSide(pPos *Position, from l03.Square, promote bool, moveEndList []MoveEnd) {
	// ２つ先のマスからの左への長い利き
	if l03.File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
		for to := from + 20; l03.File(to) != 0; to += 10 { // 左
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

	// ２つ先のマスからの右への長い利き
	if l03.File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
		for to := from - 20; l03.File(to) != 0; to -= 10 { // 右
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 15 ２つ先のマスからの斜めへの長い利き（成り手のみの生成）
func makeLongDiagonalPromotion(pPos *Position, phase l03.Phase, from l03.Square, moveEndList []MoveEnd) {
	var src_pro bool
	if (phase == l03.FIRST && l03.Rank(from) < 4) || (phase == l03.SECOND && l03.Rank(from) > 6) {
		src_pro = true
	} else {
		src_pro = false
	}

	if l03.File(from) < 8 && l03.Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; l03.File(to) != 0 && l03.Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) > 2 && l03.Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) < 8 && l03.Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; l03.File(to) != 0 && l03.Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if l03.File(from) > 2 && l03.Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; l03.File(to) != 0 && l03.Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			if src_pro || (phase == l03.FIRST && l03.Rank(to) < 4) || (phase == l03.SECOND && l03.Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 16 ２つ先のマスからの横への長い利き
func makeSidePromotion(pPos *Position, from l03.Square, promote bool, moveEndList []MoveEnd) {
	// ２つ先のマスからの左への長い利き
	if l03.File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
		for to := from + 20; l03.File(to) != 0; to += 10 { // 左
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

	// ２つ先のマスからの右への長い利き
	if l03.File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
		for to := from - 20; l03.File(to) != 0; to -= 10 { // 右
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 17 先手桂の利き
func makeFrontKnight(from l03.Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が３段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 3

	if 2 < l03.Rank(from) && l03.Rank(from) < 10 {
		if 0 < l03.File(from) && l03.File(from) < 9 { // 左上桂馬飛び
			to := from + 8
			ValidateSq(to)
			if keepGoing {
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
		}
		if 1 < l03.File(from) && l03.File(from) < 10 { // 右上桂馬飛び
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

// 18 先手桂の利き
func makeFrontKnightPromotion(from l03.Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が３段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 3

	if 2 < l03.Rank(from) && l03.Rank(from) < 10 {
		if 0 < l03.File(from) && l03.File(from) < 9 { // 左上桂馬飛び
			to := from + 8
			ValidateSq(to)
			if keepGoing {
				moveEndList = append(moveEndList, NewMoveEnd(to, false))
			}
			if promote {
				moveEndList = append(moveEndList, NewMoveEnd(to, true))
			}
		}
		if 1 < l03.File(from) && l03.File(from) < 10 { // 右上桂馬飛び
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

// 19 後手桂の利き
func makeBackKnight(from l03.Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が７段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 7

	if to := from + 12; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 左下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
	if to := from - 8; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 右下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

// 20
func makeBackKnightPromotion(from l03.Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が７段目のときは必ずならなければならない
	var keepGoing = l03.File(from) != 7

	if to := from + 12; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 左下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
	if to := from - 8; l03.File(to) != 0 && l03.Rank(to) != 0 && l03.Rank(to) != 9 { // 右下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEnd(to, true))
		}
	}
}

func makeDrop(pPos *Position, droppableFiles [10]bool, rank l03.Square, moveEndList []MoveEnd) {
	for file := l03.Square(9); file > 0; file -= 1 {
		if droppableFiles[file] {
			to := SquareFrom(file, rank)
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEnd(to, false))
		}
	}
}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file l03.Square) bool {
	for rank := l03.Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l03.PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file l03.Square) bool {
	for rank := l03.Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == l03.PIECE_P2 {
			return true
		}
	}

	return false
}

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPosSys *PositionSystem, pPos *Position) []l13.Move {

	move_list := []l13.Move{}

	// 王手をされているときは、自玉を逃がす必要があります
	friend := pPosSys.GetPhase()
	var friendKingSq l03.Square
	var hand_start l03.HandIdx
	var hand_end l03.HandIdx
	// var opponentKingSq l03.Square
	var pOpponentSumCB *ControlBoard
	if friend == l03.FIRST {
		friendKingSq = pPos.GetPieceLocation(l11.PCLOC_K1)
		hand_start = l03.HAND_IDX_START
		pOpponentSumCB = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2]
	} else if friend == l03.SECOND {
		friendKingSq = pPos.GetPieceLocation(l11.PCLOC_K2)
		hand_start = l03.HAND_IDX_START + l03.HAND_TYPE_SIZE
		pOpponentSumCB = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1]
	} else {
		panic(App.LogNotEcho.Fatal("unknown phase=%d", friend))
	}
	hand_end = hand_start + l03.HAND_TYPE_SIZE

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
				from := l03.Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := l03.What(piece)

					if pieceType == l03.PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							// 敵の長い駒の利きは、玉が逃げても伸びてくる方向があるので、
							// いったん玉を動かしてから 再チェックするぜ（＾～＾）
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := l13.NewMove(from, to, pro)
								pPosSys.DoMove(pPos, move)

								if pOpponentSumCB.Board1[to] == 0 {
									// よっしゃ利きから逃げ切った（＾～＾）
									// 王手が解除されてるから採用（＾～＾）
									move_list = append(move_list, move)
								}

								pPosSys.UndoMove(pPos)
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := l13.NewMove(from, to, pro)
								pPosSys.DoMove(pPos, move)

								if pOpponentSumCB.Board1[friendKingSq] == 0 {
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
				hand_sq := l03.Square(hand_index) + l03.SQ_HAND_START
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, pro := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move := l13.NewMove(hand_sq, to, pro)
						pPosSys.DoMove(pPos, move)

						if pOpponentSumCB.Board1[friendKingSq] == 0 {
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
				from := l03.Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					moveEndList := GenMoveEnd(pPos, from)

					piece := pPos.Board[from]
					pieceType := l03.What(piece)

					if pieceType == l03.PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) && pOpponentSumCB.Board1[to] == 0 { // 自駒の上、敵の利きには移動できません
								move_list = append(move_list, l13.NewMove(from, to, pro))
							}
						}
					} else {
						for _, moveEnd := range moveEndList {
							to, pro := moveEnd.Destructure()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move_list = append(move_list, l13.NewMove(from, to, pro))
							}
						}
					}
				}
			}
		}

		// 自分の駒台もスキャンしよ（＾～＾）
		for hand_index := hand_start; hand_index < hand_end; hand_index += 1 {
			if pPos.Hands1[hand_index] > 0 {
				hand_sq := l03.Square(hand_index) + l03.SQ_HAND_START
				moveEndList := GenMoveEnd(pPos, hand_sq)

				for _, moveEnd := range moveEndList {
					to, pro := moveEnd.Destructure()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move_list = append(move_list, l13.NewMove(hand_sq, to, pro))
					}
				}
			}
		}
	}

	return move_list
}
