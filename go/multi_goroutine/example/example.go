package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fishwin/mylib/go/multi_goroutine"
)

func operation(goroutineName string, args interface{}) error {
	fmt.Println(goroutineName, args)
	time.Sleep(100 * time.Millisecond)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if r.Intn(10) == 6 {
		return errors.New(fmt.Sprintf("%v error %v", goroutineName, time.Now()))
	}
	return nil
}

func main() {
	errChan := make(chan error, 10)

	mg := multi_goroutine.NewMultiGoroutine(
		10,
		operation,
		multi_goroutine.WithErrChan(errChan),
		multi_goroutine.WithGoroutinePrefixName("gogogo_"),
		multi_goroutine.WithSubmitStrategy(multi_goroutine.Random),
	)
	mg.Run()

	go func() {
		for {
			select {
			case err := <-errChan:
				fmt.Println(err)
			}
		}
	}()

	for i := 0; i < 1e9; i++ {
		mg.Submit(i)
	}

	mg.CLose()
}
