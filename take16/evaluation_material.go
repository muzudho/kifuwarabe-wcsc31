// 駒の価値
package take16

import (
	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// EvalMaterial - 駒の価値。開発者のむずでょが勝手に決めた（＾～＾）
func EvalMaterial(piece p.Piece) p.Value {
	switch piece {
	case p.PIECE_EMPTY: // 空きマス
		return 0
	case p.PIECE_K1, p.PIECE_K2: // 玉
		return 15000
	case p.PIECE_R1, p.PIECE_PR1, p.PIECE_R2, p.PIECE_PR2: // 飛、竜
		return 1000
	case p.PIECE_B1, p.PIECE_PB1, p.PIECE_B2, p.PIECE_PB2: // 角、馬
		return 900
	case p.PIECE_G1, p.PIECE_G2: // 金
		return 600
	case p.PIECE_S1, p.PIECE_PS1, p.PIECE_S2, p.PIECE_PS2: // 銀、全
		return 500
	case p.PIECE_N1, p.PIECE_PN1, p.PIECE_N2, p.PIECE_PN2: // 桂、圭
		return 250
	case p.PIECE_L1, p.PIECE_PL1, p.PIECE_L2, p.PIECE_PL2: // 香、杏
		return 200
	case p.PIECE_P1, p.PIECE_PP1, p.PIECE_P2, p.PIECE_PP2: // 歩、と
		return 100
	default:
		panic(G.Log.Fatal("Error: Unknown piece=[%d]", piece))
	}
}
