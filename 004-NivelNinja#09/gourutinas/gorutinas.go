package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("rutina principal")
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("funcion foo")
	wg.Done()
}

func bar() {
	fmt.Println("funcion bar")
	wg.Done()
}
