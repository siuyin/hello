package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/siuyin/dflt"
)

func main() {
	connStr := dflt.EnvString("POSTGRES_CONN_STR", "postgresql://localhost/mydb?user=siuyin&password=secret")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("database opened")

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ping OK")

	rows, err := db.QueryContext(ctx, "SELECT name FROM names order by name asc")
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
