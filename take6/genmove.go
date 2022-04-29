package take6

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
)

// GenMoveList - 現局面の指し手のリスト。合法手とは限らないし、全ての合法手を含むとも限らないぜ（＾～＾）
func GenMoveList(pPos *Position) []l04.Move {

	move_list := []l04.Move{}

	// 盤面スキャンしたくないけど、駒の位置インデックスを作ってないから 仕方ない（＾～＾）
	for rank := 1; rank < 10; rank += 1 {
		for file := 1; file < 10; file += 1 {
			from := l04.Square(file*10 + rank)
			if pPos.Homo(from) {
				piece := pPos.Board[from]

				switch piece {
				case l03.PIECE_K1.ToCodeOfPc(), l03.PIECE_K2.ToCodeOfPc(): // 先手玉, 後手玉
					if to := l04.Square(from + 9); to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 左上
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from - 1; to%10 != 0 && pPos.Hetero(to) { // 上
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from - 11; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 右上
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from + 10; to/10%10 != 0 && pPos.Hetero(to) { // 左
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from - 10; to/10%10 != 0 && pPos.Hetero(to) { // 右
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from + 11; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 左下
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from + 1; to%10 != 0 && pPos.Hetero(to) { // 下
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
					if to := from - 9; to/10%10 != 0 && to%10 != 0 && pPos.Hetero(to) { // 右下
						move_list = append(move_list, l04.NewMove(from, to, false))
					}
				}
			}
		}
	}

	return move_list
}

// Homo - 手番と移動元の駒を持つプレイヤーが等しければ真。移動先が空なら偽
func (pPos *Position) Homo(to l04.Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return pPos.Phase == Who(pPos.Board[to])
}

// Hetero - 手番と移動先の駒を持つプレイヤーが異なれば真。移動先が空マスでも真
// Homo の逆だぜ（＾～＾）片方ありゃいいんだけど（＾～＾）
func (pPos *Position) Hetero(to l04.Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return pPos.Phase != Who(pPos.Board[to])
}
