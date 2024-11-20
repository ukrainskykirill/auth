package model

import "time"

type RepoUserIn struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password []byte `db:"password"`
}

type RepoUserInUpdate struct {
	ID    int64   `db:"id"`
	Name  *string `db:"name"`
	Email *string `db:"email"`
	Role  *string `db:"role"`
}

type RepoUser struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type RepoUserAuthInfo struct {
	Password string `db:"password"`
	Role     string `db:"role"`
}
