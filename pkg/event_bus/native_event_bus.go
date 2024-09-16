package eventbus

import (
	"sync"
)

type Message struct {
	Data  interface{}
	Topic Topic
}

type NativeEventBus struct {
	subscribers map[Topic][]func(Message)
	mu          sync.RWMutex
}

var instance *NativeEventBus
var once sync.Once

func NewNativeEventBus() *NativeEventBus {
	once.Do(func() {
		instance = &NativeEventBus{
			subscribers: make(map[Topic][]func(Message)),
		}
	})
	return instance
}

func (e *NativeEventBus) Subscribe(topic Topic, callback func(Message)) {
	e.mu.Lock()
	e.subscribers[topic] = append(e.subscribers[topic], callback)
	e.mu.Unlock()
}

func (e *NativeEventBus) Emit(topic Topic, data interface{}) {
	e.mu.RLock()
	if callbacks, found := e.subscribers[topic]; found {
		for _, callback := range callbacks {
			go func(cb func(Message)) {
				cb(Message{Data: data, Topic: topic})
			}(callback)
		}
	}
	e.mu.RUnlock()
}
