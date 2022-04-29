// 利きのテスト
package take16

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l10 "github.com/muzudho/kifuwarabe-wcsc31/take10"
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l13 "github.com/muzudho/kifuwarabe-wcsc31/take13"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// TestControl
func TestControl(pNerve *Nerve, pPos *l15.Position) (bool, string) {
	pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_COPY1].Clear()
	pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_COPY2].Clear()

	pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_ERROR1].Clear()
	pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_ERROR2].Clear()

	// 利きをコピー
	copyCb1 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_COPY1]
	sumCb1 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1]
	copyCb2 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_COPY2]
	sumCb2 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2]
	for sq := 0; sq < l03.BOARD_SIZE; sq += 1 {
		copyCb1.Board1[sq] = sumCb1.Board1[sq]
		copyCb2.Board1[sq] = sumCb2.Board1[sq]
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pNerve, pPos)
	move_total := len(move_list)

	for move_seq, move := range move_list {
		// その手を指してみるぜ（＾～＾）
		pNerve.DoMove(pPos, move)

		// すぐ戻すぜ（＾～＾）
		pNerve.UndoMove(pPos)

		// 元に戻っていればOK（＾～＾）
		is_error := checkControl(pNerve, move_seq, move_total, move)
		if is_error {
			return is_error, fmt.Sprintf("Error! move_seq=(%d/%d) move=%s", move_seq, move_total, move.ToCodeOfM())
		}
	}

	return false, ""
}

// Check - 元に戻っていればOK（＾～＾）
func checkControl(pNerve *Nerve, move_seq int, move_total int, move l13.Move) bool {

	is_error := false

	// 誤差調べ
	copyCB1 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_COPY1]
	sumCB1 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1]
	errorCB1 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_ERROR1]
	copyCB2 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_COPY2]
	sumCB2 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2]
	errorCB2 := pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_TEST_ERROR2]
	for sq := 0; sq < l03.BOARD_SIZE; sq += 1 {
		diff1 := copyCB1.Board1[sq] - sumCB1.Board1[sq]
		errorCB1.Board1[sq] = diff1
		if diff1 != 0 {
			is_error = true
			break
		}

		diff2 := copyCB2.Board1[sq] - sumCB2.Board1[sq]
		errorCB2.Board1[sq] = diff2
		if diff2 != 0 {
			is_error = true
			break
		}
	}

	return is_error
}

// SumAbsControl - 利きテーブルの各マスを絶対値にし、その総和を返します
func SumAbsControl(pNerve *Nerve, ph1_c ControlLayerT, ph2_c ControlLayerT) [2]int {

	sumList := [2]int{0, 0}

	cb1 := pNerve.PCtrlBrdSys.PBoards[ph1_c]
	for from := l03.Square(11); from < l03.BOARD_SIZE; from += 1 {
		if l03.File(from) != 0 && l03.Rank(from) != 0 {

			sumList[l06.FIRST-1] += int(math.Abs(float64(cb1.Board1[from])))

		}
	}

	cb2 := pNerve.PCtrlBrdSys.PBoards[ph2_c]
	for from := l03.Square(11); from < l03.BOARD_SIZE; from += 1 {
		if l03.File(from) != 0 && l03.Rank(from) != 0 {

			sumList[l06.SECOND-1] += int(math.Abs(float64(cb2.Board1[from])))

		}
	}

	return sumList
}

// ShuffleBoard - 盤上の駒、持ち駒をシャッフルします
// ゲーム中にはできない動きをするので、利きの計算は無視します。
// 最後に利きは再計算します
func ShuffleBoard(pNerve *Nerve, pPos *l15.Position) {

	// 駒の数を数えます
	countList1 := CountAllPieces(pPos)

	// 盤と駒台との移動
	// 適当な回数
	for i := 0; i < 200; i += 1 {

		// 盤から駒台の方向
		for rank := l03.Square(1); rank < 10; rank += 1 {
			for file := l03.Square(9); file > 0; file -= 1 {
				sq := l15.SquareFrom(file, rank)

				// 10マスに1マスは駒台へ
				change := l03.Square(rand.Intn(10))
				if change == 0 {
					piece := pPos.Board[sq]
					if piece != l03.PIECE_EMPTY {
						phase := Who(piece)
						pieceType := l11.What(piece)

						ok := false
						switch phase {
						case l06.FIRST:
							switch pieceType {
							case l11.PIECE_TYPE_K:
								pPos.Hands1[l03.HAND_K1] += 1
								ok = true
							case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
								pPos.Hands1[l03.HAND_R1] += 1
								ok = true
							case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
								pPos.Hands1[l03.HAND_B1] += 1
								ok = true
							case l11.PIECE_TYPE_G:
								pPos.Hands1[l03.HAND_G1] += 1
								ok = true
							case l11.PIECE_TYPE_S, l11.PIECE_TYPE_PS:
								pPos.Hands1[l03.HAND_S1] += 1
								ok = true
							case l11.PIECE_TYPE_N, l11.PIECE_TYPE_PN:
								pPos.Hands1[l03.HAND_N1] += 1
								ok = true
							case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL:
								pPos.Hands1[l03.HAND_L1] += 1
								ok = true
							case l11.PIECE_TYPE_P, l11.PIECE_TYPE_PP:
								pPos.Hands1[l03.HAND_P1] += 1
								ok = true
							default:
								// Ignored
							}
						case l06.SECOND:
							switch pieceType {
							case l11.PIECE_TYPE_K:
								pPos.Hands1[l03.HAND_K2] += 1
								ok = true
							case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
								pPos.Hands1[l03.HAND_R2] += 1
								ok = true
							case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
								pPos.Hands1[l03.HAND_B2] += 1
								ok = true
							case l11.PIECE_TYPE_G:
								pPos.Hands1[l03.HAND_G2] += 1
								ok = true
							case l11.PIECE_TYPE_S, l11.PIECE_TYPE_PS:
								pPos.Hands1[l03.HAND_S2] += 1
								ok = true
							case l11.PIECE_TYPE_N, l11.PIECE_TYPE_PN:
								pPos.Hands1[l03.HAND_N2] += 1
								ok = true
							case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL:
								pPos.Hands1[l03.HAND_L2] += 1
								ok = true
							case l11.PIECE_TYPE_P, l11.PIECE_TYPE_PP:
								pPos.Hands1[l03.HAND_P2] += 1
								ok = true
							default:
								// Ignored
							}
						default:
							panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
						}

						if ok {
							pPos.Board[sq] = l03.PIECE_EMPTY
						}
					}

				}
			}
		}

		// 駒の数を数えます
		countList2 := CountAllPieces(pPos)
		countError := CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(App.LogNotEcho.Fatal("shuffle: (1) countError=%d", countError))
		}

		// 駒台から盤の方向
		for hand_index := l03.HAND_IDX_START; hand_index < l03.HAND_IDX_END; hand_index += 1 {
			num := pPos.Hands1[hand_index]
			if num > 0 {
				sq := l03.Square(rand.Intn(100))
				// うまく空マスなら移動成功
				if l15.OnBoard(sq) && pPos.IsEmptySq(sq) {
					pPos.Board[sq] = l10.HandPieceArray[hand_index]
					pPos.Hands1[hand_index] -= 1
				}
			}
		}

		// 駒の数を数えます
		countList2 = CountAllPieces(pPos)
		countError = CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(App.LogNotEcho.Fatal("shuffle: (2) countError=%d", countError))
		}
	}

	// 盤上での移動
	// 適当に大きな回数
	for i := 0; i < 81*80; i += 1 {
		sq1 := l03.Square(rand.Intn(100))
		sq2 := l03.Square(rand.Intn(100))
		if l15.OnBoard(sq1) && l15.OnBoard(sq2) && !pPos.IsEmptySq(sq1) {
			piece := pPos.Board[sq1]
			// 位置スワップ
			pPos.Board[sq1] = pPos.Board[sq2]
			pPos.Board[sq2] = piece

			// 成／不成 変更
			promote := l03.Square(rand.Intn(10))
			if promote == 0 {
				pPos.Board[sq2] = l09.Promote(pPos.Board[sq2])
			} else if promote == 1 {
				pPos.Board[sq2] = l09.Demote(pPos.Board[sq2])
			}

			// 駒の先後変更（玉除く）
			piece = pPos.Board[sq2]
			switch l11.What(piece) {
			case l11.PIECE_TYPE_K, l11.PIECE_TYPE_EMPTY:
				// Ignored
			default:
				phase := Who(piece)
				pieceType := l11.What(piece)

				change := l03.Square(rand.Intn(10))
				if change == 0 {
					phase = FlipPhase(phase)
				}

				pPos.Board[sq2] = l11.PieceFromPhPt(phase, pieceType)
			}
		}

		// 駒の数を数えます
		countList2 := CountAllPieces(pPos)
		countError := CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(App.LogNotEcho.Fatal("shuffle: (3) countError=%d", countError))
		}
	}

	// 手番のシャッフル
	switch rand.Intn(2) {
	case 0:
		pNerve.PPosSys.phase = l06.FIRST
	default:
		pNerve.PPosSys.phase = l06.SECOND
	}

	// 手目は 1 に戻します
	pNerve.PRecord.StartMovesNum = 1
	pNerve.PRecord.OffsetMovesIndex = 0

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	App.Out.Debug(pPos.SprintBoardHeader(
		pNerve.PPosSys.phase,
		pNerve.PRecord.StartMovesNum,
		pNerve.PRecord.OffsetMovesIndex))
	App.Out.Debug(pPos.SprintBoard())
	App.Out.Debug(pNerve.SprintBoardFooter())

	if false {
		var countList [8]int

		if true {
			countList = [8]int{}

			// 盤上
			for rank := l03.Square(1); rank < 10; rank += 1 {
				for file := l03.Square(9); file > 0; file -= 1 {
					sq := l15.SquareFrom(file, rank)

					fmt.Printf("%s,", pPos.Board[sq].ToCodeOfPc())

					piece := l11.What(pPos.Board[sq])
					switch piece {
					case l11.PIECE_TYPE_K:
						countList[0] += 1
					case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
						countList[1] += 1
					case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
						countList[2] += 1
					case l11.PIECE_TYPE_G:
						countList[3] += 1
					case l11.PIECE_TYPE_S, l11.PIECE_TYPE_PS:
						countList[4] += 1
					case l11.PIECE_TYPE_N, l11.PIECE_TYPE_PN:
						countList[5] += 1
					case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL:
						countList[6] += 1
					case l11.PIECE_TYPE_P, l11.PIECE_TYPE_PP:
						countList[7] += 1
					default:
						// Ignore
					}
				}
				fmt.Printf("\n")
			}

			// 駒台
			countList[0] += pPos.Hands1[l03.HAND_K1] + pPos.Hands1[l03.HAND_K2]
			countList[1] += pPos.Hands1[l03.HAND_R1] + pPos.Hands1[l03.HAND_R2]
			countList[2] += pPos.Hands1[l03.HAND_B1] + pPos.Hands1[l03.HAND_B2]
			countList[3] += pPos.Hands1[l03.HAND_G1] + pPos.Hands1[l03.HAND_G2]
			countList[4] += pPos.Hands1[l03.HAND_S1] + pPos.Hands1[l03.HAND_S2]
			countList[5] += pPos.Hands1[l03.HAND_N1] + pPos.Hands1[l03.HAND_N2]
			countList[6] += pPos.Hands1[l03.HAND_L1] + pPos.Hands1[l03.HAND_L2]
			countList[7] += pPos.Hands1[l03.HAND_P1] + pPos.Hands1[l03.HAND_P2]
		} else {
			countList = CountAllPieces(pPos)
		}

		App.Out.Debug("#Count\n")
		App.Out.Debug("#-----\n")
		App.Out.Debug("#King  :%3d\n", countList[0])
		App.Out.Debug("#Rook  :%3d\n", countList[1])
		App.Out.Debug("#Bishop:%3d\n", countList[2])
		App.Out.Debug("#Gold  :%3d\n", countList[3])
		App.Out.Debug("#Silver:%3d\n", countList[4])
		App.Out.Debug("#Knight:%3d\n", countList[5])
		App.Out.Debug("#Lance :%3d\n", countList[6])
		App.Out.Debug("#Pawn  :%3d\n", countList[7])
		App.Out.Debug("#----------\n")
		App.Out.Debug("#Total :%3d\n", countList[0]+countList[1]+countList[2]+countList[3]+countList[4]+countList[5]+countList[6]+countList[7])
	} else {
		ShowAllPiecesCount(pPos)
	}

	// position sfen 文字列を取得
	command := sprintSfenResignation(pNerve.PPosSys, pPos, pNerve.PRecord)
	App.Out.Debug("#command=%s", command)

	// 利きの再計算もやってくれる
	pNerve.ReadPosition(pPos, command)

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	App.Out.Debug(pPos.SprintBoardHeader(
		pNerve.PPosSys.phase,
		pNerve.PRecord.StartMovesNum,
		pNerve.PRecord.OffsetMovesIndex))
	App.Out.Debug(pPos.SprintBoard())
	App.Out.Debug(pNerve.SprintBoardFooter())
	ShowAllPiecesCount(pPos)
	command2 := sprintSfenResignation(pNerve.PPosSys, pPos, pNerve.PRecord)
	App.Out.Debug("#command2=%s", command2)

	// 駒の数を数えます
	countList2 := CountAllPieces(pPos)
	countError := CountErrorCountLists(countList1, countList2)
	if countError != 0 {
		panic(App.LogNotEcho.Fatal("shuffle: (4) countError=%d", countError))
	}
}

// CountAllPieces - 駒の数を確認するぜ（＾～＾）
func CountAllPieces(pPos *l15.Position) [8]int {

	countList := [8]int{}

	// 盤上
	for rank := l03.Square(1); rank < 10; rank += 1 {
		for file := l03.Square(9); file > 0; file -= 1 {
			sq := l15.SquareFrom(file, rank)

			piece := l11.What(pPos.Board[sq])
			switch piece {
			case l11.PIECE_TYPE_K:
				countList[0] += 1
			case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
				countList[1] += 1
			case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
				countList[2] += 1
			case l11.PIECE_TYPE_G:
				countList[3] += 1
			case l11.PIECE_TYPE_S, l11.PIECE_TYPE_PS:
				countList[4] += 1
			case l11.PIECE_TYPE_N, l11.PIECE_TYPE_PN:
				countList[5] += 1
			case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL:
				countList[6] += 1
			case l11.PIECE_TYPE_P, l11.PIECE_TYPE_PP:
				countList[7] += 1
			default:
				// Ignore
			}
		}
	}

	// 駒台
	countList[0] += pPos.Hands1[l03.HAND_K1] + pPos.Hands1[l03.HAND_K2]
	countList[1] += pPos.Hands1[l03.HAND_R1] + pPos.Hands1[l03.HAND_R2]
	countList[2] += pPos.Hands1[l03.HAND_B1] + pPos.Hands1[l03.HAND_B2]
	countList[3] += pPos.Hands1[l03.HAND_G1] + pPos.Hands1[l03.HAND_G2]
	countList[4] += pPos.Hands1[l03.HAND_S1] + pPos.Hands1[l03.HAND_S2]
	countList[5] += pPos.Hands1[l03.HAND_N1] + pPos.Hands1[l03.HAND_N2]
	countList[6] += pPos.Hands1[l03.HAND_L1] + pPos.Hands1[l03.HAND_L2]
	countList[7] += pPos.Hands1[l03.HAND_P1] + pPos.Hands1[l03.HAND_P2]

	return countList
}

// CountErrorCountLists - 数えた駒の枚数を比較します
func CountErrorCountLists(countList1 [8]int, countList2 [8]int) int {
	sum := 0
	for i := 0; i < 8; i += 1 {
		sum += int(math.Abs(float64(countList1[i] - countList2[i])))
	}
	return sum
}

// copyBoard - 盤[b0] を 盤[b1] にコピーします
func copyBoard(pPos0 *l15.Position, pPos1 *l15.Position) {
	for sq := 0; sq < 100; sq += 1 {
		pPos1.Board[sq] = pPos0.Board[sq]
	}

	pPos1.Hands1 = pPos0.Hands1
	for i := l11.PCLOC_START; i < l11.PCLOC_END; i += 1 {
		pPos1.PieceLocations[i] = pPos0.PieceLocations[i]
	}
}

// copyBoard - 盤[0] を 盤[1] で異なるマスを 盤[2] 盤[3] にセットします
func diffBoard(pPos0 *l15.Position, pPos1 *l15.Position, pPos2 *l15.Position, pPos3 *l15.Position) {
	// 盤上
	for sq := 0; sq < 100; sq += 1 {
		if pPos1.Board[sq] == pPos0.Board[sq] {
			// 等しければ空マス
			pPos2.Board[sq] = l03.PIECE_EMPTY
			pPos3.Board[sq] = l03.PIECE_EMPTY

		} else {
			// 異なったら
			pPos2.Board[sq] = pPos0.Board[sq]
			pPos3.Board[sq] = pPos1.Board[sq]
		}
	}

	// 駒台
	for i := l03.HAND_IDX_START; i < l03.HAND_IDX_END; i += 1 {
		if pPos0.Hands1[i] == pPos1.Hands1[i] {
			// 等しければゼロ
			pPos2.Hands1[i] = 0
			pPos3.Hands1[i] = 0
		} else {
			// 異なればその数
			pPos2.Hands1[i] = pPos0.Hands1[i]
			pPos3.Hands1[i] = pPos1.Hands1[i]
		}
	}

	// 位置
	for i := l11.PCLOC_START; i < l11.PCLOC_END; i += 1 {
		if pPos0.PieceLocations[i] == pPos1.PieceLocations[i] {
			// 等しければゼロ
			pPos2.PieceLocations[i] = 0
			pPos3.PieceLocations[i] = 0
		} else {
			// 異なればその数
			pPos2.PieceLocations[i] = pPos0.PieceLocations[i]
			pPos3.PieceLocations[i] = pPos1.PieceLocations[i]
		}
	}
}

// ２つのボードの違いを数えるぜ（＾～＾）
func errorBoard(pPos0 *l15.Position, pPos1 *l15.Position, pPos2 *l15.Position, pPos3 *l15.Position) int {
	diffBoard(pPos0, pPos1, pPos2, pPos3)

	errorNum := 0

	// 盤上
	for sq := 0; sq < 100; sq += 1 {
		if pPos2.Board[sq] != pPos3.Board[sq] {
			errorNum += 1
		}
	}

	// 駒台
	for i := l03.HAND_IDX_START; i < l03.HAND_IDX_END; i += 1 {
		if pPos2.Hands1[i] != pPos3.Hands1[i] {
			errorNum += 1
		}
	}

	// 位置
	if pPos2.PieceLocations[l11.PCLOC_K1] != pPos3.PieceLocations[l11.PCLOC_K1] {
		errorNum += 1
	}
	if pPos2.PieceLocations[l11.PCLOC_K2] != pPos3.PieceLocations[l11.PCLOC_K2] {
		errorNum += 1
	}

	// 位置（不安定注意）
	rook2 := []int{}
	rook3 := []int{}
	for i := l11.PCLOC_R1; i < l11.PCLOC_R2+1; i += 1 {
		rook2 = append(rook2, int(pPos2.PieceLocations[i]))
		rook3 = append(rook3, int(pPos2.PieceLocations[i]))
	}
	sort.Ints(rook2)
	sort.Ints(rook3)
	for i := 0; i < len(rook2); i += 1 {
		if rook2[i] != rook3[i] {
			errorNum += 1
		}
	}

	// 位置（不安定注意）
	bishop2 := []int{}
	bishop3 := []int{}
	for i := l11.PCLOC_B1; i < l11.PCLOC_B2+1; i += 1 {
		bishop2 = append(bishop2, int(pPos2.PieceLocations[i]))
		bishop3 = append(bishop3, int(pPos2.PieceLocations[i]))
	}
	sort.Ints(bishop2)
	sort.Ints(bishop3)
	for i := 0; i < len(bishop2); i += 1 {
		if bishop2[i] != bishop3[i] {
			errorNum += 1
		}
	}

	// 位置（不安定注意）
	lance2 := []int{}
	lance3 := []int{}
	for i := l11.PCLOC_L1; i < l11.PCLOC_L4+1; i += 1 {
		lance2 = append(lance2, int(pPos2.PieceLocations[i]))
		lance3 = append(lance3, int(pPos2.PieceLocations[i]))
	}
	sort.Ints(lance2)
	sort.Ints(lance3)
	for i := 0; i < len(lance2); i += 1 {
		if lance2[i] != lance3[i] {
			errorNum += 1
		}
	}

	// 駒割り評価値
	if pPos2.MaterialValue != pPos3.MaterialValue {
		errorNum += 1
	}

	return errorNum
}
