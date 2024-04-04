package mychannel

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func RunSum() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// by default, on non-buffered channels, both sends and receives block until the other side is ready
// buffered channel
// sends to a buffered channel block only when the buffer is full
// receives block only when the buffer is empty
func RunBufferChannel() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 // creates a deadlock bc sends to a buffered channel block when the buffer is full; leaving the channel empty also blocks bc receives from a buffered channel block when empty
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fib(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // only need to close when there's a need to tell the receiver that there are no more values coming, like when a loop terminates
}

func RunFib() {
	c := make(chan int, 10)

	go fib(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
