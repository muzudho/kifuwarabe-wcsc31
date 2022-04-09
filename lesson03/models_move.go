package lesson03

// Move - 指し手
type Move struct {
	// [0]移動元 [1]移動先
	// 持ち駒は仕方ないから 100～113 を使おうぜ（＾～＾）
	Squares []Square
	// 成
	Promotion bool
}

func NewMove(from Square, to Square, promotion bool) *Move {
	move := new(Move)
	move.Squares = []Square{from, to}
	move.Promotion = promotion
	return move
}

// ToCode - SFEN の moves の後に続く指し手に使える文字列を返します
func (move *Move) ToCode() string {
	str := make([]byte, 0, 5)
	count := 0

	from, _, pro := move.Destructure()

	switch from {
	case HAND_R1, HAND_R2:
		str = append(str, 'R')
		count = 1
	case HAND_B1, HAND_B2:
		str = append(str, 'B')
		count = 1
	case HAND_G1, HAND_G2:
		str = append(str, 'G')
		count = 1
	case HAND_S1, HAND_S2:
		str = append(str, 'S')
		count = 1
	case HAND_N1, HAND_N2:
		str = append(str, 'N')
		count = 1
	case HAND_L1, HAND_L2:
		str = append(str, 'L')
		count = 1
	case HAND_P1, HAND_P2:
		str = append(str, 'P')
		count = 1
	default:
		// Ignored
	}

	if count == 1 {
		str = append(str, '+')
	}

	for count < 2 {
		// 正常時は必ず２桁（＾～＾）
		file := byte(move.Squares[count] / 10)
		rank := byte(move.Squares[count] % 10)
		// ASCII Code
		// '0'=48, '9'=57, 'a'=97, 'i'=105
		str = append(str, file+48)
		str = append(str, rank+96)
		// fmt.Printf("Debug: file=%d rank=%d\n", file, rank)
		count += 1
	}

	if pro {
		str = append(str, '+')
	}

	return string(str)
}

// Destructure - 移動元マス、移動先マス、成りの有無
func (move Move) Destructure() (Square, Square, bool) {
	var from = move.Squares[0]
	var to = move.Squares[1]
	var pro = move.Promotion
	return from, to, pro
}
