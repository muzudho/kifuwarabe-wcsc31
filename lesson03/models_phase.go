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
