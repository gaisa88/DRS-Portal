package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	"log"
	"time"
)

var db *sql.DB

var server = "172.27.5.214"
var port = 1433
var user = "sa"
var password = "Test1K8rR"
var database = "CigSuite_v3"

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	count, err := ReadDocuments()
	if err != nil {
		log.Fatal("Error reading Documents: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

}

func ReadDocuments() (int, error) {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT Id, message, date FROM dbo.test2;")

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		var date time.Time
		var message string
		var id int

		err := rows.Scan(&id, &message, &date)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, message, date)
		count++
	}

	return count, nil
}
