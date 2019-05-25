package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Sistema Operativo: %v\nArquitectura: %v\n", runtime.GOOS, runtime.GOARCH)
}
