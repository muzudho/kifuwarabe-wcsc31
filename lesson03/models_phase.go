package lesson03

// 1:先手 2:後手
type Phase byte

const (
	// 空マス
	ZEROTH Phase = iota
	// 先手
	FIRST
	// 後手
	SECOND
)

// ToCodeOfPh - 文字列
func (ph Phase) ToCodeOfPh() string {
	switch ph {
	case ZEROTH:
		return "Z"
	case FIRST:
		return "F"
	case SECOND:
		return "S"
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", ph))
	}
}
