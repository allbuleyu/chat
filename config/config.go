package config

import (
	"fmt"
	"runtime"
	"strings"

	"gopkg.in/ini.v1"
)

var conf *ini.File

func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-2], "/")
	return dir
}

func init() {
	fpath := getCurrentDir()
	cfg, err := ini.Load(fpath + "/conf.ini")
	if err != nil {
		// 如果配置加载出错,则 panic
		panic(fmt.Sprintf("load conf:%s", err))
	}

	conf = cfg
}

type MysqlConf struct {
	Host, Port, User, Password, Database string
}

func GetMysqlConf() *MysqlConf {
	my := conf.Section("mysql")
	return &MysqlConf{
		Host:     my.Key("host").String(),
		Port:     my.Key("port").String(),
		User:     my.Key("user").String(),
		Password: my.Key("password").String(),
		Database: my.Key("database").String(),
	}
}

type RedisConf struct {
	Host, Port, User, Password, Database string
}

func GetRedisConf() *RedisConf {
	my := conf.Section("redis")
	return &RedisConf{
		Host: my.Key("host").String(),
		Port: my.Key("port").String(),

		Database: "0",
	}
}

type RPCConf struct {
	Port string
}

func GetRpcConf() *RPCConf {
	rpc := conf.Section("rpc")

	return &RPCConf{
		Port: rpc.Key("port").String(),
	}
}

type ApiConf struct {
	Port string
}

func GetApiConf() *ApiConf {
	rpc := conf.Section("api")

	return &ApiConf{
		Port: rpc.Key("port").String(),
	}
}

type SiteConf struct {
	Port string
}

func GetSiteConf() *SiteConf {
	rpc := conf.Section("site")

	return &SiteConf{
		Port: rpc.Key("port").String(),
	}
}

type WsConf struct {
	Port string
}

func GetWsConf() *WsConf {
	rpc := conf.Section("websocket")

	return &WsConf{
		Port: rpc.Key("port").String(),
	}
}
