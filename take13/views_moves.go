package take13

import l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"

// CreateMovesList - " 7g7f 3c3d" みたいな部分を返します。最初は半角スペースです
func (pPosSys *PositionSystem) createMovesText() string {
	moves_text := make([]byte, 0, l02.MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for i := 0; i < pPosSys.OffsetMovesIndex; i += 1 {
		moves_text = append(moves_text, ' ')
		moves_text = append(moves_text, pPosSys.Moves[i].ToCodeOfM()...)
	}
	return string(moves_text)
}
