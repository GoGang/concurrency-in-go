package main

import "time"

func producer(c chan int) {
	i := 0
	for {
		c <- i
		i++
		time.Sleep(time.Millisecond)
	}
}

func consumer(c chan int) {
	ticker := time.Tick(time.Millisecond)
	result := 0
	for {
		select {
		case val := <-c:
			result = result + val
			println(result)
		case <-ticker:
			println("\nFini:", result)
			return
		}
	}
}

func main() {
	ones := make(chan int)
	go consumer(ones)
	producer(ones)
	time.Sleep(2 * time.Millisecond)
}
