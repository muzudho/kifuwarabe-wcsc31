package take12

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

// 将棋盤の内側をスキャンします。
var centerScanningLine = []l04.Square{
	82, 72, 62, 52, 42, 32, 22,
	83, 73, 63, 53, 43, 33, 23,
	84, 74, 64, 54, 44, 34, 24,
	85, 75, 65, 55, 45, 35, 25,
	86, 76, 66, 56, 46, 36, 26,
	87, 77, 67, 57, 47, 37, 27,
	88, 78, 68, 58, 48, 38, 28,
}

// ブラシの太さ（８近傍）中心を含まない
var centerBrushingArea = []int32{
	9, -1, -11,
	10, -10,
	11, 1, -9}

// 上辺用
var topScanningLine = []l04.Square{
	81, 71, 61, 51, 41, 31, 21,
}
var topBrushingArea = []int32{
	10, -10,
	11, 1, -9}

// 右上用
var rightTopScanningLine = []l04.Square{
	11,
}
var rightTopBrushingArea = []int32{
	10,
	11, 1}

// 右辺用
var rightScanningLine = []l04.Square{
	12,
	13,
	14,
	15,
	16,
	17,
	18,
}
var rightBrushingArea = []int32{
	9, -1,
	10,
	11, 1}

// 右下用
var rightBottomScanningLine = []l04.Square{
	19,
}
var rightBottomBrushingArea = []int32{
	9, -1,
	10}

// 下辺用
var bottomScanningLine = []l04.Square{
	89, 79, 69, 59, 49, 39, 29,
}
var bottomBrushingArea = []int32{
	9, -1, -11,
	10, -10}

// 左下用
var leftBottomScanningLine = []l04.Square{
	99,
}
var leftBottomBrushingArea = []int32{
	-1, -11,
	-10}

// 左辺用
var leftScanningLine = []l04.Square{
	92,
	93,
	94,
	95,
	96,
	97,
	98,
}
var leftBrushingArea = []int32{
	-1, -11,
	-10,
	1, -9}

// 左上用
var leftTopScanningLine = []l04.Square{
	91,
}
var leftTopBrushingArea = []int32{
	-1, -11,
	-10}

// WaterColor - 水で薄めたような評価値にします
// pCB3 = 0
// pCB4 = 0
// pCB5 = 0
// pCB1 - pCB2 = pCB3
// pCB3 - pCB4 = pCB5
func WaterColor(pCB1 *ControlBoard, pCB2 *ControlBoard, pCB3 *ControlBoard, pCB4 *ControlBoard, pCB5 *ControlBoard) {
	// 将棋盤の内側をスキャンします。

	pCB3.Clear()
	pCB4.Clear()
	pCB5.Clear()

	pW := pCB1
	pX := pCB2
	pY := pCB3
	waterColor2(centerScanningLine, centerBrushingArea, pW, pX, pY)
	waterColor2(topScanningLine, topBrushingArea, pW, pX, pY)
	waterColor2(rightTopScanningLine, rightTopBrushingArea, pW, pX, pY)
	waterColor2(rightScanningLine, rightBrushingArea, pW, pX, pY)
	waterColor2(rightBottomScanningLine, rightBottomBrushingArea, pW, pX, pY)
	waterColor2(bottomScanningLine, bottomBrushingArea, pW, pX, pY)
	waterColor2(leftBottomScanningLine, leftBottomBrushingArea, pW, pX, pY)
	waterColor2(leftScanningLine, leftBrushingArea, pW, pX, pY)
	waterColor2(leftTopScanningLine, leftTopBrushingArea, pW, pX, pY)

	pW = pCB3
	pX = pCB4
	pY = pCB5
	waterColor2(centerScanningLine, centerBrushingArea, pW, pX, pY)
	waterColor2(topScanningLine, topBrushingArea, pW, pX, pY)
	waterColor2(rightTopScanningLine, rightTopBrushingArea, pW, pX, pY)
	waterColor2(rightScanningLine, rightBrushingArea, pW, pX, pY)
	waterColor2(rightBottomScanningLine, rightBottomBrushingArea, pW, pX, pY)
	waterColor2(bottomScanningLine, bottomBrushingArea, pW, pX, pY)
	waterColor2(leftBottomScanningLine, leftBottomBrushingArea, pW, pX, pY)
	waterColor2(leftScanningLine, leftBrushingArea, pW, pX, pY)
	waterColor2(leftTopScanningLine, leftTopBrushingArea, pW, pX, pY)
}

func waterColor2(scanningLine []l04.Square, brushingArea []int32, pCB1 *ControlBoard, pCB2 *ControlBoard, pCB3 *ControlBoard) {
	brushAreaSize := int8(len(brushingArea))

	// 真ん中
	for _, sq1 := range scanningLine {
		// ブラシの面積分の利きを総和します
		var sum int8 = 0
		for _, rel := range brushingArea {
			sq2 := l04.Square(int32(sq1) + rel)
			sum += pCB1.Board[sq2] - pCB2.Board[sq2]
		}
		sum /= brushAreaSize
		// 総和したものを平均し、結果表に上乗せします
		for _, rel := range brushingArea {
			sq2 := l04.Square(int32(sq1) + rel)
			pCB3.Board[sq2] += sum
		}
	}
}
