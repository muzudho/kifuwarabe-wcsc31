package lesson01

import (
	l "github.com/muzudho/go-logger"
)

// OutNode - グローバル変数。
type OutNode struct {
	// Log - ロガー。
	Log l.Logger
	// Chat - チャッター。 標準出力とロガーを一緒にしただけです。
	Chat l.Chatter
	// StderrChat - チャッター。 標準エラー出力とロガーを一緒にしただけです。
	StderrChat l.StderrChatter
}

// Out - グローバル変数。思い切った名前。
var Out OutNode
