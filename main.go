package main

import (
	"fmt"
	"time"
)

var fibonacci = func() chan uint64 {
	c := make(chan uint64)
	go func() {
		var x, y uint64 = 0, 1
		for ; y < (1 << 63); c <- y {
			x, y = y, x+y
		}
		close(c)
	}()
	return c
}

func main() {
	c := fibonacci()
	for x, ok := <-c; ok; x, ok = <-c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
