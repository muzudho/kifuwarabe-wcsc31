package take13

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	l02 "github.com/muzudho/kifuwarabe-wcsc31/lesson02"
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l07 "github.com/muzudho/kifuwarabe-wcsc31/take7"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

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
func FlipPhase(phase l03.Phase) l03.Phase {
	return phase%2 + 1
}

// From - 筋と段からマス番号を作成します
func SquareFrom(file l03.Square, rank l03.Square) l03.Square {
	return l03.Square(file*10 + rank)
}

// OnHands - 持ち駒なら真
func OnHands(sq l03.Square) bool {
	return l03.SQ_HAND_START <= sq && sq < l03.SQ_HAND_END
}

// OnBoard - 盤上なら真
func OnBoard(sq l03.Square) bool {
	return 10 < sq && sq < 100 && l03.File(sq) != 0 && l03.Rank(sq) != 0
}

// [0], [1]
const PHASE_ARRAY_SIZE = 2

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

	// 利きボード・システム
	PControlBoardSystem *ControlBoardSystem

	// 先手が1、後手が2（＾～＾）
	phase l03.Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [l02.MOVES_SIZE]l03.Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [l02.MOVES_SIZE]l03.Piece
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
func (pPosSys *PositionSystem) GetPhase() l03.Phase {
	return pPosSys.phase
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	pPosSys.PControlBoardSystem = NewControlBoardSystem()

	// 先手の局面
	pPosSys.phase = l03.FIRST
	// 何手目か
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.Moves = [l02.MOVES_SIZE]l03.Move{}
	// 取った駒のリスト
	pPosSys.CapturedList = [l02.MOVES_SIZE]l03.Piece{}
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pPosSys *PositionSystem) ReadPosition(pPos *Position, command string) {
	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面をセット（＾～＾）
		pPos.clearBoard()
		pPosSys.resetPosition()
		pPos.setToStartpos()
		i = 17

		if i < len && command[i] == ' ' {
			i += 1
		}
		// moves へ続くぜ（＾～＾）

	} else if strings.HasPrefix(command, "position sfen ") {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		pPos.clearBoard()
		pPosSys.resetPosition()
		i = 14
		var rank = l03.Square(1)
		var file = l03.Square(9)

	BoardLoop:
		for {
			promoted := false
			switch pc := command[i]; pc {
			case 'K', 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'k', 'r', 'b', 'g', 's', 'n', 'l', 'p':
				pPos.Board[file*10+rank] = l03.FromCodeToPiece(string(pc))
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = l03.PIECE_EMPTY
					file -= 1
				}
				i += 1
			case '+':
				i += 1
				promoted = true
			case '/':
				file = 9
				rank += 1
				i += 1
			case ' ':
				i += 1
				break BoardLoop
			default:
				panic("Undefined sfen board")
			}

			if promoted {
				switch pc2 := command[i]; pc2 {
				case 'R', 'B', 'S', 'N', 'L', 'P', 'r', 'b', 's', 'n', 'l', 'p':
					pPos.Board[file*10+rank] = l03.FromCodeToPiece("+" + string(pc2))
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.PieceLocations[l07.PCLOC_K1] = l03.Square((file+1)*10 + rank)
			case 'k':
				pPos.PieceLocations[l07.PCLOC_K2] = l03.Square((file+1)*10 + rank)
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i := l07.PCLOC_R1; i < l07.PCLOC_R2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'B', 'b':
				for i := l07.PCLOC_B1; i < l07.PCLOC_B2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'L', 'l':
				for i := l07.PCLOC_L1; i < l07.PCLOC_L4+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			}
		}

		// 手番
		switch command[i] {
		case 'b':
			pPosSys.phase = l03.FIRST
			i += 1
		case 'w':
			pPosSys.phase = l03.SECOND
			i += 1
		default:
			panic("fatal: unknown phase")
		}

		if command[i] != ' ' {
			// 手番の後ろにスペースがない（＾～＾）
			panic("fatal: Nothing space")
		}
		i += 1

		// 持ち駒
		if command[i] == '-' {
			i += 1
			if command[i] != ' ' {
				// 持ち駒 - の後ろにスペースがない（＾～＾）
				panic("fatal: Nothing space after -")
			}
			i += 1
		} else {

			// R なら竜1枚
			// R2 なら竜2枚
			// P10 なら歩10枚。数が2桁になるのは歩だけ（＾～＾）
			// {アルファベット１文字}{数字1～2文字} になっている
			// アルファベットまたは半角スペースを見つけた時点で、以前の取り込み分が確定する
			var handIndex l03.HandIdx = 999 //存在しない数
			var number = 0

		HandLoop:
			for {
				var piece = command[i]

				if unicode.IsLetter(rune(piece)) || piece == ' ' {

					if handIndex == 999 {
						// ループの１週目は無視します

					} else {
						// 数字が書いてなかったら１個
						if number == 0 {
							number = 1
						}

						pPos.Hands1[handIndex] = number
						number = 0

						// 長い利きの駒は位置を覚えておくぜ（＾～＾）
						switch handIndex {
						case l03.HAND_R1, l03.HAND_R2:
							for i := l07.PCLOC_R1; i < l07.PCLOC_R2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == l03.SQ_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = l03.Square(handIndex) + l03.SQ_HAND_START
									break
								}
							}
						case l03.HAND_B1, l03.HAND_B2:
							for i := l07.PCLOC_B1; i < l07.PCLOC_B2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == l03.SQ_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = l03.Square(handIndex) + l03.SQ_HAND_START
									break
								}
							}
						case l03.HAND_L1, l03.HAND_L2:
							for i := l07.PCLOC_L1; i < l07.PCLOC_L4+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == l03.SQ_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = l03.Square(handIndex) + l03.SQ_HAND_START
									break
								}
							}
						}
					}
					i += 1

					var isBreak = false
					var convertAlternativeValue = func(code byte) l03.HandIdx {
						if code == ' ' {
							isBreak = true
							return l03.HAND_SIZE // この値は使いません
						} else {
							panic(App.LogNotEcho.Fatal("fatal: unknown piece=%c", piece))
						}
					}

					handIndex = l03.FromCodeToHandIndex(byte(piece), &convertAlternativeValue)

					if isBreak {
						// ループを抜けます
						break HandLoop
					}

				} else if unicode.IsNumber(rune(piece)) {
					switch piece {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						num, err := strconv.Atoi(string(piece))
						if err != nil {
							panic(err)
						}
						i += 1
						number *= 10
						number += num
					default:
						panic(App.LogNotEcho.Fatal("fatal: Unknown number character=%c", piece))
					}

				} else {
					panic(App.LogNotEcho.Fatal("fatal: unknown piece=%c", piece))
				}
			}
		}

		// 手数
		pPosSys.StartMovesNum = 0
	MovesNumLoop:
		for i < len {
			switch figure := command[i]; figure {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num, err := strconv.Atoi(string(figure))
				if err != nil {
					panic(err)
				}
				i += 1
				pPosSys.StartMovesNum *= 10
				pPosSys.StartMovesNum += num
			case ' ':
				i += 1
				break MovesNumLoop
			default:
				break MovesNumLoop
			}
		}

	} else {
		fmt.Printf("unknown command=[%s]", command)
	}

	// fmt.Printf("command[i:]=[%s]\n", command[i:])

	start_phase := pPosSys.GetPhase()
	if strings.HasPrefix(command[i:], "moves") {
		i += 5

		// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
		for i < len {
			if command[i] != ' ' {
				break
			}
			i += 1

			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			var move, err = l03.ParseMove(command, &i, pPosSys.GetPhase())
			if err != nil {
				fmt.Println(err)
				fmt.Println(SprintBoard(
					pPos,
					pPosSys.phase,
					pPosSys.StartMovesNum,
					pPosSys.OffsetMovesIndex,
					pPosSys.createMovesText()))
				panic(err)
			}
			pPosSys.Moves[pPosSys.OffsetMovesIndex] = move
			pPosSys.OffsetMovesIndex += 1
			pPosSys.FlipPhase()
		}
	}

	if pPosSys.BuildType == BUILD_DEV {
		// 利きの差分テーブルをクリアー（＾～＾）
		pPosSys.PControlBoardSystem.ClearControlDiff(pPosSys.BuildType)
	}

	// 開始局面の利きを計算（＾～＾）
	//fmt.Printf("Debug: 開始局面の利きを計算（＾～＾）\n")
	for sq := l03.Square(11); sq < 100; sq += 1 {
		if l03.File(sq) != 0 && l03.Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				//fmt.Printf("Debug: sq=%d\n", sq)
				// あとですぐクリアーするので、どのレイヤー使ってても関係ないんで、仮で PUTレイヤーを使っているぜ（＾～＾）

				// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
				piece := pPos.Board[sq]
				ValidateThereArePieceIn(pPos, sq)
				phase := l03.Who(piece)
				// fmt.Printf("Debug: ph=%d\n", ph)
				var pCB7 *ControlBoard
				if pPosSys.BuildType == BUILD_DEV {
					pCB7 = ControllBoardFromPhase(phase,
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_PUT],
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_PUT])
				} else {
					pCB7 = ControllBoardFromPhase(phase,
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
				}
				pCB7.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, sq)), sq, 1)
			}
		}
	}
	if pPosSys.BuildType == BUILD_DEV {
		//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
		pPosSys.PControlBoardSystem.MergeControlDiff(pPosSys.BuildType)
	}

	// 読込んだ l03.Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPosSys.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPosSys.OffsetMovesIndex = 0
	pPosSys.phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPosSys.DoMove(pPos, pPosSys.Moves[i])
	}
}

// DoMove - 一手指すぜ（＾～＾）
func (pPosSys *PositionSystem) DoMove(pPos *Position, move l03.Move) {
	before_move_phase := pPosSys.GetPhase()

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := l03.PIECE_TYPE_EMPTY
	cap_piece_type := l03.PIECE_TYPE_EMPTY

	// 移動元マス、移動先マス、成りの有無
	from, to, pro := move.Destructure()
	if pPos.IsEmptySq(from) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", from)
	}
	var cap_src_sq l03.Square
	var cap_dst_sq = l03.SQ_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff(pPosSys.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	AddControlRook(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], -1, from)
	AddControlBishop(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], -1, from)
	AddControlLance(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], -1, from)

	// まず、打かどうかで処理を分けます
	sq_hand := from
	var piece l03.Piece
	switch from {
	case l03.SQ_K1:
		piece = l03.PIECE_K1
	case l03.SQ_R1:
		piece = l03.PIECE_R1
	case l03.SQ_B1:
		piece = l03.PIECE_B1
	case l03.SQ_G1:
		piece = l03.PIECE_G1
	case l03.SQ_S1:
		piece = l03.PIECE_S1
	case l03.SQ_N1:
		piece = l03.PIECE_N1
	case l03.SQ_L1:
		piece = l03.PIECE_L1
	case l03.SQ_P1:
		piece = l03.PIECE_P1
	case l03.SQ_K2:
		piece = l03.PIECE_K2
	case l03.SQ_R2:
		piece = l03.PIECE_R2
	case l03.SQ_B2:
		piece = l03.PIECE_B2
	case l03.SQ_G2:
		piece = l03.PIECE_G2
	case l03.SQ_S2:
		piece = l03.PIECE_S2
	case l03.SQ_N2:
		piece = l03.PIECE_N2
	case l03.SQ_L2:
		piece = l03.PIECE_L2
	case l03.SQ_P2:
		piece = l03.PIECE_P2
	default:
		// Not hand
		sq_hand = l03.SQ_EMPTY
	}

	if sq_hand != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands1[sq_hand-l03.SQ_HAND_START] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		mov_piece_type = l03.What(piece)

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		ValidateThereArePieceIn(pPos, to)
		var pCB *ControlBoard
		if pPosSys.BuildType == BUILD_DEV {
			pCB = ControllBoardFromPhase(before_move_phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB = ControllBoardFromPhase(before_move_phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[to]
		if captured != l03.PIECE_EMPTY {
			pieceType := l03.What(captured)
			switch pieceType {
			case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR, l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB, l03.PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				piece := pPos.Board[to]

				// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
				ValidateThereArePieceIn(pPos, to)
				phase := l03.Who(piece)
				var pCB *ControlBoard
				if pPosSys.BuildType == BUILD_DEV {
					pCB = ControllBoardFromPhase(phase,
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_CAPTURED],
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_CAPTURED])
				} else {
					pCB = ControllBoardFromPhase(phase,
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
						pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
				}
				pCB.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
			}
			cap_piece_type = l03.What(captured)
			cap_src_sq = to

			// 駒得評価値の計算（＾ｑ＾）
			material_val := EvalMaterial(captured)
			if before_move_phase != l03.FIRST {
				material_val = -material_val
			}
			pPos.MaterialValue += material_val
		}

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase := l03.Who(piece)
		var pCB1 *ControlBoard
		if pPosSys.BuildType == BUILD_DEV {
			pCB1 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_REMOVE],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_REMOVE])
		} else {
			pCB1 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
		}
		// 元位置の駒の利きを除去
		pCB1.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, from)), from, -1)

		// 行き先の駒の上書き
		if pro {
			// 駒を成りに変換します
			pPos.Board[to] = l09.Promote(pPos.Board[from])
		} else {
			pPos.Board[to] = pPos.Board[from]
		}
		mov_piece_type = l03.What(pPos.Board[to])
		// 元位置の駒を削除してから、移動先の駒の利きを追加
		pPos.Board[from] = l03.PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase = l03.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB2 *ControlBoard
		if pPosSys.BuildType == BUILD_DEV {
			pCB2 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB2 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB2.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)

		switch captured {
		case l03.PIECE_EMPTY: // Ignored
		case l03.PIECE_K1: // Second player win
			cap_dst_sq = l03.SQ_K2
		case l03.PIECE_R1, l03.PIECE_PR1:
			cap_dst_sq = l03.SQ_R2
		case l03.PIECE_B1, l03.PIECE_PB1:
			cap_dst_sq = l03.SQ_B2
		case l03.PIECE_G1:
			cap_dst_sq = l03.SQ_G2
		case l03.PIECE_S1, l03.PIECE_PS1:
			cap_dst_sq = l03.SQ_S2
		case l03.PIECE_N1, l03.PIECE_PN1:
			cap_dst_sq = l03.SQ_N2
		case l03.PIECE_L1, l03.PIECE_PL1:
			cap_dst_sq = l03.SQ_L2
		case l03.PIECE_P1, l03.PIECE_PP1:
			cap_dst_sq = l03.SQ_P2
		case l03.PIECE_K2: // l03.FIRST player win
			cap_dst_sq = l03.SQ_K1
		case l03.PIECE_R2, l03.PIECE_PR2:
			cap_dst_sq = l03.SQ_R1
		case l03.PIECE_B2, l03.PIECE_PB2:
			cap_dst_sq = l03.SQ_B1
		case l03.PIECE_G2:
			cap_dst_sq = l03.SQ_G1
		case l03.PIECE_S2, l03.PIECE_PS2:
			cap_dst_sq = l03.SQ_S1
		case l03.PIECE_N2, l03.PIECE_PN2:
			cap_dst_sq = l03.SQ_N1
		case l03.PIECE_L2, l03.PIECE_PL2:
			cap_dst_sq = l03.SQ_L1
		case l03.PIECE_P2, l03.PIECE_PP2:
			cap_dst_sq = l03.SQ_P1
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		if cap_dst_sq != l03.SQ_EMPTY {
			pPosSys.CapturedList[pPosSys.OffsetMovesIndex] = captured
			pPos.Hands1[cap_dst_sq-l03.SQ_HAND_START] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pPosSys.CapturedList[pPosSys.OffsetMovesIndex] = l03.PIECE_EMPTY
		}
	}

	// DoMoveでフェーズを１つ進めます
	pPosSys.Moves[pPosSys.OffsetMovesIndex] = move
	pPosSys.OffsetMovesIndex += 1
	pPosSys.FlipPhase()

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []l03.PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []l03.Square{from, cap_src_sq}
	dst_sq_list := []l03.Square{to, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case l03.PIECE_TYPE_K:
			if j == 0 {
				switch before_move_phase {
				case l03.FIRST:
					pPos.PieceLocations[l07.PCLOC_K1] = dst_sq_list[j]
				case l03.SECOND:
					pPos.PieceLocations[l07.PCLOC_K2] = dst_sq_list[j]
				default:
					panic(App.LogNotEcho.Fatal("unknown before_move_phase=%d", before_move_phase))
				}
			} else {
				// 取った時
				switch before_move_phase {
				case l03.FIRST:
					// 相手玉
					pPos.PieceLocations[l07.PCLOC_K2] = dst_sq_list[j]
				case l03.SECOND:
					pPos.PieceLocations[l07.PCLOC_K1] = dst_sq_list[j]
				default:
					panic(App.LogNotEcho.Fatal("unknown before_move_phase=%d", before_move_phase))
				}
			}
		case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR:
			for i := l07.PCLOC_R1; i < l07.PCLOC_R2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB:
			for i := l07.PCLOC_B1; i < l07.PCLOC_B2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case l03.PIECE_TYPE_L, l03.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i := l07.PCLOC_L1; i < l07.PCLOC_L4+1; i += 1 {
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
		pPos, pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], 1, to)
	AddControlBishop(
		pPos, pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], 1, to)
	AddControlRook(
		pPos, pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], 1, to)

	pPosSys.PControlBoardSystem.MergeControlDiff(pPosSys.BuildType)
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pPosSys *PositionSystem) UndoMove(pPos *Position) {

	// App.Log.Trace(pPosSys.Sprint())

	if pPosSys.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := l03.PIECE_TYPE_EMPTY

	// 先に 手目 を１つ戻すぜ（＾～＾）UndoMoveでフェーズもひっくり返すぜ（＾～＾）
	pPosSys.OffsetMovesIndex -= 1
	move := pPosSys.Moves[pPosSys.OffsetMovesIndex]
	// next_phase := pPosSys.GetPhase()
	pPosSys.FlipPhase()

	from, to, pro := move.Destructure()

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff(pPosSys.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlRook(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], -1, to)
	AddControlBishop(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], -1, to)
	AddControlLance(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], -1, to)

	// 打かどうかで分けます
	switch from {
	case l03.SQ_K1, l03.SQ_R1, l03.SQ_B1, l03.SQ_G1, l03.SQ_S1, l03.SQ_N1, l03.SQ_L1, l03.SQ_P1, l03.SQ_K2, l03.SQ_R2, l03.SQ_B2, l03.SQ_G2, l03.SQ_S2, l03.SQ_N2, l03.SQ_L2, l03.SQ_P2:
		// 打なら
		hand := from
		// 行き先から駒を除去します
		mov_piece_type = l03.What(pPos.Board[to])

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := l03.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB3 *ControlBoard
		if pPosSys.BuildType == BUILD_DEV {
			pCB3 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB3 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB3.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
		pPos.Board[to] = l03.PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands1[hand-l03.SQ_HAND_START] += 1
	default:
		// 打でないなら

		// 行き先に進んでいた自駒の利きの除去
		mov_piece_type = l03.What(pPos.Board[to])

		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := l03.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB4 *ControlBoard
		if pPosSys.BuildType == BUILD_DEV {
			pCB4 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB4 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB4.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)

		// 自駒を移動元へ戻します
		if pro {
			// 成りを元に戻します
			pPos.Board[from] = l09.Demote(pPos.Board[to])
		} else {
			pPos.Board[from] = pPos.Board[to]
		}

		pPos.Board[to] = l03.PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase = l03.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB5 *ControlBoard
		if pPosSys.BuildType == BUILD_DEV {
			pCB5 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_REMOVE],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_REMOVE])
		} else {
			pCB5 = ControllBoardFromPhase(phase,
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1],
				pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2])
		}
		// 元の場所に戻した自駒の利きを復元します
		pCB5.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, from)), from, 1)
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch mov_piece_type {
	case l03.PIECE_TYPE_K:
		// 玉を動かした
		switch pPosSys.phase { // next_phase
		case l03.FIRST:
			pPos.PieceLocations[l07.PCLOC_K1] = from
		case l03.SECOND:
			pPos.PieceLocations[l07.PCLOC_K2] = from
		default:
			panic(App.LogNotEcho.Fatal("unknown p_pos_sys.phase=%d", pPosSys.phase))
		}
	case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR:
		for i := l07.PCLOC_R1; i < l07.PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB:
		for i := l07.PCLOC_B1; i < l07.PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case l03.PIECE_TYPE_L, l03.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := l07.PCLOC_L1; i < l07.PCLOC_L4+1; i += 1 {
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
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], 1, from)
	AddControlBishop(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], 1, from)
	AddControlRook(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], 1, from)

	pPosSys.PControlBoardSystem.MergeControlDiff(pPosSys.BuildType)

	// 取った駒を戻すぜ（＾～＾）
	pPosSys.undoCapture(pPos)
}

// undoCapture - 取った駒を戻すぜ（＾～＾）
func (pPosSys *PositionSystem) undoCapture(pPos *Position) {
	// App.Log.Trace(pPosSys.Sprint())

	// 取った駒だぜ（＾～＾）
	cap_piece_type := l03.PIECE_TYPE_EMPTY

	// 手目もフェーズもすでに１つ戻っているとするぜ（＾～＾）
	move := pPosSys.Moves[pPosSys.OffsetMovesIndex]

	// 取った駒
	captured := pPosSys.CapturedList[pPosSys.OffsetMovesIndex]
	// fmt.Printf("Debug: CapturedPiece=%s\n", captured.ToCode())

	// 取った駒に関係するのは行き先だけ（＾～＾）
	from, to, _ := move.Destructure()
	// fmt.Printf("Debug: to=%d\n", to)

	var hand_sq = l03.SQ_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff(pPosSys.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlRook(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], -1, to)
	AddControlBishop(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], -1, to)
	AddControlLance(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], -1, to)

	// 打かどうかで分けます
	switch from {
	case l03.SQ_K1, l03.SQ_R1, l03.SQ_B1, l03.SQ_G1, l03.SQ_S1, l03.SQ_N1, l03.SQ_L1, l03.SQ_P1, l03.SQ_K2, l03.SQ_R2, l03.SQ_B2, l03.SQ_G2, l03.SQ_S2, l03.SQ_N2, l03.SQ_L2, l03.SQ_P2:
		// 打で取れる駒はないぜ（＾～＾）
		// fmt.Printf("Debug: Drop from=%d\n", from)
	default:
		// 打でないなら
		// fmt.Printf("Debug: Not hand from=%d\n", from)

		// 取った相手の駒があれば、自分の駒台から下ろします
		switch captured {
		case l03.PIECE_EMPTY: // Ignored
		case l03.PIECE_K1: // Second player win
			hand_sq = l03.SQ_K2
		case l03.PIECE_R1, l03.PIECE_PR1:
			hand_sq = l03.SQ_R2
		case l03.PIECE_B1, l03.PIECE_PB1:
			hand_sq = l03.SQ_B2
		case l03.PIECE_G1:
			hand_sq = l03.SQ_G2
		case l03.PIECE_S1, l03.PIECE_PS1:
			hand_sq = l03.SQ_S2
		case l03.PIECE_N1, l03.PIECE_PN1:
			hand_sq = l03.SQ_N2
		case l03.PIECE_L1, l03.PIECE_PL1:
			hand_sq = l03.SQ_L2
		case l03.PIECE_P1, l03.PIECE_PP1:
			hand_sq = l03.SQ_P2
		case l03.PIECE_K2: // l03.FIRST player win
			hand_sq = l03.SQ_K1
		case l03.PIECE_R2, l03.PIECE_PR2:
			hand_sq = l03.SQ_R1
		case l03.PIECE_B2, l03.PIECE_PB2:
			hand_sq = l03.SQ_B1
		case l03.PIECE_G2:
			hand_sq = l03.SQ_G1
		case l03.PIECE_S2, l03.PIECE_PS2:
			hand_sq = l03.SQ_S1
		case l03.PIECE_N2, l03.PIECE_PN2:
			hand_sq = l03.SQ_N1
		case l03.PIECE_L2, l03.PIECE_PL2:
			hand_sq = l03.SQ_L1
		case l03.PIECE_P2, l03.PIECE_PP2:
			hand_sq = l03.SQ_P1
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		// fmt.Printf("Debug: hand_sq=%d\n", hand_sq)

		if hand_sq != l03.SQ_EMPTY {
			pPos.Hands1[hand_sq-l03.SQ_HAND_START] -= 1

			// 取っていた駒を行き先に戻します
			cap_piece_type = l03.What(captured)
			pPos.Board[to] = captured

			ValidateThereArePieceIn(pPos, to)
			// fmt.Printf("Debug: ph=%d\n", ph)
			var pCB6 *ControlBoard
			if pPosSys.BuildType == BUILD_DEV {
				pCB6 = ControllBoardFromPhase(pPosSys.phase,
					pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_CAPTURED],
					pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_CAPTURED])
			} else {
				pCB6 = ControllBoardFromPhase(pPosSys.phase,
					pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM2],
					pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_SUM1])
			}
			// 取った駒は盤上になかったので、ここで利きを復元させます
			// 行き先にある取られていた駒の利きの復元
			pCB6.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)
		}
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch cap_piece_type {
	case l03.PIECE_TYPE_K:
		// 玉を取っていた
		switch pPosSys.phase { // next_phase
		case l03.FIRST:
			// 後手の玉
			pPos.PieceLocations[l07.PCLOC_K2] = to
		case l03.SECOND:
			// 先手の玉
			pPos.PieceLocations[l07.PCLOC_K1] = to
		default:
			panic(App.LogNotEcho.Fatal("unknown p_pos_sys.phase=%d", pPosSys.phase))
		}
	case l03.PIECE_TYPE_R, l03.PIECE_TYPE_PR:
		for i := l07.PCLOC_R1; i < l07.PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case l03.PIECE_TYPE_B, l03.PIECE_TYPE_PB:
		for i := l07.PCLOC_B1; i < l07.PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case l03.PIECE_TYPE_L, l03.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := l07.PCLOC_L1; i < l07.PCLOC_L4+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	}

	// 駒得評価値の計算（＾ｑ＾）
	material_val := EvalMaterial(captured)
	if pPosSys.phase != l03.FIRST {
		material_val = -material_val
	}
	pPos.MaterialValue -= material_val

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlLance(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], 1, from)
	AddControlBishop(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], 1, from)
	AddControlRook(
		pPos,
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pPosSys.PControlBoardSystem.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], 1, from)

	pPosSys.PControlBoardSystem.MergeControlDiff(pPosSys.BuildType)
}
