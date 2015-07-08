package main

import (
	"flag"
	"runtime"
	"sync"
)

var (
	cores = flag.Int("cores", 3, "nombre de cœurs")
	n     = flag.Int("n", 10, "nombre de goroutines")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*cores)
	var wg sync.WaitGroup
	wg.Add(*n)
	for i := 0; i < *n; i++ {
		go func() {
			a := 0
			for j := 0; j < 10e8; j++ {
				a++
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
