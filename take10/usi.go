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
			G.Chat.Debug(pPos.Sprint())
		case "do":
			// １手指すぜ（＾～＾）
			// 前の空白を読み飛ばしたところから、指し手文字列の終わりまで読み進めるぜ（＾～＾）
			i := 3
			var move, err = ParseMove(command, &i, pPos.GetPhase())
			if err != nil {
				fmt.Println(pPos.Sprint())
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
				G.Chat.Debug(pPos.SprintControl(FIRST, 0))
				G.Chat.Debug(pPos.SprintControl(SECOND, 0))
				ok = true
			} else if length == 2 && tokens[1] == "test" {
				// 利きのテスト
				// 現局面の利きを覚え、ムーブ、アンドゥを行って
				// 元の利きに戻るか確認
				is_error, message := TestControl(pPos)
				if is_error {
					G.Chat.Debug("ControlTest: error=%s\n", message)
					G.Chat.Debug(pPos.SprintControl(FIRST, CONTROL_LAYER_TEST_ERROR))
					G.Chat.Debug(pPos.SprintControl(SECOND, CONTROL_LAYER_TEST_ERROR))
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
					G.Chat.Debug(pPos.SprintControl(FIRST, layer))
					G.Chat.Debug(pPos.SprintControl(SECOND, layer))
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
				G.Chat.Debug("ControlTest: SumAbs=%d,%d\n", sumList[0], sumList[1])
				ok = true
			}

			if !ok {
				G.Chat.Debug("Format\n")
				G.Chat.Debug("------\n")
				G.Chat.Debug("control\n")
				G.Chat.Debug("control layer {number}\n")
				G.Chat.Debug("control recalc {number}\n")
				G.Chat.Debug("control diff {layer_number} {layer_number} {layer_number}\n")
				G.Chat.Debug("control sumabs {number}\n")
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
			moveList(pPos)
		case "dump":
			// 変数を全部出力してくれだぜ（＾～＾）
			G.Chat.Debug("Position.Dump()\n")
			G.Chat.Debug("---------------\n%s", pPos.Dump())
		case "playout":
			// とにかく１００手進めるぜ（＾～＾）
			G.Chat.Debug("Playout start\n")

			for i := 0; i < 100; i += 1 {
				G.Chat.Debug(pPos.Sprint())
				// あの駒、どこにいんの（＾～＾）？
				// G.Chat.Debug(pPos.SprintLocation())

				// moveList(pPos)
				bestmove := Search(pPos)
				G.Chat.Print("bestmove %s\n", bestmove.ToCode())

				if bestmove == Move(SQUARE_EMPTY) {
					// 投了
					break
				}

				pPos.DoMove(bestmove)
			}

			G.Chat.Debug("Playout finished\n")
		case "shuffle":
			ShuffleBoard(pPos)
		case "count":
			ShowAllPiecesCount(pPos)
		}

		G.Log.FlushAllLogs()
	}

	G.Log.Trace("Finished\n")
	G.Log.FlushAllLogs()
}

// moveList - 指し手リスト出力
func moveList(pPos *Position) {
	G.Chat.Debug("MoveList\n")
	G.Chat.Debug("--------\n")
	move_list := GenMoveList(pPos)
	for i, move := range move_list {
		G.Chat.Debug("(%d) %s\n", i, move.ToCode())
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
