package ejercicios

import (
	"log"
	"strconv"
)

func Ejercicio01(texto string) (int, string) {

	numero, err := strconv.Atoi(texto)
	if err != nil {
		log.Fatal(err)
	}
	return numero, texto
}

func Ejercicio01Resuelto(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error"
	}
	if numero > 100 {
		return numero, "El numero es mayor a 100"
	} else {
		return numero, "El numero es menor a 100"
	}
}
