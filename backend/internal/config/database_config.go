package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SqlDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	addr := os.Getenv("DB_ADDR")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	var dsn string

	if pass == "" {
		dsn = fmt.Sprintf("postgres://%s:@%s/%s?sslmode=disable", user, addr, name)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, addr, name)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
