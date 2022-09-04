package main

import (
	"fmt"
	"myapp/app/db"
	"myapp/app/models"
	"strconv"

	"gorm.io/gorm"
)

func seedTodos(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		todo := models.Todo{
			Title: "title " + strconv.Itoa(i+1),
			Body:  "body " + strconv.Itoa(i+1),
		}

		if err := db.Create(&todo).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func main() {
	dbCon := db.Init()
	defer db.CloseDB(dbCon)

	if err := seedTodos(dbCon); err != nil {
		fmt.Printf("%+v", err)
	}
}
