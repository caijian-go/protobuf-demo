package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	userProto "protobuf-demo/proto/pb/users"
	"sync/atomic"
)

func main() {

	pool := NewUserClientPool("localhost:8000", 20)

	cli := pool.Get()
	resp, err := cli.Login(context.Background(), &userProto.LoginRequest{
		Username: "admin",
		Password: "123456",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Token, resp.User.Id)

}

type UserClientPool struct {
	clients []userProto.AuthServiceClient
	index   int64
}

func NewUserClientPool(addr string, size int) *UserClientPool {

	var clients []userProto.AuthServiceClient

	for i := 0; i < size; i++ {
		cc, err := grpc.Dial(addr, grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		clients = append(clients, userProto.NewAuthServiceClient(cc))
	}

	return &UserClientPool{
		clients: clients,
		index:   0,
	}
}

// 0 1 2 3 4 0 1 ...
func (p UserClientPool) Get() userProto.AuthServiceClient {
	//新增index
	index := atomic.AddInt64(&p.index, 1)
	return p.clients[int(index)%len(p.clients)]
}

func (p UserClientPool) Release(c userProto.AuthServiceClient) {

}
