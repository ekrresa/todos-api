package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	sqlMigrate "github.com/rubenv/sql-migrate"
)

func ConnectToDB() *sqlx.DB {
	dbConnString, exists := os.LookupEnv("CONN_STRING")
	if !exists {
		log.Fatal("CONN_STRING not set")
	}

	var db, err = sqlx.Connect("postgres", dbConnString)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	fmt.Println("Connected to DB!")

	return db
}

func ApplyMigrations(db *sql.DB) {
	var migrations = &sqlMigrate.FileMigrationSource{Dir: "migrations/"}
	var n, err = sqlMigrate.Exec(db, "postgres", migrations, sqlMigrate.Up)

	if err != nil {
		log.Fatal("Error applying migrations:", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
