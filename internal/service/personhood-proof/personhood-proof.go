package personhood_proof

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type PersonhoodProofService struct {
	db *sqlx.DB
}

func NewPersonhoodProofService(db *sqlx.DB) *PersonhoodProofService {
	return &PersonhoodProofService{
		db: db,
	}
}

func (p *PersonhoodProofService) Start(ctx context.Context, interval time.Duration) {
	if interval <= 0 {
		interval = 30 * time.Second
	}

	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			if err := p.runEverything(ctx); err != nil {
				fmt.Println("something happened")
			}
		}
	}
}

func (p *PersonhoodProofService) runEverything(ctx context.Context) error {
	return nil
}
