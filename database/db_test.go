package database

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetMysql(t *testing.T) {
	db, err := GetMysql()
	if err != nil {
		fmt.Println(err)
		return
	}

	var uu = struct {
		Username string
		Password string
	}{}

	db.Table("user").Find(&uu)

	fmt.Println(uu)
}

type FuckingType struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Password string
	CreateAt time.Time

	UpdateAt time.Time
}

func (FuckingType) TableName() string {
	return "user"
}

func TestCreate(t *testing.T) {
	db, err := GetMysql()
	if err != nil {
		fmt.Println(err)
		return
	}

	var uu = FuckingType{
		Username: "qqwe1",
		Password: "qqwe1",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	db.Create(&uu)

	fmt.Println(uu)
}

func TestRedis(t *testing.T) {
	rdb := GetRedis()
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	err := rdb.Set(ctx, "mychat", "hello redis", time.Hour).Err()
	if err != nil {
		t.Error(err)
	}
}
