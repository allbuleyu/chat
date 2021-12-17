package rpc

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/allbuleyu/chat/database"
	"github.com/allbuleyu/chat/proto"
	"github.com/bwmarrin/snowflake"
)

type User struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Password string
	Token    string

	CreateAt time.Time
	UpdateAt time.Time
}

func newError(msg string) error {
	return errors.New(msg)
}

func (User) TableName() string {
	return "user"
}

func register(in *proto.RegisterRequest) error {
	db, err := database.GetMysql()
	if err != nil {
		return err
	}

	var user User
	db.Where("username=?", in.UserName).Find(&user)
	if user.Id != 0 {
		return newError("username has exist!")
	}

	user.Password = in.PassWord
	user.Token = generateToken(GetSnowflakeId())
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	db.Create(&user)

	return err
}

func login(in *proto.LoginRequest) error {
	db, err := database.GetMysql()
	if err != nil {
		return err
	}

	var user User
	user.Username = in.UserName
	err = db.Find(&user).Error
	if err != nil {
		return err
	}

	if user.Password != in.PassWord {
		return newError("pwd is wrong!")
	}

	user.Token = generateToken(GetSnowflakeId())
	err = db.Select("token").Update(&user).Error

	return err
}

func logout(in *proto.LogoutRequest) error {
	db, err := database.GetMysql()
	if err != nil {
		return err
	}

	var user User
	user.Id = int(in.Uid)
	err = db.Find(&user).Error
	if err != nil {
		return err
	}

	if user.Token != in.Token {
		return newError("wrong token!")
	}

	user.Token = ""
	user.UpdateAt = time.Now()
	err = db.Select("token").Update(&user).Error

	return err
}

func checkToken(in *proto.LogoutRequest) error {
	db, err := database.GetMysql()
	if err != nil {
		return err
	}

	var user User
	user.Id = int(in.Uid)
	err = db.Find(&user).Error
	if err != nil {
		return err
	}

	if user.Token != in.Token {
		return newError("wrong token!")
	}

	return nil
}

func generateToken(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func GetSnowflakeId() string {
	//default node id eq 1,this can modify to different serverId node
	node, _ := snowflake.NewNode(1)
	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}

func Sha1(s string) (str string) {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
