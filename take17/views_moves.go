package take17

// createMovesText - " 7g7f 3c3d" みたいな部分を返します。最初は半角スペースです
func createMovesText(pRecord *DifferenceRecord) string {
	moves_text := make([]byte, 0, pRecord.OffsetMovesIndex*6) // スペース含めて１手最大6文字（＾～＾）
	for i := 0; i < pRecord.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pRecord.Moves[i].ToCodeOfM()...)
	}
	return string(moves_text)
}
