package take9

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

func (pPos *Position) GetKingLocation(index int) l04.Square {
	return pPos.KingLocations[index]
}

func (pPos *Position) GetRookLocation(index int) l04.Square {
	return pPos.RookLocations[index]
}

func (pPos *Position) GetBishopLocation(index int) l04.Square {
	return pPos.BishopLocations[index]
}

func (pPos *Position) GetLanceLocation(index int) l04.Square {
	return pPos.LanceLocations[index]
}
