// 駒の価値
package take14

import (
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
)

// EvalMaterial - 駒の価値。開発者のむずでょが勝手に決めた（＾～＾）
func EvalMaterial(piece l03.Piece) int16 {
	switch piece {
	case l03.PIECE_EMPTY: // 空きマス
		return 0
	case l03.PIECE_K1, l03.PIECE_K2: // 玉
		return 15000
	case l03.PIECE_R1, l03.PIECE_PR1, l03.PIECE_R2, l03.PIECE_PR2: // 飛、竜
		return 1000
	case l03.PIECE_B1, l03.PIECE_PB1, l03.PIECE_B2, l03.PIECE_PB2: // 角、馬
		return 900
	case l03.PIECE_G1, l03.PIECE_G2: // 金
		return 600
	case l03.PIECE_S1, l03.PIECE_PS1, l03.PIECE_S2, l03.PIECE_PS2: // 銀、全
		return 500
	case l03.PIECE_N1, l03.PIECE_PN1, l03.PIECE_N2, l03.PIECE_PN2: // 桂、圭
		return 250
	case l03.PIECE_L1, l03.PIECE_PL1, l03.PIECE_L2, l03.PIECE_PL2: // 香、杏
		return 200
	case l03.PIECE_P1, l03.PIECE_PP1, l03.PIECE_P2, l03.PIECE_PP2: // 歩、と
		return 100
	default:
		panic(App.LogNotEcho.Fatal("unknown piece=[%d]", piece))
	}
}
