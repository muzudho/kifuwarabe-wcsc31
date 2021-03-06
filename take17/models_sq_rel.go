package take17

import (
	"fmt"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// GetSqNorthOf - 自分から見て手前を南としたときの、１つ北のマス。無ければ空マス
func GetSqNorthOf(turn l03.Phase, srcSq l03.Square) l03.Square {
	var latitude int8

	switch turn {
	case l03.FIRST:
		latitude = -1
	case l03.SECOND:
		latitude = 1
	default:
		panic(App.Log.Fatal(fmt.Sprintf("turn=[%d]", turn)))
	}

	var newRank = l03.Rank(srcSq) + latitude

	if 1 <= newRank && newRank < 10 {
		// 盤内
		var newFile = l03.File(srcSq)
		return l03.FromFileRankToSq(newFile, newRank)
	}

	return l03.SQ_EMPTY
}

// GetSqSouthOf - 自分から見て手前を南としたときの、１つ南のマス。無ければ空マス
func GetSqSouthOf(turn l03.Phase, src l03.Square) l03.Square {
	var latitude int8

	switch turn {
	case l03.FIRST:
		latitude = 1
	case l03.SECOND:
		latitude = -1
	default:
		panic(App.Log.Fatal(fmt.Sprintf("turn=[%d]", turn)))
	}

	var newRank = l03.Rank(src) + latitude

	if 1 <= newRank && newRank < 10 {
		// 盤内
		var newFile = l03.File(src)
		// if App.IsDebug {
		// 	App.Out.Print("# newFile=%d newRank=%d\n", newFile, newRank)
		// }

		// 移動先マスの南の座標
		return l03.FromFileRankToSq(newFile, newRank)
	}

	return l03.SQ_EMPTY
}

// GetSqWestSouthAndEastSouthOf - 自分から見て手前を南としたときの、１つ南西と南東のマス。無ければ空マス
func GetSqWestSouthAndEastSouthOf(turn l03.Phase, src l03.Square) [2]l03.Square {
	var latitude int8
	var longitude1 int8
	var longitude2 int8

	switch turn {
	case l03.FIRST:
		latitude = 1
		longitude1 = -1
		longitude2 = 1
	case l03.SECOND:
		latitude = -1
		longitude1 = 1
		longitude2 = -1
	default:
		panic(App.Log.Fatal(fmt.Sprintf("turn=[%d]", turn)))
	}

	var newRank = l03.Rank(src) + latitude
	var newFile1 = l03.File(src + l03.Square(longitude1))
	var newFile2 = l03.File(src + l03.Square(longitude2))

	if 1 <= newRank && newRank < 10 {
		// 盤内
		// if App.IsDebug {
		// 	App.Out.Print("# newFile=%d newRank=%d\n", newFile, newRank)
		// }
		var sq1 l03.Square
		var sq2 l03.Square

		if 1 <= newFile1 && newFile1 < 10 {
			sq1 = l03.FromFileRankToSq(newFile1, newRank)
		} else {
			sq1 = l03.SQ_EMPTY
		}

		if 1 <= newFile2 && newFile2 < 10 {
			sq2 = l03.FromFileRankToSq(newFile2, newRank)
		} else {
			sq2 = l03.SQ_EMPTY
		}

		// 移動先マスの南の座標
		return [2]l03.Square{sq1, sq2}
	}

	return [2]l03.Square{l03.SQ_EMPTY, l03.SQ_EMPTY}
}

// GetSqOfOpponentKnightFrom - 自分から見て手前を南としたときの、指定の升に利いている相手の桂馬の位置
func GetSqOfOpponentKnightFrom(turn l03.Phase, src l03.Square) [2]l03.Square {
	var latitude int8
	var longitude1 int8
	var longitude2 int8

	switch turn {
	case l03.FIRST:
		latitude = -2
		longitude1 = 1
		longitude2 = -1
	case l03.SECOND:
		latitude = 2
		longitude1 = -1
		longitude2 = 1
	default:
		panic(App.Log.Fatal(fmt.Sprintf("turn=[%d]", turn)))
	}

	var newRank = l03.Rank(src) + latitude
	var newFile1 = l03.File(src) + longitude1
	var newFile2 = l03.File(src) + longitude2

	if 1 <= newRank && newRank < 10 {
		// 盤内
		var sq1 l03.Square
		var sq2 l03.Square

		if 1 <= newFile1 && newFile1 < 10 {
			sq1 = l03.FromFileRankToSq(newFile1, newRank)
		} else {
			sq1 = l03.SQ_EMPTY
		}

		if 1 <= newFile2 && newFile2 < 10 {
			sq2 = l03.FromFileRankToSq(newFile2, newRank)
		} else {
			sq2 = l03.SQ_EMPTY
		}

		// 移動先マスの南の座標
		return [2]l03.Square{sq1, sq2}
	}

	return [2]l03.Square{l03.SQ_EMPTY, l03.SQ_EMPTY}
}
