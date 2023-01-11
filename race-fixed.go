package main

import "sync/atomic"

type WaitGroup struct {
	state int64
}

func (wg *WaitGroup) Add(n int64) {
	atomic.AddInt64(&wg.state, n)
}

func (wg *WaitGroup) Done() {
	atomic.AddInt64(&wg.state, -1)
}

func (wg *WaitGroup) Wait() {
	for wg.state != 0 {
		//runs forever
	}
}

func main() {
	wg := WaitGroup{0}
	wg.Add(2)
	go func() {
		wg.Done()
	}()
	wg.Wait()
}
