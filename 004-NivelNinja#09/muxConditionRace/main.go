package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var counter = 0

func main() {
	fmt.Println("Numero de CPUs", runtime.NumCPU())
	fmt.Println("Numero de Goroutines inicio", runtime.NumGoroutine())
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Numero de Goroutines dentro del for", runtime.NumGoroutine())
	}
	fmt.Println("Numero de Goroutines final", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(counter)
	//para ver la condicion race usa el comando: go run -race main.go
}
