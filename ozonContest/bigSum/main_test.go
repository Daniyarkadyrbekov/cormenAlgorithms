package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConcurency(t *testing.T) {
	timeStart := time.Now()
	f := func(i int) int {
		fmt.Printf("i = %d time since Start = %fsec\n", i, time.Since(timeStart).Seconds())
		time.Sleep(time.Duration(rand.Int63()%10) * time.Second)
		return i
	}
	in1, in2, out := make(chan int), make(chan int), make(chan int)
	n := 100
	Merge2Channels(f, in1, in2, out, n)

	for i := 0; i < n; i++ {
		in1 <- i
		in2 <- -2 * i
	}

	for i := 0; i < n; i++ {
		res := <-out

		fmt.Printf("i = %d res = %d\n", i, res)
		//require.Equal(t, f(i)+f(-i), res)
		require.Equal(t, -i, res)
	}
}
