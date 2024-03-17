package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/siuyin/dflt"
)

func main() {
	db := openDB()
	defer db.Close()
	fmt.Println("database opened")

	ping(db)

	query(db, "SELECT name FROM names order by name asc")
}

func openDB() *sql.DB {
	connStr := dflt.EnvString("POSTGRES_CONN_STR", "postgresql://localhost/mydb?user=siuyin&password=secret")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ping(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ping OK")
}

func query(db *sql.DB, qryStr string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Printf("query: %s\n", qryStr)
	rows, err := db.QueryContext(ctx, qryStr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
