package user

import (
	"context"
	"strconv"

	"github.com/ukrainskykirill/auth/internal/cache/user/converter"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (c *Implementation) Create(ctx context.Context, user *model.User) error {
	userCache := converter.ToUserCacheFromModel(user)
	idStr := strconv.FormatInt(userCache.ID, 10)
	err := c.cl.HashSet(ctx, idStr, userCache)
	if err != nil {
		return err
	}
	return nil
}
