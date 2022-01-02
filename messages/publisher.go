package messages

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var __consumers map[uuid.UUID]*Consumer

func init() {
	__consumers = make(map[uuid.UUID]*Consumer)
}

type Publisher struct {
	consumers map[uuid.UUID]*Consumer
}

func (p *Publisher) Consume(consumer *Consumer) {
	p.consumers[consumer.id] = consumer
}

func (p *Publisher) Unconsume(consumer *Consumer) {
	delete(p.consumers, consumer.id)
}

func (p *Publisher) Broadcast(msg *Message) {
	if !msg.validate() {
		return
	}
	now := time.Now()
	for _, consumer := range p.consumers {
		// TODO: configurable connstant
		if now.After(consumer.lastReaded.Add(60 * time.Second)) {
			p.Unconsume(consumer)
			consumer.err <- errors.New("timeout exceeded")
			continue
		}
		// FIXME: Mb need here some sign of open chan; potential deadlock
		consumer.msgChan <- msg
	}
}

func NewPublsher() *Publisher {
	return &Publisher{
		consumers: __consumers,
	}
}
