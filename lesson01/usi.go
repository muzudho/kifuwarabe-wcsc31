package lesson01

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	l "github.com/muzudho/go-logger"
)

const (
	// Author - 囲碁思考エンジンの作者名だぜ☆（＾～＾）
	Author = "Satoshi Takahashi"
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
	My = *new(Lesson01My)

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
	My.LogNotEcho = *l.NewLogger(
		tracePath,
		debugPath,
		infoPath,
		noticePath,
		warnPath,
		errorPath,
		fatalPath,
		printPath)

	// 既存のログ・ファイルを削除
	My.LogNotEcho.RemoveAllOldLogs()

	// ログ・ファイルの開閉
	err = My.LogNotEcho.OpenAllLogs()
	if err != nil {
		// ログ・ファイルを開くのに失敗したのだから、ログ・ファイルへは書き込めません
		panic(err)
	}
	defer My.LogNotEcho.CloseAllLogs()

	// チャッターの作成。 標準出力とロガーを一緒にしただけです。
	My.Out = *l.NewChatter(My.LogNotEcho)
	My.Log = *l.NewStderrChatter(My.LogNotEcho)

	My.Log.Trace("Start Take1\n")
	My.Log.Trace("engineConfPath=%s\n", engineConfPath)

	// 設定ファイル読込。ファイルが存在しなければ強制終了してしまうので注意！
	config, err := LoadEngineConf(engineConfPath)
	if err != nil {
		panic(My.Log.Fatal(fmt.Sprintf("engineConfPath=[%s] err=[%s]", engineConfPath, err)))
	}

	// 何か標準入力しろだぜ☆（＾～＾）
	scanner := bufio.NewScanner(os.Stdin)

MainLoop:
	for scanner.Scan() {
		My.LogNotEcho.FlushAllLogs()

		command := scanner.Text()
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "usi":
			My.Out.Print("id name %s\n", config.Profile.Name)
			My.Out.Print("id author %s\n", config.Profile.Author)
			My.Out.Print("usiok\n")
		case "isready":
			My.Out.Print("readyok\n")
		case "usinewgame":
		case "position":
		case "go":
			My.Out.Print("bestmove resign\n")
		case "quit":
			break MainLoop
		}
	}

	My.Log.Trace("Finished\n")
	My.LogNotEcho.FlushAllLogs()
}
