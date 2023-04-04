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
	Cfg      Config
	conns    []net.Conn
	messages [][]byte
}

func (s *Server) GetMessages() [][]byte {
	return s.messages
}

func (s *Server) AddMessage(msg []byte) {
	s.messages = append(s.messages, msg)
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

		s.SendOldMessages(conn)

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

func (s *Server) SendOldMessages(conn net.Conn) {
	for _, msg := range s.messages {
		if _, err := conn.Write(msg); err != nil {
			logger.Error().Msgf("error while sending old messages: %s", err)
		}
	}
}

func (s *Server) SendMessageToAll(msg []byte) {
	s.AddMessage(msg)

	for _, conn := range s.conns {
		if _, err := conn.Write(msg); err != nil {
			logger.Error().Msgf("errow while sending message to all users: %s", err)
		}
	}
}
