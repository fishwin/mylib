package multi_goroutine

import (
	"context"
	"log"
	"strconv"
	"sync"
)

type SubmitStrategyType string

const (
	Random SubmitStrategyType = "random"
)

const (
	defaultGoroutinePrefixName = "goroutine_"
)

type OperateFunc func(goroutineName string, args interface{}) error

type MultiGoroutine struct {
	wg             sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
	goroutineCount int
	chanMap        map[string]chan interface{}
	operation      OperateFunc
	option         *Option
}

func NewMultiGoroutine(goroutineCount int, operation OperateFunc, options ...OptionFunc) *MultiGoroutine {
	if operation == nil {
		panic("operation must not be nil")
	}
	if goroutineCount <= 0 {
		panic("goroutine count must be positive integer")
	}

	ctx, cancel := context.WithCancel(context.Background())
	m := &MultiGoroutine{
		ctx:            ctx,
		cancel:         cancel,
		goroutineCount: goroutineCount,
		operation:      operation,
	}

	option := &Option{}
	for _, op := range options {
		op(option)
	}
	m.option = option

	m.initChan()

	return m
}

func (m *MultiGoroutine) initChan() {
	m.chanMap = make(map[string]chan interface{})

	prefixName := defaultGoroutinePrefixName
	if m.option.goroutinePrefixName != "" {
		prefixName = m.option.goroutinePrefixName
	}

	for i := 0; i < m.goroutineCount; i++ {
		m.chanMap[prefixName+strconv.Itoa(i)] = make(chan interface{})
	}
}

func (m *MultiGoroutine) Run() {
	m.wg.Add(m.goroutineCount)

	for name, ch := range m.chanMap {
		go func(name string, ch chan interface{}) {
			defer func() {
				m.wg.Done()
				if e := recover(); e != nil {
					log.Fatalf("%s exit unexpected\n", name)
				}
				log.Printf("%s exit\n", name)
			}()

			for {
				select {
				case args := <-ch:
					if err := m.operation(name, args); err != nil {
						log.Printf("%s error(%v) happen\n", name, err)
						m.writeError(err)
					}
				case <-m.ctx.Done():
					return
				}
			}
		}(name, ch)
	}
}

func (m *MultiGoroutine) writeError(err error) {
	if m.option.errChan == nil {
		return
	}
	select {
	case m.option.errChan <- err:
	default:
		log.Printf("error chan bolocked\n")
	}
}

func (m *MultiGoroutine) Submit(args interface{}) {
	switch m.option.submitStrategy {
	case Random:
		for _, ch := range m.chanMap {
			ch <- args
			break
		}
	default:
		for _, ch := range m.chanMap {
			ch <- args
			break
		}
	}
}

func (m *MultiGoroutine) CLose() {
	for name, ch := range m.chanMap {
		log.Printf("closing %s channel\n", name)
		close(ch)
	}
	m.cancel()
	m.wg.Wait()
}
