package main

import (
	"context"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type WithParallel interface {
	DoWithParallel(wg *sync.WaitGroup)
}

func main() {
	// wg := &sync.WaitGroup{}

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	cur := A{
	// 		AA: strconv.Itoa(i) + "'s home.",
	// 	}
	// 	go cur.DoWithParallel(wg)
	// }

	// wg.Wait()

	var total int = 12
	mutex := &sync.Mutex{}
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < total; i++ {
		var t = i
		go func() {
			mutex.Lock()
			cancel()
			println("the " + strconv.Itoa(t))
			mutex.Unlock()
		}()
	}
	<-ctx.Done()
}

type A struct {
	AA string
}

func (a A) DoWithParallel(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
	println(a.AA)
	wg.Done()
}
