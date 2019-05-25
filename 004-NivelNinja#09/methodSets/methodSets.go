package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("esta hablando", p.nombre)
	//return
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {

	p1 := persona{
		nombre: "daniel",
	}
	//diAlgo(&p1)
	p1.hablar()
}
