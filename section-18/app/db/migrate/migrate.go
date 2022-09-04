package main

import (
	"myapp/app/db"
	"myapp/app/models"

	"gorm.io/gorm"
)

func migrate(gormDb *gorm.DB) {
	gormDb.AutoMigrate(&models.Todo{})
}

func main() {
	gormDb := db.Init()
	defer db.CloseDB(gormDb)

	migrate(gormDb)
}
