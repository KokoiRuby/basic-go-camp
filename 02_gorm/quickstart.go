package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Product struct maps to RDB
type Product struct {
	gorm.Model // extra fields
	Code       string
	Price      uint
}

func main() {
	// create conn to db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/my_db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = db.Debug()

	// migrate schema
	// Create table
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	// Create = Insert
	db.Create(&Product{Code: "D42", Price: 100})

	// Read = Query
	var product Product
	db.First(&product, 1)                 // by primary key
	db.First(&product, "code = ?", "D42") // by cond

	// Update
	db.Model(&product).Update("Price", 200)
	// Update multiple cols, Price & Code only
	db.Model(&product).Updates(&Product{Price: 200, Code: "F42"})
	// Update by map
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete
	db.Delete(&product, 1)
}
