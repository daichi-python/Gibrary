package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBController struct {
	db *sql.DB
}

func NewDBController(db *sql.DB) *DBController {
	return &DBController{
		db: db,
	}
}

func Setup() *DBController {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")

	dbEngine, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/gibrary?charset=utf8mb4", user, pass, host))
	if err != nil {
		log.Println("Error: Failed to connect MySQL server.")
		return nil
	}

	return NewDBController(dbEngine)
}
