package take16

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// 開発 or リリース モード
type BuildT int

const (
	BUILD_DEV     = BuildT(0)
	BUILD_RELEASE = BuildT(1)
)

// Nerve - 局面システムと、利き盤システムの２つを持つもの
type Nerve struct {
	// 開発モードフラグ。デフォルト値：真。 'usi' コマンドで解除
	BuildType BuildT
	// 局面システム
	PPosSys *PositionSystem
	// 利きボード・システム
	PCtrlBrdSys *ControlBoardSystem
	// 差分での連続局面記録。つまり、ふつうの棋譜（＾～＾）
	PRecord *DifferenceRecord
	// 時間管理用
	OneMoveSec int
	// 時間管理用
	IsStopSearch bool
	// 時間管理用
	PStopwatchSearch *Stopwatch
	// エンジン・オプション
	MaxDepth int
}

func NewNerve() *Nerve {
	var pNerve = new(Nerve)
	pNerve.BuildType = BUILD_DEV
	pNerve.PPosSys = NewPositionSystem()
	pNerve.PRecord = NewDifferenceRecord()
	pNerve.PCtrlBrdSys = NewControlBoardSystem()
	pNerve.OneMoveSec = 0
	pNerve.IsStopSearch = false
	pNerve.PStopwatchSearch = NewStopwatch()
	pNerve.MaxDepth = 2
	return pNerve
}

func (pNerve *Nerve) ClearBySearchEntry() {
	pNerve.IsStopSearch = false
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pNerve *Nerve) ReadPosition(pPos *Position, command string) {
	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面をセット（＾～＾）
		pPos.ClearBoard()
		pNerve.PCtrlBrdSys = NewControlBoardSystem()
		pNerve.PPosSys.ResetPosition()
		pNerve.PRecord.ResetDifferenceRecord()
		pPos.SetToStartpos()
		i = 17

		if i < len && command[i] == ' ' {
			i += 1
		}
		// moves へ続くぜ（＾～＾）

	} else if strings.HasPrefix(command, "position sfen ") {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		pPos.ClearBoard()
		pNerve.PCtrlBrdSys = NewControlBoardSystem()
		pNerve.PPosSys.ResetPosition()
		pNerve.PRecord.ResetDifferenceRecord()
		i = 14
		var rank = Square(1)
		var file = Square(9)

	BoardLoop:
		for {
			promoted := false
			switch pc := command[i]; pc {
			case 'K', 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'k', 'r', 'b', 'g', 's', 'n', 'l', 'p':
				pPos.Board[file*10+rank] = l09.FromStringToPiece(string(pc))
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = l09.PIECE_EMPTY
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
					pPos.Board[file*10+rank] = l09.FromStringToPiece("+" + string(pc2))
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.PieceLocations[PCLOC_K1] = Square((file+1)*10 + rank)
			case 'k':
				pPos.PieceLocations[PCLOC_K2] = Square((file+1)*10 + rank)
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == SQUARE_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'B', 'b':
				for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == SQUARE_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'L', 'l':
				for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == SQUARE_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			}
		}

		// 手番
		switch command[i] {
		case 'b':
			pNerve.PPosSys.phase = FIRST
			i += 1
		case 'w':
			pNerve.PPosSys.phase = SECOND
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
			var hand_index int = 999 //存在しない数
			var number = 0

		HandLoop:
			for {
				var piece = command[i]

				if unicode.IsLetter(rune(piece)) || piece == ' ' {

					if hand_index == 999 {
						// ループの１週目は無視します

					} else {
						// 数字が書いてなかったら１個
						if number == 0 {
							number = 1
						}

						pPos.Hands1[hand_index] = number
						number = 0

						// 長い利きの駒は位置を覚えておくぜ（＾～＾）
						switch hand_index {
						case l11.HAND_R1, l11.HAND_R2:
							for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						case l11.HAND_B1, l11.HAND_B2:
							for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						case l11.HAND_L1, l11.HAND_L2:
							for i := PCLOC_L1; i < PCLOC_L4+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						}
					}
					i += 1

					switch piece {
					case 'R':
						hand_index = l11.HAND_R1
					case 'B':
						hand_index = l11.HAND_B1
					case 'G':
						hand_index = l11.HAND_G1
					case 'S':
						hand_index = l11.HAND_S1
					case 'N':
						hand_index = l11.HAND_N1
					case 'L':
						hand_index = l11.HAND_L1
					case 'P':
						hand_index = l11.HAND_P1
					case 'r':
						hand_index = l11.HAND_R2
					case 'b':
						hand_index = l11.HAND_B2
					case 'g':
						hand_index = l11.HAND_G2
					case 's':
						hand_index = l11.HAND_S2
					case 'n':
						hand_index = l11.HAND_N2
					case 'l':
						hand_index = l11.HAND_L2
					case 'p':
						hand_index = l11.HAND_P2
					case ' ':
						// ループを抜けます
						break HandLoop
					default:
						panic(App.LogNotEcho.Fatal("fatal: unknown piece=%c", piece))
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
		pNerve.PRecord.StartMovesNum = 0
	MovesNumLoop:
		for i < len {
			switch figure := command[i]; figure {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num, err := strconv.Atoi(string(figure))
				if err != nil {
					panic(err)
				}
				i += 1
				pNerve.PRecord.StartMovesNum *= 10
				pNerve.PRecord.StartMovesNum += num
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

	start_phase := pNerve.PPosSys.GetPhase()
	if strings.HasPrefix(command[i:], "moves") {
		i += 5

		// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
		for i < len {
			if command[i] != ' ' {
				break
			}
			i += 1

			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			var move, err = ParseMove(command, &i, pNerve.PPosSys.GetPhase())
			if err != nil {
				fmt.Println(err)
				fmt.Println(pPos.SprintBoardHeader(
					pNerve.PPosSys.phase,
					pNerve.PRecord.StartMovesNum,
					pNerve.PRecord.OffsetMovesIndex))
				fmt.Println(pPos.SprintBoard())
				fmt.Println(pNerve.SprintBoardFooter())
				panic(err)
			}
			pNerve.PRecord.Moves[pNerve.PRecord.OffsetMovesIndex] = move
			pNerve.PRecord.OffsetMovesIndex += 1
			pNerve.PPosSys.FlipPhase()
		}
	}

	if pNerve.BuildType == BUILD_DEV {
		// 利きの差分テーブルをクリアー（＾～＾）
		pNerve.PCtrlBrdSys.ClearControlDiff(pNerve.BuildType)
	}

	// 開始局面の利きを計算（＾～＾）
	//fmt.Printf("Debug: 開始局面の利きを計算（＾～＾）\n")
	for sq := Square(11); sq < 100; sq += 1 {
		if File(sq) != 0 && Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				//fmt.Printf("Debug: sq=%d\n", sq)
				// あとですぐクリアーするので、どのレイヤー使ってても関係ないんで、仮で PUTレイヤーを使っているぜ（＾～＾）

				// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
				piece := pPos.Board[sq]
				ValidateThereArePieceIn(pPos, sq)
				phase := Who(piece)
				// fmt.Printf("Debug: ph=%d\n", ph)
				var pCB7 *ControlBoard
				if pNerve.BuildType == BUILD_DEV {
					pCB7 = ControllBoardFromPhase(phase,
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
				} else {
					pCB7 = ControllBoardFromPhase(phase,
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
				}
				pCB7.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, sq)), sq, 1)
			}
		}
	}
	if pNerve.BuildType == BUILD_DEV {
		//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
		pNerve.PCtrlBrdSys.MergeControlDiff(pNerve.BuildType)
	}

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pNerve.PRecord.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pNerve.PRecord.OffsetMovesIndex = 0
	pNerve.PPosSys.phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pNerve.DoMove(pPos, pNerve.PRecord.Moves[i])
	}
}

// 長い利きの駒から王手を受けていないかチェック（＾～＾）
func (pNerve *Nerve) IsCheckmate(phase Phase) bool {
	switch phase {
	case FIRST:
		// 先手玉への王手を調べます
		// 先手玉の位置を調べます
		var k1 = pNerve.PPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[PCLOC_K1]
		// 後手の角の利きボードの、先手玉の位置のマスの数を調べます
		var b2 = pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON].Board1[k1] + pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF].Board1[k1]
		if 0 < b2 {
			// 1以上なら王手を受けています
			return true
		}
		// 飛
		var r2 = pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON].Board1[k1] + pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF].Board1[k1]
		if 0 < r2 {
			return true
		}
		// 香
		var l2 = pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON].Board1[k1] + pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF].Board1[k1]
		if 0 < l2 {
			return true
		}
	case SECOND:
		// 後手玉の王手を調べます
		// 後手玉の位置を調べます
		var k2 = pNerve.PPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[PCLOC_K2]
		// 先手の角の利きボードの、先手玉の位置のマスの数を調べます
		var b1 = pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON].Board1[k2] + pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF].Board1[k2]
		if 0 < b1 {
			// 1以上なら王手を受けています
			return true
		}
		// 飛
		var r1 = pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON].Board1[k2] + pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF].Board1[k2]
		if 0 < r1 {
			return true
		}
		// 香
		var l2 = pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON].Board1[k2] + pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF].Board1[k2]
		if 0 < l2 {
			return true
		}
	default:
		panic(App.LogNotEcho.Fatal("unknown phase=%d", phase))
	}

	// 王手は受けていなかったぜ（＾～＾）
	return false
}

// DoMove - 一手指すぜ（＾～＾）
func (pNerve *Nerve) DoMove(pPos *Position, move Move) {
	before_move_phase := pNerve.PPosSys.GetPhase()

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY
	cap_piece_type := PIECE_TYPE_EMPTY

	// 移動元マス、移動先マス、成りの有無
	from, to, pro := Destructure(move)
	if pPos.IsEmptySq(from) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", from)
	}
	var cap_src_sq Square
	var cap_dst_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pNerve.PCtrlBrdSys.ClearControlDiff(pNerve.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	AddControlRook(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], -1, from)
	AddControlBishop(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], -1, from)
	AddControlLance(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], -1, from)

	// まず、打かどうかで処理を分けます
	sq_hand := from
	var piece l09.Piece
	switch from {
	case SQ_K1:
		piece = l09.PIECE_K1
	case SQ_R1:
		piece = l09.PIECE_R1
	case SQ_B1:
		piece = l09.PIECE_B1
	case SQ_G1:
		piece = l09.PIECE_G1
	case SQ_S1:
		piece = l09.PIECE_S1
	case SQ_N1:
		piece = l09.PIECE_N1
	case SQ_L1:
		piece = l09.PIECE_L1
	case SQ_P1:
		piece = l09.PIECE_P1
	case SQ_K2:
		piece = l09.PIECE_K2
	case SQ_R2:
		piece = l09.PIECE_R2
	case SQ_B2:
		piece = l09.PIECE_B2
	case SQ_G2:
		piece = l09.PIECE_G2
	case SQ_S2:
		piece = l09.PIECE_S2
	case SQ_N2:
		piece = l09.PIECE_N2
	case SQ_L2:
		piece = l09.PIECE_L2
	case SQ_P2:
		piece = l09.PIECE_P2
	default:
		// Not hand
		sq_hand = SQUARE_EMPTY
	}

	if sq_hand != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands1[sq_hand-SQ_HAND_START] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		mov_piece_type = What(piece)

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		ValidateThereArePieceIn(pPos, to)
		var pCB *ControlBoard
		if pNerve.BuildType == BUILD_DEV {
			pCB = ControllBoardFromPhase(before_move_phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB = ControllBoardFromPhase(before_move_phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[to]
		if captured != l09.PIECE_EMPTY {
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
				if pNerve.BuildType == BUILD_DEV {
					pCB = ControllBoardFromPhase(phase,
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_CAPTURED],
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_CAPTURED])
				} else {
					pCB = ControllBoardFromPhase(phase,
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
						pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
				}
				pCB.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
			}
			cap_piece_type = What(captured)
			cap_src_sq = to

			// 駒得評価値。駒取って得したあと、相手の手番になるからひっくり返せだぜ（＾～＾）
			pPos.MaterialValue += EvalMaterial(captured)
			pPos.MaterialValue = -pPos.MaterialValue
		}

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase := Who(piece)
		var pCB1 *ControlBoard
		if pNerve.BuildType == BUILD_DEV {
			pCB1 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_REMOVE],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_REMOVE])
		} else {
			pCB1 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
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
		mov_piece_type = What(pPos.Board[to])
		// 元位置の駒を削除してから、移動先の駒の利きを追加
		pPos.Board[from] = l09.PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase = Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB2 *ControlBoard
		if pNerve.BuildType == BUILD_DEV {
			pCB2 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB2 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB2.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, 1)

		switch captured {
		case l09.PIECE_EMPTY: // Ignored
		case l09.PIECE_K1: // Second player win
			cap_dst_sq = SQ_K2
		case l09.PIECE_R1, l09.PIECE_PR1:
			cap_dst_sq = SQ_R2
		case l09.PIECE_B1, l09.PIECE_PB1:
			cap_dst_sq = SQ_B2
		case l09.PIECE_G1:
			cap_dst_sq = SQ_G2
		case l09.PIECE_S1, l09.PIECE_PS1:
			cap_dst_sq = SQ_S2
		case l09.PIECE_N1, l09.PIECE_PN1:
			cap_dst_sq = SQ_N2
		case l09.PIECE_L1, l09.PIECE_PL1:
			cap_dst_sq = SQ_L2
		case l09.PIECE_P1, l09.PIECE_PP1:
			cap_dst_sq = SQ_P2
		case l09.PIECE_K2: // First player win
			cap_dst_sq = SQ_K1
		case l09.PIECE_R2, l09.PIECE_PR2:
			cap_dst_sq = SQ_R1
		case l09.PIECE_B2, l09.PIECE_PB2:
			cap_dst_sq = SQ_B1
		case l09.PIECE_G2:
			cap_dst_sq = SQ_G1
		case l09.PIECE_S2, l09.PIECE_PS2:
			cap_dst_sq = SQ_S1
		case l09.PIECE_N2, l09.PIECE_PN2:
			cap_dst_sq = SQ_N1
		case l09.PIECE_L2, l09.PIECE_PL2:
			cap_dst_sq = SQ_L1
		case l09.PIECE_P2, l09.PIECE_PP2:
			cap_dst_sq = SQ_P1
		default:
			fmt.Printf("unknown captured=[%d]", captured)
		}

		if cap_dst_sq != SQUARE_EMPTY {
			pNerve.PRecord.CapturedList[pNerve.PRecord.OffsetMovesIndex] = captured
			pPos.Hands1[cap_dst_sq-SQ_HAND_START] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pNerve.PRecord.CapturedList[pNerve.PRecord.OffsetMovesIndex] = l09.PIECE_EMPTY
		}
	}

	// DoMoveでフェーズを１つ進めます
	pNerve.PRecord.Moves[pNerve.PRecord.OffsetMovesIndex] = move
	pNerve.PRecord.OffsetMovesIndex += 1
	pNerve.PPosSys.FlipPhase()

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
					panic(App.LogNotEcho.Fatal("Unknown before_move_phase=%d", before_move_phase))
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
					panic(App.LogNotEcho.Fatal("Unknown before_move_phase=%d", before_move_phase))
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
		pPos, pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], 1, to)
	AddControlBishop(
		pPos, pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], 1, to)
	AddControlRook(
		pPos, pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], 1, to)

	pNerve.PCtrlBrdSys.MergeControlDiff(pNerve.BuildType)
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pNerve *Nerve) UndoMove(pPos *Position) {

	// App.Log.Trace(pNerve.PPosSys.Sprint())

	if pNerve.PRecord.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := PIECE_TYPE_EMPTY

	// 先に 手目 を１つ戻すぜ（＾～＾）UndoMoveでフェーズもひっくり返すぜ（＾～＾）
	pNerve.PRecord.OffsetMovesIndex -= 1
	move := pNerve.PRecord.Moves[pNerve.PRecord.OffsetMovesIndex]
	// next_phase := pNerve.PPosSys.GetPhase()
	pNerve.PPosSys.FlipPhase()

	from, to, pro := Destructure(move)

	// 利きの差分テーブルをクリアー（＾～＾）
	pNerve.PCtrlBrdSys.ClearControlDiff(pNerve.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlRook(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], -1, to)
	AddControlBishop(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], -1, to)
	AddControlLance(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], -1, to)

	// 打かどうかで分けます
	switch from {
	case SQ_K1, SQ_R1, SQ_B1, SQ_G1, SQ_S1, SQ_N1, SQ_L1, SQ_P1, SQ_K2, SQ_R2, SQ_B2, SQ_G2, SQ_S2, SQ_N2, SQ_L2, SQ_P2:
		// 打なら
		hand := from
		// 行き先から駒を除去します
		mov_piece_type = What(pPos.Board[to])

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB3 *ControlBoard
		if pNerve.BuildType == BUILD_DEV {
			pCB3 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB3 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB3.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
		pPos.Board[to] = l09.PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands1[hand-SQ_HAND_START] += 1
	default:
		// 打でないなら

		// 行き先に進んでいた自駒の利きの除去
		mov_piece_type = What(pPos.Board[to])

		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB4 *ControlBoard
		if pNerve.BuildType == BUILD_DEV {
			pCB4 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB4 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB4.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)

		// 自駒を移動元へ戻します
		if pro {
			// 成りを元に戻します
			pPos.Board[from] = l09.Demote(pPos.Board[to])
		} else {
			pPos.Board[from] = pPos.Board[to]
		}

		pPos.Board[to] = l09.PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase = Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB5 *ControlBoard
		if pNerve.BuildType == BUILD_DEV {
			pCB5 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_REMOVE],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_REMOVE])
		} else {
			pCB5 = ControllBoardFromPhase(phase,
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		// 元の場所に戻した自駒の利きを復元します
		pCB5.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, from)), from, 1)
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch mov_piece_type {
	case PIECE_TYPE_K:
		// 玉を動かした
		switch pNerve.PPosSys.phase { // next_phase
		case FIRST:
			pPos.PieceLocations[PCLOC_K1] = from
		case SECOND:
			pPos.PieceLocations[PCLOC_K2] = from
		default:
			panic(App.LogNotEcho.Fatal("Unknown pNerve.PPosSys.phase=%d", pNerve.PPosSys.phase))
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
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], 1, from)
	AddControlBishop(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], 1, from)
	AddControlRook(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], 1, from)

	pNerve.PCtrlBrdSys.MergeControlDiff(pNerve.BuildType)

	// 取った駒を戻すぜ（＾～＾）
	pNerve.undoCapture(pPos)
}

// undoCapture - 取った駒を戻すぜ（＾～＾）
func (pNerve *Nerve) undoCapture(pPos *Position) {
	// App.Log.Trace(pNerve.PPosSys.Sprint())

	// 取った駒だぜ（＾～＾）
	cap_piece_type := PIECE_TYPE_EMPTY

	// 手目もフェーズもすでに１つ戻っているとするぜ（＾～＾）
	move := pNerve.PRecord.Moves[pNerve.PRecord.OffsetMovesIndex]

	// 取った駒
	captured := pNerve.PRecord.CapturedList[pNerve.PRecord.OffsetMovesIndex]
	// fmt.Printf("Debug: CapturedPiece=%s\n", captured.ToCode())

	// 駒得評価値。今、自分は駒を取られて減っているので、それを戻すために増やす。
	// そして相手の手番になるので　評価値をひっくり返すぜ（＾ｑ＾）
	pPos.MaterialValue += EvalMaterial(captured)
	pPos.MaterialValue = -pPos.MaterialValue

	// 取った駒に関係するのは行き先だけ（＾～＾）
	from, to, _ := Destructure(move)
	// fmt.Printf("Debug: to=%d\n", to)

	var hand_sq = SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pNerve.PCtrlBrdSys.ClearControlDiff(pNerve.BuildType)

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlRook(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON], -1, to)
	AddControlBishop(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON], -1, to)
	AddControlLance(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON], -1, to)

	// 打かどうかで分けます
	switch from {
	case SQ_K1, SQ_R1, SQ_B1, SQ_G1, SQ_S1, SQ_N1, SQ_L1, SQ_P1, SQ_K2, SQ_R2, SQ_B2, SQ_G2, SQ_S2, SQ_N2, SQ_L2, SQ_P2:
		// 打で取れる駒はないぜ（＾～＾）
		// fmt.Printf("Debug: Drop from=%d\n", from)
	default:
		// 打でないなら
		// fmt.Printf("Debug: Not hand from=%d\n", from)

		// 取った相手の駒があれば、自分の駒台から下ろします
		switch captured {
		case l09.PIECE_EMPTY: // Ignored
		case l09.PIECE_K1: // Second player win
			hand_sq = SQ_K2
		case l09.PIECE_R1, l09.PIECE_PR1:
			hand_sq = SQ_R2
		case l09.PIECE_B1, l09.PIECE_PB1:
			hand_sq = SQ_B2
		case l09.PIECE_G1:
			hand_sq = SQ_G2
		case l09.PIECE_S1, l09.PIECE_PS1:
			hand_sq = SQ_S2
		case l09.PIECE_N1, l09.PIECE_PN1:
			hand_sq = SQ_N2
		case l09.PIECE_L1, l09.PIECE_PL1:
			hand_sq = SQ_L2
		case l09.PIECE_P1, l09.PIECE_PP1:
			hand_sq = SQ_P2
		case l09.PIECE_K2: // First player win
			hand_sq = SQ_K1
		case l09.PIECE_R2, l09.PIECE_PR2:
			hand_sq = SQ_R1
		case l09.PIECE_B2, l09.PIECE_PB2:
			hand_sq = SQ_B1
		case l09.PIECE_G2:
			hand_sq = SQ_G1
		case l09.PIECE_S2, l09.PIECE_PS2:
			hand_sq = SQ_S1
		case l09.PIECE_N2, l09.PIECE_PN2:
			hand_sq = SQ_N1
		case l09.PIECE_L2, l09.PIECE_PL2:
			hand_sq = SQ_L1
		case l09.PIECE_P2, l09.PIECE_PP2:
			hand_sq = SQ_P1
		default:
			fmt.Printf("unknown captured=[%d]", captured)
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
			if pNerve.BuildType == BUILD_DEV {
				pCB6 = ControllBoardFromPhase(pNerve.PPosSys.phase,
					pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_CAPTURED],
					pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_CAPTURED])
			} else {
				pCB6 = ControllBoardFromPhase(pNerve.PPosSys.phase,
					pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2],
					pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1])
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
		switch pNerve.PPosSys.phase { // next_phase
		case FIRST:
			// 後手の玉
			pPos.PieceLocations[PCLOC_K2] = to
		case SECOND:
			// 先手の玉
			pPos.PieceLocations[PCLOC_K1] = to
		default:
			panic(App.LogNotEcho.Fatal("Unknown pNerve.PPosSys.phase=%d", pNerve.PPosSys.phase))
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

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	AddControlLance(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF], 1, from)
	AddControlBishop(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF], 1, from)
	AddControlRook(
		pPos,
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF],
		pNerve.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF], 1, from)

	pNerve.PCtrlBrdSys.MergeControlDiff(pNerve.BuildType)
}
