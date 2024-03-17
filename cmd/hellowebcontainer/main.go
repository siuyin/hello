package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/siuyin/dflt"
)

var db *sql.DB

func main() {
	fmt.Println("Hello, World!")
	db = openDB()
	defer db.Close()

	ping(db)

	runWebServer()
}

// runWebServer is a function that uses net/http to run a webserver on port 8080
func runWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})

	http.HandleFunc("/rasp", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		qryStr := "SELECT name FROM names order by name desc"
		fmt.Fprintf(w,"query: %s\n", qryStr)
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
			fmt.Fprintf(w,"%s\n", name)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}

func openDB() *sql.DB {
	connStr := dflt.EnvString("POSTGRES_CONN", "postgresql://localhost/mydb?user=siuyin&password=secret")
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
