package take15

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
	l01 "github.com/muzudho/kifuwarabe-wcsc31/lesson01"
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l08 "github.com/muzudho/kifuwarabe-wcsc31/take8"
)

// App - アプリケーション変数の宣言
var App l01.Lesson01App

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

	// アプリケーション変数の生成
	App = *new(l01.Lesson01App)
	App.IsDebug = true

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
	App.LogNotEcho = *l.NewLogger(
		tracePath,
		debugPath,
		infoPath,
		noticePath,
		warnPath,
		errorPath,
		fatalPath,
		printPath)

	// 既存のログ・ファイルを削除
	App.LogNotEcho.RemoveAllOldLogs()

	// ログ・ファイルの開閉
	err = App.LogNotEcho.OpenAllLogs()
	if err != nil {
		// ログ・ファイルを開くのに失敗したのだから、ログ・ファイルへは書き込めません
		panic(err)
	}
	defer App.LogNotEcho.CloseAllLogs()

	// チャッターの作成。 標準出力とロガーを一緒にしただけです。
	App.Out = *l.NewChatter(App.LogNotEcho)
	App.Log = *l.NewStderrChatter(App.LogNotEcho)

	App.LogNotEcho.Trace("Start Take1\n")
	App.LogNotEcho.Trace("engineConfPath=%s\n", engineConfPath)

	// 設定ファイル読込。ファイルが存在しなければ強制終了してしまうので注意！
	config, err := l01.LoadEngineConf(engineConfPath)
	if err != nil {
		panic(App.LogNotEcho.Fatal(fmt.Sprintf("engineConfPath=[%s] err=[%s]", engineConfPath, err)))
	}

	// 何か標準入力しろだぜ☆（＾～＾）
	scanner := bufio.NewScanner(os.Stdin)

	App.LogNotEcho.FlushAllLogs()

	var pBrain = NewBrain()

MainLoop:
	for scanner.Scan() {
		command := scanner.Text()
		App.LogNotEcho.Trace("command=%s\n", command)

		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "usi":
			// With Build Number
			App.Out.Print("id name %sB24\n", config.Profile.Name)
			App.Out.Print("id author %s\n", config.Profile.Author)
			pBrain.PPosSys.BuildType = BUILD_RELEASE
			// 乱数のタネを変更するぜ（＾～＾）
			rand.Seed(time.Now().UnixNano())
			App.Out.Print("usiok\n")
		case "isready":
			App.Out.Print("readyok\n")
		case "usinewgame":
		case "position":
			// position うわっ、大変だ（＾～＾）
			pBrain.ReadPosition(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], command)
		case "go":
			bestmove := Search(pBrain)
			App.Out.Print("bestmove %s\n", bestmove.ToCodeOfM())
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
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("gameover win\n")
				App.Out.Debug("gameover lose\n")
				App.Out.Debug("gameover draw\n")
			}
		// 以下、きふわらべ独自拡張コマンド
		case "pos":
			// 局面表示
			length := len(tokens)
			ok := false
			if length == 1 {
				// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
				App.Out.Debug(SprintBoard(
					pBrain.PPosSys.PPosition[POS_LAYER_MAIN],
					pBrain.PPosSys.phase,
					pBrain.PPosSys.StartMovesNum,
					pBrain.PPosSys.OffsetMovesIndex,
					pBrain.PPosSys.createMovesText()))
				ok = true
				ok = true
			} else if length == 2 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				} else {
					App.Out.Debug(SprintBoard(
						pBrain.PPosSys.PPosition[b1],
						pBrain.PPosSys.phase,
						pBrain.PPosSys.StartMovesNum,
						pBrain.PPosSys.OffsetMovesIndex,
						pBrain.PPosSys.createMovesText()))
					ok = true
				}
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("pos\n")
				App.Out.Debug("pos {boardNumber}\n")
			}
		case "do":
			// １手指すぜ（＾～＾）
			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			i := 3
			var move, err = ParseMove(command, &i, pBrain.PPosSys.GetPhase())
			if err != nil {
				App.Out.Debug(SprintBoard(
					pBrain.PPosSys.PPosition[POS_LAYER_MAIN],
					pBrain.PPosSys.phase,
					pBrain.PPosSys.StartMovesNum,
					pBrain.PPosSys.OffsetMovesIndex,
					pBrain.PPosSys.createMovesText()))
				panic(err)
			}

			pBrain.DoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], move)
		case "undo":
			// 棋譜を頼りに１手戻すぜ（＾～＾）
			pBrain.UndoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
		case "control":
			length := len(tokens)
			// fmt.Printf("length=%d", length)
			ok := false
			if length == 1 {
				// 利きの表示（＾～＾）
				App.Out.Debug(pBrain.PCtrlBrdSys.SprintControl(CONTROL_LAYER_SUM1))
				App.Out.Debug(pBrain.PCtrlBrdSys.SprintControl(CONTROL_LAYER_SUM2))
				ok = true
			} else if length == 2 && tokens[1] == "test" {
				// 利きのテスト
				// 現局面の利きを覚え、ムーブ、アンドゥを行って
				// 元の利きに戻るか確認
				is_error, message := TestControl(pBrain, pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
				if is_error {
					App.Out.Debug("ControlTest: error=%s\n", message)
					App.Out.Debug(pBrain.PCtrlBrdSys.SprintControl(CONTROL_LAYER_TEST_ERROR1))
					App.Out.Debug(pBrain.PCtrlBrdSys.SprintControl(CONTROL_LAYER_TEST_ERROR2))
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

				pBrain.PCtrlBrdSys.DiffControl(ControlLayerT(c1), ControlLayerT(c2), ControlLayerT(c3))
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

				pBrain.PCtrlBrdSys.RecalculateControl(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], ControlLayerT(c1), ControlLayerT(c2))
				ok = true
			} else if length == 3 && tokens[1] == "layer" {
				// 指定の利きテーブルの表示（＾～＾）
				c1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				} else if 0 <= c1 && ControlLayerT(c1) < CONTROL_LAYER_ALL_SIZE {
					App.Out.Debug(pBrain.PCtrlBrdSys.SprintControl(ControlLayerT(c1)))
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
				sumList := SumAbsControl(pBrain, ControlLayerT(c1), ControlLayerT(c2))
				App.Out.Debug("ControlTest: SumAbs=%d,%d\n", sumList[0], sumList[1])
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("control\n")
				App.Out.Debug("control layer {number}\n")
				App.Out.Debug("control recalc {number} {number}\n")
				App.Out.Debug("control diff {layer_number} {layer_number} {layer_number}\n")
				App.Out.Debug("control sumabs {number} {number}\n")
			}
		case "location":
			length := len(tokens)
			ok := false
			if length == 2 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				// あの駒、どこにいんの（＾～＾）？
				App.Out.Debug(l08.SprintLocation(pBrain.PPosSys.PPosition[PosLayerT(b1)]))
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("location {boardLayerIndex}\n")
			}
		case "sfen":
			// SFEN文字列返せよ（＾～＾）
			App.Out.Debug(pBrain.PPosSys.SprintSfenResignation(pBrain.PPosSys.PPosition[POS_LAYER_MAIN]))
		case "record":
			// 棋譜表示。取った駒を表示するためのもの（＾～＾）
			App.Out.Debug(pBrain.PPosSys.SprintRecord())
		case "movelist":
			// 指し手の一覧
			moveList(pBrain)
		case "dump":
			// 変数を全部出力してくれだぜ（＾～＾）
			App.Out.Debug("PositionSystem.Dump()\n")
			App.Out.Debug("---------------\n%s", pBrain.Dump())
		case "playout":
			// とにかく手を進めるぜ（＾～＾）
			// 時間の計測は リリース・モードでやれだぜ（＾～＾）
			App.Out.Debug("Playout start\n")
			start := time.Now()

		PlayoutLoop:
			// 棋譜を書き直してさらに多く続けるぜ（＾～＾）
			for j := 0; j < 1000; j += 1 {
				// 512手が最大だが（＾～＾）
				for i := 0; i < l04.MOVES_SIZE; i += 1 {
					App.Out.Debug(SprintBoard(
						pBrain.PPosSys.PPosition[POS_LAYER_MAIN],
						pBrain.PPosSys.phase,
						pBrain.PPosSys.StartMovesNum,
						pBrain.PPosSys.OffsetMovesIndex,
						pBrain.PPosSys.createMovesText()))
					// あの駒、どこにいんの（＾～＾）？
					// App.Out.Debug(SprintLocation(pBrain.PPosSys))

					// moveList(pBrain.PPosSys)
					bestmove := Search(pBrain)
					App.Out.Print("bestmove %s\n", bestmove.ToCodeOfM())

					if bestmove == l03.Move(l03.SQ_EMPTY) {
						// 投了
						break PlayoutLoop
					}

					pBrain.DoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], bestmove)
				}

				sfen1 := pBrain.PPosSys.SprintSfenResignation(pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
				pBrain.ReadPosition(pBrain.PPosSys.PPosition[POS_LAYER_MAIN], sfen1)

				// ここを開始局面ということにするぜ（＾～＾）
				// pBrain.PPosSys.StartMovesNum = 0
			}

			end := time.Now()
			App.Out.Debug("Playout finished。%f seconds\n", (end.Sub(start)).Seconds())
		case "shuffle":
			ShuffleBoard(pBrain, pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
		case "count":
			ShowAllPiecesCount(pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
		case "board":
			length := len(tokens)
			ok := false
			if length == 4 && tokens[1] == "copy" {
				// 盤番号
				b1, err := strconv.Atoi(tokens[2])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[3])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				copyBoard(pBrain.PPosSys.PPosition[b1], pBrain.PPosSys.PPosition[b2])
				ok = true
			} else if length == 2 && tokens[1] == "diff" {
				diffBoard(pBrain.PPosSys.PPosition[0], pBrain.PPosSys.PPosition[1], pBrain.PPosSys.PPosition[2], pBrain.PPosSys.PPosition[3])
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("board copy {boardLayerIndex} {boardLayerIndex}\n")
			}
		case "posdiff":
			length := len(tokens)
			ok := false
			if length == 3 {
				// 盤番号
				b1, err := strconv.Atoi(tokens[1])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[2])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				App.Out.Debug(pBrain.PPosSys.SprintDiff(PosLayerT(b1), PosLayerT(b2)))
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("posdiff {boardIndex1} {boardIndex2}\n")
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
					App.Out.Debug("Error: %s", err)
				}

				b1, err := strconv.Atoi(tokens[3])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[4])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b3, err := strconv.Atoi(tokens[5])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				errorNum := errorBoard(pBrain.PPosSys.PPosition[b0], pBrain.PPosSys.PPosition[b1], pBrain.PPosSys.PPosition[b2], pBrain.PPosSys.PPosition[b3])
				if errorNum == 0 {
					App.Out.Debug("ok\n")
				} else {
					App.Out.Debug("error=%d\n", errorNum)
				}
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("error board {*1} {*2} {*3} {*4}\n")
				App.Out.Debug("    *1 boardLayerIndex Compare 1\n")
				App.Out.Debug("    *2 boardLayerIndex Compare 2\n")
				App.Out.Debug("    *3 boardLayerIndex Temp\n")
				App.Out.Debug("    *4 boardLayerIndex Temp\n")
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
					App.Out.Debug("Error: %s", err)
				}

				b2, err := strconv.Atoi(tokens[2])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b3, err := strconv.Atoi(tokens[3])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b4, err := strconv.Atoi(tokens[4])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				b5, err := strconv.Atoi(tokens[5])
				if err != nil {
					App.Out.Debug("Error: %s", err)
				}

				WaterColor(
					pBrain.PCtrlBrdSys.PBoards[b1],
					pBrain.PCtrlBrdSys.PBoards[b2],
					pBrain.PCtrlBrdSys.PBoards[b3],
					pBrain.PCtrlBrdSys.PBoards[b4],
					pBrain.PCtrlBrdSys.PBoards[b5])
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("watercolor {control1} {control2} {control3} {control4} {control5}\n")
			}
		case "dev":
			// 乱数のタネを0固定（＾～＾）
			rand.Seed(0)
		case "value":
			// 現局面の評価値を表示（＾～＾）
			App.Out.Debug("Value\n")
			App.Out.Debug("-----\n")
			App.Out.Debug("MaterialValue=%d\n", pBrain.PPosSys.PPosition[POS_LAYER_MAIN].MaterialValue)
		case "":
			// Ignored
		default:
			// 将棋所からいろいろメッセージ飛んでくるから、リリースモードでは無視しろだぜ（＾～＾）
			if pBrain.PPosSys.BuildType == BUILD_DEV {
				fmt.Printf("Unknown command=%s\n", command)
			}
		}

		App.LogNotEcho.FlushAllLogs()
	}

	App.LogNotEcho.Trace("Finished\n")
	App.LogNotEcho.FlushAllLogs()
}

// moveList - 指し手リスト出力
func moveList(pBrain *Brain) {
	App.Out.Debug("MoveList\n")
	App.Out.Debug("--------\n")
	move_list := GenMoveList(pBrain, pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
	for i, move := range move_list {
		var pPos = pBrain.PPosSys.PPosition[POS_LAYER_MAIN]
		pBrain.DoMove(pPos, move)
		App.Out.Debug("(%3d) %-5s . %11d value\n", i, move.ToCodeOfM(), pPos.MaterialValue)
		pBrain.UndoMove(pBrain.PPosSys.PPosition[POS_LAYER_MAIN])
		// App.Out.Debug("(%3d) Undo  . %11d value\n", i, pPos.MaterialValue) // Debug
	}
	App.Out.Debug("* Except for those to be removed during the search\n")
}

// ShowAllPiecesCount - 駒の枚数表示
func ShowAllPiecesCount(pPos *Position) {
	countList := CountAllPieces(pPos)
	App.Out.Debug("Count\n")
	App.Out.Debug("-----\n")
	App.Out.Debug("King  :%3d\n", countList[0])
	App.Out.Debug("Rook  :%3d\n", countList[1])
	App.Out.Debug("Bishop:%3d\n", countList[2])
	App.Out.Debug("Gold  :%3d\n", countList[3])
	App.Out.Debug("Silver:%3d\n", countList[4])
	App.Out.Debug("Knight:%3d\n", countList[5])
	App.Out.Debug("Lance :%3d\n", countList[6])
	App.Out.Debug("Pawn  :%3d\n", countList[7])
	App.Out.Debug("----------\n")
	App.Out.Debug("Total :%3d\n", countList[0]+countList[1]+countList[2]+countList[3]+countList[4]+countList[5]+countList[6]+countList[7])
}
