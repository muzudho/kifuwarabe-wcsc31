package take5

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	l "github.com/muzudho/go-logger"
	l01 "github.com/muzudho/kifuwarabe-wcsc31/lesson01"
	l03 "github.com/muzudho/kifuwarabe-wcsc31/lesson03"
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
	// TODO ディレクトリが存在しなければ、強制終了します。
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

	var pos = NewPosition()

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
			// TODO position うわっ、大変だ（＾～＾）
			pos.ReadPosition(command)
		case "go":
			App.Out.Print("bestmove resign\n")
		case "quit":
			break MainLoop
		case "pos":
			// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
			App.Out.Debug(SprintBoard(pos))
		case "do":
			// １手指すぜ（＾～＾）
			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			i := 3
			var move, err = l03.ParseMove(command, &i, pos.Phase)
			if err != nil {
				fmt.Println(SprintBoard(pos))
				panic(err)
			}

			pos.DoMove(move)
		case "undo":
			// 棋譜を頼りに１手戻すぜ（＾～＾）
			pos.UndoMove()
		}

		App.LogNotEcho.FlushAllLogs()
	}

	App.LogNotEcho.Trace("Finished\n")
	App.LogNotEcho.FlushAllLogs()
}
