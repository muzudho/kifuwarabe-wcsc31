package take8

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
	l06 "github.com/muzudho/kifuwarabe-wcsc31/take6"
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
			var move, err = ParseMove(command, &i, pPos.Phase)
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
			fmt.Printf("length=%d", length)
			ok := false
			if length == 1 {
				// 利きの表示（＾～＾）
				App.Out.Debug(pPos.SprintControl(l06.FIRST, 0))
				App.Out.Debug(pPos.SprintControl(l06.SECOND, 0))
				ok = true
			} else if length == 3 && tokens[1] == "diff" {
				// 利きの差分の表示（＾～＾）
				layer, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				} else if 0 <= layer && layer < 5 {
					App.Out.Debug(pPos.SprintControl(l06.FIRST, layer+1))
					App.Out.Debug(pPos.SprintControl(l06.SECOND, layer+1))
					ok = true
				}
			}

			if !ok {
				App.Out.Debug("Format\n")
				App.Out.Debug("------\n")
				App.Out.Debug("control\n")
				App.Out.Debug("control diff {0-4}\n")
			}
		case "location":
			// あの駒、どこにいんの（＾～＾）？
			App.Out.Debug(SprintLocation(pPos))
		case "sfen":
			// SFEN文字列返せよ（＾～＾）
			App.Out.Debug(pPos.SprintSfen())
		case "record":
			// 棋譜表示。取った駒を表示するためのもの（＾～＾）
			App.Out.Debug(pPos.SprintRecord())
		case "movelist":
			App.Out.Debug("MoveList\n")
			App.Out.Debug("--------\n")
			move_list := GenMoveList(pPos)
			for i, move := range move_list {
				App.Out.Debug("(%d) %s\n", i, move.ToCodeOfM())
			}
			App.Out.Debug("* Except for those to be removed during the search\n")
		}

		App.LogNotEcho.FlushAllLogs()
	}

	App.LogNotEcho.Trace("Finished\n")
	App.LogNotEcho.FlushAllLogs()
}
