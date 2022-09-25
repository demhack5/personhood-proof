package personhood_proof

import (
	"context"
	"fmt"
	"personhood-proof/internal/client/cdn"
	"personhood-proof/internal/client/telegram"
	"personhood-proof/internal/repository/user"
	"time"

	"github.com/jmoiron/sqlx"
)

type PersonhoodProofService struct {
	db       *sqlx.DB
	cdn      cdn.Client
	tg       telegram.Client
	userRepo *user.UserRepository
}

func NewPersonhoodProofService(db *sqlx.DB, cdn cdn.Client, tg telegram.Client, userRepo *user.UserRepository) *PersonhoodProofService {
	return &PersonhoodProofService{
		db:       db,
		cdn:      cdn,
		tg:       tg,
		userRepo: userRepo,
	}
}

func (pp *PersonhoodProofService) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := pp.Analyze(ctx); err != nil {
				fmt.Println("cannot Analyze ", err)
			}
		}
	}
}

func (pp *PersonhoodProofService) Analyze(ctx context.Context) error {
	_ = ctx
	fmt.Println("Analyzing")
	fmt.Println(pp.tg.GetUpdates())
	fmt.Println("End of analyzing")
	return nil
}
