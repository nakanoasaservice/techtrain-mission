package model

// TODO: gormではforeign key制約がかけれない
type UserCharacter struct {
	ID          uint `db:"ID, primarykey, autoincrement" json:"id"`
	UserID      uint `db:"UserID" json:"userId"`
	CharacterID uint `db:"CharacterID" json:"characterId"`
}
