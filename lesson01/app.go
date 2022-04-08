package lesson01

import (
	l "github.com/muzudho/go-logger"
)

// Lesson01App - グローバル変数。
type Lesson01App struct {
	// Out - チャッター。 標準出力とロガーを一緒にしただけです。
	Out l.Chatter
	// Log - チャッター。 標準エラー出力とロガーを一緒にしただけです。
	Log l.StderrChatter
	// LogNotEcho - ロガー。
	LogNotEcho l.Logger
}
