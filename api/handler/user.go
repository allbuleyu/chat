package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/allbuleyu/chat/proto"
	"github.com/allbuleyu/chat/rpc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 用于 rpc 请求的 ctx
var RPCctx = context.Background()

type FormRegister struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"passWord" json:"passWord" binding:"required"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FailResponse(c *gin.Context, err error) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    1,
		Message: fmt.Sprintf("%s", err),
	})
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: 0,
		Data: data,
	})
}

func Login(c *gin.Context) {
	var formRegister FormRegister
	if err := c.ShouldBindBodyWith(&formRegister, binding.JSON); err != nil {
		FailResponse(c, err)
		return
	}

	conn, client, err := rpc.GetClientConn()
	if err != nil {
		FailResponse(c, err)
		return
	}
	defer conn.Close()

	rd := &proto.LoginRequest{
		UserName: formRegister.UserName,
		PassWord: formRegister.Password,
	}

	// err 永远是 nil 所以不判断
	rep, _ := client.RPCLogin(RPCctx, rd)
	if rep.Code != 0 {
		FailResponse(c, errors.New(rep.Msg))
		return
	}

	SuccessResponse(c, rep.Token)
}

func Register(c *gin.Context) {
	var formRegister FormRegister
	if err := c.ShouldBindBodyWith(&formRegister, binding.JSON); err != nil {
		FailResponse(c, err)
		return
	}

	// grpc
	conn, client, err := rpc.GetClientConn()
	if err != nil {
		FailResponse(c, err)
		return
	}
	defer conn.Close()

	rd := &proto.RegisterRequest{
		UserName: formRegister.UserName,
		PassWord: formRegister.Password,
	}

	// err 永远是 nil 所以不判断
	rep, _ := client.RPCRegister(RPCctx, rd)
	if rep.Code != 0 {
		FailResponse(c, errors.New(rep.Msg))
		return
	}

	SuccessResponse(c, rep.Token)
}

func CheckAuth(c *gin.Context) {
	type Token struct {
		AuthToken string
	}
	var token Token
	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		FailResponse(c, err)
		return
	}

	if token.AuthToken == "" {
		FailResponse(c, errors.New("token is empty!"))
		c.Next()
	}

	// grpc
	conn, client, err := rpc.GetClientConn()
	if err != nil {
		FailResponse(c, err)
		return
	}
	defer conn.Close()

	rd := &proto.LogoutRequest{
		Token: token.AuthToken,
	}

	// err 永远是 nil 所以不判断
	rep, _ := client.RPCCheckToken(RPCctx, rd)
	if rep.Code != 0 {
		FailResponse(c, errors.New(rep.Msg))
		return
	}

	c.Set("userName", rep.Token)

	res := make(map[string]interface{})
	res["userName"] = rep.Token
	SuccessResponse(c, res)
}

func Logout(c *gin.Context) {
	type Token struct {
		AuthToken string
	}
	var token Token
	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		FailResponse(c, err)
		return
	}

	// grpc
	conn, client, err := rpc.GetClientConn()
	if err != nil {
		FailResponse(c, err)
		return
	}
	defer conn.Close()

	rd := &proto.LogoutRequest{
		Token: token.AuthToken,
	}

	// err 永远是 nil 所以不判断
	rep, _ := client.RPCLogout(RPCctx, rd)
	if rep.Code != 0 {
		FailResponse(c, errors.New(rep.Msg))
		return
	}

	SuccessResponse(c, rd.Token)
}
