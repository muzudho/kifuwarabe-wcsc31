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
		// 以下、きふわらべ独自拡張コマンド
		case "pos":
			// 局面表示しないと、データが合ってんのか分からないからな（＾～＾）
			G.Chat.Debug(Sprint(pPos))
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
				G.Chat.Debug(pPos.SprintControl(FIRST, 0))
				G.Chat.Debug(pPos.SprintControl(SECOND, 0))
				ok = true
			} else if length == 3 && tokens[1] == "diff" {
				// 利きの差分の表示（＾～＾）
				layer, err := strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Printf("Error: %s", err)
				} else if 0 <= layer && layer < 5 {
					G.Chat.Debug(pPos.SprintControl(FIRST, layer+1))
					G.Chat.Debug(pPos.SprintControl(SECOND, layer+1))
					ok = true
				}
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("control\n")
				G.Chat.Debug("control diff {0-4}\n")
			}
		case "location":
			// あの駒、どこにいんの（＾～＾）？
			G.Chat.Debug(pPos.SprintLocation())
		case "sfen":
			// SFEN文字列返せよ（＾～＾）
			G.Chat.Debug(pPos.SprintSfen())
		case "record":
			// 棋譜表示。取った駒を表示するためのもの（＾～＾）
			G.Chat.Debug(pPos.SprintRecord())
		case "movelist":
			G.Chat.Debug("MoveList\n")
			G.Chat.Debug("--------\n")
			move_list := GenMoveList(pPos)
			for i, move := range move_list {
				G.Chat.Debug("(%d) %s\n", i, move.ToCode())
			}
			G.Chat.Debug("* Except for those to be removed during the search\n")
		}

		G.Log.FlushAllLogs()
	}

	G.Log.Trace("Finished\n")
	G.Log.FlushAllLogs()
}
