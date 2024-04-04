package myselect

import (
	"fmt"
	"time"
)

func fib(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("          .")
			time.Sleep((50 * time.Millisecond))
		}
	}
}

func RunSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit)
}

func Tick() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("          .")
			fmt.Println(50 * time.Millisecond)
		}
	}
}
