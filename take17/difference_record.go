package take17

import (
	l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// 差分での連続局面記録。つまり、ふつうの棋譜（＾～＾）
type DifferenceRecord struct {
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [l02.MOVES_SIZE]l03.Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [l02.MOVES_SIZE]l03.Piece
}

func NewDifferenceRecord() *DifferenceRecord {
	var pRecord = new(DifferenceRecord)
	pRecord.ResetDifferenceRecord()
	return pRecord
}

func (pRecord *DifferenceRecord) ResetDifferenceRecord() {
	pRecord.OffsetMovesIndex = 0
	// 何手目か
	pRecord.StartMovesNum = 1
	pRecord.OffsetMovesIndex = 0
	// 指し手のリスト
	pRecord.Moves = [l02.MOVES_SIZE]l03.Move{}
	// 取った駒のリスト
	pRecord.CapturedList = [l02.MOVES_SIZE]l03.Piece{}
}
