package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema Operativo:\t\t", runtime.GOOS)
	fmt.Println("Arquitectura del Equipo:\t", runtime.GOARCH)
	fmt.Println("Numero de Nucleos del CPU:\t", runtime.NumCPU())
	fmt.Println("Numero de Gorutines:\t\t", runtime.NumGoroutine())
}
