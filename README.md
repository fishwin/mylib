# mylib
golang utils

## features
+ magic
    + multi goroutines
    + spin lock
+ data struct
    + single list
    + double list
    + stack
    + queue
## examples
### magic
#### multi goroutines
```go
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fishwin/mylib/multi_goroutine"
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

```
#### spin lock
```go
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

```

### data struct
#### single list
```go
package main

import (
	"fmt"

	"github.com/fishwin/mylib/data_struct/single_list"
)

func main() {
	sl := single_list.NewSingleList()
	sl.Append(&single_list.SingleNode{Data: 1})
	sl.Append(&single_list.SingleNode{Data: 2})
	sl.Append(&single_list.SingleNode{Data: 3})
	sl.Append(&single_list.SingleNode{Data: 4})
	sl.Append(&single_list.SingleNode{Data: 5})

	fmt.Println(sl.Length())
	sl.Display()
	fmt.Println(sl.Get(2))
	sl.Insert(3, &single_list.SingleNode{Data: 99})
	sl.Display()
	sl.Delete(5)
	sl.Display()
}
```
### double list
```go
package main

import (
	"fmt"

	"github.com/fishwin/mylib/data_struct/double_list"
)

func main() {
	dl := double_list.NewDoubleList()
	dl.Append(&double_list.DoubleNode{Data: 1})
	dl.Append(&double_list.DoubleNode{Data: 2})
	dl.Append(&double_list.DoubleNode{Data: 3})
	dl.Append(&double_list.DoubleNode{Data: 4})
	dl.Append(&double_list.DoubleNode{Data: 5})

	dl.Display()
	dl.Reverse()

	dl.Delete(2)

	dl.Display()
	dl.Reverse()

	dl.Insert(4, &double_list.DoubleNode{Data: 111})

	dl.Display()
	dl.Reverse()
	
	fmt.Println(dl.Get(3))
}

```
#### stack
```go
package main

import (
	"fmt"

	"github.com/fishwin/mylib/stack"
)

func main() {
	s := stack.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Length())
	fmt.Println(s.Pop())
	fmt.Println(s.Length())
}


```
### queue
```go
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

```



