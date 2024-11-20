package converter

import (
	"github.com/ukrainskykirill/auth/internal/model"
	repoModel "github.com/ukrainskykirill/auth/internal/repository/user/model"
)

func ToUserFromRepo(user repoModel.RepoUser) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserAuthInfoFromRepo(info repoModel.RepoUserAuthInfo) *model.UserAuthInfo {
	return &model.UserAuthInfo{
		Password: info.Password,
		Role:     info.Role,
	}
}
