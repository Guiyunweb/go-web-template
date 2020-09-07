package server

import (
	"fmt"
	"github.com/Guiyunweb/go-web-template/server/http"
)

type Server struct {
	Post int `toml:"post"`
}

var Conf = &Server{}

func Run() {
	// 装载路由
	r := http.NewRouter()
	var adder = fmt.Sprintf(":%d", Conf.Post)
	r.Run(adder)
}

func Info(s *Server) {
	Conf = s
}
