package ctl

import (
	"context"
	userProto "protobuf-demo/proto/pb/users"
)

type AuthController struct {
}

func (a AuthController) Login(ctx context.Context, request *userProto.LoginRequest) (*userProto.LoginResponse, error) {

}

func (a AuthController) Register(ctx context.Context, request *userProto.RegisterRequest) (*userProto.RegisterResponse, error) {

}

func (a AuthController) mustEmbedUnimplementedAuthServiceServer() {
	
}

