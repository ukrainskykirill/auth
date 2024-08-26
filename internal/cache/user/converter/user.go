package converter

import (
	"time"

	modelCache "github.com/ukrainskykirill/auth/internal/cache/user/model"
	"github.com/ukrainskykirill/auth/internal/model"
)

func ToUserCacheFromModel(user *model.User) *modelCache.User {
	return &modelCache.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		CreatedAtNs: user.CreatedAt.Unix(),
		UpdatedAtNs: user.UpdatedAt.Unix(),
	}
}

func ToUserFromCache(user *modelCache.User) *model.User {
	return &model.User{
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: time.Unix(user.CreatedAtNs, 0),
		UpdatedAt: time.Unix(user.UpdatedAtNs, 0),
	}
}
