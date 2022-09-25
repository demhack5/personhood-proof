package main 

import (
	//"context"
	"database/sql"
	"log"
	//"os"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func main() {
	//ctx := context.Background()

	db, err := sql.Open("postgres", "postgres://alsoalgo:@127.0.0.1:5432/personhood-proof?sslmode=disable&binary_parameters=yes")
	if err != nil {
		log.Fatalf("err during open: %v", err)
	}
	defer db.Close()

	if err = goose.Up(db, "./db/migrations", goose.WithAllowMissing()); err != nil {
		log.Fatalf("err during up: %v", err)
	}
}