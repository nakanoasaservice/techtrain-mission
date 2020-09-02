package db

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"techtrain-mission/go/helper"
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

	db.AddTable(model.User{})
	db.AddTable(model.Character{})
	db.AddTable(model.UserCharacter{})
	db.AddTable(model.Gacha{})
	db.AddTable(model.GachaCharactersOdds{})

	_ = db.DropTablesIfExists()

	err = db.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	trans, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	if err := InsertSeeds(trans); err != nil {
		if err2 := trans.Rollback(); err2 != nil {
			log.Print(err)
			log.Fatal(err2)
		}
		log.Fatal(err)
	}
	if err := trans.Commit(); err != nil {
		log.Fatal(err)
	}
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

func InsertSeeds(trans *gorp.Transaction) error {
	user := model.User{Name: "ユウキ"}
	err := trans.Insert(&user)
	if err != nil {
		return err
	}
	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	log.Print("first users token is ", token)

	characters := [...] model.Character{
		{Name: "コロ"}, {Name: "ペコ"}, {Name: "キャル"},
	}

	// TODO: charactersを[]interface{}にキャストできない理由が知りたい
	err = trans.Insert(&characters[0], &characters[1], &characters[2])
	if err != nil {
		return err
	}

	ownership := model.UserCharacter{UserID: user.ID, CharacterID: characters[0].ID}
	err = trans.Insert(&ownership)
	if err != nil {
		return err
	}

	gacha := model.Gacha{Name: "通常ガチャ"}
	err = trans.Insert(&gacha)
	if err != nil {
		return err
	}

	err = trans.Insert(
		&model.GachaCharactersOdds{GachaID: gacha.ID, CharacterID: characters[0].ID, Odds: 1.0 / 3},
		&model.GachaCharactersOdds{GachaID: gacha.ID, CharacterID: characters[1].ID, Odds: 1.0 / 3},
		&model.GachaCharactersOdds{GachaID: gacha.ID, CharacterID: characters[2].ID, Odds: 1.0 / 3},
	)
	if err != nil {
		return err
	}

	return nil
}
