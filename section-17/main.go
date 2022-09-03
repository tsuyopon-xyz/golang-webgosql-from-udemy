package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj.String())

	uuidObj2, _ := uuid.NewUUID()
	fmt.Println(" ", uuidObj2.String())
}

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var Db *sql.DB

// func main() {
// 	Db, _ := sql.Open("sqlite3", "./example.db")

// 	defer Db.Close()

// 	cmd := `CREATE TABLE IF NOT EXISTS people(
// 		name STRING,
// 		age INT
// 	)`

// 	_, err := Db.Exec(cmd)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	Insert(Db)
// }

// func Insert(db *sql.DB) {
// 	cmd := "INSERT INTO people (name, age) VALUES (?, ?)"
// 	_, err := db.Exec(cmd, "Taro", 20)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
