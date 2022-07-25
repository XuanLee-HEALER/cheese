package main

import (
	"fmt"
	"math/rand"
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

	// var total int = 12
	// mutex := &sync.Mutex{}
	// ctx, cancel := context.WithCancel(context.Background())
	// for i := 0; i < total; i++ {
	// 	var t = i
	// 	go func() {
	// 		mutex.Lock()
	// 		cancel()
	// 		println("the " + strconv.Itoa(t))
	// 		mutex.Unlock()
	// 	}()
	// }
	// <-ctx.Done()
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	var b int32 = int32(byte('a'))
	var e = b + 25
	ch := make(chan int32)
	go func() {
		for b != e {
			mu.Lock()
			for b%2 != 1 {
				cond.Wait()
			}
			// println("A")
			ch <- b
			if b == e {
				close(ch)
			}
			b++
			cond.Signal()
			mu.Unlock()
			if b >= e {
				break
			}
		}
	}()
	go func() {
		for b != e {
			mu.Lock()
			for b%2 != 0 {
				cond.Wait()
			}
			// println("B")
			ch <- b
			if b == e {
				close(ch)
			}
			b++
			cond.Signal()
			mu.Unlock()
			if b >= e {
				break
			}
		}
	}()
	for ele := range ch {
		fmt.Println(string(ele))
	}
}

type A struct {
	AA string
}

func (a A) DoWithParallel(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)))
	println(a.AA)
	wg.Done()
}
