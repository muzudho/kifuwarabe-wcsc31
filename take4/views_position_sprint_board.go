package take4

import (
	"fmt"

	l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"
)

// Print - 局面出力（＾ｑ＾）
func SprintBoard(pos *Position) string {
	var phase_str = "First"
	if pos.Phase == 2 {
		phase_str = "Second"
	}

	var s1 = "\n" +
		//
		fmt.Sprintf("[%d moves / %s / ? repeats]\n", pos.MovesNum, phase_str) +
		//
		"\n" +
		//
		"  r  b  g  s  n  l  p\n" +
		"+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pos.Hands[7], pos.Hands[8], pos.Hands[9], pos.Hands[10], pos.Hands[11], pos.Hands[12], pos.Hands[13]) +
		//
		"+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		fmt.Sprintf(" %2s %2s %2s %2s %2s %2s %2s %2s %2s %2s\n", pos.GetPieceAtIndex(90).ToCodeOfPc(), pos.GetPieceAtIndex(80).ToCodeOfPc(), pos.GetPieceAtIndex(70).ToCodeOfPc(), pos.GetPieceAtIndex(60).ToCodeOfPc(), pos.GetPieceAtIndex(50).ToCodeOfPc(), pos.GetPieceAtIndex(40).ToCodeOfPc(), pos.GetPieceAtIndex(30).ToCodeOfPc(), pos.GetPieceAtIndex(20).ToCodeOfPc(), pos.GetPieceAtIndex(10).ToCodeOfPc(), pos.GetPieceAtIndex(0).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(91).ToCodeOfPc(), pos.GetPieceAtIndex(81).ToCodeOfPc(), pos.GetPieceAtIndex(71).ToCodeOfPc(), pos.GetPieceAtIndex(61).ToCodeOfPc(), pos.GetPieceAtIndex(51).ToCodeOfPc(), pos.GetPieceAtIndex(41).ToCodeOfPc(), pos.GetPieceAtIndex(31).ToCodeOfPc(), pos.GetPieceAtIndex(21).ToCodeOfPc(), pos.GetPieceAtIndex(11).ToCodeOfPc(), pos.GetPieceAtIndex(1).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(92).ToCodeOfPc(), pos.GetPieceAtIndex(82).ToCodeOfPc(), pos.GetPieceAtIndex(72).ToCodeOfPc(), pos.GetPieceAtIndex(62).ToCodeOfPc(), pos.GetPieceAtIndex(52).ToCodeOfPc(), pos.GetPieceAtIndex(42).ToCodeOfPc(), pos.GetPieceAtIndex(32).ToCodeOfPc(), pos.GetPieceAtIndex(22).ToCodeOfPc(), pos.GetPieceAtIndex(12).ToCodeOfPc(), pos.GetPieceAtIndex(2).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(93).ToCodeOfPc(), pos.GetPieceAtIndex(83).ToCodeOfPc(), pos.GetPieceAtIndex(73).ToCodeOfPc(), pos.GetPieceAtIndex(63).ToCodeOfPc(), pos.GetPieceAtIndex(53).ToCodeOfPc(), pos.GetPieceAtIndex(43).ToCodeOfPc(), pos.GetPieceAtIndex(33).ToCodeOfPc(), pos.GetPieceAtIndex(23).ToCodeOfPc(), pos.GetPieceAtIndex(13).ToCodeOfPc(), pos.GetPieceAtIndex(3).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(94).ToCodeOfPc(), pos.GetPieceAtIndex(84).ToCodeOfPc(), pos.GetPieceAtIndex(74).ToCodeOfPc(), pos.GetPieceAtIndex(64).ToCodeOfPc(), pos.GetPieceAtIndex(54).ToCodeOfPc(), pos.GetPieceAtIndex(44).ToCodeOfPc(), pos.GetPieceAtIndex(34).ToCodeOfPc(), pos.GetPieceAtIndex(24).ToCodeOfPc(), pos.GetPieceAtIndex(14).ToCodeOfPc(), pos.GetPieceAtIndex(4).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(95).ToCodeOfPc(), pos.GetPieceAtIndex(85).ToCodeOfPc(), pos.GetPieceAtIndex(75).ToCodeOfPc(), pos.GetPieceAtIndex(65).ToCodeOfPc(), pos.GetPieceAtIndex(55).ToCodeOfPc(), pos.GetPieceAtIndex(45).ToCodeOfPc(), pos.GetPieceAtIndex(35).ToCodeOfPc(), pos.GetPieceAtIndex(25).ToCodeOfPc(), pos.GetPieceAtIndex(15).ToCodeOfPc(), pos.GetPieceAtIndex(5).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(96).ToCodeOfPc(), pos.GetPieceAtIndex(86).ToCodeOfPc(), pos.GetPieceAtIndex(76).ToCodeOfPc(), pos.GetPieceAtIndex(66).ToCodeOfPc(), pos.GetPieceAtIndex(56).ToCodeOfPc(), pos.GetPieceAtIndex(46).ToCodeOfPc(), pos.GetPieceAtIndex(36).ToCodeOfPc(), pos.GetPieceAtIndex(26).ToCodeOfPc(), pos.GetPieceAtIndex(16).ToCodeOfPc(), pos.GetPieceAtIndex(6).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(97).ToCodeOfPc(), pos.GetPieceAtIndex(87).ToCodeOfPc(), pos.GetPieceAtIndex(77).ToCodeOfPc(), pos.GetPieceAtIndex(67).ToCodeOfPc(), pos.GetPieceAtIndex(57).ToCodeOfPc(), pos.GetPieceAtIndex(47).ToCodeOfPc(), pos.GetPieceAtIndex(37).ToCodeOfPc(), pos.GetPieceAtIndex(27).ToCodeOfPc(), pos.GetPieceAtIndex(17).ToCodeOfPc(), pos.GetPieceAtIndex(7).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(98).ToCodeOfPc(), pos.GetPieceAtIndex(88).ToCodeOfPc(), pos.GetPieceAtIndex(78).ToCodeOfPc(), pos.GetPieceAtIndex(68).ToCodeOfPc(), pos.GetPieceAtIndex(58).ToCodeOfPc(), pos.GetPieceAtIndex(48).ToCodeOfPc(), pos.GetPieceAtIndex(38).ToCodeOfPc(), pos.GetPieceAtIndex(28).ToCodeOfPc(), pos.GetPieceAtIndex(18).ToCodeOfPc(), pos.GetPieceAtIndex(8).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s|%2s\n", pos.GetPieceAtIndex(99).ToCodeOfPc(), pos.GetPieceAtIndex(89).ToCodeOfPc(), pos.GetPieceAtIndex(79).ToCodeOfPc(), pos.GetPieceAtIndex(69).ToCodeOfPc(), pos.GetPieceAtIndex(59).ToCodeOfPc(), pos.GetPieceAtIndex(49).ToCodeOfPc(), pos.GetPieceAtIndex(39).ToCodeOfPc(), pos.GetPieceAtIndex(29).ToCodeOfPc(), pos.GetPieceAtIndex(19).ToCodeOfPc(), pos.GetPieceAtIndex(9).ToCodeOfPc()) +
		//
		"+--+--+--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"        R  B  G  S  N  L  P\n" +
		"      +--+--+--+--+--+--+--+\n" +
		//
		fmt.Sprintf("      |%2d|%2d|%2d|%2d|%2d|%2d|%2d|\n", pos.Hands[0], pos.Hands[1], pos.Hands[2], pos.Hands[3], pos.Hands[4], pos.Hands[5], pos.Hands[6]) +
		//
		"      +--+--+--+--+--+--+--+\n" +
		//
		"\n" +
		//
		"moves"

	moves_list := make([]byte, 0, l02.MOVES_SIZE*6) // 6文字 512手分で ほとんどの大会で大丈夫だろ（＾～＾）
	for _, pMove := range pos.Moves {
		moves_list = append(moves_list, ' ')
		moves_list = append(moves_list, pMove.ToCodeOfM()...)
	}

	// unsafe使うと速いみたいなんだが、読みにくくなるしな（＾～＾）
	// return s1 + *(*string)(unsafe.Pointer(&moves_list)) + "\n"
	return s1 + string(moves_list) + "\n"
}
