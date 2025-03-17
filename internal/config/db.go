package config

import (
	"fmt"
	"sync"

	"stockhive-server/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

//jenkins testtttttt banget ke 9
func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=postgres password=root dbname=stockhive port=5432 sslmode=disable TimeZone=Asia/Jakarta"
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Gagal koneksi ke database: %v", err))
		}
		fmt.Println("Berhasil koneksi ke database")

		//migration put here
		DB.AutoMigrate(&models.User{}, &models.Location{}, &models.Item{})
		fmt.Println("Database Migrated")
	})

	return DB
}