// 駒の価値
package take16

import l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"

// EvalMaterial - 駒の価値。開発者のむずでょが勝手に決めた（＾～＾）
func EvalMaterial(piece l09.Piece) Value {
	switch piece {
	case PIECE_EMPTY: // 空きマス
		return 0
	case PIECE_K1, PIECE_K2: // 玉
		return 15000
	case PIECE_R1, PIECE_PR1, PIECE_R2, PIECE_PR2: // 飛、竜
		return 1000
	case PIECE_B1, PIECE_PB1, PIECE_B2, PIECE_PB2: // 角、馬
		return 900
	case PIECE_G1, PIECE_G2: // 金
		return 600
	case PIECE_S1, PIECE_PS1, PIECE_S2, PIECE_PS2: // 銀、全
		return 500
	case PIECE_N1, PIECE_PN1, PIECE_N2, PIECE_PN2: // 桂、圭
		return 250
	case PIECE_L1, PIECE_PL1, PIECE_L2, PIECE_PL2: // 香、杏
		return 200
	case PIECE_P1, PIECE_PP1, PIECE_P2, PIECE_PP2: // 歩、と
		return 100
	default:
		panic(App.LogNotEcho.Fatal("Error: Unknown piece=[%d]", piece))
	}
}
