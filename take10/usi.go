package take10

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	l "github.com/muzudho/go-logger"
	l01 "github.com/muzudho/kifuwarabe-wcsc31/lesson01"
	l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
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

	var pPos = NewPosition()

MainLoop:
	for scanner.Scan() {
		command := scanner.Text()
		App.LogNotEcho.Trace("command=%s\n", command)

		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "usi":
			App.Out.Print("id name %s\n", config.Profile.Name)
			App.Out.Print("id author %s\n", config.Profile.Author)
			App.Out.Print("usiok\n")
		case "isready":
			App.Out.Print("readyok\n")
		case "usinewgame":
		case "position":
			// position うわっ、大変だ（＾～＾）
			pPos.ReadPosition(command)
		case "go":
			bestmove := Search(pPos)
			App.Out.Print("bestmove %s\n", bestmove.ToCodeOfM())
		case "quit":
			break MainLoop
		// 以下、きふわらべ独自拡張コマンド
		case "pos":
			// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
			App.Out.Debug(Sprint(pPos))
		case "do":
			// １手指すぜ（＾～＾）
			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			i := 3
			var move, err = ParseMove(command, &i, pPos.GetPhase())
			if err != nil {
				fmt.Println(Sprint(pPos))
				panic(err)
			}

			pPos.DoMove(move)
		case "undo":
			// 棋譜を頼りに１手戻すぜ（＾～＾）
			pPos.UndoMove()
		case "control":
			length := len(tokens)
			// fmt.Printf("length=%d", length)
			ok := false
			if length == 1 {
				// 利きの表示（＾～＾）
				App.Out.Debug(pPos.SprintControl(l06.FIRST, 0))
				App.Out.Debug(pPos.SprintControl(l06.SECOND, 0))
				ok = true
			} else if length == 2 && tokens[1] == "test" {
				// 利きのテスト
				// 現局面の利きを覚え、ムーブ、アンドゥを行って
				// 元の利きに戻るか確認
				is_error, message := TestControl(pPos)
				if is_error {
					App.Out.Debug("ControlTest: error=%s\n", message)
					App.Out.Debug(pPos.SprintControl(l06.FIRST, CONTROL_LAYER_TEST_ERROR))
					App.Out.Debug(pPos.SprintControl(l06.SECOND, CONTROL_LAYER_TEST_ERROR))
				}
				ok = true
			} else if length == 5 && tokens[1] == "diff" {
				// control diff 11 0 12
				// 利きボード 11番 から 0番を引いた結果を 12番へ入れる。
				layer1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
				layer2, err := strconv.Atoi(tokens[3])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
				layer3, err := strconv.Atoi(tokens[4])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				pPos.DiffControl(layer1, layer2, layer3)
				ok = true
			} else if length == 3 && tokens[1] == "recalc" {
				// control recalc 12
				// 利きの再計算
				layer1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
				pPos.RecalculateControl(layer1)
				ok = true
			} else if length == 3 && tokens[1] == "layer" {
				// 利きテーブルの表示（＾～＾）
				layer, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				} else if 0 <= layer && layer < CONTROL_LAYER_ALL_SIZE {
					App.Out.Debug(pPos.SprintControl(l06.FIRST, layer))
					App.Out.Debug(pPos.SprintControl(l06.SECOND, layer))
					ok = true
				}
			} else if length == 3 && tokens[1] == "sumabs" {
				layer1, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
				// 利きのテスト
				// 現局面の利きを覚え、ムーブ、アンドゥを行って
				// 元の利きに戻るか確認
				sumList := SumAbsControl(pPos, layer1)
				App.Out.Debug("ControlTest: SumAbs=%d,%d\n", sumList[0], sumList[1])
				ok = true
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("control\n")
				App.Out.Debug("control layer {number}\n")
				App.Out.Debug("control recalc {number}\n")
				App.Out.Debug("control diff {layer_number} {layer_number} {layer_number}\n")
				App.Out.Debug("control sumabs {number}\n")
			}
		case "location":
			// あの駒、どこにいんの（＾～＾）？
			App.Out.Debug(l08.SprintLocation(pPos))
		case "sfen":
			// SFEN文字列返せよ（＾～＾）
			App.Out.Debug(pPos.SprintSfen())
		case "record":
			// 棋譜表示。取った駒を表示するためのもの（＾～＾）
			App.Out.Debug(pPos.SprintRecord())
		case "movelist":
			moveList(pPos)
		case "dump":
			// 変数を全部出力してくれだぜ（＾～＾）
			App.Out.Debug("Position.Dump()\n")
			App.Out.Debug("---------------\n%s", pPos.Dump())
		case "playout":
			// とにかく１００手進めるぜ（＾～＾）
			App.Out.Debug("Playout start\n")

			for i := 0; i < 100; i += 1 {
				App.Out.Debug(Sprint(pPos))
				// あの駒、どこにいんの（＾～＾）？
				// App.Out.Debug(SprintLocation(pPos))

				// moveList(pPos)
				bestmove := Search(pPos)
				App.Out.Print("bestmove %s\n", bestmove.ToCodeOfM())

				if bestmove == Move(l04.SQ_EMPTY) {
					// 投了
					break
				}

				pPos.DoMove(bestmove)
			}

			App.Out.Debug("Playout finished\n")
		case "shuffle":
			ShuffleBoard(pPos)
		case "count":
			ShowAllPiecesCount(pPos)
		}

		App.LogNotEcho.FlushAllLogs()
	}

	App.LogNotEcho.Trace("Finished\n")
	App.LogNotEcho.FlushAllLogs()
}

// moveList - 指し手リスト出力
func moveList(pPos *Position) {
	App.Out.Debug("MoveList\n")
	App.Out.Debug("--------\n")
	move_list := GenMoveList(pPos)
	for i, move := range move_list {
		App.Out.Debug("(%d) %s\n", i, move.ToCodeOfM())
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
