package main

import (
	"flag"
	"tcp_chat/client"
	"tcp_chat/server"
)

func main() {
	port := flag.String("port", "9090", "server port")
	addr := flag.String("addr", "localhost", "server address")
	isServer := flag.Bool("server", false, "choose between client/server")
	username := flag.String("username", "", "client username")

	flag.Parse()

	cfg := server.Config{
		Port: *port,
		Addr: *addr,
	}

	if *isServer {
		s := server.Server{Cfg: cfg}
		s.Start()
	}

	c := client.New(cfg, *username)
	c.Start()

}
