package main

import (
	"fmt"
)

func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

//The pass of the channel type, channel is also a built-in type in go so it can be passed through in channel too.

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func main() {
chs:
	make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}

}
