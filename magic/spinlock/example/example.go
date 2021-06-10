package main

import (
	"fmt"
	"sync"

	"github.com/fishwin/mylib/magic/spinlock"
)

var m map[int]int

var l = spinlock.NewSpinLock()

func get(key int) int {
	l.Lock()
	defer l.Unlock()
	return m[key]
}

func set(key, value int) {
	l.Lock()
	defer l.Unlock()
	m[key] = value
}

func main() {
	m = make(map[int]int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1e9; i++ {
			set(i, i*10)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1e9; i++ {
			fmt.Println(get(i))
		}
	}()

	wg.Wait()
}
