package take14

import "fmt"

// 盤上の駒、駒台の駒に対して、ルールを実装すればいいはず（＾～＾）
var genmv_k1 = []int{1, 2, 3, 4, 5}
var genmv_k2 = genmv_k1
var genmv_r1 = []int{5, 6, 7, 9, 10, 12, 16}
var genmv_r2 = genmv_r1
var genmv_pr1 = []int{1, 2, 3, 4, 5, 9, 10, 12}
var genmv_pr2 = genmv_pr1
var genmv_b1 = []int{1, 2, 8, 11, 15}
var genmv_b2 = genmv_b1
var genmv_pb1 = []int{1, 2, 3, 4, 5, 11}
var genmv_pb2 = genmv_pb1
var genmv_g1 = []int{1, 3, 4, 5}
var genmv_ps1 = genmv_g1
var genmv_pn1 = genmv_g1
var genmv_pl1 = genmv_g1
var genmv_pp1 = genmv_g1
var genmv_g2 = []int{2, 3, 4, 5}
var genmv_ps2 = genmv_g2
var genmv_pn2 = genmv_g2
var genmv_pl2 = genmv_g2
var genmv_pp2 = genmv_g2
var genmv_s1 = []int{1, 2, 3, 6, 8}
var genmv_s2 = []int{1, 2, 4, 7, 8}
var genmv_n1 = []int{17}
var genmv_n2 = []int{18}
var genmv_l1 = []int{6, 9, 13}
var genmv_l2 = []int{7, 10, 14}
var genmv_p1 = []int{6, 13}
var genmv_p2 = []int{7, 14}

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

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}

// GenControl - 利いているマスの一覧を返します。動けるマスではありません。
// 成らないと移動できないが、成れば移動できるマスがあるので、移動先と成りの２つセットで返します。
// TODO 成る、成らないも入れたいぜ（＾～＾）
func GenControl(pPos *Position, from Square) []MoveEnd {
	moveEndList := []MoveEnd{}

	var genmv_list []int

	if from == SQUARE_EMPTY {
		panic(fmt.Errorf("GenControl has empty square"))
	} else if OnHands(from) {
		// 打なら
		switch from {
		case SQ_R1:
			genmv_list = genmv_dr1
		case SQ_B1:
			genmv_list = genmv_db1
		case SQ_G1:
			genmv_list = genmv_dg1
		case SQ_S1:
			genmv_list = genmv_ds1
		case SQ_N1:
			genmv_list = genmv_dn1
		case SQ_L1:
			genmv_list = genmv_dl1
		case SQ_P1:
			genmv_list = genmv_dp1
		case SQ_R2:
			genmv_list = genmv_dr1
		case SQ_B2:
			genmv_list = genmv_db1
		case SQ_G2:
			genmv_list = genmv_dg1
		case SQ_S2:
			genmv_list = genmv_ds1
		case SQ_N2:
			genmv_list = genmv_dn1
		case SQ_L2:
			genmv_list = genmv_dl2
		case SQ_P2:
			genmv_list = genmv_dp2
		default:
			panic(fmt.Errorf("Unknown hand from=%d", from))
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
			// 	for file := Square(9); file > 0; file -= 1 {
			// 		if NifuFirst(pPos, file) { // ２歩禁止
			// 			droppableFiles[file] = false
			// 		}
			// 	}

			// 	genmv_list = append(genmv_list, 34)
			// case 37:
			// 	for file := Square(9); file > 0; file -= 1 {
			// 		if NifuSecond(pPos, file) { // ２歩禁止
			// 			droppableFiles[file] = false
			// 		}
			// 	}
			// 	genmv_list = append(genmv_list, 35)
			default:
				panic(fmt.Errorf("Unknown step=%d", step))
			}
		}

	} else {
		// 打でないなら
		piece := pPos.Board[from]
		phase := Who(piece)

		switch piece {
		case PIECE_EMPTY:
			panic(fmt.Errorf("Piece empty"))
		case PIECE_K1:
			genmv_list = genmv_k1
		case PIECE_R1:
			genmv_list = genmv_r1
		case PIECE_B1:
			genmv_list = genmv_b1
		case PIECE_G1:
			genmv_list = genmv_g1
		case PIECE_S1:
			genmv_list = genmv_s1
		case PIECE_N1:
			genmv_list = genmv_n1
		case PIECE_L1:
			genmv_list = genmv_l1
		case PIECE_P1:
			genmv_list = genmv_p1
		case PIECE_PR1:
			genmv_list = genmv_pr1
		case PIECE_PB1:
			genmv_list = genmv_pb1
		case PIECE_PS1:
			genmv_list = genmv_ps1
		case PIECE_PN1:
			genmv_list = genmv_pn1
		case PIECE_PL1:
			genmv_list = genmv_pl1
		case PIECE_PP1:
			genmv_list = genmv_pp1
		case PIECE_K2:
			genmv_list = genmv_k2
		case PIECE_R2:
			genmv_list = genmv_r2
		case PIECE_B2:
			genmv_list = genmv_b2
		case PIECE_G2:
			genmv_list = genmv_g2
		case PIECE_S2:
			genmv_list = genmv_s2
		case PIECE_N2:
			genmv_list = genmv_n2
		case PIECE_L2:
			genmv_list = genmv_l2
		case PIECE_P2:
			genmv_list = genmv_p2
		case PIECE_PR2:
			genmv_list = genmv_pr2
		case PIECE_PB2:
			genmv_list = genmv_pb2
		case PIECE_PS2:
			genmv_list = genmv_ps2
		case PIECE_PN2:
			genmv_list = genmv_pn2
		case PIECE_PL2:
			genmv_list = genmv_pl2
		case PIECE_PP2:
			genmv_list = genmv_pp2
		default:
			panic(fmt.Errorf("Unknown piece=%d", piece))
		}

		for _, step := range genmv_list {
			switch step {
			case 1, 2, 3, 4, 7, 11, 12:
				panic(fmt.Errorf("Don't execute step=%d", step))
			case 5:
				makeFrontLong(pPos, from, NOT_PROMOTE, moveEndList) // ２つ先のマスからの上への長い利き
				makeBackLong(pPos, from, NOT_PROMOTE, moveEndList)  // ２つ先のマスからの下への長い利き
				makeSideLong(pPos, from, NOT_PROMOTE, moveEndList)  // ２つ先のマスからの横への長い利き
			case 6:
				makeDiagonalLongNotPromote(pPos, from, moveEndList) // ２つ先のマスからの斜めへの長い利き
			case 8:
				makeFrontDiagonal(from, NOT_PROMOTE, moveEndList) // 先手から見て斜め前の利き
			case 9:
				makeBackDiagonal(from, NOT_PROMOTE, moveEndList) // 先手から見て斜め後ろの利き
			case 10:
				makeFront(from, NOT_PROMOTE, moveEndList) // 先手から見て１つ上への利き
				makeBack(from, NOT_PROMOTE, moveEndList)  // 先手から見て１つ後ろへの利き
				makeSide(from, NOT_PROMOTE, moveEndList)  // 先手から見て１つ横への利き
			case 13:
				// 移動元または移動先が敵陣なら成れる
				var promote bool
				switch phase {
				case FIRST:
					promote = Rank(from) < 4 // 移動元の段が敵陣ならOK
				case SECOND:
					promote = Rank(from) > 6 // 移動元の段が敵陣ならOK
				default:
					panic(fmt.Errorf("Unknown phase=%d", phase))
				}
				makeSideLong(pPos, from, promote, moveEndList) // ２つ先のマスからの横への長い利き
				makeSide(from, promote, moveEndList)           // 先手から見て１つ横への利き
			case 14:
				makeDiagonalLongPromote(pPos, phase, from, moveEndList) // ２つ先のマスからの斜めへの長い利き
			case 15:
				var promote = File(from) < 6                // 移動元または移動先が敵陣なら成れる
				makeFrontKnight(from, promote, moveEndList) // 先手桂の利き
			case 16:
				var promote = File(from) > 4               // 移動元または移動先が敵陣なら成れる
				makeBackKnight(from, promote, moveEndList) // 後手桂の利き
			case 17:
				var promote = File(from) < 5          // 移動元または移動先が敵陣なら成れる
				makeFront(from, promote, moveEndList) // 先手から見て１つ上への利き
			case 18:
				var promote = File(from) > 5         // 移動元または移動先が敵陣なら成れる
				makeBack(from, promote, moveEndList) // 先手から見て１つ後ろへの利き
			case 19:
				makeFront(from, moveEndList) // 先手から見て１つ上への利き
				genmv_list = append(genmv_list, 13)
			case 20:
				makeBack(from, NOT_PROMOTE, moveEndList) // 先手から見て１つ後ろへの利き
				genmv_list = append(genmv_list, 13)
			case 21:
				makeFrontDiagonal(from, true, moveEndList) // 先手から見て斜め前の利き
				makeBackDiagonal(from, true, moveEndList)  // 先手から見て斜め後ろの利き
				genmv_list = append(genmv_list, 13)
			case 22:
				makeFrontLong(pPos, from, CAN_PROMOTE, moveEndList) // ２つ先のマスからの上への長い利き
				genmv_list = append(genmv_list, 13)
			case 23:
				genmv_list = append(genmv_list, 13)
			default:
				panic(fmt.Errorf("Unknown step=%d", step))
			}
		}
	}

	return moveEndList
}

// 1 先手から見て斜め前の利き
func makeFrontDiagonal(from Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 9; File(to) != 0 && Rank(to) != 0 { // 左上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
	if to := from - 11; File(to) != 0 && Rank(to) != 0 { // 右上
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 2 先手から見て斜め後ろの利き
func makeBackDiagonal(from Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 11; File(to) != 0 && Rank(to) != 0 { // 左下
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
	if to := from - 9; File(to) != 0 && Rank(to) != 0 { // 右下
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 3 先手から見て１つ前への利き
func makeFront(from Square, moveEndList []MoveEnd) {
	//  promote bool,
	var keepGoing = File(from) != 2

	if to := from - 1; Rank(to) != 0 { // 上
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		// if promote {
		// 	moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		// }
	}
}

// 4 先手から見て１つ後ろへの利き
func makeBack(from Square, moveEndList []MoveEnd) {
	// promote bool,
	// 移動元が８段目のときは必ずならなければならない
	var keepGoing = File(from) != 8

	if to := from + 1; Rank(to) != 0 { // 下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		// if promote {
		// 	moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		// }
	}
}

// 5 先手から見て１つ横への利き
func makeSide(from Square, promote bool, moveEndList []MoveEnd) {
	if to := from + 10; File(to) != 0 { // 左
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
	if to := from - 10; File(to) != 0 { // 右
		ValidateSq(to)
		moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 6 先手から見て１つ上への利き（成りの動き、制約なし）
func makeFrontPromotion(from Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が２段目のときは必ずならなければならない
	var keepGoing = File(from) != 2

	if to := from - 1; Rank(to) != 0 { // 上
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 7
func makeBackPromotion(from Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が２段目のときは必ずならなければならない
	var keepGoing = File(from) != 2

	if to := from + 1; Rank(to) != 0 { // 下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

// 8 ２つ先のマスからの斜めへの長い利き（成り手のみの生成）
func makeDiagonalPromotion(pPos *Position, phase Phase, from Square, moveEndList []MoveEnd) {
	var src_pro bool
	if (phase == FIRST && Rank(from) < 4) || (phase == SECOND && Rank(from) > 6) {
		src_pro = true
	} else {
		src_pro = false
	}

	if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 9 ２つ先のマスからの上への長い利き
func makeFrontLong(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	if Rank(from) > 2 && pPos.IsEmptySq(from-1) { // 1～2段目にある駒でもなく、１つ上が空マスなら
		for to := from - 2; Rank(to) != 0; to -= 1 { // 上
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 10 ２つ先のマスからの下への長い利き
func makeBackLong(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	if Rank(from) < 8 && pPos.IsEmptySq(from+1) { // 8～9段目にある駒でもなく、１つ下が空マスなら
		for to := from + 2; Rank(to) != 0; to += 1 { // 下
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 11 ２つ先のマスからの斜めへの長い利き（成らず）
func makeDiagonalLongNotPromote(pPos *Position, from Square, moveEndList []MoveEnd) {
	if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

}

// 12 ２つ先のマスからの横への長い利き
func makeSideLong(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	// ２つ先のマスからの左への長い利き
	if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
		for to := from + 20; File(to) != 0; to += 10 { // 左
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

	// ２つ先のマスからの右への長い利き
	if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
		for to := from - 20; File(to) != 0; to -= 10 { // 右
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 13
// 14

// 15 ２つ先のマスからの斜めへの長い利き（成り手のみの生成）
func makeDiagonalLongPromote(pPos *Position, phase Phase, from Square, moveEndList []MoveEnd) {
	var src_pro bool
	if (phase == FIRST && Rank(from) < 4) || (phase == SECOND && Rank(from) > 6) {
		src_pro = true
	} else {
		src_pro = false
	}

	if File(from) < 8 && Rank(from) > 2 && pPos.IsEmptySq(from+9) { // 8～9筋にある駒でもなく、1～2段目でもなく、１つ左上が空マスなら
		for to := from + 18; File(to) != 0 && Rank(to) != 0; to += 9 { // ２つ左上から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) > 2 && pPos.IsEmptySq(from-11) { // 1～2筋にある駒でもなく、1～2段目でもなく、１つ右上が空マスなら
		for to := from - 22; File(to) != 0 && Rank(to) != 0; to -= 11 { // ２つ右上から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) < 8 && Rank(from) < 8 && pPos.IsEmptySq(from+11) { // 8～9筋にある駒でもなく、8～9段目でもなく、１つ左下が空マスなら
		for to := from + 22; File(to) != 0 && Rank(to) != 0; to += 11 { // ２つ左下から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
	if File(from) > 2 && Rank(from) < 8 && pPos.IsEmptySq(from-9) { // 1～2筋にある駒でもなく、8～9段目でもなく、１つ右下が空マスなら
		for to := from - 18; File(to) != 0 && Rank(to) != 0; to -= 9 { // ２つ右下から
			ValidateSq(to)
			if src_pro || (phase == FIRST && Rank(to) < 4) || (phase == SECOND && Rank(to) > 6) {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 16 ２つ先のマスからの横への長い利き
func makeSidePromotion(pPos *Position, from Square, promote bool, moveEndList []MoveEnd) {
	// ２つ先のマスからの左への長い利き
	if File(from) < 8 && pPos.IsEmptySq(from+10) { // 8～9筋にある駒でもなく、１つ左が空マスなら
		for to := from + 20; File(to) != 0; to += 10 { // 左
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}

	// ２つ先のマスからの右への長い利き
	if File(from) > 2 && pPos.IsEmptySq(from-10) { // 1～2筋にある駒でもなく、１つ右が空マスなら
		for to := from - 20; File(to) != 0; to -= 10 { // 右
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
			if !pPos.IsEmptySq(to) {
				break
			}
		}
	}
}

// 17 先手桂の利き
func makeFrontKnight(from Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が３段目のときは必ずならなければならない
	var keepGoing = File(from) != 3

	if 2 < Rank(from) && Rank(from) < 10 {
		if 0 < File(from) && File(from) < 9 { // 左上桂馬飛び
			to := from + 8
			ValidateSq(to)
			if keepGoing {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			}
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
		}
		if 1 < File(from) && File(from) < 10 { // 右上桂馬飛び
			to := from - 12
			ValidateSq(to)
			if keepGoing {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
			}
			if promote {
				moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
			}
		}
	}
}

// 18 後手桂の利き
func makeBackKnight(from Square, promote bool, moveEndList []MoveEnd) {
	// 移動元が７段目のときは必ずならなければならない
	var keepGoing = File(from) != 7

	if to := from + 12; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 左下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
	if to := from - 8; File(to) != 0 && Rank(to) != 0 && Rank(to) != 9 { // 右下
		ValidateSq(to)
		if keepGoing {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
		if promote {
			moveEndList = append(moveEndList, NewMoveEndValue2(to, true))
		}
	}
}

func makeDrop(pPos *Position, droppableFiles [10]bool, rank Square, moveEndList []MoveEnd) {
	for file := Square(9); file > 0; file -= 1 {
		if droppableFiles[file] {
			to := SquareFrom(file, rank)
			ValidateSq(to)
			moveEndList = append(moveEndList, NewMoveEndValue2(to, false))
		}
	}
}

// NifuFirst - 先手で二歩になるか筋調べ
func NifuFirst(pPos *Position, file Square) bool {
	for rank := Square(2); rank < 10; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == PIECE_P1 {
			return true
		}
	}

	return false
}

// NifuSecond - 後手で二歩になるか筋調べ
func NifuSecond(pPos *Position, file Square) bool {
	for rank := Square(1); rank < 9; rank += 1 {
		if pPos.Board[SquareFrom(file, rank)] == PIECE_P2 {
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
	var friendKingSq Square
	var hand_start int
	var hand_end int
	// var opponentKingSq Square
	var pOpponentSumCB *ControlBoard
	if friend == FIRST {
		friendKingSq = pPos.GetPieceLocation(PCLOC_K1)
		hand_start = HAND_IDX_START
		pOpponentSumCB = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2]
	} else if friend == SECOND {
		friendKingSq = pPos.GetPieceLocation(PCLOC_K2)
		hand_start = HAND_IDX_START + HAND_TYPE_SIZE
		pOpponentSumCB = pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1]
	} else {
		panic(fmt.Errorf("Unknown phase=%d", friend))
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
					control_list := GenControl(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							// 敵の長い駒の利きは、玉が逃げても伸びてくる方向があるので、
							// いったん玉を動かしてから 再チェックするぜ（＾～＾）
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMoveValue2(from, to)
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
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move := NewMoveValue2(from, to)
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
				hand_sq := Square(hand_index) + SQ_HAND_START
				control_list := GenControl(pPos, hand_sq)

				for _, moveEnd := range control_list {
					to := moveEnd.GetDestination()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move := NewMoveValue2(hand_sq, to)
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
				from := Square(file*10 + rank)
				if pPos.Homo(from, friendKingSq) { // 自玉と同じプレイヤーの駒を動かします
					control_list := GenControl(pPos, from)

					piece := pPos.Board[from]
					pieceType := What(piece)

					if pieceType == PIECE_TYPE_K {
						// 玉は自殺手を省きます
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							if pPos.Hetero(from, to) && pOpponentSumCB.Board1[to] == 0 { // 自駒の上、敵の利きには移動できません
								move_list = append(move_list, NewMoveValue2(from, to))
							}
						}
					} else {
						for _, moveEnd := range control_list {
							to := moveEnd.GetDestination()
							if pPos.Hetero(from, to) { // 自駒の上には移動できません
								move_list = append(move_list, NewMoveValue2(from, to))
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
				control_list := GenControl(pPos, hand_sq)

				for _, moveEnd := range control_list {
					to := moveEnd.GetDestination()
					if pPos.IsEmptySq(to) { // 駒の上には打てません
						move_list = append(move_list, NewMoveValue2(hand_sq, to))
					}
				}
			}
		}
	}

	return move_list
}
