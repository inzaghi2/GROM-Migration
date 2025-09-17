package main

import (
	"GORM-Migration/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func main() {
	// koneksi ke database
	var err error
	DB, err = gorm.Open("postgres", "postgres://postgres:Zaghi08@@127.0.0.1/postgres?sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()
	// koneksi berhasil
	fmt.Println("Successfully connected to database")
	//memanggil function migrate
	Migrate()
}

// migrasi database
func Migrate() {
	DB.AutoMigrate(&models.User{})

}
