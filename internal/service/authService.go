package service

import (
	"context"
	"protobuf-demo/internal/model"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) Login(ctx context.Context, username string, password string) (*model.UserModel, error) {
	if username == "admin" && password == "123456" {
		return &model.UserModel{
			Id:       1,
			Username: username,
			Password: password,
		}, nil
	}

	return &model.UserModel{
		Id:       2,
		Username: username,
		Password: password,
	}, nil
}
