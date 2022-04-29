package take12

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l11 "github.com/muzudho/kifuwarabe-wcsc31/take11"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
	l09 "github.com/muzudho/kifuwarabe-wcsc31/take9"
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
func FlipPhase(phase l06.Phase) l06.Phase {
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

// PositionSystem - 局面にいろいろな機能を付けたもの
type PositionSystem struct {
	// 局面
	PPosition [POS_LAYER_SIZE]*Position

	// 利きボード・システム
	PControlBoardSystem *ControlBoardSystem

	// 先手が1、後手が2（＾～＾）
	phase l06.Phase
	// 開始局面の時点で何手目か（＾～＾）これは表示のための飾りのようなものだぜ（＾～＾）
	StartMovesNum int
	// 開始局面から数えて何手目か（＾～＾）0から始まるぜ（＾～＾）
	OffsetMovesIndex int
	// 指し手のリスト（＾～＾）
	// 1手目は[0]へ、512手目は[511]へ入れろだぜ（＾～＾）
	Moves [MOVES_SIZE]Move
	// 取った駒のリスト（＾～＾）アンドゥ ムーブするときに使うだけ（＾～＾）指し手のリストと同じ添え字を使うぜ（＾～＾）
	CapturedList [l09.MOVES_SIZE]l03.Piece
}

func NewPositionSystem() *PositionSystem {
	var pPosSys = new(PositionSystem)

	pPosSys.PPosition = [POS_LAYER_SIZE]*Position{NewPosition(), NewPosition(), NewPosition(), NewPosition()}

	pPosSys.resetPosition()
	return pPosSys
}

// FlipPhase - フェーズをひっくり返すぜ（＾～＾）
func (pPosSys *PositionSystem) FlipPhase() {
	pPosSys.phase = FlipPhase(pPosSys.phase)
}

// GetPhase - フェーズ
func (pPosSys *PositionSystem) GetPhase() l06.Phase {
	return pPosSys.phase
}

// ResetToStartpos - 駒を置いていな状態でリセットします
func (pPosSys *PositionSystem) resetPosition() {
	pPosSys.PControlBoardSystem = NewControlBoardSystem()

	// 先手の局面
	pPosSys.phase = l06.FIRST
	// 何手目か
	pPosSys.StartMovesNum = 1
	pPosSys.OffsetMovesIndex = 0
	// 指し手のリスト
	pPosSys.Moves = [MOVES_SIZE]Move{}
	// 取った駒のリスト
	pPosSys.CapturedList = [MOVES_SIZE]l03.Piece{}
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
				pPos.Board[file*10+rank] = l03.FromStringToPiece(string(pc))
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
					pPos.Board[file*10+rank] = l03.FromStringToPiece("+" + string(pc2))
					file -= 1
					i += 1
				default:
					panic("Undefined sfen board+")
				}
			}

			// 玉と、長い利きの駒は位置を覚えておくぜ（＾～＾）
			switch command[i-1] {
			case 'K':
				pPos.PieceLocations[l11.PCLOC_K1] = l03.Square((file+1)*10 + rank)
			case 'k':
				pPos.PieceLocations[l11.PCLOC_K2] = l03.Square((file+1)*10 + rank)
			case 'R', 'r': // 成も兼ねてる（＾～＾）
				for i := l11.PCLOC_R1; i < l11.PCLOC_R2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'B', 'b':
				for i := l11.PCLOC_B1; i < l11.PCLOC_B2+1; i += 1 {
					sq := pPos.PieceLocations[i]
					if sq == l03.SQ_EMPTY {
						pPos.PieceLocations[i] = SquareFrom(file+1, rank)
						break
					}
				}
			case 'L', 'l':
				for i := l11.PCLOC_L1; i < l11.PCLOC_L4+1; i += 1 {
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
			pPosSys.phase = l06.FIRST
			i += 1
		case 'w':
			pPosSys.phase = l06.SECOND
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
							for i := l11.PCLOC_R1; i < l11.PCLOC_R2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == l03.SQ_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = l03.Square(handIndex) + l03.SQ_HAND_START
									break
								}
							}
						case l03.HAND_B1, l03.HAND_B2:
							for i := l11.PCLOC_B1; i < l11.PCLOC_B2+1; i += 1 {
								sq := pPos.PieceLocations[i]
								if sq == l03.SQ_EMPTY { // 空いているところから埋めていくぜ（＾～＾）
									pPos.PieceLocations[i] = l03.Square(handIndex) + l03.SQ_HAND_START
									break
								}
							}
						case l03.HAND_L1, l03.HAND_L2:
							for i := l11.PCLOC_L1; i < l11.PCLOC_L4+1; i += 1 {
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
						panic(fmt.Errorf("fatal: Unknown number character=%c", piece))
					}

				} else {
					panic(fmt.Errorf("fatal: unknown piece=%c", piece))
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
			var move, err = ParseMove(command, &i, pPosSys.GetPhase())
			if err != nil {
				fmt.Println(err)
				fmt.Println(Sprint(
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

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff()

	// 開始局面の利きを計算（＾～＾）
	//fmt.Printf("Debug: 開始局面の利きを計算（＾～＾）\n")
	for sq := l03.Square(11); sq < 100; sq += 1 {
		if l03.File(sq) != 0 && l03.Rank(sq) != 0 {
			if !pPos.IsEmptySq(sq) {
				//fmt.Printf("Debug: sq=%d\n", sq)
				// あとですぐクリアーするので、どのレイヤー使ってても関係ないんで、仮で PUTレイヤーを使っているぜ（＾～＾）
				pPosSys.PControlBoardSystem.AddControlDiff(
					pPos, CONTROL_LAYER_DIFF1_PUT, CONTROL_LAYER_DIFF2_PUT, sq, 1)
			}
		}
	}
	//fmt.Printf("Debug: 開始局面の利き計算おわり（＾～＾）\n")
	pPosSys.PControlBoardSystem.MergeControlDiff()

	// 読込んだ Move を、上書きする感じで、もう一回 全て実行（＾～＾）
	moves_size := pPosSys.OffsetMovesIndex
	// 一旦 0 リセットするぜ（＾～＾）
	pPosSys.OffsetMovesIndex = 0
	pPosSys.phase = start_phase
	for i = 0; i < moves_size; i += 1 {
		pPosSys.DoMove(pPos, pPosSys.Moves[i])
	}
}

// ParseMove - 指し手コマンドを解析
func ParseMove(command string, i *int, phase l06.Phase) (Move, error) {
	var len = len(command)
	var move = RESIGN_MOVE

	var hand_sq = l03.SQ_EMPTY

	// file
	switch ch := command[*i]; ch {
	case 'R':
		hand_sq = l03.SQ_R1
	case 'B':
		hand_sq = l03.SQ_B1
	case 'G':
		hand_sq = l03.SQ_G1
	case 'S':
		hand_sq = l03.SQ_S1
	case 'N':
		hand_sq = l03.SQ_N1
	case 'L':
		hand_sq = l03.SQ_L1
	case 'P':
		hand_sq = l03.SQ_P1
	default:
		// Ignored
	}

	// 0=移動元 1=移動先
	var count = 0

	if hand_sq != l03.SQ_EMPTY {
		*i += 1
		switch phase {
		case l06.FIRST:
			move = move.ReplaceSource(hand_sq)
		case l06.SECOND:
			move = move.ReplaceSource(hand_sq + l03.HAND_TYPE_SIZE_SQ)
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

			sq := l03.Square(file*10 + rank)
			if count == 0 {
				move = move.ReplaceSource(sq)
			} else if count == 1 {
				move = move.ReplaceDestination(sq)
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
		move = move.ReplacePromotion(true)
	}

	return move, nil
}

// DoMove - 一手指すぜ（＾～＾）
func (pPosSys *PositionSystem) DoMove(pPos *Position, move Move) {
	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := l11.PIECE_TYPE_EMPTY
	cap_piece_type := l11.PIECE_TYPE_EMPTY

	from, to, pro := move.Destructure()

	if pPos.IsEmptySq(from) {
		// 人間の打鍵ミスか（＾～＾）
		fmt.Printf("Error: %d square is empty\n", from)
	}

	var cap_src_sq l03.Square
	var cap_dst_sq = l03.SQ_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただし今から動かす駒を除きます。
	pPosSys.PControlBoardSystem.AddControlRook(
		pPos, CONTROL_LAYER_DIFF1_ROOK_OFF, CONTROL_LAYER_DIFF2_ROOK_OFF, -1, from)
	pPosSys.PControlBoardSystem.AddControlBishop(
		pPos, CONTROL_LAYER_DIFF1_BISHOP_OFF, CONTROL_LAYER_DIFF2_BISHOP_OFF, -1, from)
	pPosSys.PControlBoardSystem.AddControlLance(
		pPos, CONTROL_LAYER_DIFF1_LANCE_OFF, CONTROL_LAYER_DIFF2_LANCE_OFF, -1, from)

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
		pPosSys.PControlBoardSystem.AddControlDiff(
			pPos, CONTROL_LAYER_DIFF1_PUT, CONTROL_LAYER_DIFF2_PUT, to, 1)
		mov_piece_type = l11.What(piece)
	} else {
		// 打でないなら

		// 移動先に駒があれば、その駒の利きを除外します。
		captured := pPos.Board[to]
		if captured != l03.PIECE_EMPTY {
			pieceType := l11.What(captured)
			switch pieceType {
			case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR, l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB, l11.PIECE_TYPE_L:
				// Ignored: 長い利きの駒は 既に除外しているので無視します
			default:
				pPosSys.PControlBoardSystem.AddControlDiff(
					pPos, CONTROL_LAYER_DIFF1_CAPTURED, CONTROL_LAYER_DIFF2_CAPTURED, to, -1)
			}
			cap_piece_type = l11.What(captured)
			cap_src_sq = to
		}

		// 元位置の駒の利きを除去
		pPosSys.PControlBoardSystem.AddControlDiff(
			pPos, CONTROL_LAYER_DIFF1_REMOVE, CONTROL_LAYER_DIFF2_REMOVE, from, -1)

		// 行き先の駒の上書き
		if pro {
			// 駒を成りに変換します
			pPos.Board[to] = l09.Promote(pPos.Board[from])
		} else {
			pPos.Board[to] = pPos.Board[from]
		}
		mov_piece_type = l11.What(pPos.Board[to])
		// 元位置の駒を削除してから、移動先の駒の利きを追加
		pPos.Board[from] = l03.PIECE_EMPTY
		pPosSys.PControlBoardSystem.AddControlDiff(
			pPos, CONTROL_LAYER_DIFF1_PUT, CONTROL_LAYER_DIFF2_PUT, to, 1)

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
		case l03.PIECE_K2: // l06.FIRST player win
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
	prev_phase := pPosSys.GetPhase()
	pPosSys.FlipPhase()

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	piece_type_list := []l11.PieceType{mov_piece_type, cap_piece_type}
	src_sq_list := []l03.Square{from, cap_src_sq}
	dst_sq_list := []l03.Square{to, cap_dst_sq}
	for j, piece_type := range piece_type_list {
		switch piece_type {
		case l11.PIECE_TYPE_K:
			if j == 0 {
				switch prev_phase {
				case l06.FIRST:
					pPos.PieceLocations[l11.PCLOC_K1] = dst_sq_list[j]
				case l06.SECOND:
					pPos.PieceLocations[l11.PCLOC_K2] = dst_sq_list[j]
				default:
					panic(fmt.Errorf("unknown prev_phase=%d", prev_phase))
				}
			} else {
				// 取った時
				switch prev_phase {
				case l06.FIRST:
					// 相手玉
					pPos.PieceLocations[l11.PCLOC_K2] = dst_sq_list[j]
				case l06.SECOND:
					pPos.PieceLocations[l11.PCLOC_K1] = dst_sq_list[j]
				default:
					panic(fmt.Errorf("unknown prev_phase=%d", prev_phase))
				}
			}
		case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
			for i := l11.PCLOC_R1; i < l11.PCLOC_R2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
			for i := l11.PCLOC_B1; i < l11.PCLOC_B2+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
			for i := l11.PCLOC_L1; i < l11.PCLOC_L4+1; i += 1 {
				sq := pPos.PieceLocations[i]
				if sq == src_sq_list[j] {
					pPos.PieceLocations[i] = dst_sq_list[j]
					break
				}
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし動かした駒を除きます
	pPosSys.PControlBoardSystem.AddControlLance(
		pPos, CONTROL_LAYER_DIFF1_LANCE_ON, CONTROL_LAYER_DIFF2_LANCE_ON, 1, to)
	pPosSys.PControlBoardSystem.AddControlBishop(
		pPos, CONTROL_LAYER_DIFF1_BISHOP_ON, CONTROL_LAYER_DIFF2_BISHOP_ON, 1, to)
	pPosSys.PControlBoardSystem.AddControlRook(
		pPos, CONTROL_LAYER_DIFF1_ROOK_ON, CONTROL_LAYER_DIFF2_ROOK_ON, 1, to)

	pPosSys.PControlBoardSystem.MergeControlDiff()
}

// UndoMove - 棋譜を頼りに１手戻すぜ（＾～＾）
func (pPosSys *PositionSystem) UndoMove(pPos *Position) {

	// App.Log.Trace(pPosSys.Sprint())

	if pPosSys.OffsetMovesIndex < 1 {
		return
	}

	// １手指すと１～２の駒が動くことに着目してくれだぜ（＾～＾）
	// 動かしている駒と、取った駒だぜ（＾～＾）
	mov_piece_type := l11.PIECE_TYPE_EMPTY

	// 先に 手目 を１つ戻すぜ（＾～＾）UndoMoveでフェーズもひっくり返すぜ（＾～＾）
	pPosSys.OffsetMovesIndex -= 1
	move := pPosSys.Moves[pPosSys.OffsetMovesIndex]
	// next_phase := pPosSys.GetPhase()
	pPosSys.FlipPhase()

	from, to, pro := move.Destructure()

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	pPosSys.PControlBoardSystem.AddControlRook(
		pPos, CONTROL_LAYER_DIFF1_ROOK_ON, CONTROL_LAYER_DIFF2_ROOK_ON, -1, to)
	pPosSys.PControlBoardSystem.AddControlBishop(
		pPos, CONTROL_LAYER_DIFF1_BISHOP_ON, CONTROL_LAYER_DIFF2_BISHOP_ON, -1, to)
	pPosSys.PControlBoardSystem.AddControlLance(
		pPos, CONTROL_LAYER_DIFF1_LANCE_ON, CONTROL_LAYER_DIFF2_LANCE_ON, -1, to)

	// 打かどうかで分けます
	switch from {
	case l03.SQ_K1, l03.SQ_R1, l03.SQ_B1, l03.SQ_G1, l03.SQ_S1, l03.SQ_N1, l03.SQ_L1, l03.SQ_P1, l03.SQ_K2, l03.SQ_R2, l03.SQ_B2, l03.SQ_G2, l03.SQ_S2, l03.SQ_N2, l03.SQ_L2, l03.SQ_P2:
		// 打なら
		hand := from
		// 行き先から駒を除去します
		mov_piece_type = l11.What(pPos.Board[to])
		pPosSys.PControlBoardSystem.AddControlDiff(
			pPos, CONTROL_LAYER_DIFF1_PUT, CONTROL_LAYER_DIFF2_PUT, to, -1)
		pPos.Board[to] = l03.PIECE_EMPTY

		// 駒台に駒を戻します
		pPos.Hands1[hand-l03.SQ_HAND_START] += 1
	default:
		// 打でないなら

		// 行き先に進んでいた自駒の利きの除去
		mov_piece_type = l11.What(pPos.Board[to])
		pPosSys.PControlBoardSystem.AddControlDiff(
			pPos, CONTROL_LAYER_DIFF1_PUT, CONTROL_LAYER_DIFF2_PUT, to, -1)

		// 自駒を移動元へ戻します
		if pro {
			// 成りを元に戻します
			pPos.Board[from] = l09.Demote(pPos.Board[to])
		} else {
			pPos.Board[from] = pPos.Board[to]
		}

		pPos.Board[to] = l03.PIECE_EMPTY

		// 元の場所に戻した自駒の利きを復元します
		pPosSys.PControlBoardSystem.AddControlDiff(
			pPos, CONTROL_LAYER_DIFF1_REMOVE, CONTROL_LAYER_DIFF2_REMOVE, from, 1)
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch mov_piece_type {
	case l11.PIECE_TYPE_K:
		// 玉を動かした
		switch pPosSys.phase { // next_phase
		case l06.FIRST:
			pPos.PieceLocations[l11.PCLOC_K1] = from
		case l06.SECOND:
			pPos.PieceLocations[l11.PCLOC_K2] = from
		default:
			panic(fmt.Errorf("unknown p_pos_sys.phase=%d", pPosSys.phase))
		}
	case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
		for i := l11.PCLOC_R1; i < l11.PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
		for i := l11.PCLOC_B1; i < l11.PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := l11.PCLOC_L1; i < l11.PCLOC_L4+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == to {
				pPos.PieceLocations[i] = from
				break
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	pPosSys.PControlBoardSystem.AddControlLance(
		pPos, CONTROL_LAYER_DIFF1_LANCE_OFF, CONTROL_LAYER_DIFF2_LANCE_OFF, 1, from)
	pPosSys.PControlBoardSystem.AddControlBishop(
		pPos, CONTROL_LAYER_DIFF1_BISHOP_OFF, CONTROL_LAYER_DIFF2_BISHOP_OFF, 1, from)
	pPosSys.PControlBoardSystem.AddControlRook(
		pPos, CONTROL_LAYER_DIFF1_ROOK_OFF, CONTROL_LAYER_DIFF2_ROOK_OFF, 1, from)

	pPosSys.PControlBoardSystem.MergeControlDiff()

	// 取った駒を戻すぜ（＾～＾）
	pPosSys.undoCapture(pPos)
}

// undoCapture - 取った駒を戻すぜ（＾～＾）
func (pPosSys *PositionSystem) undoCapture(pPos *Position) {
	// App.Log.Trace(pPosSys.Sprint())

	// 取った駒だぜ（＾～＾）
	cap_piece_type := l11.PIECE_TYPE_EMPTY

	// 手目もフェーズもすでに１つ戻っているとするぜ（＾～＾）
	move := pPosSys.Moves[pPosSys.OffsetMovesIndex]

	// 取った駒
	captured := pPosSys.CapturedList[pPosSys.OffsetMovesIndex]
	// fmt.Printf("Debug: CapturedPiece=%s\n", captured.ToCode())

	from, to, _ := move.Destructure()

	// 取った駒に関係するのは行き先だけ（＾～＾）
	var hand_sq = l03.SQ_EMPTY

	// 利きの差分テーブルをクリアー（＾～＾）
	pPosSys.PControlBoardSystem.ClearControlDiff()

	// 作業前に、長い利きの駒の利きを -1 します。ただしこれから動かす駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	pPosSys.PControlBoardSystem.AddControlRook(
		pPos, CONTROL_LAYER_DIFF1_ROOK_ON, CONTROL_LAYER_DIFF2_ROOK_ON, -1, to)
	pPosSys.PControlBoardSystem.AddControlBishop(
		pPos, CONTROL_LAYER_DIFF1_BISHOP_ON, CONTROL_LAYER_DIFF2_BISHOP_ON, -1, to)
	pPosSys.PControlBoardSystem.AddControlLance(
		pPos, CONTROL_LAYER_DIFF1_LANCE_ON, CONTROL_LAYER_DIFF2_LANCE_ON, -1, to)

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
		case l03.PIECE_K2: // l06.FIRST player win
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
			cap_piece_type = l11.What(captured)
			pPos.Board[to] = captured

			// 取った駒は盤上になかったので、ここで利きを復元させます
			// 行き先にある取られていた駒の利きの復元
			pPosSys.PControlBoardSystem.AddControlDiff(
				pPos, CONTROL_LAYER_DIFF1_CAPTURED, CONTROL_LAYER_DIFF2_CAPTURED, to, 1)
		}
	}

	// 玉と、長い利きの駒が動いたときは、位置情報更新
	switch cap_piece_type {
	case l11.PIECE_TYPE_K:
		// 玉を取っていた
		switch pPosSys.phase { // next_phase
		case l06.FIRST:
			// 後手の玉
			pPos.PieceLocations[l11.PCLOC_K2] = to
		case l06.SECOND:
			// 先手の玉
			pPos.PieceLocations[l11.PCLOC_K1] = to
		default:
			panic(fmt.Errorf("unknown p_pos_sys.phase=%d", pPosSys.phase))
		}
	case l11.PIECE_TYPE_R, l11.PIECE_TYPE_PR:
		for i := l11.PCLOC_R1; i < l11.PCLOC_R2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case l11.PIECE_TYPE_B, l11.PIECE_TYPE_PB:
		for i := l11.PCLOC_B1; i < l11.PCLOC_B2+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	case l11.PIECE_TYPE_L, l11.PIECE_TYPE_PL: // 成香も一応、位置を覚えておかないと存在しない香を監視してしまうぜ（＾～＾）
		for i := l11.PCLOC_L1; i < l11.PCLOC_L4+1; i += 1 {
			sq := pPos.PieceLocations[i]
			if sq == hand_sq {
				pPos.PieceLocations[i] = to
				break
			}
		}
	}

	// 作業後に、長い利きの駒の利きをプラス１します。ただし、今動かした駒を除きます
	// アンドゥなので逆さになっているぜ（＾～＾）
	pPosSys.PControlBoardSystem.AddControlLance(
		pPos, CONTROL_LAYER_DIFF1_LANCE_OFF, CONTROL_LAYER_DIFF2_LANCE_OFF, 1, from)
	pPosSys.PControlBoardSystem.AddControlBishop(
		pPos, CONTROL_LAYER_DIFF1_BISHOP_OFF, CONTROL_LAYER_DIFF2_BISHOP_OFF, 1, from)
	pPosSys.PControlBoardSystem.AddControlRook(
		pPos, CONTROL_LAYER_DIFF1_ROOK_OFF, CONTROL_LAYER_DIFF2_ROOK_OFF, 1, from)

	pPosSys.PControlBoardSystem.MergeControlDiff()
}
