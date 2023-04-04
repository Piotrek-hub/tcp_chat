package main

import (
	"flag"

	"os"
	"tcp_chat/client"
	"tcp_chat/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

}

func main() {
	port := flag.String("port", "9090", "server port")
	addr := flag.String("addr", "localhost", "server address")
	isServer := flag.Bool("server", false, "choose between client/server")
	isClient := flag.Bool("client", false, "choose between client/server")
	username := flag.String("username", "", "client username")
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	cfg := server.Config{
		Port: *port,
		Addr: *addr,
	}

	if *isServer {
		s := server.Server{Cfg: cfg}
		s.Start()
	}

	if *isClient {
		c := client.New(cfg, *username)
		c.Start()
	}

}
