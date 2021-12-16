package rpc

import (
	"errors"
	"time"

	"github.com/allbuleyu/chat/database"
	"github.com/allbuleyu/chat/proto"
)

type User struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Password string
	Token    string

	CreateAt time.Time
	UpdateAt time.Time
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
	user.Username = in.UserName
	user.Password = in.PassWord
	user.Token = ""
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
		user.Token = ""
		db.Select("token").Update(&user)
		return errors.New("pwd is wrong!")
	}

	user.Token = user.Username
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
		return nil
	}

	user.Token = ""
	user.UpdateAt = time.Now()
	err = db.Select("token").Update(&user).Error

	return err
}
