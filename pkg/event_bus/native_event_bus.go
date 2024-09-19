package eventbus

import (
	systemdomain "digital-bank/internal/system/domain"
	"sync"
)

type Message struct {
	Data  interface{}
	Topic systemdomain.Topic
}

type NativeEventBus struct {
	subscribers map[systemdomain.Topic][]func(Message)
	mu          sync.RWMutex
}

var instance *NativeEventBus
var once sync.Once

func NewNativeEventBus() *NativeEventBus {
	once.Do(func() {
		instance = &NativeEventBus{
			subscribers: make(map[systemdomain.Topic][]func(Message)),
		}
	})
	return instance
}

func (e *NativeEventBus) Emit(data interface{}, topic systemdomain.Topic) error {
	e.mu.RLock()
	if callbacks, found := e.subscribers[topic]; found {
		for _, callback := range callbacks {
			go func(cb func(Message)) {
				cb(Message{Data: data, Topic: topic})
			}(callback)
		}
	}
	e.mu.RUnlock()

	return nil
}

func (e *NativeEventBus) Subscribe(topic systemdomain.Topic, callback func(Message)) {
	e.mu.Lock()
	e.subscribers[topic] = append(e.subscribers[topic], callback)
	e.mu.Unlock()
}
