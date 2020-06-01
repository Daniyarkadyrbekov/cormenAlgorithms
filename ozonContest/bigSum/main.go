package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {

		var outInd = 0
		mu := sync.Mutex{}
		wg := sync.WaitGroup{}
		defer close(out)

		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(ind, x1, x2 int) {
				ch1 := make(chan int)
				ch2 := make(chan int)

				go func() {
					ch1 <- f(x1)
				}()
				go func() {
					ch2 <- f(x2)
				}()

				res := <-ch1 + <-ch2
				for {
					mu.Lock()
					if outInd == ind {
						out <- res
						outInd += 1
						mu.Unlock()
						break
					}
					mu.Unlock()
					runtime.Gosched()
				}
				wg.Done()
			}(i, <-in1, <-in2)
		}
		wg.Wait()
	}()
}

func main() {
	f := func(i int) int {
		return i
	}
	in1, in2, out := make(chan int), make(chan int), make(chan int)
	n := 100
	Merge2Channels(f, in1, in2, out, n)

	for i := 0; i < n; i++ {
		in1 <- i
		in2 <- i * 2
	}

	for i := 0; i < n; i++ {
		res := <-out

		fmt.Printf("i = %d res = %d\n", i, res)
		if res != f(i)+f(i*2) {
			fmt.Printf("res = %d exp = %d\n", res, f(i)+f(i*2))
			panic("olololo")
		}
	}
}
