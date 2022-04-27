package take10

import l04 "github.com/muzudho/kifuwarabe-wcsc31/take4"

func (pPos *Position) GetKingLocation(index int) l04.Square {
	return pPos.PieceLocations[PCLOC_K1:PCLOC_K2][index]
}

func (pPos *Position) GetRookLocation(index int) l04.Square {
	return pPos.PieceLocations[PCLOC_R1:PCLOC_R2][index]
}

func (pPos *Position) GetBishopLocation(index int) l04.Square {
	return pPos.PieceLocations[PCLOC_B1:PCLOC_B2][index]
}

func (pPos *Position) GetLanceLocation(index int) l04.Square {
	return pPos.PieceLocations[PCLOC_L1:PCLOC_L4][index]
}
