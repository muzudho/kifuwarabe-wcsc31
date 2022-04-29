package take16

// SprintBoardFooter - 局面出力（＾ｑ＾）
func (pNerve *Nerve) SprintBoardFooter() string {
	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return "moves" + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return "moves" + createMovesText(pNerve.PRecord) + "\n"
}
