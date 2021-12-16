package database

import (
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
