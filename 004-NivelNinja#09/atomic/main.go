package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var counter int64 //la variable debe ser tipo int64 para invocar la funcion

func main() {
	fmt.Println("Numero de CPUs", runtime.NumCPU())
	fmt.Println("Numero de Goroutines inicio", runtime.NumGoroutine())
	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//v := counter
			runtime.Gosched()
			//v++
			//counter = v
			atomic.LoadInt64(&counter)
			wg.Done()
		}()
		fmt.Println("Numero de Goroutines dentro del for", runtime.NumGoroutine())
	}
	fmt.Println("Numero de Goroutines final", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(counter)
	//para ver la condicion race usa el comando: go run -race main.go
}
