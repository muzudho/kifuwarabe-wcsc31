package take15

// Brain - 局面システムと、利き盤システムの２つを持つもの
type Brain struct {
	// 局面システム
	PPosSys *PositionSystem
}

func NewBrain() *Brain {
	var pBrain = new(Brain)
	pBrain.PPosSys = NewPositionSystem()
	return pBrain
}
