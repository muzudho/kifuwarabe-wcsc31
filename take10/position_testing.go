// 利きのテスト
package take10

import (
	"fmt"
	"math"
	"math/rand"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// TestControl
func TestControl(pPos *Position) (bool, string) {
	pPos.ClearControlLayer(CONTROL_LAYER_TEST_COPY)
	pPos.ClearControlLayer(CONTROL_LAYER_TEST_ERROR)

	// 利きをコピー
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			pPos.ControlBoards[phase][CONTROL_LAYER_TEST_COPY][sq] = pPos.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
		}
	}

	// 指し手生成
	// 探索中に削除される指し手も入ってるかも
	move_list := GenMoveList(pPos)
	move_total := len(move_list)

	for move_seq, move := range move_list {
		// その手を指してみるぜ（＾～＾）
		pPos.DoMove(move)

		// すぐ戻すぜ（＾～＾）
		pPos.UndoMove()

		// 元に戻っていればOK（＾～＾）
		is_error := checkControl(pPos, move_seq, move_total, move)
		if is_error {
			return is_error, fmt.Sprintf("Error! move_seq=(%d/%d) move=%s", move_seq, move_total, move.ToCodeOfM())
		}
	}

	return false, ""
}

// Check - 元に戻っていればOK（＾～＾）
func checkControl(pPos *Position, move_seq int, move_total int, move Move) bool {

	is_error := false

	// 誤差調べ
	for phase := 0; phase < 2; phase += 1 {
		for sq := 0; sq < BOARD_SIZE; sq += 1 {
			diff := pPos.ControlBoards[phase][CONTROL_LAYER_TEST_COPY][sq] - pPos.ControlBoards[phase][CONTROL_LAYER_SUM][sq]
			pPos.ControlBoards[phase][CONTROL_LAYER_TEST_ERROR][sq] = diff
			if diff != 0 {
				is_error = true
			}
		}
	}

	return is_error
}

// SumAbsControl - 利きテーブルの各マスを絶対値にし、その総和を返します
func SumAbsControl(pPos *Position, layer1 int) [2]int {

	sumList := [2]int{0, 0}

	for phase := 0; phase < 2; phase += 1 {
		for from := l03.Square(11); from < BOARD_SIZE; from += 1 {
			if l03.File(from) != 0 && l03.Rank(from) != 0 {

				sumList[phase] += int(math.Abs(float64(pPos.ControlBoards[phase][layer1][from])))

			}
		}
	}

	return sumList
}

// ShuffleBoard - 盤上の駒、持ち駒をシャッフルします
// ゲーム中にはできない動きをするので、利きの計算は無視します。
// 最後に利きは再計算します
func ShuffleBoard(pPos *Position) {

	// 駒の数を数えます
	countList1 := CountAllPieces(pPos)

	// 盤と駒台との移動
	// 適当な回数
	for i := 0; i < 200; i += 1 {

		// 盤から駒台の方向
		for rank := l03.Square(1); rank < 10; rank += 1 {
			for file := l03.Square(9); file > 0; file -= 1 {
				sq := SquareFrom(file, rank)

				// 10マスに1マスは駒台へ
				change := l03.Square(rand.Intn(10))
				if change == 0 {
					piece := pPos.Board[sq]
					if piece != l03.PIECE_EMPTY {
						phase := Who(piece)
						pieceType := What(piece)

						ok := false
						switch phase {
						case l06.FIRST:
							switch pieceType {
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPos.Hands[l03.HANDSQ_R1] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPos.Hands[l03.HANDSQ_B1] += 1
								ok = true
							case PIECE_TYPE_G:
								pPos.Hands[l03.HANDSQ_G1] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPos.Hands[l03.HANDSQ_S1] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPos.Hands[l03.HANDSQ_N1] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPos.Hands[l03.HANDSQ_L1] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPos.Hands[l03.HANDSQ_P1] += 1
								ok = true
							default:
								// Ignored
							}
						case l06.SECOND:
							switch pieceType {
							case PIECE_TYPE_R, PIECE_TYPE_PR:
								pPos.Hands[l03.HANDSQ_R2] += 1
								ok = true
							case PIECE_TYPE_B, PIECE_TYPE_PB:
								pPos.Hands[l03.HANDSQ_B2] += 1
								ok = true
							case PIECE_TYPE_G:
								pPos.Hands[l03.HANDSQ_G2] += 1
								ok = true
							case PIECE_TYPE_S, PIECE_TYPE_PS:
								pPos.Hands[l03.HANDSQ_S2] += 1
								ok = true
							case PIECE_TYPE_N, PIECE_TYPE_PN:
								pPos.Hands[l03.HANDSQ_N2] += 1
								ok = true
							case PIECE_TYPE_L, PIECE_TYPE_PL:
								pPos.Hands[l03.HANDSQ_L2] += 1
								ok = true
							case PIECE_TYPE_P, PIECE_TYPE_PP:
								pPos.Hands[l03.HANDSQ_P2] += 1
								ok = true
							default:
								// Ignored
							}
						default:
							panic(fmt.Errorf("uknown phase=%d", phase))
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
			panic(fmt.Errorf("shuffle: (1) countError=%d", countError))
		}

		// 駒台から盤の方向
		for hand_index := HAND_IDX_START; hand_index < HAND_IDX_END; hand_index += 1 {
			num := pPos.Hands[hand_index]
			if num > 0 {
				sq := l03.Square(rand.Intn(100))
				// うまく空マスなら移動成功
				if OnBoard(sq) && pPos.IsEmptySq(sq) {
					pPos.Board[sq] = HandPieceArray[hand_index]
					pPos.Hands[hand_index] -= 1
				}
			}
		}

		// 駒の数を数えます
		countList2 = CountAllPieces(pPos)
		countError = CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(fmt.Errorf("shuffle: (2) countError=%d", countError))
		}
	}

	// 盤上での移動
	// 適当に大きな回数
	for i := 0; i < 81*80; i += 1 {
		sq1 := l03.Square(rand.Intn(100))
		sq2 := l03.Square(rand.Intn(100))
		if OnBoard(sq1) && OnBoard(sq2) && !pPos.IsEmptySq(sq1) {
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
			switch What(piece) {
			case PIECE_TYPE_K, PIECE_TYPE_EMPTY:
				// Ignored
			default:
				phase := Who(piece)
				pieceType := What(piece)

				change := l03.Square(rand.Intn(10))
				if change == 0 {
					phase = FlipPhase(phase)
				}

				pPos.Board[sq2] = PieceFromPhPt(phase, pieceType)
			}
		}

		// 駒の数を数えます
		countList2 := CountAllPieces(pPos)
		countError := CountErrorCountLists(countList1, countList2)
		if countError != 0 {
			panic(fmt.Errorf("shuffle: (3) countError=%d", countError))
		}
	}

	// 手番のシャッフル
	switch rand.Intn(2) {
	case 0:
		pPos.phase = l06.FIRST
	default:
		pPos.phase = l06.SECOND
	}

	// 手目は 1 に戻します
	pPos.StartMovesNum = 1
	pPos.OffsetMovesIndex = 0

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	App.Out.Debug(Sprint(pPos))

	if false {
		var countList [8]int

		if true {
			countList = [8]int{}

			// 盤上
			for rank := l03.Square(1); rank < 10; rank += 1 {
				for file := l03.Square(9); file > 0; file -= 1 {
					sq := SquareFrom(file, rank)

					fmt.Printf("%s,", pPos.Board[sq].ToCodeOfPc())

					piece := What(pPos.Board[sq])
					switch piece {
					case PIECE_TYPE_K:
						countList[0] += 1
					case PIECE_TYPE_R, PIECE_TYPE_PR:
						countList[1] += 1
					case PIECE_TYPE_B, PIECE_TYPE_PB:
						countList[2] += 1
					case PIECE_TYPE_G:
						countList[3] += 1
					case PIECE_TYPE_S, PIECE_TYPE_PS:
						countList[4] += 1
					case PIECE_TYPE_N, PIECE_TYPE_PN:
						countList[5] += 1
					case PIECE_TYPE_L, PIECE_TYPE_PL:
						countList[6] += 1
					case PIECE_TYPE_P, PIECE_TYPE_PP:
						countList[7] += 1
					default:
						// Ignore
					}
				}
				fmt.Printf("\n")
			}

			// 駒台
			countList[1] += pPos.Hands[0] + pPos.Hands[7]
			countList[2] += pPos.Hands[1] + pPos.Hands[8]
			countList[3] += pPos.Hands[2] + pPos.Hands[9]
			countList[4] += pPos.Hands[3] + pPos.Hands[10]
			countList[5] += pPos.Hands[4] + pPos.Hands[11]
			countList[6] += pPos.Hands[5] + pPos.Hands[12]
			countList[7] += pPos.Hands[6] + pPos.Hands[13]
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
	command := pPos.SprintSfen()
	App.Out.Debug("#command=%s", command)

	// 利きの再計算もやってくれる
	pPos.ReadPosition(command)

	// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
	App.Out.Debug(Sprint(pPos))
	ShowAllPiecesCount(pPos)
	command2 := pPos.SprintSfen()
	App.Out.Debug("#command2=%s", command2)

	// 駒の数を数えます
	countList2 := CountAllPieces(pPos)
	countError := CountErrorCountLists(countList1, countList2)
	if countError != 0 {
		panic(fmt.Errorf("shuffle: (4) countError=%d", countError))
	}
}

// CountAllPieces - 駒の数を確認するぜ（＾～＾）
func CountAllPieces(pPos *Position) [8]int {

	countList := [8]int{}

	// 盤上
	for rank := l03.Square(1); rank < 10; rank += 1 {
		for file := l03.Square(9); file > 0; file -= 1 {
			sq := SquareFrom(file, rank)

			piece := What(pPos.Board[sq])
			switch piece {
			case PIECE_TYPE_K:
				countList[0] += 1
			case PIECE_TYPE_R, PIECE_TYPE_PR:
				countList[1] += 1
			case PIECE_TYPE_B, PIECE_TYPE_PB:
				countList[2] += 1
			case PIECE_TYPE_G:
				countList[3] += 1
			case PIECE_TYPE_S, PIECE_TYPE_PS:
				countList[4] += 1
			case PIECE_TYPE_N, PIECE_TYPE_PN:
				countList[5] += 1
			case PIECE_TYPE_L, PIECE_TYPE_PL:
				countList[6] += 1
			case PIECE_TYPE_P, PIECE_TYPE_PP:
				countList[7] += 1
			default:
				// Ignore
			}
		}
	}

	// 駒台
	countList[1] += pPos.Hands[0] + pPos.Hands[7]
	countList[2] += pPos.Hands[1] + pPos.Hands[8]
	countList[3] += pPos.Hands[2] + pPos.Hands[9]
	countList[4] += pPos.Hands[3] + pPos.Hands[10]
	countList[5] += pPos.Hands[4] + pPos.Hands[11]
	countList[6] += pPos.Hands[5] + pPos.Hands[12]
	countList[7] += pPos.Hands[6] + pPos.Hands[13]

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
