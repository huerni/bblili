package main

import (
	mywebsoc "bblili/api/internal/common/websocket"
	"bblili/api/internal/config"
	"bblili/api/internal/handler"
	"bblili/api/internal/svc"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/bblili-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			mywebsoc.ServerBarrageWs(w, r, ctx)
		},
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
