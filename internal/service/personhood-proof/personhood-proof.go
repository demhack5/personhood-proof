package personhood_proof

import (
	"personhood-proof/internal/client/cdn"

	"github.com/jmoiron/sqlx"
)

type PersonhoodProofService struct {
	db  *sqlx.DB
	cdn cdn.Client
}

func NewPersonhoodProofService(db *sqlx.DB, cdn cdn.Client) *PersonhoodProofService {
	return &PersonhoodProofService{
		db:  db,
		cdn: cdn,
	}
}
