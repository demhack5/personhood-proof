package main

import (
	//"context"
	"context"
	"net/http"
	"time"

	personhood_proof "internal/service/personhood-proof"

	"github.com/go-chi/chi"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	publicRouter := chi.NewRouter()

	publicRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	http.ListenAndServe(":3333", publicRouter)

	interval := time.Minute

	personhoodProofService := personhood_proof.NewPersonhoodProofService()
	defer personhoodProofService.Start(ctx, interval)
}
