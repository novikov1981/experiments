package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Synthesis struct {
	//these are tags, provide additional information for the field for some library
	Uuid      string `db:"uuid"`
	Name      string `db:"name"`
	Content   string `db:"content"`
	CreatedAt string `db:"created_at"`
}

//type Oligs struct {
//	//these are tags, provide additional information for the field for some library
//	Uuid     string `db:"uuid"`
//	Name     string `db:"name"`
//	Sequence string `db:"sequence"`
//}
//
//func main() {
//
//		database, err := sqlx.Open("sqlite3", "./synthesis1.db")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		statement, err := database.Prepare(`
//	CREATE TABLE IF NOT EXISTS oligs (
//		uuid TEXT PRIMARY KEY,
//		name TEXT,
//		sequence TEXT
//			);`)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		_, err = statement.Exec() //здесь два возвращаемых аргумента _, err узнать по подробнее
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		err = insertSnt(database, Oligs{"","synthesis cool!","96 rows"})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//
//		ss := []Oligs{}
//		err = database.Select(&ss, "SELECT uuid, name, sequence FROM oligs;")
//		if err != nil {
//			log.Fatal(err)
//		}
//		for _, s := range ss {
//			fmt.Printf("%+v\n", s)// +v идет форматирование в соответствии со всеми имеющимися типами данных?
//		}
//	}
//
//	func insertSnt(database *sqlx.DB, snt Oligs) error {//не понятно использование типа error для ввода данных в таблицу?
//		statement, err := database.Prepare(`
//		INSERT INTO oligs (uuid, name, sequence) VALUES (?, ?, ?);
//		`)
//		if err != nil {
//			log.Fatal(err)
//		}
//		u := genuuid() //я генирирую uuid прямо в функции это не удобно, потому что у нас две таблицы, а uuid должен быть?
//
//		//u := uuid.NewV4()
//		if err != nil {
//			log.Fatal(err)
//		}
//		_, err = statement.Exec(fmt.Sprintf("%s", u), snt.Name, snt.Sequence)
//		if err != nil {
//			log.Fatal(err)
//		}
//		return err
//	}

func main() {

	database, err := sqlx.Open("sqlite3", "./synthesis.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := database.Prepare(`
	CREATE TABLE IF NOT EXISTS synthesis (
		uuid TEXT PRIMARY KEY,
		name TEXT,
		content TEXT,
		created_at TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec() //здесь два возвращаемых аргумента _, err узнать по подробнее
	if err != nil {
		log.Fatal(err)
	}

	err = insertSnt(database, Synthesis{"", "synthesis cool!", "96 rows", "date"})
	if err != nil {
		log.Fatal(err)
	}

	ss := []Synthesis{}
	err = database.Select(&ss, "SELECT uuid, name, content, created_at FROM synthesis;")
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range ss {
		fmt.Printf("%+v\n", s) // +v идет форматирование в соответствии со всеми имеющимися типами данных?
	}
	err = database.Select(&ss, "SELECT FROM uuid, name, content, created_at WHERE content LIKE '%96%'")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
}

func insertSnt(database *sqlx.DB, snt Synthesis) error { //не понятно использование типа error для ввода данных в таблицу?
	statement, err := database.Prepare(`
		INSERT INTO synthesis (uuid, name, content, created_at) VALUES (?, ?, ?, ?);
		`)
	if err != nil {
		log.Fatal(err)
	}
	u := genuuid() //я генирирую uuid прямо в функции это не удобно, потому что у нас две таблицы, а uuid должен быть?

	//u := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(fmt.Sprintf("%s", u), snt.Name, snt.Content, snt.CreatedAt)
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
