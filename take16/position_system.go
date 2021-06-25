package take16

import (
	"fmt"
	"strconv"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

// 00～99
const BOARD_SIZE = 100

// 盤レイヤー・インデックス型
type PosLayerT int

const (
	POS_LAYER_MAIN  = PosLayerT(0)
	POS_LAYER_COPY  = PosLayerT(1) // テスト用
	POS_LAYER_DIFF1 = PosLayerT(2) // テスト用
	POS_LAYER_DIFF2 = PosLayerT(3) // テスト用
	POS_LAYER_SIZE  = 4
)

// position sfen の盤のスペース数に使われますN
var OneDigitNumbers = [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

// 1:先手 2:後手
type Phase byte

// FlipPhase - 先後を反転します
func FlipPhase(phase Phase) Phase {
	return phase%2 + 1
}

// マス番号 00～99,100～113
type Square uint32

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

// マス番号を指定しないことを意味するマス番号
const SQUARE_EMPTY = Square(0)

const (
	// 空マス
	ZEROTH = Phase(0)
	// 先手
	FIRST = Phase(1)
	// 後手
	SECOND = Phase(2)
)

// [0], [1]
const PHASE_ARRAY_SIZE = 2

// 先後付きの駒
type Piece uint8

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

// ToCode - 文字列
func (pc Piece) ToCode() string {
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
		panic(G.Log.Fatal("Unknown piece=%d", pc))
	}
}

// PieceFrom - 文字列
func PieceFrom(piece string) Piece {
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
		panic(G.Log.Fatal("Unknown piece=[%s]", piece))
	}
}

// PieceFromPhPt - 駒作成。空マスは作れません
func PieceFromPhPt(phase Phase, pieceType PieceType) Piece {
	switch phase {
	case FIRST:
		switch pieceType {
		case PIECE_TYPE_K:
			return PIECE_K1
		case PIECE_TYPE_R:
			return PIECE_R1
		case PIECE_TYPE_B:
			return PIECE_B1
		case PIECE_TYPE_G:
			return PIECE_G1
		case PIECE_TYPE_S:
			return PIECE_S1
		case PIECE_TYPE_N:
			return PIECE_N1
		case PIECE_TYPE_L:
			return PIECE_L1
		case PIECE_TYPE_P:
			return PIECE_P1
		case PIECE_TYPE_PR:
			return PIECE_PR1
		case PIECE_TYPE_PB:
			return PIECE_PB1
		case PIECE_TYPE_PS:
			return PIECE_PS1
		case PIECE_TYPE_PN:
			return PIECE_PN1
		case PIECE_TYPE_PL:
			return PIECE_PL1
		case PIECE_TYPE_PP:
			return PIECE_PP1
		default:
			panic(G.Log.Fatal("Unknown pieceType=%d", pieceType))
		}
	case SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return PIECE_K2
		case PIECE_TYPE_R:
			return PIECE_R2
		case PIECE_TYPE_B:
			return PIECE_B2
		case PIECE_TYPE_G:
			return PIECE_G2
		case PIECE_TYPE_S:
			return PIECE_S2
		case PIECE_TYPE_N:
			return PIECE_N2
		case PIECE_TYPE_L:
			return PIECE_L2
		case PIECE_TYPE_P:
			return PIECE_P2
		case PIECE_TYPE_PR:
			return PIECE_PR2
		case PIECE_TYPE_PB:
			return PIECE_PB2
		case PIECE_TYPE_PS:
			return PIECE_PS2
		case PIECE_TYPE_PN:
			return PIECE_PN2
		case PIECE_TYPE_PL:
			return PIECE_PL2
		case PIECE_TYPE_PP:
			return PIECE_PP2
		default:
			panic(G.Log.Fatal("Unknown pieceType=%d", pieceType))
		}
	default:
		panic(G.Log.Fatal("Unknown phase=%d", phase))
	}
}

const (
	// 持ち駒を打つ 0～15 (Index)
	HAND_K1_IDX    = 0
	HAND_R1_IDX    = 1 // 先手飛打
	HAND_B1_IDX    = 2
	HAND_G1_IDX    = 3
	HAND_S1_IDX    = 4
	HAND_N1_IDX    = 5
	HAND_L1_IDX    = 6
	HAND_P1_IDX    = 7
	HAND_K2_IDX    = 8
	HAND_R2_IDX    = 9
	HAND_B2_IDX    = 10
	HAND_G2_IDX    = 11
	HAND_S2_IDX    = 12
	HAND_N2_IDX    = 13
	HAND_L2_IDX    = 14
	HAND_P2_IDX    = 15
	HAND_SIZE      = 16
	HAND_TYPE_SIZE = 8
	HAND_IDX_START = HAND_K1_IDX
	HAND_IDX_END   = HAND_SIZE // この数を含まない
)

var HandPieceMap1 = [HAND_SIZE]Piece{
	PIECE_K1, PIECE_R1, PIECE_B1, PIECE_G1, PIECE_S1, PIECE_N1, PIECE_L1, PIECE_P1,
	PIECE_K2, PIECE_R2, PIECE_B2, PIECE_G2, PIECE_S2, PIECE_N2, PIECE_L2, PIECE_P2}

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

// 開発 or リリース モード
type BuildT int

const (
	BUILD_DEV     = BuildT(0)
	BUILD_RELEASE = BuildT(1)
)

// PositionSystem - 局面にいろいろな機能を付けたもの
type PositionSystem struct {
	// 開発モードフラグ。デフォルト値：真。 'usi' コマンドで解除
	BuildType BuildT
	// 局面
	PPosition [POS_LAYER_SIZE]*Position

	// 先手が1、後手が2（＾～＾）
	phase Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [MOVES_SIZE]Piece
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)
	pPosSys.BuildType = BUILD_DEV

	pPosSys.PPosition = [POS_LAYER_SIZE]*Position{NewPosition(), NewPosition(), NewPosition(), NewPosition()}

	pPosSys.resetPosition()
	return pPosSys
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() Phase {
	return pPosSys.phase
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	// 先手の局面
	pPosSys.phase = FIRST
	// 何手目か
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPosSys.CapturedList = [MOVES_SIZE]Piece{}
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
			return *new(Move), fmt.Errorf("Fatal: Unknown phase=%d", phase)
		}

		if command[*i] != '*' {
			return *new(Move), fmt.Errorf("Fatal: not *")
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
				return *new(Move), fmt.Errorf("Fatal: Unknown file or rank. ch2='%c'", ch2)
			}
			*i += 1

			sq := Square(file*10 + rank)
			if count == 0 {
				from = sq
			} else if count == 1 {
				to = sq
			} else {
				return *new(Move), fmt.Errorf("Fatal: Unknown count='%c'", count)
			}
		default:
			return *new(Move), fmt.Errorf("Fatal: Unknown move. ch='%c' i='%d'", ch, *i)
		}

		count += 1
	}

	if *i < len && command[*i] == '+' {
		*i += 1
		pro = true
	}

	return NewMove(from, to, pro), nil
}
