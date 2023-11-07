package ctl

import (
	"context"
	"log"
	"protobuf-demo/internal/service"
	userProto "protobuf-demo/proto/pb/users"
)

type AuthController struct {
	userProto.UnimplementedAuthServiceServer
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func validateRequest(request *userProto.LoginRequest) error {
	return nil
}

func (a AuthController) Login(ctx context.Context, request *userProto.LoginRequest) (*userProto.LoginResponse, error) {

	log.Println(" user login ->", request.Username, request.Password)

	if err := validateRequest(request); err != nil {
		return nil, err
	}

	// business SERVICE 层
	svc := service.NewAuthService()

	user, err := svc.Login(ctx, request.Username, request.Password)
	if err != nil {
		//错误日志
		return nil, err
	}

	// 做完了业务，组装一些响应数据
	resp := &userProto.LoginResponse{
		Token: "token",
		User: &userProto.User{
			Id:   user.Id,
			Name: user.Username,
		},
	}

	// 以 proto 的 response 格式，进行响应
	return resp, nil
}
