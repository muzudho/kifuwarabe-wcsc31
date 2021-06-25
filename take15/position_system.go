package take15

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
		panic(fmt.Errorf("Unknown piece=%d", pc))
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
		panic(fmt.Errorf("Unknown piece=[%s]", piece))
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
			panic(fmt.Errorf("Unknown pieceType=%d", pieceType))
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
			panic(fmt.Errorf("Unknown pieceType=%d", pieceType))
		}
	default:
		panic(fmt.Errorf("Unknown phase=%d", phase))
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

// DoMove - 一手指すぜ（＾～＾）
func (pBrain *Brain) DoMove(pPos *Position, move Move) {
	before_move_phase := pBrain.PPosSys.GetPhase()

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	// 移動元マス、移動先マス、成りの有無
	from, to, pro := move.Destructure()
	if pPos.IsEmptySq(from) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", from)
	}
	var cap_src_sq Square
	var cap_dst_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.PPosSys.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	AddControlRook(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], -1, from)
	AddControlBishop(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], -1, from)
	AddControlLance(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], -1, from)

	// まず、打かどうかで処理を分けます
	sq_drop := from
	var piece Piece
	switch from {
	case SQ_K1:
		piece = PIECE_K1
	case SQ_R1:
		piece = PIECE_R1
	case SQ_B1:
		piece = PIECE_B1
	case SQ_G1:
		piece = PIECE_G1
	case SQ_S1:
		piece = PIECE_S1
	case SQ_N1:
		piece = PIECE_N1
	case SQ_L1:
		piece = PIECE_L1
	case SQ_P1:
		piece = PIECE_P1
	case SQ_K2:
		piece = PIECE_K2
	case SQ_R2:
		piece = PIECE_R2
	case SQ_B2:
		piece = PIECE_B2
	case SQ_G2:
		piece = PIECE_G2
	case SQ_S2:
		piece = PIECE_S2
	case SQ_N2:
		piece = PIECE_N2
	case SQ_L2:
		piece = PIECE_L2
	case SQ_P2:
		piece = PIECE_P2
	default:
		// Not drop
		sq_drop = SQUARE_EMPTY
	}

	if sq_drop != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands1[sq_drop-SQ_HAND_START] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		mov_piece_type = What(piece)

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		ValidateThereArePieceIn(pPos, to)
		var pCB *ControlBoard
		if pBrain.PPosSys.BuildType == BUILD_DEV {
			pCB = ControllBoardFromPhase(before_move_phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB = ControllBoardFromPhase(before_move_phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[to]
		if captured != PIECE_EMPTY {
			pieceType := What(captured)
			switch pieceType {
			case PIECE_TYPE_R, PIECE_TYPE_PR, PIECE_TYPE_B, PIECE_TYPE_PB, PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				piece := pPos.Board[to]

				// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
				ValidateThereArePieceIn(pPos, to)
				phase := Who(piece)
				var pCB *ControlBoard
				if pBrain.PPosSys.BuildType == BUILD_DEV {
					pCB = ControllBoardFromPhase(phase,
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_CAPTURED],
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_CAPTURED])
				} else {
					pCB = ControllBoardFromPhase(phase,
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
				}
				pCB.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
			}
			cap_piece_type = What(captured)
			cap_src_sq = to

			// 駒得評価値の計算（＾ｑ＾）
			material_val := EvalMaterial(captured)
			pPos.MaterialValue += material_val
		}

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase := Who(piece)
		var pCB1 *ControlBoard
		if pBrain.PPosSys.BuildType == BUILD_DEV {
			pCB1 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_REMOVE],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_REMOVE])
		} else {
			pCB1 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		// 元位置の駒の利きを除去
		pCB1.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, from)), from, -1)

		// 行き先の駒の上書き
		if pro {
			// 駒を成りに変換します
			pPos.Board[to] = Promote(pPos.Board[from])
		} else {
			pPos.Board[to] = pPos.Board[from]
		}
		mov_piece_type = What(pPos.Board[to])
		// 元位置の駒を削除してから、移動先の駒の利きを追加
		pPos.Board[from] = PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase = Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB2 *ControlBoard
		if pBrain.PPosSys.BuildType == BUILD_DEV {
			pCB2 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB2 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB2.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)

		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			cap_dst_sq = SQ_K2
		case PIECE_R1, PIECE_PR1:
			cap_dst_sq = SQ_R2
		case PIECE_B1, PIECE_PB1:
			cap_dst_sq = SQ_B2
		case PIECE_G1:
			cap_dst_sq = SQ_G2
		case PIECE_S1, PIECE_PS1:
			cap_dst_sq = SQ_S2
		case PIECE_N1, PIECE_PN1:
			cap_dst_sq = SQ_N2
		case PIECE_L1, PIECE_PL1:
			cap_dst_sq = SQ_L2
		case PIECE_P1, PIECE_PP1:
			cap_dst_sq = SQ_P2
		case PIECE_K2: // First player win
			cap_dst_sq = SQ_K1
		case PIECE_R2, PIECE_PR2:
			cap_dst_sq = SQ_R1
		case PIECE_B2, PIECE_PB2:
			cap_dst_sq = SQ_B1
		case PIECE_G2:
			cap_dst_sq = SQ_G1
		case PIECE_S2, PIECE_PS2:
			cap_dst_sq = SQ_S1
		case PIECE_N2, PIECE_PN2:
			cap_dst_sq = SQ_N1
		case PIECE_L2, PIECE_PL2:
			cap_dst_sq = SQ_L1
		case PIECE_P2, PIECE_PP2:
			cap_dst_sq = SQ_P1
		default:
			fmt.Printf("Error: Unknown captured=[%d]", captured)
		}

		if cap_dst_sq != SQUARE_EMPTY {
			pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex] = captured
			pPos.Hands1[cap_dst_sq-SQ_HAND_START] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex] = PIECE_EMPTY
		}
	}

	// DoMoveでフェーズを１つ進めます
	pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex] = move
	pBrain.PPosSys.OffsetMovesIndex += 1
	pBrain.PPosSys.FlipPhase()

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []Square{from, cap_src_sq}
	dst_sq_list := []Square{to, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case PIECE_TYPE_K:
			if j == 0 {
				switch before_move_phase {
				case FIRST:
					pPos.PieceLocations[PCLOC_K1] = dst_sq_list[j]
				case SECOND:
					pPos.PieceLocations[PCLOC_K2] = dst_sq_list[j]
				default:
					panic(fmt.Errorf("Unknown before_move_phase=%d", before_move_phase))
				}
			} else {
				// 取った時
				switch before_move_phase {
				case FIRST:
					// 相手玉
					pPos.PieceLocations[PCLOC_K2] = dst_sq_list[j]
				case SECOND:
					pPos.PieceLocations[PCLOC_K1] = dst_sq_list[j]
				default:
					panic(fmt.Errorf("Unknown before_move_phase=%d", before_move_phase))
				}
			}
		case PIECE_TYPE_R, PIECE_TYPE_PR:
			for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case PIECE_TYPE_B, PIECE_TYPE_PB:
			for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし動かした駒を除きます
	AddControlLance(
		pPos, pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], 1, to)
	AddControlBishop(
		pPos, pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], 1, to)
	AddControlRook(
		pPos, pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], 1, to)

	pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.PPosSys.BuildType)
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pBrain *Brain) UndoMove(pPos *Position) {

	// G.StderrChat.Trace(pBrain.PPosSys.Sprint())

	if pBrain.PPosSys.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY

	// 先に 手目 を１つ戻すぜ（＾～＾）UndoMoveでフェーズもひっくり返すぜ（＾～＾）
	pBrain.PPosSys.OffsetMovesIndex -= 1
	move := pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex]
	// next_phase := pBrain.PPosSys.GetPhase()
	pBrain.PPosSys.FlipPhase()

	from, to, pro := move.Destructure()

	// 利きの差分テーブルをクリアー（＾～＾）
	pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.PPosSys.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlRook(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], -1, to)
	AddControlBishop(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], -1, to)
	AddControlLance(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], -1, to)

	// 打かどうかで分けます
	switch from {
	case SQ_K1, SQ_R1, SQ_B1, SQ_G1, SQ_S1, SQ_N1, SQ_L1, SQ_P1, SQ_K2, SQ_R2, SQ_B2, SQ_G2, SQ_S2, SQ_N2, SQ_L2, SQ_P2:
		// 打なら
		drop := from
		// 行き先から駒を除去します
		mov_piece_type = What(pPos.Board[to])

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB3 *ControlBoard
		if pBrain.PPosSys.BuildType == BUILD_DEV {
			pCB3 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB3 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB3.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
		pPos.Board[to] = PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands1[drop-SQ_HAND_START] += 1
	default:
		// 打でないなら

		// 行き先に進んでいた自駒の利きの除去
		mov_piece_type = What(pPos.Board[to])

		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB4 *ControlBoard
		if pBrain.PPosSys.BuildType == BUILD_DEV {
			pCB4 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB4 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB4.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)

		// 自駒を移動元へ戻します
		if pro {
			// 成りを元に戻します
			pPos.Board[from] = Demote(pPos.Board[to])
		} else {
			pPos.Board[from] = pPos.Board[to]
		}

		pPos.Board[to] = PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase = Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB5 *ControlBoard
		if pBrain.PPosSys.BuildType == BUILD_DEV {
			pCB5 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_REMOVE],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_REMOVE])
		} else {
			pCB5 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		// 元の場所に戻した自駒の利きを復元します
		pCB5.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, from)), from, 1)
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch mov_piece_type {
	case PIECE_TYPE_K:
		// 玉を動かした
		switch pBrain.PPosSys.phase { // next_phase
		case FIRST:
			pPos.PieceLocations[PCLOC_K1] = from
		case SECOND:
			pPos.PieceLocations[PCLOC_K2] = from
		default:
			panic(fmt.Errorf("Unknown pBrain.PPosSys.phase=%d", pBrain.PPosSys.phase))
		}
	case PIECE_TYPE_R, PIECE_TYPE_PR:
		for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case PIECE_TYPE_B, PIECE_TYPE_PB:
		for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlLance(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], 1, from)
	AddControlBishop(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], 1, from)
	AddControlRook(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], 1, from)

	pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.PPosSys.BuildType)

	// 取った駒を戻すぜ（＾～＾）
	pBrain.undoCapture(pPos)
}

// undoCapture - 取った駒を戻すぜ（＾～＾）
func (pBrain *Brain) undoCapture(pPos *Position) {
	// G.StderrChat.Trace(pBrain.PPosSys.Sprint())

	// 取った駒だぜ（＾～＾）
	cap_piece_type := PIECE_TYPE_EMPTY

	// 手目もフェーズもすでに１つ戻っているとするぜ（＾～＾）
	move := pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex]

	// 取った駒
	captured := pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex]
	// fmt.Printf("Debug: CapturedPiece=%s\n", captured.ToCode())

	// 取った駒に関係するのは行き先だけ（＾～＾）
	from, to, _ := move.Destructure()
	// fmt.Printf("Debug: to=%d\n", to)

	var hand_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.PPosSys.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlRook(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], -1, to)
	AddControlBishop(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], -1, to)
	AddControlLance(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], -1, to)

	// 打かどうかで分けます
	switch from {
	case SQ_K1, SQ_R1, SQ_B1, SQ_G1, SQ_S1, SQ_N1, SQ_L1, SQ_P1, SQ_K2, SQ_R2, SQ_B2, SQ_G2, SQ_S2, SQ_N2, SQ_L2, SQ_P2:
		// 打で取れる駒はないぜ（＾～＾）
		// fmt.Printf("Debug: Drop from=%d\n", from)
	default:
		// 打でないなら
		// fmt.Printf("Debug: Not drop from=%d\n", from)

		// 取った相手の駒があれば、自分の駒台から下ろします
		switch captured {
		case PIECE_EMPTY: // Ignored
		case PIECE_K1: // Second player win
			hand_sq = SQ_K2
		case PIECE_R1, PIECE_PR1:
			hand_sq = SQ_R2
		case PIECE_B1, PIECE_PB1:
			hand_sq = SQ_B2
		case PIECE_G1:
			hand_sq = SQ_G2
		case PIECE_S1, PIECE_PS1:
			hand_sq = SQ_S2
		case PIECE_N1, PIECE_PN1:
			hand_sq = SQ_N2
		case PIECE_L1, PIECE_PL1:
			hand_sq = SQ_L2
		case PIECE_P1, PIECE_PP1:
			hand_sq = SQ_P2
		case PIECE_K2: // First player win
			hand_sq = SQ_K1
		case PIECE_R2, PIECE_PR2:
			hand_sq = SQ_R1
		case PIECE_B2, PIECE_PB2:
			hand_sq = SQ_B1
		case PIECE_G2:
			hand_sq = SQ_G1
		case PIECE_S2, PIECE_PS2:
			hand_sq = SQ_S1
		case PIECE_N2, PIECE_PN2:
			hand_sq = SQ_N1
		case PIECE_L2, PIECE_PL2:
			hand_sq = SQ_L1
		case PIECE_P2, PIECE_PP2:
			hand_sq = SQ_P1
		default:
			fmt.Printf("Error: Unknown captured=[%d]", captured)
		}

		// fmt.Printf("Debug: hand_sq=%d\n", hand_sq)

		if hand_sq != SQUARE_EMPTY {
			pPos.Hands1[hand_sq-SQ_HAND_START] -= 1

			// 取っていた駒を行き先に戻します
			cap_piece_type = What(captured)
			pPos.Board[to] = captured

			ValidateThereArePieceIn(pPos, to)
			// fmt.Printf("Debug: ph=%d\n", ph)
			var pCB6 *ControlBoard
			if pBrain.PPosSys.BuildType == BUILD_DEV {
				pCB6 = ControllBoardFromPhase(pBrain.PPosSys.phase,
					pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_CAPTURED],
					pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_CAPTURED])
			} else {
				pCB6 = ControllBoardFromPhase(pBrain.PPosSys.phase,
					pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2],
					pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1])
			}
			// 取った駒は盤上になかったので、ここで利きを復元させます
			// 行き先にある取られていた駒の利きの復元
			pCB6.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)
		}
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch cap_piece_type {
	case PIECE_TYPE_K:
		// 玉を取っていた
		switch pBrain.PPosSys.phase { // next_phase
		case FIRST:
			// 後手の玉
			pPos.PieceLocations[PCLOC_K2] = to
		case SECOND:
			// 先手の玉
			pPos.PieceLocations[PCLOC_K1] = to
		default:
			panic(fmt.Errorf("Unknown pBrain.PPosSys.phase=%d", pBrain.PPosSys.phase))
		}
	case PIECE_TYPE_R, PIECE_TYPE_PR:
		for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case PIECE_TYPE_B, PIECE_TYPE_PB:
		for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	}

	// 駒得評価値の計算（＾ｑ＾）
	material_val := EvalMaterial(captured)
	pPos.MaterialValue -= material_val

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlLance(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], 1, from)
	AddControlBishop(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], 1, from)
	AddControlRook(
		pPos,
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], 1, from)

	pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.PPosSys.BuildType)
}
