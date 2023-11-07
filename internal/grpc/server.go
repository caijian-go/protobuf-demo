package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"protobuf-demo/internal/ctl"
	userProto "protobuf-demo/proto/pb/users"
)

func StartGRPCServer() {

	ln, err := net.Listen("tcp", ":8000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	// 注册路由和处理器
	authCtl := ctl.NewAuthController()

	userProto.RegisterAuthServiceServer(s, authCtl)

	err = s.Serve(ln)
	if err != nil {
		fmt.Println(err)
	}

}
