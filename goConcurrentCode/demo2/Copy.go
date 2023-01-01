package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
}

// go vet Copy.go
func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c)
}

func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
