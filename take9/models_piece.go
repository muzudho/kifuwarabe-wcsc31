package take9 // same lesson03

import "fmt"

// 先後付きの駒
type Piece uint8

// 駒
const (
	PIECE_EMPTY Piece = iota // 0: 駒なし
	PIECE_K1                 // 1: ▲玉
	PIECE_R1                 // 2: ▲飛
	PIECE_B1                 // 3: ▲角
	PIECE_G1                 // 4: ▲金
	PIECE_S1                 // 5: ▲銀
	PIECE_N1                 // 6: ▲桂
	PIECE_L1                 // 7: ▲香
	PIECE_P1                 // 8: ▲歩
	PIECE_PR1                // 9: ▲竜
	PIECE_PB1                //10: ▲馬
	PIECE_PS1                //11: ▲全
	PIECE_PN1                //12: ▲圭
	PIECE_PL1                //13: ▲杏
	PIECE_PP1                //14: ▲と
	PIECE_K2                 //15: ▽玉
	PIECE_R2                 //16: ▽飛
	PIECE_B2                 //17: ▽角
	PIECE_G2                 //18: ▽金
	PIECE_S2                 //19: ▽銀
	PIECE_N2                 //20: ▽桂
	PIECE_L2                 //21: ▽香
	PIECE_P2                 //22: ▽歩
	PIECE_PR2                //23: ▽竜
	PIECE_PB2                //24: ▽馬
	PIECE_PS2                //25: ▽全
	PIECE_PN2                //26: ▽圭
	PIECE_PL2                //27: ▽杏
	PIECE_PP2                //28: ▽と
)

// ToCodeOfPc - 文字列
func (pc Piece) ToCodeOfPc() string {
	switch pc {
	case PIECE_EMPTY:
		return ""
	case PIECE_K1:
		return "K"
	case PIECE_R1:
		return "R"
	case PIECE_B1:
		return "B"
	case PIECE_G1:
		return "G"
	case PIECE_S1:
		return "S"
	case PIECE_N1:
		return "N"
	case PIECE_L1:
		return "L"
	case PIECE_P1:
		return "P"
	case PIECE_PR1:
		return "+R"
	case PIECE_PB1:
		return "+B"
	case PIECE_PS1:
		return "+S"
	case PIECE_PN1:
		return "+N"
	case PIECE_PL1:
		return "+L"
	case PIECE_PP1:
		return "+P"
	case PIECE_K2:
		return "k"
	case PIECE_R2:
		return "r"
	case PIECE_B2:
		return "b"
	case PIECE_G2:
		return "g"
	case PIECE_S2:
		return "s"
	case PIECE_N2:
		return "n"
	case PIECE_L2:
		return "l"
	case PIECE_P2:
		return "p"
	case PIECE_PR2:
		return "+r"
	case PIECE_PB2:
		return "+b"
	case PIECE_PS2:
		return "+s"
	case PIECE_PN2:
		return "+n"
	case PIECE_PL2:
		return "+l"
	case PIECE_PP2:
		return "+p"
	default:
		panic(fmt.Errorf("unknown piece=%d", pc))
	}
}
