package mymutex

import (
	"fmt"
	"sync"
	"time"
)

// safeCounter is safe to use concurrently
type safeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// inc increments the counter for the given key
func (c *safeCounter) inc(key string) {
	c.mu.Lock() // lock so only 1 goroutine at a time can access the map c.v
	c.v[key]++
	c.mu.Unlock()
}

// value returns the current value of the counter for the given key
func (c *safeCounter) value(key string) int {
	c.mu.Lock() // lock so only 1 goroutine at a time can access the map c.v
	defer c.mu.Unlock()
	return c.v[key]
}

func Run() {
	c := safeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}
