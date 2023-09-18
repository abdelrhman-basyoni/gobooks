package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abdelrhman-basyoni/gobooks/utils"
	_ "github.com/lib/pq"
)

func connectSql() *sql.DB {
	db, err := sql.Open("postgres", utils.ReadEnv("SQL_URI"))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// Attempt to connect to the database
	err = db.Ping()
	if err != nil {

		fmt.Println(err)
	}

	fmt.Println("Connected to  Postgres")
	return db
}

var SqlDb = connectSql()
