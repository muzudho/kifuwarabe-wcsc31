package take7

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	engineConfPath := filepath.Join(*workdir, "input/take1/engine.conf.toml")

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
	// TODO ディレクトリが存在しなければ、強制終了します。
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

	var pPos = NewPosition()

MainLoop:
	for scanner.Scan() {
		command := scanner.Text()
		G.Log.Trace("command=%s\n", command)

		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "usi":
			G.Chat.Print("id name %s\n", config.Profile.Name)
			G.Chat.Print("id author %s\n", config.Profile.Author)
			G.Chat.Print("usiok\n")
		case "isready":
			G.Chat.Print("readyok\n")
		case "usinewgame":
		case "position":
			// position うわっ、大変だ（＾～＾）
			pPos.ReadPosition(command)
		case "go":
			bestmove := Search(pPos)
			G.Chat.Print("bestmove %s\n", bestmove.ToCode())
		case "quit":
			break MainLoop
		case "pos":
			// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
			G.Chat.Debug(pPos.Sprint())
		case "do":
			// １手指すぜ（＾～＾）
			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			i := 3
			var move, err = ParseMove(command, &i, pPos.Phase)
			if err != nil {
				fmt.Println(pPos.Sprint())
				panic(err)
			}

			pPos.DoMove(move)
		case "undo":
			// 棋譜を頼りに１手戻すぜ（＾～＾）
			pPos.UndoMove()
		case "control":
			// 利きの表示（＾～＾）
			G.Chat.Debug(pPos.SprintControl(FIRST))
			G.Chat.Debug(pPos.SprintControl(SECOND))
		}

		G.Log.FlushAllLogs()
	}

	G.Log.Trace("Finished\n")
	G.Log.FlushAllLogs()
}
