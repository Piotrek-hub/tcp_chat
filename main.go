package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"tcp_chat/client"
	"tcp_chat/server"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	port := flag.String("port", "9090", "server port")
	addr := flag.String("addr", "localhost", "server address")
	isServer := flag.Bool("server", false, "choose between client/server")
	isClient := flag.Bool("client", false, "choose between client/server")
	username := flag.String("username", "", "client username")

	flag.Parse()

	cfg := server.Config{
		Port: *port,
		Addr: *addr,
	}

	if *isServer {
		s := server.New(cfg)
		s.Start()
	}

	if *isClient {
		c := client.New(cfg, *username)
		c.Start()
	}

}
