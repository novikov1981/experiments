package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"

	//"go/types"
	"log"
)

//type Synthesis struct {
//these are tags, provide additional information for the field for some library
//Uuid string `db:"uuid"`
//Name string `db:"name"`
//Content string `db:"content"`
//Size int `db:"size"`
//CreatedAt string `db:"created_at"`
//}
type Oligs struct {
	//these are tags, provide additional information for the field for some library
	Uuid     string `db:"uuid"`
	Name     string `db:"name"`
	Sequence string `db:"content"`
}

func main() {
	database, err := sqlx.Open("sqlite3", "./synthesis1.db")
	if err != nil {
		log.Fatal(err)
	}

	//statement1, err := database.Prepare(`
	//CREATE TABLE IF NOT EXISTS synthesis (
	//	uuid TEXT PRIMARY KEY,
	//	name TEXT,
	//	content TEXT,
	//	size INTEGER,
	//	created_at TEXT
	//);`)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	_, err = statement1.Exec()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//		err = insertSnt1(database, Synthesis{"","synthesis cool!\n","96 rows",0,"date 555555"})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	ss := []Synthesis{}
	//	err = database.Select(&ss, "SELECT uuid, name, content,size, created_at FROM synthesis;")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	for _, s := range ss {
	//		fmt.Printf("%+v\n", s)
	//	}
	statement, err := database.Prepare(`
CREATE TABLE IF NOT EXISTS oligs (
	uuid TEXT PRIMARY KEY,
	name TEXT,
	sequence TEXT
);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err = insertSnt2(database, Oligs{"", "Olig", "ACGT"})
	if err != nil {
		log.Fatal(err)
	}
	oo := []Oligs{}
	err = database.Select(&oo, "SELECT uuid, name, sequence;")
	if err != nil {
		log.Fatal(err)
	}
	for _, o := range oo {
		fmt.Printf("%+v\n", o)
	}

}

//func insertSnt1(database *sqlx.DB, snt Synthesis) error {
//	statement1, err := database.Prepare(`
//	INSERT INTO synthesis (uuid, name, content,size, created_at) VALUES (?, ?, ?, ?, ?);
//	`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	u := genuuid()
//	u := uuid.NewV4()
//if err != nil {
//	log.Fatal(err)
//}
//_, err = statement1.Exec(fmt.Sprintf("%s", u), snt.Name, snt.Content, snt.Size, snt.CreatedAt)
//if err != nil {
//	log.Fatal(err)
//}
//return err
//}
func insertSnt2(database *sqlx.DB, snt Oligs) error {
	statement, err := database.Prepare(`
	INSERT INTO oligs (uuid, name, sequence) VALUES (?, ?, ?);
	`)
	if err != nil {
		log.Fatal(err)
	}
	u := genuuid()
	//	u := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(fmt.Sprintf("%s", u), snt.Name, snt.Sequence)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func genuuid() uuid.UUID {
	u := uuid.NewV4()
	//if err != nil {
	//	log.Fatal(err)
	//}
	return u
}
