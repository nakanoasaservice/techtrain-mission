package model

type User struct {
	ID    uint   `db:"ID, primarykey, autoincrement" json:"id"`
	Name  string `db:"Name" json:"name"`
	Token string `db:"Token" json:"token"`
}
