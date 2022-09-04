package main

import (
	"fmt"

	"myapp/app/db"
)

func main() {
	db := db.Init()
	fmt.Printf("%+v @@@@@@@@@\n", db)

	// log.Println("test")
}
