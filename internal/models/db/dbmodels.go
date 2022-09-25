package dbmodels

type User struct {
	ID int64 `db:"id"`
	UserName string `db:"username"`
}