package database

import (
	"fmt"
	"time"

	"github.com/allbuleyu/chat/config"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

func GetMysql(database ...string) (*gorm.DB, error) {
	cfg := config.GetMysqlConf()
	if len(database) == 0 {
		database = append(database, cfg.Database)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, database[0])
	dbb, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	dbb.DB().SetMaxIdleConns(4)
	dbb.DB().SetMaxOpenConns(20)
	dbb.DB().SetConnMaxIdleTime(10 * time.Second)

	return dbb, nil
}

func GetRedis(database ...int) *redis.Client {
	cfg := config.GetRedisConf()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
