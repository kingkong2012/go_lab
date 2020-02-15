package gorm_lab

import (
"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/sqlite"
"testing"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func TestCRUD(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "P001", Price: 1000})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?",  "P001")

	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}

