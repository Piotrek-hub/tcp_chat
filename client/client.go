package client

import (
	"bufio"
	"fmt"
	logger "github.com/rs/zerolog/log"
	"log"

	"net"
	"os"
	"tcp_chat/server"
)

const BufferSize = 1024

type Client struct {
	serverCfg server.Config
	username  string
	conn      net.Conn
	messages  [][]byte
}

func New(scfg server.Config, username string) *Client {
	return &Client{
		serverCfg: scfg,
		username:  username,
	}
}

func (c *Client) Start() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.serverCfg.Addr, c.serverCfg.Port))
	if err != nil {
		logger.Error().Err(err).Msg("error while trying to connect to server")
	}

	log.Print("\033[H\033[2J")

	c.conn = conn

	go c.ReadMessages()

	for {
		var input string
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input = scanner.Text()

			msg := fmt.Sprintf("%s: %s", c.username, input)
			c.SendMessage([]byte(msg))
		}
	}
}

func (c *Client) SendMessage(msg []byte) {
	if _, err := c.conn.Write(msg); err != nil {
		logger.Error().Err(err).Msg("error while sending message")
	}
}

func (c *Client) ReadMessages() {
	for {
		msgBuf := make([]byte, BufferSize)
		if _, err := c.conn.Read(msgBuf); err != nil {
			if err == net.ErrClosed {
				logger.Error().Msg("server down")
			}
			logger.Error().Err(err).Msg("error while reading message")
			return
		}

		c.messages = append(c.messages, msgBuf)

		log.Print("\033[H\033[2J")
		for _, msg := range c.messages {
			log.Println(string(msg))
		}
	}
}
