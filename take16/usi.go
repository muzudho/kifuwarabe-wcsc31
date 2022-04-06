package take16

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	l "github.com/muzudho/go-logger"
)

// MainLoop - 開始。
func MainLoop() {
	// Working directory
	dwd, err := os.Getwd()
	if err != nil {
		// ここでは、ログはまだ設定できてない
		panic(fmt.Sprintf("...Engine DefaultWorkingDirectory=%s", dwd))
	}

	// コマンドライン引数登録
	workdir := flag.String("workdir", dwd, "Working directory path.")
	// コマンドライン引数解析
	flag.Parse()

	engineConfPath := filepath.Join(*workdir, "input/lesson01/engine.conf.toml")

	// グローバル変数の作成
	G = *new(Variables)

	tracePath := filepath.Join(*workdir, "output/trace.log")
	debugPath := filepath.Join(*workdir, "output/debug.log")
	infoPath := filepath.Join(*workdir, "output/info.log")
	noticePath := filepath.Join(*workdir, "output/notice.log")
	warnPath := filepath.Join(*workdir, "output/warn.log")
	errorPath := filepath.Join(*workdir, "output/error.log")
	fatalPath := filepath.Join(*workdir, "output/fatal.log")
	printPath := filepath.Join(*workdir, "output/print.log")

	// ロガーの作成。
	// TODO ディレクトリが存在しなければ、強制終了してしまいます。
	G.Log = *l.NewLogger(
		tracePath,
		debugPath,
		infoPath,
		noticePath,
		warnPath,
		errorPath,
		fatalPath,
		printPath)

	// 既存のログ・ファイルを削除
	G.Log.RemoveAllOldLogs()

	// ログ・ファイルの開閉
	err = G.Log.OpenAllLogs()
	if err != nil {
		// ログ・ファイルを開くのに失敗したのだから、ログ・ファイルへは書き込めません
		panic(err)
	}
	defer G.Log.CloseAllLogs()

	G.Log.Trace("Start Take1\n")
	G.Log.Trace("engineConfPath=%s\n", engineConfPath)

	// チャッターの作成。 標準出力とロガーを一緒にしただけです。
	G.Chat = *l.NewChatter(G.Log)
	G.StderrChat = *l.NewStderrChatter(G.Log)

	// 設定ファイル読込。ファイルが存在しなければ強制終了してしまうので注意！
	config, err := LoadEngineConf(engineConfPath)
	if err != nil {
		panic(G.Log.Fatal(fmt.Sprintf("engineConfPath=[%s] err=[%s]", engineConfPath, err)))
	}

	// 何か標準入力しろだぜ☆（＾～＾）
	scanner := bufio.NewScanner(os.Stdin)

	G.Log.FlushAllLogs()

	var pNerve = NewNerve()

MainLoop:
	for scanner.Scan() {
		command := scanner.Text()
		G.Log.Trace("command=%s\n", command)

		if command == "position startpos moves *0" {
			// 将棋所の連続対局中に
			// 相手が 時間切れを知らずに bestmove を返すと、
			// 将棋所は `isready` など次の対局が始まっている最中に
			// `position startpos moves *0` を返してくる。
			// この `*0` をパースできずに落ちることがあるので、無視するぜ（＾～＾）
			continue
		}

		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "usi":
			// With Build Number
			G.Chat.Print("id name %sB30\n", config.Profile.Name)
			G.Chat.Print("id author %s\n", config.Profile.Author)
			G.Chat.Print("option name MaxDepth type spin default %d min 1 max 15\n", pNerve.MaxDepth)
			// 大会モード
			pNerve.BuildType = BUILD_RELEASE
			// 乱数のタネを変更するぜ（＾～＾）
			rand.Seed(time.Now().UnixNano())
			G.Chat.Print("usiok\n")
		case "isready":
			G.Chat.Print("readyok\n")
		case "usinewgame":
			pNerve = NewNerve()
		case "position":
			// position うわっ、大変だ（＾～＾）
			pNerve.ReadPosition(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], command)
		case "setoption":
			// TODO
			if tokens[1] == "name" {
				// # Example:
				//
				// ```
				// setoption name USI_Ponder value true
				// ```
				name := tokens[2]
				if 5 <= len(tokens) {
					value := tokens[4]
					switch name {
					case "MaxDepth":
						// TODO
						// 0 にすると 1手読み（＾～＾）
						// 1 の 2手読みにしておくと、玉を取りに行くぜ（＾～＾）
						// 2 の 3手読みだと駒を取らない（＾～＾）駒のただ捨てをする（＾～＾）駒をとりかえさない（＾～＾）
						// 3 の 4手読みは、まだ遅い（＾～＾）
						pNerve.MaxDepth, _ = strconv.Atoi(value)
					}
				}
			}
		case "go":
			pNerve.PStopwatchSearch.StartStopwatch()
			bestmove := IterativeDeepeningSearch(pNerve, tokens)
			G.Chat.Print("bestmove %s\n", ToMoveCode(bestmove))
		case "quit":
			break MainLoop
		case "gameover":
			// 時間切れのときなど、将棋所から このメッセージがくるぜ（＾～＾）
			// gameover win
			// gameover lose
			// gameover draw
			length := len(tokens)
			// fmt.Printf("length=%d", length)
			ok := false
			if length == 2 {
				switch tokens[1] {
				case "win":
					ok = true
				case "lose":
					ok = true
				case "draw":
					ok = true
				}
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("gameover win\n")
				G.Chat.Debug("gameover lose\n")
				G.Chat.Debug("gameover draw\n")
			}
		// 以下、きふわらべ独自拡張コマンド
		case "pos":
			// 局面表示
			length := len(tokens)
			ok := false
			if length == 1 {
				// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
				G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoardHeader(
					pNerve.PPosSys.phase,
					pNerve.PRecord.StartMovesNum,
					pNerve.PRecord.OffsetMovesIndex))
				G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoard())
				G.Chat.Debug(pNerve.SprintBoardFooter())
				ok = true
				ok = true
			} else if length == 2 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				} else {
					G.Chat.Debug(pNerve.PPosSys.PPosition[b1].SprintBoardHeader(
						pNerve.PPosSys.phase,
						pNerve.PRecord.StartMovesNum,
						pNerve.PRecord.OffsetMovesIndex))
					G.Chat.Debug(pNerve.PPosSys.PPosition[b1].SprintBoard())
					G.Chat.Debug(pNerve.SprintBoardFooter())
					ok = true
				}
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("pos\n")
				G.Chat.Debug("pos {boardNumber}\n")
			}
		case "do":
			// １手指すぜ（＾～＾）
			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			i := 3
			var move, err = ParseMove(command, &i, pNerve.PPosSys.GetPhase())
			if err != nil {
				G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoardHeader(
					pNerve.PPosSys.phase,
					pNerve.PRecord.StartMovesNum,
					pNerve.PRecord.OffsetMovesIndex))
				G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoard())
				G.Chat.Debug(pNerve.SprintBoardFooter())
				panic(err)
			}

			pNerve.DoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], move)
		case "undo":
			// 棋譜を頼りに１手戻すぜ（＾～＾）
			pNerve.UndoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
		case "control":
			length := len(tokens)
			// fmt.Printf("length=%d", length)
			ok := false
			if length == 1 {
				// 利きの表示（＾～＾）
				G.Chat.Debug(pNerve.PCtrlBrdSys.SprintControl(CONTROL_LAYER_SUM1))
				G.Chat.Debug(pNerve.PCtrlBrdSys.SprintControl(CONTROL_LAYER_SUM2))
				ok = true
			} else if length == 2 && tokens[1] == "test" {
				// 利きのテスト
				// 現局面の利きを覚え、ムーブ、アンドゥを行って
				// 元の利きに戻るか確認
				is_error, message := TestControl(pNerve, pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
				if is_error {
					G.Chat.Debug("ControlTest: error=%s\n", message)
					G.Chat.Debug(pNerve.PCtrlBrdSys.SprintControl(CONTROL_LAYER_TEST_ERROR1))
					G.Chat.Debug(pNerve.PCtrlBrdSys.SprintControl(CONTROL_LAYER_TEST_ERROR2))
				}
				ok = true
			} else if length == 5 && tokens[1] == "diff" {
				// control diff 11 0 12
				// 利きボード 11番 から 0番を引いた結果を 12番へ入れる。
				c1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
				c2, err := strconv.Atoi(tokens[3])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
				c3, err := strconv.Atoi(tokens[4])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				pNerve.PCtrlBrdSys.DiffControl(ControlLayerT(c1), ControlLayerT(c2), ControlLayerT(c3))
				ok = true
			} else if length == 4 && tokens[1] == "recalc" {
				// control recalc 22 25
				// 利きの再計算
				c1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				c2, err := strconv.Atoi(tokens[3])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				pNerve.PCtrlBrdSys.RecalculateControl(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], ControlLayerT(c1), ControlLayerT(c2))
				ok = true
			} else if length == 3 && tokens[1] == "layer" {
				// 指定の利きテーブルの表示（＾～＾）
				c1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				} else if 0 <= c1 && ControlLayerT(c1) < CONTROL_LAYER_ALL_SIZE {
					G.Chat.Debug(pNerve.PCtrlBrdSys.SprintControl(ControlLayerT(c1)))
					ok = true
				}
			} else if length == 4 && tokens[1] == "sumabs" {
				c1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				c2, err := strconv.Atoi(tokens[3])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				// 利きのテスト
				// 現局面の利きを覚え、ムーブ、アンドゥを行って
				// 元の利きに戻るか確認
				sumList := SumAbsControl(pNerve, ControlLayerT(c1), ControlLayerT(c2))
				G.Chat.Debug("ControlTest: SumAbs=%d,%d\n", sumList[0], sumList[1])
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("control\n")
				G.Chat.Debug("control layer {number}\n")
				G.Chat.Debug("control recalc {number} {number}\n")
				G.Chat.Debug("control diff {layer_number} {layer_number} {layer_number}\n")
				G.Chat.Debug("control sumabs {number} {number}\n")
			}
		case "location":
			length := len(tokens)
			ok := false
			if length == 2 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				// あの駒、どこにいんの（＾～＾）？
				G.Chat.Debug(pNerve.PPosSys.PPosition[PosLayerT(b1)].SprintLocation())
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("location {boardLayerIndex}\n")
			}
		case "sfen":
			// SFEN文字列返せよ（＾～＾）
			G.Chat.Debug(sprintSfenResignation(pNerve.PPosSys, pNerve.PPosSys.PPosition[POS_LAYER_MAIN], pNerve.PRecord))
		case "record":
			// 棋譜表示。取った駒を表示するためのもの（＾～＾）
			G.Chat.Debug(sprintRecord(pNerve.PRecord))
		case "movelist":
			// 指し手の一覧
			moveList(pNerve)
		case "dump":
			// 変数を全部出力してくれだぜ（＾～＾）
			G.Chat.Debug("PositionSystem.Dump()\n")
			G.Chat.Debug("---------------\n%s", pNerve.Dump())
		case "playout":
			// とにかく手を進めるぜ（＾～＾）
			// 時間の計測は リリース・モードでやれだぜ（＾～＾）
			G.Chat.Debug("Playout start\n")
			start := time.Now()

		PlayoutLoop:
			// 棋譜を書き直してさらに多く続けるぜ（＾～＾）
			for j := 0; j < 1000; j += 1 {
				// 512手が最大だが（＾～＾）
				for i := 0; i < MOVES_SIZE; i += 1 {
					G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoardHeader(
						pNerve.PPosSys.phase,
						pNerve.PRecord.StartMovesNum,
						pNerve.PRecord.OffsetMovesIndex))
					G.Chat.Debug(pNerve.PPosSys.PPosition[POS_LAYER_MAIN].SprintBoard())
					G.Chat.Debug(pNerve.SprintBoardFooter())
					// あの駒、どこにいんの（＾～＾）？
					// G.Chat.Debug(pNerve.PPosSys.SprintLocation())

					// moveList(pNerve.PPosSys)
					bestmove := IterativeDeepeningSearch(pNerve, []string{"go"})
					G.Chat.Print("bestmove %s\n", ToMoveCode(bestmove))

					if bestmove == Move(SQUARE_EMPTY) {
						// 投了
						break PlayoutLoop
					}

					pNerve.DoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], bestmove)
				}

				sfen1 := sprintSfenResignation(pNerve.PPosSys, pNerve.PPosSys.PPosition[POS_LAYER_MAIN], pNerve.PRecord)
				pNerve.ReadPosition(pNerve.PPosSys.PPosition[POS_LAYER_MAIN], sfen1)

				// ここを開始局面ということにするぜ（＾～＾）
				// pNerve.PPosSys.StartMovesNum = 0
			}

			end := time.Now()
			G.Chat.Debug("Playout finished。%f seconds\n", (end.Sub(start)).Seconds())
		case "shuffle":
			ShuffleBoard(pNerve, pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
		case "count":
			ShowAllPiecesCount(pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
		case "board":
			length := len(tokens)
			ok := false
			if length == 4 && tokens[1] == "copy" {
				// 盤番号
				b1, err := strconv.Atoi(tokens[2])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[3])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				copyBoard(pNerve.PPosSys.PPosition[b1], pNerve.PPosSys.PPosition[b2])
				ok = true
			} else if length == 2 && tokens[1] == "diff" {
				diffBoard(pNerve.PPosSys.PPosition[0], pNerve.PPosSys.PPosition[1], pNerve.PPosSys.PPosition[2], pNerve.PPosSys.PPosition[3])
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("board copy {boardLayerIndex} {boardLayerIndex}\n")
			}
		case "posdiff":
			length := len(tokens)
			ok := false
			if length == 3 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[2])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				G.Chat.Debug(sprintPositionDiff(pNerve.PPosSys, PosLayerT(b1), PosLayerT(b2), pNerve.PRecord))
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("posdiff {boardIndex1} {boardIndex2}\n")
			}
		case "error":
			// 2つのものを比較して、違いが何個あったか返すぜ（＾ｑ＾）
			length := len(tokens)
			ok := false
			if length == 6 && tokens[1] == "board" {
				// 2つの盤を比較するぜ（＾～＾）これを使うには一時テーブルとしてさらに２つ指定しろだぜ（＾～＾）
				// 盤番号
				b0, err := strconv.Atoi(tokens[2])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b1, err := strconv.Atoi(tokens[3])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[4])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b3, err := strconv.Atoi(tokens[5])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				errorNum := errorBoard(pNerve.PPosSys.PPosition[b0], pNerve.PPosSys.PPosition[b1], pNerve.PPosSys.PPosition[b2], pNerve.PPosSys.PPosition[b3])
				if errorNum == 0 {
					G.Chat.Debug("ok\n")
				} else {
					G.Chat.Debug("error=%d\n", errorNum)
				}
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("error board {*1} {*2} {*3} {*4}\n")
				G.Chat.Debug("    *1 boardLayerIndex Compare 1\n")
				G.Chat.Debug("    *2 boardLayerIndex Compare 2\n")
				G.Chat.Debug("    *3 boardLayerIndex Temp\n")
				G.Chat.Debug("    *4 boardLayerIndex Temp\n")
			}
		case "watercolor":
			// 水彩絵の具でにじませたような、利きボード作り
			// watercolor 0 10 26 27 28
			length := len(tokens)
			ok := false
			if length == 6 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[2])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b3, err := strconv.Atoi(tokens[3])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b4, err := strconv.Atoi(tokens[4])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				b5, err := strconv.Atoi(tokens[5])
				if err != nil {
					G.Chat.Debug("Error: %s", err)
				}

				WaterColor(
					pNerve.PCtrlBrdSys.PBoards[b1],
					pNerve.PCtrlBrdSys.PBoards[b2],
					pNerve.PCtrlBrdSys.PBoards[b3],
					pNerve.PCtrlBrdSys.PBoards[b4],
					pNerve.PCtrlBrdSys.PBoards[b5])
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("watercolor {control1} {control2} {control3} {control4} {control5}\n")
			}
		case "dev":
			// 乱数のタネを0固定（＾～＾）
			rand.Seed(0)
		case "value":
			// 現局面の評価値を表示（＾～＾）
			G.Chat.Debug("Value\n")
			G.Chat.Debug("-----\n")
			G.Chat.Debug("MaterialValue=%d\n", pNerve.PPosSys.PPosition[POS_LAYER_MAIN].MaterialValue)
		case "":
			// Ignored
		default:
			// 将棋所からいろいろメッセージ飛んでくるから、リリースモードでは無視しろだぜ（＾～＾）
			if pNerve.BuildType == BUILD_DEV {
				fmt.Printf("Unknown command=%s\n", command)
			}
		}

		G.Log.FlushAllLogs()
	}

	G.Log.Trace("Finished\n")
	G.Log.FlushAllLogs()
}

// moveList - 指し手リスト出力
func moveList(pNerve *Nerve) {
	G.Chat.Debug("MoveList\n")
	G.Chat.Debug("--------\n")
	move_list := GenMoveList(pNerve, pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
	for i, move := range move_list {
		var pPos = pNerve.PPosSys.PPosition[POS_LAYER_MAIN]
		pNerve.DoMove(pPos, move)
		G.Chat.Debug("(%3d) %-5s . %11d value\n", i, ToMoveCode(move), pPos.MaterialValue)
		pNerve.UndoMove(pNerve.PPosSys.PPosition[POS_LAYER_MAIN])
		// G.Chat.Debug("(%3d) Undo  . %11d value\n", i, pPos.MaterialValue) // Debug
	}
	G.Chat.Debug("* Except for those to be removed during the search\n")
}

// ShowAllPiecesCount - 駒の枚数表示
func ShowAllPiecesCount(pPos *Position) {
	countList := CountAllPieces(pPos)
	G.Chat.Debug("Count\n")
	G.Chat.Debug("-----\n")
	G.Chat.Debug("King  :%3d\n", countList[0])
	G.Chat.Debug("Rook  :%3d\n", countList[1])
	G.Chat.Debug("Bishop:%3d\n", countList[2])
	G.Chat.Debug("Gold  :%3d\n", countList[3])
	G.Chat.Debug("Silver:%3d\n", countList[4])
	G.Chat.Debug("Knight:%3d\n", countList[5])
	G.Chat.Debug("Lance :%3d\n", countList[6])
	G.Chat.Debug("Pawn  :%3d\n", countList[7])
	G.Chat.Debug("----------\n")
	G.Chat.Debug("Total :%3d\n", countList[0]+countList[1]+countList[2]+countList[3]+countList[4]+countList[5]+countList[6]+countList[7])
}
