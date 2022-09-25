package telegram_worker

import (
	"context"
	"fmt"

	"personhood-proof/internal/client/telegram"
	"time"

	"github.com/jmoiron/sqlx"
)

type Worker struct {
	db             *sqlx.DB
	telegramClient telegram.Client
}

func NewWorker(db *sqlx.DB, telegramClient telegram.Client) *Worker {
	return &Worker{
		db:             db,
		telegramClient: telegramClient,
	}
}

func (w *Worker) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			updates, err := w.GetUpdates(ctx)
			if err != nil {
				fmt.Printf("error during GetUpdates: %v", err)
			}
			err = w.RecordUpdates(ctx, updates)
			if err != nil {
				fmt.Printf("error during RecordUpdates: %v", err)
			}
		}

	}
}

func (w *Worker) GetUpdates(ctx context.Context) (map[int64][]string, error) {
	_ = ctx
	return nil, nil
}

func (w *Worker) RecordUpdates(ctx context.Context, updates map[int64][]string) error {
	
	return nil
}
