package api

import (
	"fmt"

	"github.com/allbuleyu/chat/api/router"
	"github.com/allbuleyu/chat/config"
)

type api struct {
}

func New() *api {
	return &api{}
}

func (*api) Run() {
	r := router.Register()

	cfg := config.GetApiConf()
	port := cfg.Port
	r.Run(fmt.Sprintf(":%s", port))
}
