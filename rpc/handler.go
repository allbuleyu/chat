package rpc

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/allbuleyu/chat/database"
	"github.com/allbuleyu/chat/proto"
	"github.com/bwmarrin/snowflake"
)

const (
	redisPrefix string = "Sess_"
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

func register(in *proto.RegisterRequest) (string, error) {
	db, err := database.GetMysql()
	if err != nil {
		return "", err
	}

	var user User
	db.Where("username=?", in.UserName).Find(&user)
	if user.Id != 0 {
		return "", newError("username has exist!")
	}

	user.Username = in.UserName
	user.Password = in.PassWord
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	db.Create(&user)

	return redisSet(in.UserName, user)
}

func login(in *proto.LoginRequest) (string, error) {
	db, err := database.GetMysql()
	if err != nil {
		return "", err
	}

	var user User

	err = db.Where("username=?", in.UserName).Find(&user).Error
	if err != nil {
		return "", err
	}

	if user.Password != in.PassWord {
		return "", newError("pwd is wrong!")
	}

	return redisSet(in.UserName, user)
}

func logout(in *proto.LogoutRequest) error {
	rdb := database.GetRedis()

	return rdb.Del(context.Background(), redisPrefix+in.Token).Err()
}

func checkToken(in *proto.LogoutRequest) (User, error) {
	return redisGet(redisPrefix + in.Token)
}

func redisSet(k string, v interface{}) (string, error) {
	rdb := database.GetRedis()
	token := GetSnowflakeId()
	key := redisPrefix + token

	vv, err := json.Marshal(&v)
	if err != nil {
		return "", err
	}
	err = rdb.Set(context.Background(), key, vv, time.Hour).Err()

	return token, err
}

func redisGet(k string) (User, error) {
	rdb := database.GetRedis()
	s, err := rdb.Get(context.Background(), k).Result()
	user := User{}
	if err != nil {
		return user, err
	}

	err = json.Unmarshal([]byte(s), &user)

	return user, err
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
