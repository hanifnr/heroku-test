package models

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	var (
		dbUser         = os.Getenv("DB_USER")              // e.g. 'my-db-user'
		dbPwd          = os.Getenv("DB_PASS")              // e.g. 'my-db-password'
		dbName         = os.Getenv("DB_NAME")              // e.g. 'my-database'
		unixSocketPath = os.Getenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
	)
	dbURI := fmt.Sprintf("%s:%s@unix(/%s)/%s?parseTime=true",
		dbUser, dbPwd, unixSocketPath, dbName)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error open connection: %s", err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
