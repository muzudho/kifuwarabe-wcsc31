package take7

import (
	"fmt"

	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
)

// Print - 局面出力（＾ｑ＾）
func SprintBoard(pPos *Position) string {
	var phase_str = "?"
	if pPos.Phase == l06.FIRST {
		phase_str = "First"
	} else if pPos.Phase == l06.SECOND {
		phase_str = "Second"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d -> %d moves / %s / ? repeats]\n", pPos.StartMovesNum, (pPos.StartMovesNum+pPos.OffsetMovesIndex), phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[7], pPos.Hands[8], pPos.Hands[9], pPos.Hands[10], pPos.Hands[11], pPos.Hands[12], pPos.Hands[13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pPos.GetPieceAtIndex(90).ToCodeOfPc(), pPos.GetPieceAtIndex(80).ToCodeOfPc(), pPos.GetPieceAtIndex(70).ToCodeOfPc(), pPos.GetPieceAtIndex(60).ToCodeOfPc(), pPos.GetPieceAtIndex(50).ToCodeOfPc(), pPos.GetPieceAtIndex(40).ToCodeOfPc(), pPos.GetPieceAtIndex(30).ToCodeOfPc(), pPos.GetPieceAtIndex(20).ToCodeOfPc(), pPos.GetPieceAtIndex(10).ToCodeOfPc(), pPos.GetPieceAtIndex(0).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(91).ToCodeOfPc(), pPos.GetPieceAtIndex(81).ToCodeOfPc(), pPos.GetPieceAtIndex(71).ToCodeOfPc(), pPos.GetPieceAtIndex(61).ToCodeOfPc(), pPos.GetPieceAtIndex(51).ToCodeOfPc(), pPos.GetPieceAtIndex(41).ToCodeOfPc(), pPos.GetPieceAtIndex(31).ToCodeOfPc(), pPos.GetPieceAtIndex(21).ToCodeOfPc(), pPos.GetPieceAtIndex(11).ToCodeOfPc(), pPos.GetPieceAtIndex(1).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(92).ToCodeOfPc(), pPos.GetPieceAtIndex(82).ToCodeOfPc(), pPos.GetPieceAtIndex(72).ToCodeOfPc(), pPos.GetPieceAtIndex(62).ToCodeOfPc(), pPos.GetPieceAtIndex(52).ToCodeOfPc(), pPos.GetPieceAtIndex(42).ToCodeOfPc(), pPos.GetPieceAtIndex(32).ToCodeOfPc(), pPos.GetPieceAtIndex(22).ToCodeOfPc(), pPos.GetPieceAtIndex(12).ToCodeOfPc(), pPos.GetPieceAtIndex(2).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(93).ToCodeOfPc(), pPos.GetPieceAtIndex(83).ToCodeOfPc(), pPos.GetPieceAtIndex(73).ToCodeOfPc(), pPos.GetPieceAtIndex(63).ToCodeOfPc(), pPos.GetPieceAtIndex(53).ToCodeOfPc(), pPos.GetPieceAtIndex(43).ToCodeOfPc(), pPos.GetPieceAtIndex(33).ToCodeOfPc(), pPos.GetPieceAtIndex(23).ToCodeOfPc(), pPos.GetPieceAtIndex(13).ToCodeOfPc(), pPos.GetPieceAtIndex(3).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(94).ToCodeOfPc(), pPos.GetPieceAtIndex(84).ToCodeOfPc(), pPos.GetPieceAtIndex(74).ToCodeOfPc(), pPos.GetPieceAtIndex(64).ToCodeOfPc(), pPos.GetPieceAtIndex(54).ToCodeOfPc(), pPos.GetPieceAtIndex(44).ToCodeOfPc(), pPos.GetPieceAtIndex(34).ToCodeOfPc(), pPos.GetPieceAtIndex(24).ToCodeOfPc(), pPos.GetPieceAtIndex(14).ToCodeOfPc(), pPos.GetPieceAtIndex(4).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(95).ToCodeOfPc(), pPos.GetPieceAtIndex(85).ToCodeOfPc(), pPos.GetPieceAtIndex(75).ToCodeOfPc(), pPos.GetPieceAtIndex(65).ToCodeOfPc(), pPos.GetPieceAtIndex(55).ToCodeOfPc(), pPos.GetPieceAtIndex(45).ToCodeOfPc(), pPos.GetPieceAtIndex(35).ToCodeOfPc(), pPos.GetPieceAtIndex(25).ToCodeOfPc(), pPos.GetPieceAtIndex(15).ToCodeOfPc(), pPos.GetPieceAtIndex(5).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(96).ToCodeOfPc(), pPos.GetPieceAtIndex(86).ToCodeOfPc(), pPos.GetPieceAtIndex(76).ToCodeOfPc(), pPos.GetPieceAtIndex(66).ToCodeOfPc(), pPos.GetPieceAtIndex(56).ToCodeOfPc(), pPos.GetPieceAtIndex(46).ToCodeOfPc(), pPos.GetPieceAtIndex(36).ToCodeOfPc(), pPos.GetPieceAtIndex(26).ToCodeOfPc(), pPos.GetPieceAtIndex(16).ToCodeOfPc(), pPos.GetPieceAtIndex(6).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(97).ToCodeOfPc(), pPos.GetPieceAtIndex(87).ToCodeOfPc(), pPos.GetPieceAtIndex(77).ToCodeOfPc(), pPos.GetPieceAtIndex(67).ToCodeOfPc(), pPos.GetPieceAtIndex(57).ToCodeOfPc(), pPos.GetPieceAtIndex(47).ToCodeOfPc(), pPos.GetPieceAtIndex(37).ToCodeOfPc(), pPos.GetPieceAtIndex(27).ToCodeOfPc(), pPos.GetPieceAtIndex(17).ToCodeOfPc(), pPos.GetPieceAtIndex(7).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(98).ToCodeOfPc(), pPos.GetPieceAtIndex(88).ToCodeOfPc(), pPos.GetPieceAtIndex(78).ToCodeOfPc(), pPos.GetPieceAtIndex(68).ToCodeOfPc(), pPos.GetPieceAtIndex(58).ToCodeOfPc(), pPos.GetPieceAtIndex(48).ToCodeOfPc(), pPos.GetPieceAtIndex(38).ToCodeOfPc(), pPos.GetPieceAtIndex(28).ToCodeOfPc(), pPos.GetPieceAtIndex(18).ToCodeOfPc(), pPos.GetPieceAtIndex(8).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pPos.GetPieceAtIndex(99).ToCodeOfPc(), pPos.GetPieceAtIndex(89).ToCodeOfPc(), pPos.GetPieceAtIndex(79).ToCodeOfPc(), pPos.GetPieceAtIndex(69).ToCodeOfPc(), pPos.GetPieceAtIndex(59).ToCodeOfPc(), pPos.GetPieceAtIndex(49).ToCodeOfPc(), pPos.GetPieceAtIndex(39).ToCodeOfPc(), pPos.GetPieceAtIndex(29).ToCodeOfPc(), pPos.GetPieceAtIndex(19).ToCodeOfPc(), pPos.GetPieceAtIndex(9).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pPos.Hands[0], pPos.Hands[1], pPos.Hands[2], pPos.Hands[3], pPos.Hands[4], pPos.Hands[5], pPos.Hands[6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	var moves_text = pPos.createMovesText()

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	//return s1 + *(*string)(unsafe.Pointer(&moves_text)) + "\n"
	return s1 + string(moves_text) + "\n"
}
