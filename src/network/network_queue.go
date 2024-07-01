package network

import (
	"sync"
)

// dont need to add further safety checks a few checks as its only got one usage, will save abit of performance

type NetworkQueue struct {
	items []IncomingMessage
	lock  sync.Mutex
}

func NewNetworkQueue() *NetworkQueue {
	return &NetworkQueue{
		items: make([]IncomingMessage, 0),
		lock:  sync.Mutex{},
	}
}

func (q *NetworkQueue) Lock() {
	q.lock.Lock()
}

func (q *NetworkQueue) Push(item IncomingMessage) {
	q.items = append(q.items, item)
}

func (q *NetworkQueue) Pop() IncomingMessage {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *NetworkQueue) Size() int {
	return len(q.items)
}

func (q *NetworkQueue) Unlock() {
	q.lock.Unlock()
}
