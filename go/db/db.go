package db

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"techtrain-mission/go/model"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

//Init ...
func Init() {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_SERVER"),
		os.Getenv("MYSQL_DATABASE"),
	)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	table := db.AddTable(model.User{})
	table.ColMap("Token").SetUnique(true).SetNotNull(true)

	table = db.AddTable(model.Character{})

	table = db.AddTable(model.UserCharacter{})

	err = db.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	InsertSeeds(db)
}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	if os.Getenv("GIN_MODE") != "release" {
		dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	}

	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}

func InsertSeeds(db *gorp.DbMap) {
	user := model.User{Name: "ユウキ", Token: "ea4fa3dd-ddaf-418f-a936-2cf0f6e43c58"}
	err := db.Insert(&user)
	if err != nil {
		log.Fatal(err)
	}

	characters := [...] model.Character{
		{Name: "コロ"}, {Name: "ペコ"}, {Name: "キャル"},
	}
	for i := range characters {
		err = db.Insert(&characters[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println(characters)

	ownership := model.UserCharacter{UserID: user.ID, CharacterID: characters[0].ID}
	err = db.Insert(&ownership)
	if err != nil {
		log.Fatal(err)
	}
}
