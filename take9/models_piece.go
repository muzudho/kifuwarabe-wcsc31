package take9 // not same lesson03

// 先後付きの駒
type Piece uint8

// 駒
const (
	PIECE_EMPTY = iota // 0: 駒なし
	PIECE_K1           // 1: ▲玉
	PIECE_R1           // 2: ▲飛
	PIECE_B1           // 3: ▲角
	PIECE_G1           // 4: ▲金
	PIECE_S1           // 5: ▲銀
	PIECE_N1           // 6: ▲桂
	PIECE_L1           // 7: ▲香
	PIECE_P1           // 8: ▲歩
	PIECE_PR1          // 9: ▲竜
	PIECE_PB1          //10: ▲馬
	PIECE_PS1          //11: ▲全
	PIECE_PN1          //12: ▲圭
	PIECE_PL1          //13: ▲杏
	PIECE_PP1          //14: ▲と
	PIECE_K2           //15: ▽玉
	PIECE_R2           //16: ▽飛
	PIECE_B2           //17: ▽角
	PIECE_G2           //18: ▽金
	PIECE_S2           //19: ▽銀
	PIECE_N2           //20: ▽桂
	PIECE_L2           //21: ▽香
	PIECE_P2           //22: ▽歩
	PIECE_PR2          //23: ▽竜
	PIECE_PB2          //24: ▽馬
	PIECE_PS2          //25: ▽全
	PIECE_PN2          //26: ▽圭
	PIECE_PL2          //27: ▽杏
	PIECE_PP2          //28: ▽と
)
