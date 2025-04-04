package config

import (
	"fmt"
	"os"
	"stockhive-server/internal/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := os.Getenv("DB_CONFIG")
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Gagal koneksi ke database: %v", err))
		}
		fmt.Println("Berhasil koneksi ke database")

		//migration put here
		DB.AutoMigrate(&models.User{})
		DB.AutoMigrate(&models.Location{})
		DB.AutoMigrate(&models.Item{})
		DB.AutoMigrate(&models.Vendor{})
		DB.AutoMigrate(&models.Category{})
		DB.AutoMigrate(&models.AuditLog{})

		fmt.Println("Database Migrated")
	})

	return DB
}