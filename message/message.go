package message

import (
	"fmt"
	logger "github.com/rs/zerolog/log"
	"strings"
	"time"
)

type Message struct {
	Content string
	Author  string
	Date    string
}

func New(author, content string) *Message {
	t := time.Now()

	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	return &Message{
		Content: content,
		Author:  author,
		Date:    fmt.Sprintf("%d:%d:%d", h, m, s),
	}
}

func (m *Message) Print() string {
	return fmt.Sprintf("%-5s> %s:%s", m.Date, m.Author, m.Content)
}

func (m *Message) PrintByte() []byte {
	return []byte(m.Print())
}

func NewFromBuffer(buf []byte) *Message {
	splittedMsg := strings.Split(string(buf), ":")

	author := splittedMsg[0]
	content := splittedMsg[1]

	return New(author, content)
}

func NewFromBufferWithTime(buf []byte) *Message {

	splittedMessage := strings.Split(string(buf), ">")
	logger.Error().Msgf("essa %s", len(buf))
	msg := NewFromBuffer([]byte(splittedMessage[1]))
	msg.Date = splittedMessage[0]

	logger.Error().Msgf("msg: %v", msg)

	return msg
}
