package take8

import (
	"fmt"

	l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

type positionForRecord interface {
	GetOffsetMoveIndex() int
	GetCapturedPieceAtMovesIndex(movesIndex int) l03.Piece
	GetMoveAtMovesIndex(movesIndex int) l03.Move
}

// SprintRecord - 棋譜表示（＾～＾）
func SprintRecord(pPos positionForRecord) string {

	// "8h2b+ b \n" 1行9byteぐらいを想定（＾～＾）
	record_text := make([]byte, 0, l02.MOVES_SIZE*9)
	max := pPos.GetOffsetMoveIndex()
	for i := 0; i < max; i += 1 {
		record_text = append(record_text, pPos.GetMoveAtMovesIndex(i).ToCodeOfM()...)
		record_text = append(record_text, ' ')
		record_text = append(record_text, []byte(pPos.GetCapturedPieceAtMovesIndex(i).ToCodeOfPc())...)
		record_text = append(record_text, '\n')
	}

	return fmt.Sprintf("record\n------\n%s", record_text)
}
