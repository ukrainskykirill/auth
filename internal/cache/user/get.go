package user

import (
	"context"
	"errors"
	"strconv"

	redigo "github.com/gomodule/redigo/redis"

	"github.com/ukrainskykirill/auth/internal/cache/user/converter"
	cacheModel "github.com/ukrainskykirill/auth/internal/cache/user/model"
	"github.com/ukrainskykirill/auth/internal/model"
)

func (c *Implementation) Get(ctx context.Context, id int64) (*model.User, error) {
	idStr := strconv.FormatInt(id, 10)
	values, err := c.cl.HGetAll(ctx, idStr)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, errors.New("user not found")
	}

	var user cacheModel.User
	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromCache(&user), nil
}
