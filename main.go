package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/allbuleyu/chat/rpc"

	"github.com/allbuleyu/chat/api"
	_ "github.com/allbuleyu/chat/config"
	"github.com/allbuleyu/chat/site"
	"github.com/allbuleyu/chat/websocket"
)

func main() {
	var module string
	// switch reloadconf 重新加载配置到内存
	flag.StringVar(&module, "m", "", "assign run module")
	flag.Parse()

	fmt.Fprintf(os.Stdout, "start run %s module \n", module)

	switch module {
	case "api":
		api.New().Run()
	case "rpc":
		rpc.New().Run()
	case "site":
		site.New().Run()
	case "ws":
		websocket.New().Run()
	default:
		fmt.Println("You are not input module type!")
	}
}
