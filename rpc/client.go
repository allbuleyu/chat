package rpc

import (
	"github.com/allbuleyu/chat/config"
	"github.com/allbuleyu/chat/proto"
	"google.golang.org/grpc"
)

func GetClientConn() (*grpc.ClientConn, proto.ChatAuthClient, error) {
	cfg := config.GetRpcConf()
	conn, err := grpc.Dial(":"+cfg.Port, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	// 新建一个客户端
	c := proto.NewChatAuthClient(conn)
	return conn, c, nil
}
