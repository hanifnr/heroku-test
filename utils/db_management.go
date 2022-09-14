package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"
)

func GetDB() *sql.DB {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}
	var (
		dbUser         = mustGetenv("DB_USER")              // e.g. 'my-db-user'
		dbPwd          = mustGetenv("DB_PASS")              // e.g. 'my-db-password'
		dbName         = mustGetenv("DB_NAME")              // e.g. 'my-database'
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
	)
	fmt.Printf("pg info %s %s %s %s", dbUser, dbPwd, dbName, unixSocketPath)
	dbURI := fmt.Sprintf("%s:%s@unix(/%s)/%s?parseTime=true",
		dbUser, dbPwd, unixSocketPath, dbName)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		fmt.Printf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxOpenConns(100)
	dbPool.SetConnMaxLifetime(time.Hour)
	dbPool.SetConnMaxIdleTime(10 * time.Minute)
	// [END_EXCLUDE]

	return dbPool
}
