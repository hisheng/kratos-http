package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hisheng/kratos-http/api"
	"github.com/hisheng/kratos-http/service"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := http.NewServer(http.Listener(ln))
	userService := service.UserService{}
	api.RegisterUserHTTPServer(server, userService)

	_ = server.Start(context.Background())
	// _ = server.Serve(ln) 这个也行，不过感觉上面的start更好，start会调用server.Serve
}
