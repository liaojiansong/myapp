package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	server := g.Server()

	server.SetAccessLogEnabled(true)
	server.Run()
}
