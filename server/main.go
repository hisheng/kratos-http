package main

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hisheng/kratos-http/api"
	"github.com/hisheng/kratos-http/service"
	"net"
)

var h = func(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("s")
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := http.NewServer(http.Listener(ln))
	userService := service.UserService{}
	api.RegisterUserHTTPServer(server, userService)

	_ = server.Serve(ln)
}
