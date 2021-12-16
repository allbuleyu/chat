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

type FormRegister struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"passWord" json:"passWord" binding:"required"`
}

type ResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func FailResponse(c *gin.Context, err error) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:    1,
		Message: fmt.Sprintf("%s", err),
	})
}

func SuccessResponse(c *gin.Context, data string) {
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
	rep, _ := client.RPCLogin(context.Background(), rd)
	if rep.Code != 0 {
		FailResponse(c, errors.New(rep.Msg))
		return
	}

	SuccessResponse(c, rd.UserName)
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
	rep, _ := client.RPCRegister(context.Background(), rd)
	if rep.Code != 0 {
		FailResponse(c, errors.New(rep.Msg))
		return
	}

	SuccessResponse(c, rd.UserName)
}

func CheckAuth(c *gin.Context) {
	token := c.Param("authToken")

	if token == "" {
		c.JSON(http.StatusOK, &ResponseData{
			Code: 0,
			Data: "",
		})
	}
}
