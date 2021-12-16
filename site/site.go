package site

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allbuleyu/chat/config"
)

type site struct {
}

func New() *site {
	return &site{}
}

func (*site) Run() {
	// r := gin.Default()

	// r.LoadHTMLFiles("site/index.html", "site/login.html", "site/register.html")
	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })

	// r.Run(":8080")

	cfg := config.GetSiteConf()
	port := cfg.Port
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), http.FileServer(http.Dir("./site/"))))
}
