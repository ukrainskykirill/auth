package access

import (
	"context"

	"github.com/ukrainskykirill/platform_common/pkg/db"

	"github.com/ukrainskykirill/auth/internal/repository"
)

type accessRepo struct {
	db db.Client
}

func NewAccessRepository(db db.Client) repository.AccessRepository {
	return &accessRepo{db: db}
}

func (r *accessRepo) Check(ctx context.Context) error {
	return nil
}
