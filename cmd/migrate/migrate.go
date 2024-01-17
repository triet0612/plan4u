package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	script_design, err := os.ReadFile("./cmd/migrate/design.sql")
	if err != nil {
		log.Println(err)
		return
	}
	db, _ := sql.Open("sqlite3", "../build/app.db")
	_, err = db.Exec(string(script_design))
	if err != nil {
		log.Fatalln(err)
		return
	}
}
