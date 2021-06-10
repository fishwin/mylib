package main

import (
	"fmt"

	"github.com/fishwin/mylib/data_struct/queue"
)

func main() {
	q := queue.NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println(q.Peek())
	fmt.Println(q.Size())
}
