package dbmodels

type UserMessage struct {
	Message string `db:"message"`
	IsNew   bool   `db:"is_new"`
}

type User struct {
	ID       int64          `db:"id"`
	UserName string         `db:"username"`
	Messages []*UserMessage `db:"messages"`
}
