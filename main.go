package main

import (
	"fmt"

	"github.com/elcorderista/godesdecero/variables"
)

func main() {
	estado, texto := variables.ConviertoTexto(1588)
	fmt.Println(texto)
	fmt.Println(estado)
}
