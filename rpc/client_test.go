package rpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/allbuleyu/chat/proto"
)

func Test_GetClientConn(t *testing.T) {
	conn, client, err := GetClientConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	rd := &proto.RegisterRequest{
		UserName: "qqqqq1",
		PassWord: "wwwww1",
	}

	// err 永远是 nil 所以不判断
	rep, err := client.RPCRegister(context.Background(), rd)
	fmt.Println("rep:", rep, err)
}
