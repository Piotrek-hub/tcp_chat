package client

import (
	"bufio"
	"fmt"
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
		log.Printf("error while trying to connect to server: %s\n", err)
	}

	c.conn = conn

	go func() {
		c.ReadMessages()

	}()

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
		log.Printf("error while sending message: %s\n", err)
	}
}

func (c *Client) ReadMessages() {
	for {
		msgBuf := make([]byte, BufferSize)
		if _, err := c.conn.Read(msgBuf); err != nil {
			log.Printf("error while reading message: %s", err)
		}

		c.messages = append(c.messages, msgBuf)

		log.Print("\033[H\033[2J")
		for _, msg := range c.messages {
			log.Printf(string(msg))
		}
	}
}
