package take17

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l15 "github.com/muzudho/kifuwarabe-wcsc31/take15"
)

// IsBadForm - 悪形なら真
func IsBadForm(pPos *l15.Position, pNerve *Nerve, move l03.Move) bool {
	from, to, promotion := move.Destructure()
	// 自分の先後は？
	var turn = pNerve.PPosSys.GetPhase()

	// 打のケースがあることに注意
	if 11 <= from && from < 100 {
		// 動く駒は？
		var movedPiece = pPos.GetPieceAtSq(from)
		var movedPieceType = l03.What(movedPiece)
		if App.IsDebug {
			App.Out.Print("# movePiece=%s\n", movedPiece.ToCodeOfPc())
		}

		var isBadForm = false
		switch movedPieceType { // 動かした駒が
		case l03.PIECE_TYPE_G: // 金
			isBadForm = isBadFormOfGold(pPos, turn, to)
		case l03.PIECE_TYPE_S: // 銀
			isBadForm = isBadFormOfSilver(pPos, turn, to)
		case l03.PIECE_TYPE_L: // 香
			isBadForm = isBadFormOfLance(pPos, turn, to, promotion)
		}

		if isBadForm {
			return true
		}

	}

	return false
}

// isBadFormOfGold - 動かした駒が金なら
func isBadFormOfGold(pPos *l15.Position, turn l03.Phase, to l03.Square) bool {
	// 駒を取る動きは、悪形とはしません
	{
		// 移動先に駒はあるか？
		var captured = pPos.GetPieceAtSq(to)
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
		var southPiece = pPos.GetPieceAtSq(southSq)
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

	/*TODO
	// 玉に近い金が、玉から離れる動きは悪形とします
	var kingSq l03.Square
	if turn == l03.FIRST {
		kingSq = pPos.PieceLocations[l15.PCLOC_K1]
	} else {
		kingSq = pPos.PieceLocations[l15.PCLOC_K2]
	}
	*/

	return false
}

// isBadFormOfSilver - 動かした駒が銀なら
func isBadFormOfSilver(pPos *l15.Position, turn l03.Phase, to l03.Square) bool {
	// 駒を取る動きは、悪形とはしません
	{
		// 移動先に駒はあるか？
		var captured = pPos.GetPieceAtSq(to)
		if captured != l03.PIECE_EMPTY {
			// 自駒は取らないので、相手の駒を取った
			return false
		}
	}

	// 自分から見て手前を南としたときの、移動先の１つ南西と南東のマス。無ければ空マス
	var squares = GetSqWestSouthAndEastSouthOf(turn, to)

	for _, xxstSouthSq := range squares {
		if xxstSouthSq == l03.SQ_EMPTY {
			return false
		}

		// 盤内
		// その座標の駒は？
		var southPiece = pPos.GetPieceAtSq(xxstSouthSq)

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

	return true
}

// isBadFormOfLance - 動かした駒が香なら
func isBadFormOfLance(pPos *l15.Position, turn l03.Phase, to l03.Square, promotion bool) bool {
	if promotion {
		return false
	}

	var rank1 l03.Square
	var rank2 l03.Square
	switch turn {
	case l03.FIRST:
		rank1 = l03.Square(1)
		rank2 = l03.Square(2)
	case l03.SECOND:
		rank1 = l03.Square(9)
		rank2 = l03.Square(8)
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
