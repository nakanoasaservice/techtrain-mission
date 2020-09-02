package model

type GachaCharactersOdds struct {
	GachaID uint `db:"GachaID, primarykey" json:"gachaId"`
	CharacterID uint `db:"CharacterID, primarykey" json:"characterId"`
	Odds float64 `db:"Odds, notnull" json:"odds"`
}
