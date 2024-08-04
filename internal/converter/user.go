package converter

import (
	"github.com/ukrainskykirill/auth/internal/model"
	guser "github.com/ukrainskykirill/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getRole(role string) guser.UserRole {
	switch role {
	case guser.UserRole_USER.String():
		return guser.UserRole_USER
	case guser.UserRole_ADMIN.String():
		return guser.UserRole_ADMIN
	default:
		return guser.UserRole_UNKNOW
	}
}

func ToUserFromService(user *model.User) *guser.GetResponse {
	return &guser.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      getRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.CreatedAt),
	}
}

func ToUserInFromGUser(user *guser.CreateRequest) *model.UserIn {
	return &model.UserIn{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            user.Role.String(),
	}
}

func ToUserInUpdateFromGUser(user *guser.UpdateRequest) *model.UserInUpdate {
	var name, email, role string

	if user.Name != nil {
		name = user.Name.GetValue()
	}
	if user.Email != nil {
		email = user.Email.GetValue()
	}
	if user.Role != nil {
		role = user.Role.String()
	}

	return &model.UserInUpdate{
		ID:    user.Id,
		Name:  name,
		Email: email,
		Role:  role,
	}
}
