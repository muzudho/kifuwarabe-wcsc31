package lesson03

// 先後のない駒種類
type PieceType byte

const (
	PIECE_TYPE_EMPTY PieceType = 0 + iota // 空マス
	PIECE_TYPE_K
	PIECE_TYPE_R
	PIECE_TYPE_B
	PIECE_TYPE_G
	PIECE_TYPE_S
	PIECE_TYPE_N
	PIECE_TYPE_L
	PIECE_TYPE_P
	PIECE_TYPE_PR
	PIECE_TYPE_PB
	PIECE_TYPE_PS
	PIECE_TYPE_PN
	PIECE_TYPE_PL
	PIECE_TYPE_PP
)

// ToCodeOfPt - 文字列
func (pt PieceType) ToCodeOfPt() string {
	switch pt {
	case PIECE_TYPE_EMPTY:
		return ""
	case PIECE_TYPE_K:
		return "K"
	case PIECE_TYPE_R:
		return "R"
	case PIECE_TYPE_B:
		return "B"
	case PIECE_TYPE_G:
		return "G"
	case PIECE_TYPE_S:
		return "S"
	case PIECE_TYPE_N:
		return "N"
	case PIECE_TYPE_L:
		return "L"
	case PIECE_TYPE_P:
		return "P"
	case PIECE_TYPE_PR:
		return "+R"
	case PIECE_TYPE_PB:
		return "+B"
	case PIECE_TYPE_PS:
		return "+S"
	case PIECE_TYPE_PN:
		return "+N"
	case PIECE_TYPE_PL:
		return "+L"
	case PIECE_TYPE_PP:
		return "+P"
	default:
		panic(App.LogNotEcho.Fatal("unknown piece type=%d", pt))
	}
}
