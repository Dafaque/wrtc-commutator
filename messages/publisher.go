package messages

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

var __consumers map[uuid.UUID]*Consumer

func init() {
	__consumers = make(map[uuid.UUID]*Consumer)
}

type Publisher struct {
	mu        sync.RWMutex
	consumers map[uuid.UUID]*Consumer
}

func (p *Publisher) Consume(consumer *Consumer) {
	p.consumers[consumer.id] = consumer
	println("consumed", consumer.id.String(), len(__consumers))
}

func (p *Publisher) Unconsume(consumer *Consumer) {
	delete(p.consumers, consumer.id)
	println("unconsumed", consumer.id.String(), len(__consumers))
}

func (p *Publisher) Broadcast(msg *Message) {
	if !msg.validate() {
		return
	}
	now := time.Now()
	println("broadcasting msg", string(msg.To), string(msg.With))
	p.mu.Lock()
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
	p.mu.Unlock()
}

func NewPublsher() *Publisher {
	return &Publisher{
		consumers: __consumers,
	}
}
