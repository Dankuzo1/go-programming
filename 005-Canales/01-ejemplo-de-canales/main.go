package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go enviar(c)

	recibir(c)

	fmt.Println("Finalizando.")
}

// send channel  (canal sólo para enviar)
func enviar(c chan<- int) {
	c <- 42
}

// receive channel (canal sólo para recibir)
func recibir(c <-chan int) {
	fmt.Println("El valor recibido desde el canal:", <-c)
}
