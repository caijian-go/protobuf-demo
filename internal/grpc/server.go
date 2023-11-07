package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"protobuf-demo/internal/ctl"
	userProto "protobuf-demo/proto/pb/users"
)

func StartGRPCServer() {

	port := ":8000"

	ln, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	// 注册路由和处理器
	authCtl := ctl.NewAuthController()

	userProto.RegisterAuthServiceServer(s, authCtl)

	fmt.Println("serve successful... ")
	log.Println("listening on " + port)

	err = s.Serve(ln)

	if err != nil {
		fmt.Println(err)
	}

}
