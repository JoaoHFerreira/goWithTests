package database

import (
	"todo/models"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBInstance() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=golang_todo_app port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func Migrate() {
	log.Println("Migration started")
	db := GetDBInstance()

	// ceate a slice of map, being key as string table name and value as model
	// models := []map[string]interface{}{
	models := []map[string]interface{}{
		{"name": "Todo", "model": models.Todo{}},
	}

	for i, m := range models {
		log.Printf("Migration started for %s\n", m["name"].(string))
		log.Printf("Migrating %d of %d\n", i+1, len(models))
		db.AutoMigrate(m["model"])
	}

	log.Println("Migration ended")

}
