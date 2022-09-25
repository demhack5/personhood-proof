package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const insertUserQuery = ``

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUpdates(ctx context.Context) (map[int64][]string, error) {
	type newMessage struct {
		UserID   int64    `db:"user_id"`
		Messages []string `db:"message"`
	}

	var getNewMessagesQuery = `
		select user_id, message from user_messages where is_new=true group by user_id 
	`

	messages := make([]*newMessage, 0)

	if err := ur.db.SelectContext(ctx, &messages, getNewMessagesQuery); err != nil {
		return nil, fmt.Errorf("cannot select context: %v", err)
	}

	ret := make(map[int64][]string)

	for _, v := range messages {
		ret[v.UserID] = v.Messages
	}
	return ret, nil
}
