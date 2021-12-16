package rpc

import (
	"context"
	"fmt"
	"net"

	"github.com/allbuleyu/chat/config"
	"github.com/allbuleyu/chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myRpc struct {
}

func New() *myRpc {
	return &myRpc{}
}

type rpcServer struct {
	proto.UnimplementedChatAuthServer
}

func (s *rpcServer) RPCLogin(ctx context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	if in.UserName == "" || len(in.UserName) > 15 {
		return &proto.AuthResponse{
			Code: 1,
			Msg:  "userName invalid!",
		}, nil
	}

	// others check, register
	rep := &proto.AuthResponse{}
	err := login(in)
	if err != nil {
		rep.Code = 1
		rep.Msg = fmt.Sprint(err)
	}

	return rep, nil
}

func (s *rpcServer) RPCRegister(ctx context.Context, in *proto.RegisterRequest) (*proto.AuthResponse, error) {
	if in.UserName == "" || len(in.UserName) > 15 {
		return &proto.AuthResponse{
			Code: 1,
			Msg:  "userName invalid!",
		}, nil
	}

	// others check, register
	rep := &proto.AuthResponse{}
	err := register(in)
	if err != nil {
		rep.Code = 1
		rep.Msg = fmt.Sprint(err)
	}

	return rep, nil
}

func (s *rpcServer) RPCLogout(ctx context.Context, in *proto.LogoutRequest) (*proto.AuthResponse, error) {
	if in.Uid == 0 || in.Token == "" {
		return &proto.AuthResponse{
			Code: 1,
			Msg:  "uid or token invalid!",
		}, nil
	}

	err := logout(in)
	rep := &proto.AuthResponse{}
	if err != nil {
		rep.Code = 1
		rep.Msg = fmt.Sprint(err)
	}

	return rep, nil
}

func (mr *myRpc) Run() {
	cfg := config.GetRpcConf()
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterChatAuthServer(s, &rpcServer{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
