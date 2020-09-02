package model

// TODO: gormではforeign key制約がかけれない
type UserCharacter struct {
	ID          uint `db:"ID, primarykey, autoincrement" json:"id"`
	UserID      uint `db:"UserID, notnull" json:"userId"`
	CharacterID uint `db:"CharacterID, notnull" json:"characterId"`
}
