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
	token, err := login(in)
	if err != nil {
		rep.Code = 1
		rep.Msg = fmt.Sprint(err)
	}
	rep.Token = token

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
	token, err := register(in)
	if err != nil {
		rep.Code = 1
		rep.Msg = fmt.Sprint(err)
	}
	rep.Token = token

	return rep, nil
}

func (s *rpcServer) RPCLogout(ctx context.Context, in *proto.LogoutRequest) (*proto.AuthResponse, error) {
	if in.Token == "" {
		return &proto.AuthResponse{
			Code: 1,
			Msg:  "token invalid!",
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

func (s *rpcServer) RPCCheckToken(ctx context.Context, in *proto.LogoutRequest) (*proto.AuthResponse, error) {
	if in.Token == "" {
		return &proto.AuthResponse{
			Code: 1,
			Msg:  "uid or token invalid!",
		}, nil
	}

	user, err := checkToken(in)
	if user.Id == 0 && err == nil {
		err = newError("wrong token!")
	}
	rep := &proto.AuthResponse{}
	if err != nil {
		rep.Code = 1
		rep.Msg = fmt.Sprint(err)
	}
	rep.Token = user.Username

	return rep, nil
}

func (mr *myRpc) Run() {
	cfg := config.GetRpcConf()
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		fmt.Printf("??????????????????: %s", err)
		return
	}

	// ??????gRPC?????????
	s := grpc.NewServer()
	// ????????????
	proto.RegisterChatAuthServer(s, &rpcServer{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("??????????????????: %s", err)
		return
	}
}
