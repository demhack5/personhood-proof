package main

import (
	//"context"
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	personhood_proof "personhood-proof/internal/service/personhood-proof"

	"github.com/go-chi/chi"
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

	publicRouter := chi.NewRouter()

	publicRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	http.ListenAndServe(":3333", publicRouter)

	interval := time.Minute

	personhoodProofService := personhood_proof.NewPersonhoodProofService(db)
	defer personhoodProofService.Start(ctx, interval)
}
