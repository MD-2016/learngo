package syncPack

import "sync"

type Counter struct {
	mut sync.Mutex
	val int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
