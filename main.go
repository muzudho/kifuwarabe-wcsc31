package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/muzudho/kifuwarabe-wcsc31/lesson01"
	"github.com/muzudho/kifuwarabe-wcsc31/lesson02"
	"github.com/muzudho/kifuwarabe-wcsc31/lesson03"
	"github.com/muzudho/kifuwarabe-wcsc31/take10"
	"github.com/muzudho/kifuwarabe-wcsc31/take11"
	"github.com/muzudho/kifuwarabe-wcsc31/take12"
	"github.com/muzudho/kifuwarabe-wcsc31/take13"
	"github.com/muzudho/kifuwarabe-wcsc31/take14"
	"github.com/muzudho/kifuwarabe-wcsc31/take15"
	"github.com/muzudho/kifuwarabe-wcsc31/take16"
	"github.com/muzudho/kifuwarabe-wcsc31/take17"
	"github.com/muzudho/kifuwarabe-wcsc31/take4"
	"github.com/muzudho/kifuwarabe-wcsc31/take5"
	"github.com/muzudho/kifuwarabe-wcsc31/take6"
	"github.com/muzudho/kifuwarabe-wcsc31/take7"
	"github.com/muzudho/kifuwarabe-wcsc31/take8"
	"github.com/muzudho/kifuwarabe-wcsc31/take9"
)

// main - 最初に実行されます
func main() {
	flag.Parse()
	lessonVer := flag.Arg(0)

	// fmt.Printf("(11-12) mod 10=%d\n", (11-12)%10)

	// fmt.Println("Hello, world!")

	switch lessonVer {
	case "lesson01":
		// 自殺手きふわらべ（＾▽＾）
		lesson01.MainLoop()
	case "lesson02":
		// ダミーの盤を表示
		lesson02.MainLoop()
	case "lesson03":
		lesson03.MainLoop()
	case "lesson04":
		take4.MainLoop()
	case "lesson05":
		take5.MainLoop()
	case "lesson06":
		// ゲーム向けの軽い乱数のタネ (take6～12)
		rand.Seed(time.Now().UnixNano())
		take6.MainLoop()
	case "lesson07":
		rand.Seed(time.Now().UnixNano())
		take7.MainLoop()
	case "lesson08":
		rand.Seed(time.Now().UnixNano())
		take8.MainLoop()
	case "lesson09":
		rand.Seed(time.Now().UnixNano())
		take9.MainLoop()
	case "lesson10":
		rand.Seed(time.Now().UnixNano())
		take10.MainLoop()
	case "lesson11":
		rand.Seed(time.Now().UnixNano())
		take11.MainLoop()
	case "lesson12":
		rand.Seed(time.Now().UnixNano())
		take12.MainLoop()
	case "lesson13":
		// 大会は take13で行くか（＾～＾）安定版（＾～＾）
		rand.Seed(time.Now().UnixNano())
		take13.MainLoop()
	case "lesson14":
		// take14は未完成 --> suspended
		rand.Seed(time.Now().UnixNano())
		take14.MainLoop()
	case "lesson15":
		// take15 は、 take13 の後継（＾～＾）
		rand.Seed(time.Now().UnixNano())
		take15.MainLoop()
	case "lesson16":
		// 2021年最終版
		rand.Seed(time.Now().UnixNano())
		take16.MainLoop()
	case "lesson17":
	default:
		// 2022年版
		rand.Seed(time.Now().UnixNano())
		take17.MainLoop()
	}
}
