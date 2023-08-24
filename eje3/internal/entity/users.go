package entity

type User struct {
	Id       int64  `db:"id"`
	Name     string `db:"name`
	Email    string `db:"email"`
	Password string `db:"password"`
}
