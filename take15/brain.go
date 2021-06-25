package take15

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Brain - 局面システムと、利き盤システムの２つを持つもの
type Brain struct {
	// 局面システム
	PPosSys *PositionSystem
	// 利きボード・システム
	PCtrlBrdSys *ControlBoardSystem
}

func NewBrain() *Brain {
	var pBrain = new(Brain)
	pBrain.PPosSys = NewPositionSystem()
	pBrain.PCtrlBrdSys = NewControlBoardSystem()
	return pBrain
}

// ReadPosition - 局面を読み取ります。マルチバイト文字は含まれていないぜ（＾ｑ＾）
func (pBrain *Brain) ReadPosition(pPos *Position, command string) {
	var len = len(command)
	var i int
	if strings.HasPrefix(command, "position startpos") {
		// 平手初期局面をセット（＾～＾）
		pPos.clearBoard()
		pBrain.PCtrlBrdSys = NewControlBoardSystem()
		pBrain.PPosSys.resetPosition()
		pPos.setToStartpos()
		i = 17

		if i < len && command[i] == ' ' {
			i += 1
		}
		// moves へ続くぜ（＾～＾）

	} else if strings.HasPrefix(command, "position sfen ") {
		// "position sfen " のはずだから 14 文字飛ばすぜ（＾～＾）
		pPos.clearBoard()
		pBrain.PCtrlBrdSys = NewControlBoardSystem()
		pBrain.PPosSys.resetPosition()
		i = 14
		var rank = Square(1)
		var file = Square(9)

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
					pPos.Board[file*10+rank] = PIECE_EMPTY
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
			pBrain.PPosSys.phase = FIRST
			i += 1
		case 'w':
			pBrain.PPosSys.phase = SECOND
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
						case HAND_R1_IDX, HAND_R2_IDX:
							for i := PCLOC_R1; i < PCLOC_R2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						case HAND_B1_IDX, HAND_B2_IDX:
							for i := PCLOC_B1; i < PCLOC_B2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == SQUARE_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = Square(hand_index) + SQ_HAND_START
									break
								}
							}
						case HAND_L1_IDX, HAND_L2_IDX:
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
						hand_index = HAND_R1_IDX
					case 'B':
						hand_index = HAND_B1_IDX
					case 'G':
						hand_index = HAND_G1_IDX
					case 'S':
						hand_index = HAND_S1_IDX
					case 'N':
						hand_index = HAND_N1_IDX
					case 'L':
						hand_index = HAND_L1_IDX
					case 'P':
						hand_index = HAND_P1_IDX
					case 'r':
						hand_index = HAND_R2_IDX
					case 'b':
						hand_index = HAND_B2_IDX
					case 'g':
						hand_index = HAND_G2_IDX
					case 's':
						hand_index = HAND_S2_IDX
					case 'n':
						hand_index = HAND_N2_IDX
					case 'l':
						hand_index = HAND_L2_IDX
					case 'p':
						hand_index = HAND_P2_IDX
					case ' ':
						// ループを抜けます
						break HandLoop
					default:
						panic(fmt.Errorf("Fatal: Unknown piece=%c", piece))
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
						panic(fmt.Errorf("Fatal: Unknown number character=%c", piece))
					}

				} else {
					panic(fmt.Errorf("Fatal: Unknown piece=%c", piece))
				}
			}
		}

		// 手数
		pBrain.PPosSys.StartMovesNum = 0
	MovesNumLoop:
		for i < len {
			switch figure := command[i]; figure {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				num, err := strconv.Atoi(string(figure))
				if err != nil {
					panic(err)
				}
				i += 1
				pBrain.PPosSys.StartMovesNum *= 10
				pBrain.PPosSys.StartMovesNum += num
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
			var move, err = ParseMove(command, &i, pBrain.PPosSys.GetPhase())
			if err != nil {
				fmt.Println(err)
				fmt.Println(pPos.Sprint(
					pBrain.PPosSys.phase,
					pBrain.PPosSys.StartMovesNum,
					pBrain.PPosSys.OffsetMovesIndex,
					pBrain.PPosSys.createMovesText()))
				panic(err)
			}
			pBrain.PPosSys.Moves[pBrain.PPosSys.OffsetMovesIndex] = move
			pBrain.PPosSys.OffsetMovesIndex += 1
			pBrain.PPosSys.FlipPhase()
		}
	}

	if pBrain.PPosSys.BuildType == BUILD_DEV {
		// 利きの差分テーブルをクリアー（＾～＾）
		pBrain.PCtrlBrdSys.ClearControlDiff(pBrain.PPosSys.BuildType)
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
				if pBrain.PPosSys.BuildType == BUILD_DEV {
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
	if pBrain.PPosSys.BuildType == BUILD_DEV {
		//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
		pBrain.PCtrlBrdSys.MergeControlDiff(pBrain.PPosSys.BuildType)
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
