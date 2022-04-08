package take16

import (
	"fmt"
	"strconv"

	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// マス番号 00～99,100～113
type Square uint32

// 評価値
type Value int32

// 00～99
const BOARD_SIZE = 100

// Piece location
const (
	PCLOC_K1 = iota
	PCLOC_K2
	PCLOC_R1
	PCLOC_R2
	PCLOC_B1
	PCLOC_B2
	PCLOC_L1
	PCLOC_L2
	PCLOC_L3
	PCLOC_L4
	PCLOC_START = 0
	PCLOC_END   = 10 //この数を含まない
	PCLOC_SIZE  = 10
)

// Hand piece type (先後付きの持ち駒の種類)
// 持ち駒を打つときに利用。 0～15
const (
	HAND_K1        = 0
	HAND_R1        = 1 // 先手飛打
	HAND_B1        = 2
	HAND_G1        = 3
	HAND_S1        = 4
	HAND_N1        = 5
	HAND_L1        = 6
	HAND_P1        = 7
	HAND_K2        = 8
	HAND_R2        = 9
	HAND_B2        = 10
	HAND_G2        = 11
	HAND_S2        = 12
	HAND_N2        = 13
	HAND_L2        = 14
	HAND_P2        = 15
	HAND_SIZE      = 16
	HAND_TYPE_SIZE = 8
	HAND_IDX_START = HAND_K1
	HAND_IDX_END   = HAND_SIZE // この数を含まない
)

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

// 駒
const (
	PIECE_EMPTY = iota
	PIECE_K1
	PIECE_R1
	PIECE_B1
	PIECE_G1
	PIECE_S1
	PIECE_N1
	PIECE_L1
	PIECE_P1
	PIECE_PR1
	PIECE_PB1
	PIECE_PS1
	PIECE_PN1
	PIECE_PL1
	PIECE_PP1
	PIECE_K2
	PIECE_R2
	PIECE_B2
	PIECE_G2
	PIECE_S2
	PIECE_N2
	PIECE_L2
	PIECE_P2
	PIECE_PR2
	PIECE_PB2
	PIECE_PS2
	PIECE_PN2
	PIECE_PL2
	PIECE_PP2
)

// 持ち駒
var HandIndexToPiece = [HAND_SIZE]l09.Piece{
	PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1,
	PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2}

// File - マス番号から筋（列）を取り出します
func File(sq Square) Square {
	return sq / 10 % 10
}

// Rank - マス番号から段（行）を取り出します
func Rank(sq Square) Square {
	return sq % 10
}

// Promote - 成ります
func Promote(piece l09.Piece) l09.Piece {
	switch piece {
	case PIECE_EMPTY, PIECE_K1, PIECE_G1, PIECE_PR1, PIECE_PB1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1,
		PIECE_K2, PIECE_G2, PIECE_PR2, PIECE_PB2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2: // 成らずにそのまま返す駒
		return piece
	case PIECE_R1:
		return PIECE_PR1
	case PIECE_B1:
		return PIECE_PB1
	case PIECE_S1:
		return PIECE_PS1
	case PIECE_N1:
		return PIECE_PN1
	case PIECE_L1:
		return PIECE_PL1
	case PIECE_P1:
		return PIECE_PP1
	case PIECE_R2:
		return PIECE_PR2
	case PIECE_B2:
		return PIECE_PB2
	case PIECE_S2:
		return PIECE_PS2
	case PIECE_N2:
		return PIECE_PN2
	case PIECE_L2:
		return PIECE_PL2
	case PIECE_P2:
		return PIECE_PP2
	default:
		panic(fmt.Errorf("error: unknown piece=[%d]", piece))
	}
}

// Demote - 成っている駒を、成っていない駒に戻します
func Demote(piece l09.Piece) l09.Piece {
	switch piece {
	case PIECE_EMPTY, PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1,
		PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2: // 裏返さずにそのまま返す駒
		return piece
	case PIECE_PR1:
		return PIECE_R1
	case PIECE_PB1:
		return PIECE_B1
	case PIECE_PS1:
		return PIECE_S1
	case PIECE_PN1:
		return PIECE_N1
	case PIECE_PL1:
		return PIECE_L1
	case PIECE_PP1:
		return PIECE_P1
	case PIECE_PR2:
		return PIECE_R2
	case PIECE_PB2:
		return PIECE_B2
	case PIECE_PS2:
		return PIECE_S2
	case PIECE_PN2:
		return PIECE_N2
	case PIECE_PL2:
		return PIECE_L2
	case PIECE_PP2:
		return PIECE_P2
	default:
		panic(fmt.Errorf("error: unknown piece=[%d]", piece))
	}
}

// Who - 駒が先手か後手か空升かを返します
func Who(piece l09.Piece) Phase {
	switch piece {
	case PIECE_EMPTY: // 空きマス
		return ZEROTH
	case PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1, PIECE_PR1, PIECE_PB1, PIECE_PS1, PIECE_PN1, PIECE_PL1, PIECE_PP1:
		return FIRST
	case PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2, PIECE_PR2, PIECE_PB2, PIECE_PS2, PIECE_PN2, PIECE_PL2, PIECE_PP2:
		return SECOND
	default:
		panic(fmt.Errorf("error: unknown piece=[%d]", piece))
	}
}

// PieceFrom - 文字列からPieceを作成
func PieceFrom(piece string) l09.Piece {
	switch piece {
	case "":
		return PIECE_EMPTY
	case "K":
		return PIECE_K1
	case "R":
		return PIECE_R1
	case "B":
		return PIECE_B1
	case "G":
		return PIECE_G1
	case "S":
		return PIECE_S1
	case "N":
		return PIECE_N1
	case "L":
		return PIECE_L1
	case "P":
		return PIECE_P1
	case "+R":
		return PIECE_PR1
	case "+B":
		return PIECE_PB1
	case "+S":
		return PIECE_PS1
	case "+N":
		return PIECE_PN1
	case "+L":
		return PIECE_PL1
	case "+P":
		return PIECE_PP1
	case "k":
		return PIECE_K2
	case "r":
		return PIECE_R2
	case "b":
		return PIECE_B2
	case "g":
		return PIECE_G2
	case "s":
		return PIECE_S2
	case "n":
		return PIECE_N2
	case "l":
		return PIECE_L2
	case "p":
		return PIECE_P2
	case "+r":
		return PIECE_PR2
	case "+b":
		return PIECE_PB2
	case "+s":
		return PIECE_PS2
	case "+n":
		return PIECE_PN2
	case "+l":
		return PIECE_PL2
	case "+p":
		return PIECE_PP2
	default:
		panic(fmt.Errorf("unknown piece=[%s]", piece))
	}
}

// ToPieceCode - 文字列
func ToPieceCode(pc l09.Piece) string {
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

// Position - 局面
type Position struct {
	// Go言語で列挙型めんどくさいんで文字列で（＾～＾）
	// [19] は １九、 [91] は ９一（＾～＾）反時計回りに９０°回転した将棋盤の状態で入ってるぜ（＾～＾）想像しろだぜ（＾～＾）
	Board [BOARD_SIZE]l09.Piece
	// 駒の場所
	// [0]先手玉 [1]後手玉 [2:3]飛 [4:5]角 [6:9]香
	PieceLocations [PCLOC_SIZE]Square
	// 持ち駒の数だぜ（＾～＾）玉もある（＾～＾） K, R, B, G, S, N, L, P, k, r, b, g, s, n, l, p
	Hands1 [HAND_SIZE]int

	// 現局面の手番から見た駒得評価値
	MaterialValue Value
}

func NewPosition() *Position {
	var pPos = new(Position)

	pPos.Board = [BOARD_SIZE]l09.Piece{
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
	}

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	return pPos
}

// SetToStartpos - 初期局面にします。利きの計算はまだ行っていません。
func (pPos *Position) SetToStartpos() {
	// 初期局面にします
	pPos.Board = [BOARD_SIZE]l09.Piece{
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_L2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_L1,
		PIECE_EMPTY, PIECE_N2, PIECE_B2, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_R1, PIECE_N1,
		PIECE_EMPTY, PIECE_S2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_S1,
		PIECE_EMPTY, PIECE_G2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_G1,
		PIECE_EMPTY, PIECE_K2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_K1,
		PIECE_EMPTY, PIECE_G2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_G1,
		PIECE_EMPTY, PIECE_S2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_S1,
		PIECE_EMPTY, PIECE_N2, PIECE_R2, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_B1, PIECE_N1,
		PIECE_EMPTY, PIECE_L2, PIECE_EMPTY, PIECE_P2, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_P1, PIECE_EMPTY, PIECE_L1,
	}
	pPos.PieceLocations = [PCLOC_SIZE]Square{59, 51, 28, 82, 22, 88, 11, 19, 91, 99}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (pPos *Position) GetPieceLocation(index int) Square {
	return pPos.PieceLocations[index]
}

// ClearBoard - 駒を置いていな状態でリセットします
func (pPos *Position) ClearBoard() {
	pPos.Board = [BOARD_SIZE]l09.Piece{
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
		PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY, PIECE_EMPTY,
	}

	// 飛角香が存在しないので、仮に 0 を入れてるぜ（＾～＾）
	pPos.PieceLocations = [PCLOC_SIZE]Square{SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY, SQUARE_EMPTY}

	// 持ち駒の数
	pPos.Hands1 = [HAND_SIZE]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

// Homo - 移動元と移動先の駒を持つプレイヤーが等しければ真。移動先が空なら偽
// 持ち駒は指定してはいけません。
func (pPos *Position) Homo(from Square, to Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return Who(pPos.Board[from]) == Who(pPos.Board[to])
}

// Hetero - 移動元と移動先の駒を持つプレイヤーが異なれば真。移動先が空マスでも真
// 持ち駒は指定してはいけません。
// Homo の逆だぜ（＾～＾）片方ありゃいいんだけど（＾～＾）
func (pPos *Position) Hetero(from Square, to Square) bool {
	// fmt.Printf("Debug: from=%d to=%d\n", from, to)
	return Who(pPos.Board[from]) != Who(pPos.Board[to])
}

// IsEmptySq - 空きマスなら真。持ち駒は偽
func (pPos *Position) IsEmptySq(sq Square) bool {
	if sq > 99 {
		return false
	}
	return pPos.Board[sq] == PIECE_EMPTY
}

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase Phase) (Move, error) {
	var len = len(command)
	var hand_sq = SQUARE_EMPTY

	var from Square
	var to Square
	var pro = false

	// file
	switch ch := command[*i]; ch {
	case 'R':
		hand_sq = SQ_R1
	case 'B':
		hand_sq = SQ_B1
	case 'G':
		hand_sq = SQ_G1
	case 'S':
		hand_sq = SQ_S1
	case 'N':
		hand_sq = SQ_N1
	case 'L':
		hand_sq = SQ_L1
	case 'P':
		hand_sq = SQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand_sq != SQUARE_EMPTY {
		*i += 1
		switch phase {
		case FIRST:
			from = hand_sq
		case SECOND:
			from = hand_sq + HAND_TYPE_SIZE
		default:
			return *new(Move), fmt.Errorf("fatal: unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(Move), fmt.Errorf("fatal: not *")
		}
		*i += 1
		count = 1
	}

	// file, rank
	for count < 2 {
		switch ch := command[*i]; ch {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			*i += 1
			file, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}

			var rank int
			switch ch2 := command[*i]; ch2 {
			case 'a':
				rank = 1
			case 'b':
				rank = 2
			case 'c':
				rank = 3
			case 'd':
				rank = 4
			case 'e':
				rank = 5
			case 'f':
				rank = 6
			case 'g':
				rank = 7
			case 'h':
				rank = 8
			case 'i':
				rank = 9
			default:
				return *new(Move), fmt.Errorf("fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := Square(file*10 + rank)
			if count == 0 {
				from = sq
			} else if count == 1 {
				to = sq
			} else {
				return *new(Move), fmt.Errorf("fatal: Unknown count='%c'", count)
			}
		default:
			return *new(Move), fmt.Errorf("fatal: Unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pro = true
	}

	return NewMove(from, to, pro), nil
}

// From - 筋と段からマス番号を作成します
func SquareFrom(file Square, rank Square) Square {
	return Square(file*10 + rank)
}

// OnHands - 持ち駒なら真
func OnHands(sq Square) bool {
	return SQ_HAND_START <= sq && sq < SQ_HAND_END
}

// OnBoard - 盤上なら真
func OnBoard(sq Square) bool {
	return 10 < sq && sq < 100 && File(sq) != 0 && Rank(sq) != 0
}
