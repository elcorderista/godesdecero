package main

import (
	"github.com/elcorderista/godesdecero/iteraciones"
)

func main() {

	/* estado, texto := variables.ConviertoTexto(1588)
	fmt.Println(texto)
	fmt.Println(estado)
	if os := runtime.GOOS; os == "Linux." || os == "darwin" {
		fmt.Println("Esto no es Windows")
		fmt.Println(os)
	} else {
		fmt.Println("Esto es Windows")
		fmt.Println(os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwing")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.Ejercicio01Resuelto("101")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()
	*/

	iteraciones.Iterar()
	iteraciones.IterarReverso()
	iteraciones.IterarSaltos()
}
