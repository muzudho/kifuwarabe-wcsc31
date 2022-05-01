package take17

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
	l07 "github.com/muzudho/kifuwarabe-wcsc31/take7"
)

// IsBadForm - 悪形なら真
func IsBadForm(pPos *l15.Position, pNerve *Nerve, move l03.Move) bool {
	from, to, promotion := move.Destructure()
	if 116 <= from {
		panic(App.Log.Fatal(fmt.Sprintf("is bad form 1: abnormal from=%d\n", from)))
	}

	// 自分の先後は？
	var turn = pNerve.PPosSys.GetPhase()

	// 打のケースがあることに注意
	if 11 <= from && from < l03.HANDSQ_BEGIN.ToSq() {
		// 動く駒は？
		var movedPiece = pPos.GetPieceOnBoardAtSq(from)
		var movedPieceType = l03.What(movedPiece)
		// if App.IsDebug {
		// 	App.Out.Print("# movePiece=%s\n", movedPiece.ToCodeOfPc())
		// }

		var isBadForm = false
		switch movedPieceType { // 動かした駒が
		case l03.PIECE_TYPE_K: // 玉
			isBadForm = isBadFormOfKing(pPos, turn, from, to)
		case l03.PIECE_TYPE_R: // 飛
			isBadForm = isBadFormOfRook(pPos, turn, from, to, promotion)
		case l03.PIECE_TYPE_B: // 角
			isBadForm = isBadFormOfBishop(pPos, turn, from, to, promotion)
		case l03.PIECE_TYPE_G: // 金
			isBadForm = isBadFormOfGold(pPos, turn, from, to)
		case l03.PIECE_TYPE_S: // 銀
			isBadForm = isBadFormOfSilver(pPos, turn, from, to)
		case l03.PIECE_TYPE_L: // 香
			isBadForm = isBadFormOfLance(pPos, turn, to, promotion)
		}

		if isBadForm {
			return true
		}

	} else {

		var fromHandSq = l03.FromSqToHandSq(from)
		if fromHandSq < l03.HANDSQ_BEGIN || l03.HANDSQ_END <= fromHandSq {
			panic(App.Log.Fatal(fmt.Sprintf("is bad form: abnormal from hand sq=%d, from=%d\n", fromHandSq, from)))
		}

		// 打
		if App.IsDebug {
			App.Log.Debug("打判定1\n")
		}
		var droppedPieceType = l03.WhatHandSq(fromHandSq)

		var isBadForm = false
		switch droppedPieceType { // 打った駒が
		case l03.PIECE_TYPE_P: // 歩
			if App.IsDebug {
				App.Log.Debug("歩打判定1\n")
			}
			isBadForm = isBadFormOfDroppedPawn(pPos, turn, to)
		}

		if isBadForm {
			return true
		}
	}

	return false
}

// isBadFormOfKing - 動かした駒が玉なら
func isBadFormOfKing(pPos *l15.Position, turn l03.Phase, from l03.Square, to l03.Square) bool {
	// 桂馬の利きに飛び込む動きは悪形
	{
		var squares = GetSqOfOpponentKnightFrom(turn, to)
		for _, sq := range squares {
			var piece = pPos.GetPieceOnBoardAtSq(sq)
			var piecetype = l03.What(piece)
			var turn2 = l03.Who(piece)
			if piecetype == l03.PIECE_TYPE_N && turn2 != turn {
				return true
			}
		}
	}

	return false
}

// isBadFormOfRook - 動かした駒が飛なら
func isBadFormOfRook(pPos *l15.Position, turn l03.Phase, from l03.Square, to l03.Square, promotion bool) bool {
	if promotion {
		return false
	}

	// 敵陣で成らないのは悪形
	{
		var minRank int8
		var overRank int8

		switch turn {
		case l03.FIRST:
			minRank = 1
			overRank = 4
		case l03.SECOND:
			minRank = 7
			overRank = 10
		default:
			panic(App.LogNotEcho.Fatal("fatal: unknown turn=%d", turn))
		}

		var newRank = l03.Rank(to)

		if minRank <= newRank && newRank < overRank {
			return true
		}
	}

	return false
}

// isBadFormOfBishop - 動かした駒が角なら
func isBadFormOfBishop(pPos *l15.Position, turn l03.Phase, from l03.Square, to l03.Square, promotion bool) bool {
	if promotion {
		return false
	}

	// 敵陣で成らないのは悪形
	{
		var minRank int8
		var overRank int8

		switch turn {
		case l03.FIRST:
			minRank = 1
			overRank = 4
		case l03.SECOND:
			minRank = 7
			overRank = 10
		default:
			panic(App.LogNotEcho.Fatal("fatal: unknown turn=%d", turn))
		}

		var newRank = l03.Rank(to)

		if minRank <= newRank && newRank < overRank {
			return true
		}
	}

	return false
}

// isBadFormOfGold - 動かした駒が金なら
func isBadFormOfGold(pPos *l15.Position, turn l03.Phase, from l03.Square, to l03.Square) bool {
	// 駒を取る動きは、悪形とはしません
	{
		// 移動先に駒はあるか？
		var captured = pPos.GetPieceOnBoardAtSq(to)
		if captured != l03.PIECE_EMPTY {
			// 自駒は取らないので、相手の駒を取った
			return false
		}
	}

	// 自分から見て手前を南としたときの、移動先の１つ南のマス
	var southSq = GetSqSouthOf(turn, to)

	if southSq != l03.SQ_EMPTY {
		// 盤内
		// その座標の駒は？
		var southPiece = pPos.GetPieceOnBoardAtSq(southSq)
		// if App.IsDebug {
		// 	App.Out.Print("# southSq=%d southPiece=%s\n", southSq, southPiece.ToCodeOfPc())
		// }

		// その駒の先後は？
		var friendPhase = l03.Who(southPiece)
		if turn == friendPhase {
			// if App.IsDebug {
			// 	App.Out.Print("# turn=%s friendPhase=%s\n", turn.ToCodeOfPh(), friendPhase.ToCodeOfPh())
			// }

			// その駒の種類は？
			var friendPieceType = l03.What(southPiece)
			// if App.IsDebug {
			// 	App.Out.Print("# friendPieceType=%s\n", friendPieceType.ToCodeOfPt())
			// }

			switch friendPieceType {
			case l03.PIECE_TYPE_S:
				// 悪形
				// +--+
				// |金|
				// +--+
				// |銀|
				// +--+
				// 自金の１つ南に自銀がある形

				if App.IsDebug {
					App.Out.Print("# Avoid Pattern 1: Gold to=%d\n", to)
				}
				return true // 悪形はスキップします

			}
		}
	}

	// 自玉に近い自金が、玉から離れる動きは悪形とします
	{
		var myKingSq l03.Square
		var yourKingSq l03.Square
		if turn == l03.FIRST {
			myKingSq = pPos.PieceLocations[l07.PCLOC_K1]
			yourKingSq = pPos.PieceLocations[l07.PCLOC_K2]
		} else {
			myKingSq = pPos.PieceLocations[l07.PCLOC_K2]
			yourKingSq = pPos.PieceLocations[l07.PCLOC_K1]
		}

		// 敵玉の近くにある自金を　敵玉に近づけるのは悪形ではありません
		{
			var manhaYkF = GetManhattanDistance(yourKingSq, from)
			if manhaYkF < 4 {
				var manhaYkT = GetManhattanDistance(yourKingSq, to)
				if manhaYkT <= manhaYkF {
					return false
				}
			}
		}

		// 自玉に近い自金が、玉から離れる動きは悪形とします
		{
			// 動かした駒とのマンハッタン距離
			var manhaMkF = GetManhattanDistance(myKingSq, from)
			if manhaMkF < 4 {

				// もともと近くにあった駒
				var manhaMkT = GetManhattanDistance(myKingSq, to)
				if manhaMkF < manhaMkT {
					// 遠ざかる動きは悪形
					return true
				}
			}
		}
	}

	return false
}

// isBadFormOfSilver - 動かした駒が銀なら
func isBadFormOfSilver(pPos *l15.Position, turn l03.Phase, from l03.Square, to l03.Square) bool {
	// 駒を取る動きは、悪形とはしません
	{
		// 移動先に駒はあるか？
		var captured = pPos.GetPieceOnBoardAtSq(to)
		if captured != l03.PIECE_EMPTY {
			// 自駒は取らないので、相手の駒を取った
			return false
		}
	}

	// 自分から見て手前を南としたときの、移動先の１つ南西と南東のマス。無ければ空マス
	var squares = GetSqWestSouthAndEastSouthOf(turn, to)

	for _, xxstSouthSq := range squares {
		if xxstSouthSq == l03.SQ_EMPTY {
			// 駒が無いが、盤外かもしれない。ループの次項へ
			continue
		}

		// 盤内
		// その座標の駒は？
		var southPiece = pPos.GetPieceOnBoardAtSq(xxstSouthSq)

		// その駒の先後は？
		var friendPhase = l03.Who(southPiece)
		if turn != friendPhase {
			// 敵の駒
			return false
		}

		// その駒の種類は？
		var friendPieceType = l03.What(southPiece)

		switch friendPieceType {
		case l03.PIECE_TYPE_R:
		case l03.PIECE_TYPE_N:
		case l03.PIECE_TYPE_L:
		case l03.PIECE_TYPE_P:
			// 悪形
			// +--+--+--+
			// |  |銀|  |
			// +--+--+--+
			// |★|  |★|
			// +--+--+--+
			// * ★は、自飛、自桂、自香、自歩のいずれか
			// * 銀に下がるところがない形。相互紐づきがある場合を除く
		default:
			return false
		}

	}

	// 自玉に近い自銀が、玉から離れる動きは悪形とします
	{
		var myKingSq l03.Square
		var yourKingSq l03.Square
		if turn == l03.FIRST {
			myKingSq = pPos.PieceLocations[l07.PCLOC_K1]
			yourKingSq = pPos.PieceLocations[l07.PCLOC_K2]
		} else {
			myKingSq = pPos.PieceLocations[l07.PCLOC_K2]
			yourKingSq = pPos.PieceLocations[l07.PCLOC_K1]
		}

		// 敵玉の近くにある自駒を　敵玉に近づけるのは悪形ではありません
		{
			var manhaYkF = GetManhattanDistance(yourKingSq, from)
			if manhaYkF < 4 {
				var manhaYkT = GetManhattanDistance(yourKingSq, to)
				if manhaYkT <= manhaYkF {
					return false
				}
			}
		}

		// 自玉に近い自駒が、玉から離れる動きは悪形とします
		{
			// 動かした駒とのマンハッタン距離
			var manhaMkF = GetManhattanDistance(myKingSq, from)
			if manhaMkF < 4 {

				// もともと近くにあった駒
				var manhaMkT = GetManhattanDistance(myKingSq, to)
				if manhaMkF < manhaMkT {
					// 遠ざかる動きは悪形
					return true
				}
			}
		}
	}

	return true
}

// isBadFormOfLance - 動かした駒が香なら
func isBadFormOfLance(pPos *l15.Position, turn l03.Phase, to l03.Square, promotion bool) bool {
	if promotion {
		return false
	}

	var rank1 int8
	var rank2 int8
	switch turn {
	case l03.FIRST:
		rank1 = 1
		rank2 = 2
	case l03.SECOND:
		rank1 = 9
		rank2 = 8
	default:
		panic(App.LogNotEcho.Fatal("fatal: unknown turn=%d", turn))
	}

	var newRank = l03.Rank(to)

	if newRank == rank1 || newRank == rank2 {
		// 1段目、2段目で成らない香は省く
		return true
	}

	return false
}

// isBadFormOfDroppedPawn - 打った駒が歩なら
func isBadFormOfDroppedPawn(pPos *l15.Position, turn l03.Phase, to l03.Square) bool {
	// 打ち歩詰め判定が大変なので、玉頭を歩で叩くのは悪形とする

	var northSq = GetSqNorthOf(turn, to)

	if northSq == l03.SQ_EMPTY {
		return false
	}

	if App.IsDebug {
		App.Log.Debug("歩打判定2\n")
	}

	var northPiece = pPos.GetPieceOnBoardAtSq(northSq)

	if l03.What(northPiece) == l03.PIECE_TYPE_K && l03.Who(northPiece) != turn {
		if App.IsDebug {
			App.Log.Debug("打ち歩詰め判定1\n")
		}
		// +---+
		// |玉v|
		// +---+
		// |歩 |
		// +---+
		return true
	}

	return false
}
