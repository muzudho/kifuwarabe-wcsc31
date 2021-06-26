package take16

import (
	"fmt"
	"strconv"

	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// 電竜戦が一番長いだろ（＾～＾）
const MOVES_SIZE = 512

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

// FlipPhase - 先後を反転します
func FlipPhase(phase p.Phase) p.Phase {
	return phase%2 + 1
}

// From - 筋と段からマス番号を作成します
func SquareFrom(file p.Square, rank p.Square) p.Square {
	return p.Square(file*10 + rank)
}

// OnHands - 持ち駒なら真
func OnHands(sq p.Square) bool {
	return SQ_HAND_START <= sq && sq < SQ_HAND_END
}

// OnBoard - 盤上なら真
func OnBoard(sq p.Square) bool {
	return 10 < sq && sq < 100 && p.File(sq) != 0 && p.Rank(sq) != 0
}

// [0], [1]
const PHASE_ARRAY_SIZE = 2

// PieceFrom - 文字列
func PieceFrom(piece string) p.Piece {
	switch piece {
	case "":
		return p.PIECE_EMPTY
	case "K":
		return p.PIECE_K1
	case "R":
		return p.PIECE_R1
	case "B":
		return p.PIECE_B1
	case "G":
		return p.PIECE_G1
	case "S":
		return p.PIECE_S1
	case "N":
		return p.PIECE_N1
	case "L":
		return p.PIECE_L1
	case "P":
		return p.PIECE_P1
	case "+R":
		return p.PIECE_PR1
	case "+B":
		return p.PIECE_PB1
	case "+S":
		return p.PIECE_PS1
	case "+N":
		return p.PIECE_PN1
	case "+L":
		return p.PIECE_PL1
	case "+P":
		return p.PIECE_PP1
	case "k":
		return p.PIECE_K2
	case "r":
		return p.PIECE_R2
	case "b":
		return p.PIECE_B2
	case "g":
		return p.PIECE_G2
	case "s":
		return p.PIECE_S2
	case "n":
		return p.PIECE_N2
	case "l":
		return p.PIECE_L2
	case "p":
		return p.PIECE_P2
	case "+r":
		return p.PIECE_PR2
	case "+b":
		return p.PIECE_PB2
	case "+s":
		return p.PIECE_PS2
	case "+n":
		return p.PIECE_PN2
	case "+l":
		return p.PIECE_PL2
	case "+p":
		return p.PIECE_PP2
	default:
		panic(G.Log.Fatal("Unknown piece=[%s]", piece))
	}
}

// PieceFromPhPt - 駒作成。空マスは作れません
func PieceFromPhPt(phase p.Phase, pieceType PieceType) p.Piece {
	switch phase {
	case p.FIRST:
		switch pieceType {
		case PIECE_TYPE_K:
			return p.PIECE_K1
		case PIECE_TYPE_R:
			return p.PIECE_R1
		case PIECE_TYPE_B:
			return p.PIECE_B1
		case PIECE_TYPE_G:
			return p.PIECE_G1
		case PIECE_TYPE_S:
			return p.PIECE_S1
		case PIECE_TYPE_N:
			return p.PIECE_N1
		case PIECE_TYPE_L:
			return p.PIECE_L1
		case PIECE_TYPE_P:
			return p.PIECE_P1
		case PIECE_TYPE_PR:
			return p.PIECE_PR1
		case PIECE_TYPE_PB:
			return p.PIECE_PB1
		case PIECE_TYPE_PS:
			return p.PIECE_PS1
		case PIECE_TYPE_PN:
			return p.PIECE_PN1
		case PIECE_TYPE_PL:
			return p.PIECE_PL1
		case PIECE_TYPE_PP:
			return p.PIECE_PP1
		default:
			panic(G.Log.Fatal("Unknown pieceType=%d", pieceType))
		}
	case p.SECOND:
		switch pieceType {
		case PIECE_TYPE_K:
			return p.PIECE_K2
		case PIECE_TYPE_R:
			return p.PIECE_R2
		case PIECE_TYPE_B:
			return p.PIECE_B2
		case PIECE_TYPE_G:
			return p.PIECE_G2
		case PIECE_TYPE_S:
			return p.PIECE_S2
		case PIECE_TYPE_N:
			return p.PIECE_N2
		case PIECE_TYPE_L:
			return p.PIECE_L2
		case PIECE_TYPE_P:
			return p.PIECE_P2
		case PIECE_TYPE_PR:
			return p.PIECE_PR2
		case PIECE_TYPE_PB:
			return p.PIECE_PB2
		case PIECE_TYPE_PS:
			return p.PIECE_PS2
		case PIECE_TYPE_PN:
			return p.PIECE_PN2
		case PIECE_TYPE_PL:
			return p.PIECE_PL2
		case PIECE_TYPE_PP:
			return p.PIECE_PP2
		default:
			panic(G.Log.Fatal("Unknown pieceType=%d", pieceType))
		}
	default:
		panic(G.Log.Fatal("Unknown phase=%d", phase))
	}
}

var HandPieceMap1 = [p.HAND_SIZE]p.Piece{
	p.PIECE_K1, p.PIECE_R1, p.PIECE_B1, p.PIECE_G1, p.PIECE_S1, p.PIECE_N1, p.PIECE_L1, p.PIECE_P1,
	p.PIECE_K2, p.PIECE_R2, p.PIECE_B2, p.PIECE_G2, p.PIECE_S2, p.PIECE_N2, p.PIECE_L2, p.PIECE_P2}

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
	PPosition [POS_LAYER_SIZE]*p.Position

	// 先手が1、後手が2（＾～＾）
	phase p.Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [MOVES_SIZE]p.Piece
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)
	pPosSys.BuildType = BUILD_DEV

	pPosSys.PPosition = [POS_LAYER_SIZE]*p.Position{p.NewPosition(), p.NewPosition(), p.NewPosition(), p.NewPosition()}

	pPosSys.resetPosition()
	return pPosSys
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() p.Phase {
	return pPosSys.phase
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	// 先手の局面
	pPosSys.phase = p.FIRST
	// 何手目か
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPosSys.CapturedList = [MOVES_SIZE]p.Piece{}
}

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase p.Phase) (Move, error) {
	var len = len(command)
	var hand_sq = p.SQUARE_EMPTY

	var from p.Square
	var to p.Square
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

	if hand_sq != p.SQUARE_EMPTY {
		*i += 1
		switch phase {
		case p.FIRST:
			from = hand_sq
		case p.SECOND:
			from = hand_sq + p.HAND_TYPE_SIZE
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

			sq := p.Square(file*10 + rank)
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
