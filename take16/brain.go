package take16

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	p "github.com/muzudho/kifuwarabe-wcsc31/take16position"
)

// 開発 or リリース モード
type BuildT int

const (
	BUILD_DEV     = BuildT(0)
	BUILD_RELEASE = BuildT(1)
)

// Brain - 局面システムと、利き盤システムの２つを持つもの
type Brain struct {
	// 開発モードフラグ。デフォルト値：真。 'usi' コマンドで解除
	BuildType BuildT
	// 局面システム
	PPosSys *PositionSystem
	// 利きボード・システム
	PCtrlBrdSys *ControlBoardSystem
}

func NewBrain() *Brain {
	var pBrain = new(Brain)
	pBrain.BuildType = BUILD_DEV
	pBrain.PPosSys = NewPositionSystem()
	pBrain.PCtrlBrdSys = NewControlBoardSystem()
	return pBrain
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pBrain *Brain) ReadPosition(pPos *p.Position, command string) {
	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面をセット（＾～＾）
		pPos.ClearBoard()
		pBrain.PCtrlBrdSys = NewControlBoardSystem()
		pBrain.PPosSys.resetPosition()
		pPos.SetToStartpos()
		i = 17

		if i < len && command[i] == ' ' {
			i += 1
		}
		// moves へ続くぜ（＾～＾）

	} else if strings.HasPrefix(command, "position sfen ") {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		pPos.ClearBoard()
		pBrain.PCtrlBrdSys = NewControlBoardSystem()
		pBrain.PPosSys.resetPosition()
		i = 14
		var rank = p.Square(1)
		var file = p.Square(9)

	BoardLoop:
		for {
			promoted := false
			switch pc := command[i]; pc {
			case 'K', 'R', 'B', 'G', 'S', 'N', 'L', 'P', 'k', 'r', 'b', 'g', 's', 'n', 'l', 'p':
				pPos.Board[file*10+rank] = PieceFrom(string(pc))
				file -= 1
				i += 1
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var spaces, _ = strconv.Atoi(string(pc))
				for sp := 0; sp < spaces; sp += 1 {
					pPos.Board[file*10+rank] = p.PIECE_EMPTY
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
					pPos.Board[file*10+rank] = PieceFrom("+" + string(pc2))
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.PieceLocations[p.PCLOC_K1] = p.Square((file+1)*10 + rank)
			case 'k':
				pPos.PieceLocations[p.PCLOC_K2] = p.Square((file+1)*10 + rank)
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i := p.PCLOC_R1; i < p.PCLOC_R2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == p.SQUARE_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'B', 'b':
				for i := p.PCLOC_B1; i < p.PCLOC_B2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == p.SQUARE_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'L', 'l':
				for i := p.PCLOC_L1; i < p.PCLOC_L4+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == p.SQUARE_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			}
		}

		// 手番
		switch command[i] {
		case 'b':
			pBrain.PPosSys.phase = p.FIRST
			i += 1
		case 'w':
			pBrain.PPosSys.phase = p.SECOND
			i += 1
		default:
			panic("Fatal: Unknown phase")
		}

		if command[i] != ' ' {
			// 手番の後ろにスペースがない（＾～＾）
			panic("Fatal: Nothing space")
		}
		i += 1

		// 持ち駒
		if command[i] == '-' {
			i += 1
			if command[i] != ' ' {
				// 持ち駒 - の後ろにスペースがない（＾～＾）
				panic("Fatal: Nothing space after -")
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
						case p.HAND_R1_IDX, p.HAND_R2_IDX:
							for i := p.PCLOC_R1; i < p.PCLOC_R2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == p.SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = p.Square(hand_index) + p.SQ_HAND_START
									break
								}
							}
						case p.HAND_B1_IDX, p.HAND_B2_IDX:
							for i := p.PCLOC_B1; i < p.PCLOC_B2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == p.SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = p.Square(hand_index) + p.SQ_HAND_START
									break
								}
							}
						case p.HAND_L1_IDX, p.HAND_L2_IDX:
							for i := p.PCLOC_L1; i < p.PCLOC_L4+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == p.SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = p.Square(hand_index) + p.SQ_HAND_START
									break
								}
							}
						}
					}
					i += 1

					switch piece {
					case 'R':
						hand_index = p.HAND_R1_IDX
					case 'B':
						hand_index = p.HAND_B1_IDX
					case 'G':
						hand_index = p.HAND_G1_IDX
					case 'S':
						hand_index = p.HAND_S1_IDX
					case 'N':
						hand_index = p.HAND_N1_IDX
					case 'L':
						hand_index = p.HAND_L1_IDX
					case 'P':
						hand_index = p.HAND_P1_IDX
					case 'r':
						hand_index = p.HAND_R2_IDX
					case 'b':
						hand_index = p.HAND_B2_IDX
					case 'g':
						hand_index = p.HAND_G2_IDX
					case 's':
						hand_index = p.HAND_S2_IDX
					case 'n':
						hand_index = p.HAND_N2_IDX
					case 'l':
						hand_index = p.HAND_L2_IDX
					case 'p':
						hand_index = p.HAND_P2_IDX
					case ' ':
						// ループを抜けます
						break HandLoop
					default:
						panic(G.Log.Fatal("Fatal: Unknown piece=%c", piece))
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
						panic(G.Log.Fatal("Fatal: Unknown number character=%c", piece))
					}

				} else {
					panic(G.Log.Fatal("Fatal: Unknown piece=%c", piece))
				}
			}
		}

		// 手数
		pBrain.PPosSys.PRecord.StartMovesNum = 0
	MovesNumLoop:
		for i < len {
			switch figure := command[i]; figure {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num, err := strconv.Atoi(string(figure))
				if err != nil {
					panic(err)
				}
				i += 1
				pBrain.PPosSys.PRecord.StartMovesNum *= 10
				pBrain.PPosSys.PRecord.StartMovesNum += num
			case ' ':
				i += 1
				break MovesNumLoop
			default:
				break MovesNumLoop
			}
		}

	} else {
		fmt.Printf("Error: Unknown command=[%s]", command)
	}

	// fmt.Printf("command[i:]=[%s]\n", command[i:])

	start_phase := pBrain.PPosSys.GetPhase()
	if strings.HasPrefix(command[i:], "moves") {
		i += 5

		// 半角スペースに始まり、文字列の終わりで終わるぜ（＾～＾）
		for i < len {
			if command[i] != ' ' {
				break
			}
			i += 1

			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			var move, err = p.ParseMove(command, &i, pBrain.PPosSys.GetPhase())
			if err != nil {
				fmt.Println(err)
				fmt.Println(pPos.SprintBoardHeader(
					pBrain.PPosSys.phase,
					pBrain.PPosSys.PRecord.StartMovesNum,
					pBrain.PPosSys.OffsetMovesIndex))
				fmt.Println(pPos.SprintBoard())
				fmt.Println(pBrain.SprintBoardFooter())
				panic(err)
			}
			pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex] = move
			pBrain.PPosSys.OffsetMovesIndex += 1
			pBrain.PPosSys.FlipPhase()
		}
	}

	if pBrain.BuildType == BUILD_DEV {
		// 利きの差分テーブルをクリアー（＾～＾）
		pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.BuildType)
	}

	// 開始局面の利きを計算（＾～＾）
	//fmt.Printf("Debug: 開始局面の利きを計算（＾～＾）\n")
	for sq := p.Square(11); sq < 100; sq += 1 {
		if p.File(sq) != 0 && p.Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				//fmt.Printf("Debug: sq=%d\n", sq)
				// あとですぐクリアーするので、どのレイヤー使ってても関係ないんで、仮で PUTレイヤーを使っているぜ（＾～＾）

				// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
				piece := pPos.Board[sq]
				ValidateThereArePieceIn(pPos, sq)
				phase := p.Who(piece)
				// fmt.Printf("Debug: ph=%d\n", ph)
				var pCB7 *ControlBoard
				if pBrain.BuildType == BUILD_DEV {
					pCB7 = ControllBoardFromPhase(phase,
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
				} else {
					pCB7 = ControllBoardFromPhase(phase,
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
						pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
				}
				pCB7.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, sq)), sq, 1)
			}
		}
	}
	if pBrain.BuildType == BUILD_DEV {
		//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
		pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.BuildType)
	}

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pBrain.PPosSys.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pBrain.PPosSys.OffsetMovesIndex = 0
	pBrain.PPosSys.phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pBrain.DoMove(pPos, pBrain.PPosSys.Moves[i])
	}
}

// 長い利きの駒から王手を受けていないかチェック（＾～＾）
func (pBrain *Brain) IsCheckmate(phase p.Phase) bool {
	switch phase {
	case p.FIRST:
		// 先手玉への王手を調べます
		// 先手玉の位置を調べます
		var k1 = pBrain.PPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[p.PCLOC_K1]
		// 後手の角の利きボードの、先手玉の位置のマスの数を調べます
		var b2 = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_ON].Board1[k1] + pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_BISHOP_OFF].Board1[k1]
		if 0 < b2 {
			// 1以上なら王手を受けています
			return true
		}
		// 飛
		var r2 = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_ON].Board1[k1] + pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_ROOK_OFF].Board1[k1]
		if 0 < r2 {
			return true
		}
		// 香
		var l2 = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_ON].Board1[k1] + pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_LANCE_OFF].Board1[k1]
		if 0 < l2 {
			return true
		}
	case p.SECOND:
		// 後手玉の王手を調べます
		// 後手玉の位置を調べます
		var k2 = pBrain.PPosSys.PPosition[POS_LAYER_MAIN].PieceLocations[p.PCLOC_K2]
		// 先手の角の利きボードの、先手玉の位置のマスの数を調べます
		var b1 = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_ON].Board1[k2] + pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_BISHOP_OFF].Board1[k2]
		if 0 < b1 {
			// 1以上なら王手を受けています
			return true
		}
		// 飛
		var r1 = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_ON].Board1[k2] + pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_ROOK_OFF].Board1[k2]
		if 0 < r1 {
			return true
		}
		// 香
		var l2 = pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_ON].Board1[k2] + pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_LANCE_OFF].Board1[k2]
		if 0 < l2 {
			return true
		}
	default:
		panic(G.Log.Fatal("Unknown phase=%d", phase))
	}

	// 王手は受けていなかったぜ（＾～＾）
	return false
}

// DoMove - 一手指すぜ（＾～＾）
func (pBrain *Brain) DoMove(pPos *p.Position, move p.Move) {
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
	var cap_src_sq p.Square
	var cap_dst_sq = p.SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.BuildType)

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
	var piece p.Piece
	switch from {
	case p.SQ_K1:
		piece = p.PIECE_K1
	case p.SQ_R1:
		piece = p.PIECE_R1
	case p.SQ_B1:
		piece = p.PIECE_B1
	case p.SQ_G1:
		piece = p.PIECE_G1
	case p.SQ_S1:
		piece = p.PIECE_S1
	case p.SQ_N1:
		piece = p.PIECE_N1
	case p.SQ_L1:
		piece = p.PIECE_L1
	case p.SQ_P1:
		piece = p.PIECE_P1
	case p.SQ_K2:
		piece = p.PIECE_K2
	case p.SQ_R2:
		piece = p.PIECE_R2
	case p.SQ_B2:
		piece = p.PIECE_B2
	case p.SQ_G2:
		piece = p.PIECE_G2
	case p.SQ_S2:
		piece = p.PIECE_S2
	case p.SQ_N2:
		piece = p.PIECE_N2
	case p.SQ_L2:
		piece = p.PIECE_L2
	case p.SQ_P2:
		piece = p.PIECE_P2
	default:
		// Not drop
		sq_drop = p.SQUARE_EMPTY
	}

	if sq_drop != 0 {
		// 打なら

		// 持ち駒の数を減らします
		pPos.Hands1[sq_drop-p.SQ_HAND_START] -= 1

		// 行き先に駒を置きます
		pPos.Board[to] = piece
		mov_piece_type = What(piece)

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		ValidateThereArePieceIn(pPos, to)
		var pCB *ControlBoard
		if pBrain.BuildType == BUILD_DEV {
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
		if captured != p.PIECE_EMPTY {
			pieceType := What(captured)
			switch pieceType {
			case PIECE_TYPE_R, PIECE_TYPE_PR, PIECE_TYPE_B, PIECE_TYPE_PB, PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				piece := pPos.Board[to]

				// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
				ValidateThereArePieceIn(pPos, to)
				phase := p.Who(piece)
				var pCB *ControlBoard
				if pBrain.BuildType == BUILD_DEV {
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

			// 駒得評価値。駒取って得したあと、相手の手番になるからひっくり返せだぜ（＾～＾）
			pPos.MaterialValue += EvalMaterial(captured)
			pPos.MaterialValue = -pPos.MaterialValue
		}

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase := p.Who(piece)
		var pCB1 *ControlBoard
		if pBrain.BuildType == BUILD_DEV {
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
		pPos.Board[from] = p.PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase = p.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB2 *ControlBoard
		if pBrain.BuildType == BUILD_DEV {
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
		case p.PIECE_EMPTY: // Ignored
		case p.PIECE_K1: // Second player win
			cap_dst_sq = p.SQ_K2
		case p.PIECE_R1, p.PIECE_PR1:
			cap_dst_sq = p.SQ_R2
		case p.PIECE_B1, p.PIECE_PB1:
			cap_dst_sq = p.SQ_B2
		case p.PIECE_G1:
			cap_dst_sq = p.SQ_G2
		case p.PIECE_S1, p.PIECE_PS1:
			cap_dst_sq = p.SQ_S2
		case p.PIECE_N1, p.PIECE_PN1:
			cap_dst_sq = p.SQ_N2
		case p.PIECE_L1, p.PIECE_PL1:
			cap_dst_sq = p.SQ_L2
		case p.PIECE_P1, p.PIECE_PP1:
			cap_dst_sq = p.SQ_P2
		case p.PIECE_K2: // First player win
			cap_dst_sq = p.SQ_K1
		case p.PIECE_R2, p.PIECE_PR2:
			cap_dst_sq = p.SQ_R1
		case p.PIECE_B2, p.PIECE_PB2:
			cap_dst_sq = p.SQ_B1
		case p.PIECE_G2:
			cap_dst_sq = p.SQ_G1
		case p.PIECE_S2, p.PIECE_PS2:
			cap_dst_sq = p.SQ_S1
		case p.PIECE_N2, p.PIECE_PN2:
			cap_dst_sq = p.SQ_N1
		case p.PIECE_L2, p.PIECE_PL2:
			cap_dst_sq = p.SQ_L1
		case p.PIECE_P2, p.PIECE_PP2:
			cap_dst_sq = p.SQ_P1
		default:
			fmt.Printf("Error: Unknown captured=[%d]", captured)
		}

		if cap_dst_sq != p.SQUARE_EMPTY {
			pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex] = captured
			pPos.Hands1[cap_dst_sq-p.SQ_HAND_START] += 1
		} else {
			// 取った駒は無かった（＾～＾）
			pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex] = p.PIECE_EMPTY
		}
	}

	// DoMoveでフェーズを１つ進めます
	pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex] = move
	pBrain.PPosSys.OffsetMovesIndex += 1
	pBrain.PPosSys.FlipPhase()

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []p.Square{from, cap_src_sq}
	dst_sq_list := []p.Square{to, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case PIECE_TYPE_K:
			if j == 0 {
				switch before_move_phase {
				case p.FIRST:
					pPos.PieceLocations[p.PCLOC_K1] = dst_sq_list[j]
				case p.SECOND:
					pPos.PieceLocations[p.PCLOC_K2] = dst_sq_list[j]
				default:
					panic(G.Log.Fatal("Unknown before_move_phase=%d", before_move_phase))
				}
			} else {
				// 取った時
				switch before_move_phase {
				case p.FIRST:
					// 相手玉
					pPos.PieceLocations[p.PCLOC_K2] = dst_sq_list[j]
				case p.SECOND:
					pPos.PieceLocations[p.PCLOC_K1] = dst_sq_list[j]
				default:
					panic(G.Log.Fatal("Unknown before_move_phase=%d", before_move_phase))
				}
			}
		case PIECE_TYPE_R, PIECE_TYPE_PR:
			for i := p.PCLOC_R1; i < p.PCLOC_R2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case PIECE_TYPE_B, PIECE_TYPE_PB:
			for i := p.PCLOC_B1; i < p.PCLOC_B2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i := p.PCLOC_L1; i < p.PCLOC_L4+1; i += 1 {
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

	pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.BuildType)
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pBrain *Brain) UndoMove(pPos *p.Position) {

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
	pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.BuildType)

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
	case p.SQ_K1, p.SQ_R1, p.SQ_B1, p.SQ_G1, p.SQ_S1, p.SQ_N1, p.SQ_L1, p.SQ_P1, p.SQ_K2, p.SQ_R2, p.SQ_B2, p.SQ_G2, p.SQ_S2, p.SQ_N2, p.SQ_L2, p.SQ_P2:
		// 打なら
		drop := from
		// 行き先から駒を除去します
		mov_piece_type = What(pPos.Board[to])

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := p.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB3 *ControlBoard
		if pBrain.BuildType == BUILD_DEV {
			pCB3 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF1_PUT],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_DIFF2_PUT])
		} else {
			pCB3 = ControllBoardFromPhase(phase,
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM1],
				pBrain.PCtrlBrdSys.PBoards[CONTROL_LAYER_SUM2])
		}
		pCB3.AddControl(MoveEndListToControlList(GenMoveEnd(pPos, to)), to, -1)
		pPos.Board[to] = p.PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands1[drop-p.SQ_HAND_START] += 1
	default:
		// 打でないなら

		// 行き先に進んでいた自駒の利きの除去
		mov_piece_type = What(pPos.Board[to])

		piece := pPos.Board[to]
		ValidateThereArePieceIn(pPos, to)
		phase := p.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB4 *ControlBoard
		if pBrain.BuildType == BUILD_DEV {
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

		pPos.Board[to] = p.PIECE_EMPTY

		// 開発中は、利き計算を差分で行うぜ（＾～＾）実戦中は、差分は取らずに 利きテーブル本体を直接編集するぜ（＾～＾）
		piece = pPos.Board[from]
		ValidateThereArePieceIn(pPos, from)
		phase = p.Who(piece)
		// fmt.Printf("Debug: ph=%d\n", ph)
		var pCB5 *ControlBoard
		if pBrain.BuildType == BUILD_DEV {
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
		case p.FIRST:
			pPos.PieceLocations[p.PCLOC_K1] = from
		case p.SECOND:
			pPos.PieceLocations[p.PCLOC_K2] = from
		default:
			panic(G.Log.Fatal("Unknown pBrain.PPosSys.phase=%d", pBrain.PPosSys.phase))
		}
	case PIECE_TYPE_R, PIECE_TYPE_PR:
		for i := p.PCLOC_R1; i < p.PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case PIECE_TYPE_B, PIECE_TYPE_PB:
		for i := p.PCLOC_B1; i < p.PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := p.PCLOC_L1; i < p.PCLOC_L4+1; i += 1 {
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

	pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.BuildType)

	// 取った駒を戻すぜ（＾～＾）
	pBrain.undoCapture(pPos)
}

// undoCapture - 取った駒を戻すぜ（＾～＾）
func (pBrain *Brain) undoCapture(pPos *p.Position) {
	// G.StderrChat.Trace(pBrain.PPosSys.Sprint())

	// 取った駒だぜ（＾～＾）
	cap_piece_type := PIECE_TYPE_EMPTY

	// 手目もフェーズもすでに１つ戻っているとするぜ（＾～＾）
	move := pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex]

	// 取った駒
	captured := pBrain.PPosSys.CapturedList[pBrain.PPosSys.OffsetMovesIndex]
	// fmt.Printf("Debug: CapturedPiece=%s\n", captured.ToCode())

	// 駒得評価値。今、自分は駒を取られて減っているので、それを戻すために増やす。
	// そして相手の手番になるので　評価値をひっくり返すぜ（＾ｑ＾）
	pPos.MaterialValue += EvalMaterial(captured)
	pPos.MaterialValue = -pPos.MaterialValue

	// 取った駒に関係するのは行き先だけ（＾～＾）
	from, to, _ := move.Destructure()
	// fmt.Printf("Debug: to=%d\n", to)

	var hand_sq = p.SQUARE_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.BuildType)

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
	case p.SQ_K1, p.SQ_R1, p.SQ_B1, p.SQ_G1, p.SQ_S1, p.SQ_N1, p.SQ_L1, p.SQ_P1, p.SQ_K2, p.SQ_R2, p.SQ_B2, p.SQ_G2, p.SQ_S2, p.SQ_N2, p.SQ_L2, p.SQ_P2:
		// 打で取れる駒はないぜ（＾～＾）
		// fmt.Printf("Debug: Drop from=%d\n", from)
	default:
		// 打でないなら
		// fmt.Printf("Debug: Not drop from=%d\n", from)

		// 取った相手の駒があれば、自分の駒台から下ろします
		switch captured {
		case p.PIECE_EMPTY: // Ignored
		case p.PIECE_K1: // Second player win
			hand_sq = p.SQ_K2
		case p.PIECE_R1, p.PIECE_PR1:
			hand_sq = p.SQ_R2
		case p.PIECE_B1, p.PIECE_PB1:
			hand_sq = p.SQ_B2
		case p.PIECE_G1:
			hand_sq = p.SQ_G2
		case p.PIECE_S1, p.PIECE_PS1:
			hand_sq = p.SQ_S2
		case p.PIECE_N1, p.PIECE_PN1:
			hand_sq = p.SQ_N2
		case p.PIECE_L1, p.PIECE_PL1:
			hand_sq = p.SQ_L2
		case p.PIECE_P1, p.PIECE_PP1:
			hand_sq = p.SQ_P2
		case p.PIECE_K2: // First player win
			hand_sq = p.SQ_K1
		case p.PIECE_R2, p.PIECE_PR2:
			hand_sq = p.SQ_R1
		case p.PIECE_B2, p.PIECE_PB2:
			hand_sq = p.SQ_B1
		case p.PIECE_G2:
			hand_sq = p.SQ_G1
		case p.PIECE_S2, p.PIECE_PS2:
			hand_sq = p.SQ_S1
		case p.PIECE_N2, p.PIECE_PN2:
			hand_sq = p.SQ_N1
		case p.PIECE_L2, p.PIECE_PL2:
			hand_sq = p.SQ_L1
		case p.PIECE_P2, p.PIECE_PP2:
			hand_sq = p.SQ_P1
		default:
			fmt.Printf("Error: Unknown captured=[%d]", captured)
		}

		// fmt.Printf("Debug: hand_sq=%d\n", hand_sq)

		if hand_sq != p.SQUARE_EMPTY {
			pPos.Hands1[hand_sq-p.SQ_HAND_START] -= 1

			// 取っていた駒を行き先に戻します
			cap_piece_type = What(captured)
			pPos.Board[to] = captured

			ValidateThereArePieceIn(pPos, to)
			// fmt.Printf("Debug: ph=%d\n", ph)
			var pCB6 *ControlBoard
			if pBrain.BuildType == BUILD_DEV {
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
		case p.FIRST:
			// 後手の玉
			pPos.PieceLocations[p.PCLOC_K2] = to
		case p.SECOND:
			// 先手の玉
			pPos.PieceLocations[p.PCLOC_K1] = to
		default:
			panic(G.Log.Fatal("Unknown pBrain.PPosSys.phase=%d", pBrain.PPosSys.phase))
		}
	case PIECE_TYPE_R, PIECE_TYPE_PR:
		for i := p.PCLOC_R1; i < p.PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case PIECE_TYPE_B, PIECE_TYPE_PB:
		for i := p.PCLOC_B1; i < p.PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case PIECE_TYPE_L, PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := p.PCLOC_L1; i < p.PCLOC_L4+1; i += 1 {
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

	pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.BuildType)
}
