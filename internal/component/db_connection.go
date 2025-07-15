package component

import (
	"fmt"
	"log"

	// "topupservice/domain"
	"topupservice/domain"
	"topupservice/internal/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.Database.Host,     // localhost
		cnf.Database.Port,     // 5432
		cnf.Database.User,     // postgres
		cnf.Database.Password, // xxxxx
		cnf.Database.Name,     // nama_database
	)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Mengecek apakah koneksi dapat di-"ping"
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error open connection %v", err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("error open connection %v", err.Error())
	}
	db.Exec("CREATE SCHEMA IF NOT EXISTS topup_schema")
	// db.Exec("CREATE TYPE level AS ENUM ('Beginner', 'Intermediate', 'Advanced');")
	// Melakukan migrasi ke database (membuat tabel user jika belum ada)
	err = db.Debug().AutoMigrate(&domain.User{}, &domain.Game{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	return db
}
