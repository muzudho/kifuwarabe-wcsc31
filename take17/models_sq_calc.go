package take17

import (
	"math"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// GetManhattanDistance - マンハッタン距離
func GetManhattanDistance(from l03.Square, to l03.Square) int {
	var x1 = l03.File(from)
	var y1 = l03.Rank(from)
	var x2 = l03.File(to)
	var y2 = l03.Rank(to)

	return int(math.Abs(float64(x2 - x1 + y2 - y1)))
}
