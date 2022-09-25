package main

import (
	//"context"
	"context"
	"database/sql"
	"log"
	"time"

	"personhood-proof/internal/client/cdn"
	personhood_proof "personhood-proof/internal/service/personhood-proof"

	"github.com/jmoiron/sqlx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sqlDb, err := sql.Open("postgres", "postgres://alsoalgo:@127.0.0.1:5432/personhood-proof?sslmode=disable&binary_parameters=yes")
	if err != nil {
		log.Fatalf("err during sqlDb: %v", err)
	}
	defer sqlDb.Close()

	db := sqlx.NewDb(sqlDb, "postgres")

	interval := time.Minute

	cdnMock := cdn.NewClient()

	personhoodProofService := personhood_proof.NewPersonhoodProofService(db, cdnMock)
	defer personhoodProofService.Start(ctx, interval)
}
