package server

import (
	"fmt"
	logger "github.com/rs/zerolog/log"
	"net"
)

const BufferSize = 1024

type Config struct {
	Port string
	Addr string
}

type Server struct {
	Cfg   Config
	conns []net.Conn
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Cfg.Addr, s.Cfg.Port))
	if err != nil {
		panic(err)
	}

	logger.Info().Msg("Server started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error().Msgf("error while receiving connection: %s\n", err)
		}

		go func() {
			s.conns = append(s.conns, conn)
			logger.Info().Msg("New connection")

			for {
				msgBuf := make([]byte, BufferSize)
				if _, err := conn.Read(msgBuf); err != nil {
					logger.Error().Msgf("error while reading message: %s\n", err)
					conn.Close()
					break
				}
				s.SendMessageToAll(msgBuf)
			}
		}()
	}
}

func (s *Server) SendMessageToAll(msg []byte) {
	for _, conn := range s.conns {
		if _, err := conn.Write(msg); err != nil {
			logger.Error().Msgf("errow while sending message to all users: %s", err)
		}
	}
}
