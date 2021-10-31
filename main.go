package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cache map[int]int
	mu    sync.Mutex
)

func op(n int) int {
	return n * n
}

func getCache(n int) int {
	mu.Lock()
	c, isCached := cache[n]
	mu.Unlock()
	if isCached {
		return c
	}
	o := op(n)
	mu.Lock()
	cache[n] = o
	mu.Unlock()
	return o
}
func setChache() {
	total := 0
	for i := 0; i < 5; i++ {
		total += getCache(i)
	}
	fmt.Printf("total:%d\n", total)
}
func main() {
	cache = make(map[int]int)
	go setChache()
	setChache()
	time.Sleep(time.Second * 3)
}
