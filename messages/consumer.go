package messages

import (
	"time"

	"github.com/google/uuid"
)

type Consumer struct {
	lastReaded time.Time
	id         uuid.UUID
	msgChan    chan *Message
	err        chan error
}

func (c *Consumer) Read() (*Message, error) {
	select {
	case m := <-c.msgChan:
		return m, nil
	case e := <-c.err:
		return nil, e
	}
}

func (c *Consumer) Confirm() {
	c.lastReaded = time.Now()
}

func NewConsumer() *Consumer {
	return &Consumer{
		id:      uuid.New(),
		msgChan: make(chan *Message),
		err:     make(chan error),
	}
}
