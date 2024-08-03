package model

import "time"

type UserIn struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password []byte `db:"password"`
}

type UserInUpdate struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  string `db:"role"`
}

type User struct {
	ID        int64
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
