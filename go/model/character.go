package model

type Character struct {
	ID    uint   `db:"ID, primarykey, autoincrement" json:"id"`
	Name  string `db:"Name, notnull" json:"name"`
}
