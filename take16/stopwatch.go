package take16

import "time"

type Stopwatch struct {
	startTime time.Time
}

/// NewStopwatch - ストップウォッチ生成
func NewStopwatch() *Stopwatch {
	var pStopwatch = new(Stopwatch)
	return pStopwatch
}

/// StartStopwatch - 計測開始
func (pStopwatch *Stopwatch) StartStopwatch() {
	pStopwatch.startTime = time.Now()
}

/// Elapsed - 経過時間取得
func (pStopwatch *Stopwatch) Elapsed() time.Duration {
	// return time.Now().Sub(pStopwatch.startTime)
	return time.Since(pStopwatch.startTime)
}

/// ElapsedSeconds - 経過時間取得（秒）
func (pStopwatch *Stopwatch) ElapsedSeconds() float64 {
	return pStopwatch.Elapsed().Seconds()
}
