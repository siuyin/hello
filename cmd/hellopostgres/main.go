package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/siuyin/dflt"
)

func openDB() *sql.DB {
	connStr := dflt.EnvString("POSTGRES_CONN_STR", "postgresql://localhost/mydb?user=siuyin&password=secret")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ping(ctx context.Context,db *sql.DB) {

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ping OK")
}

func query(ctx context.Context, db *sql.DB, qryStr string) {
	rows, err := db.QueryContext(ctx,qryStr)
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

func main() {
	db := openDB()
	defer db.Close()
	fmt.Println("database opened")

	ctx := context.Background()
	ping(ctx,db)

	query(ctx, db, "SELECT name FROM names order by name asc")
}
