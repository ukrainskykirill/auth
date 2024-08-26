package model

import "time"

type UserIn struct {
	Name            string
	Email           string
	Role            string
	Password        string
	PasswordConfirm string
}

type UserInUpdate struct {
	ID    int64
	Name  string
	Email string
	Role  string
}

type User struct {
	ID        int64
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
