package funciones

import "fmt"

func Calculos() {
	//Anonim Function Declaration
	suma := func(num01 int, num02 int) int {
		return num01 + num02
	}

	var numero03 int = 32
	var numero04 int = 243

	calculo := func(numero01 int, numero02 int) int {
		return numero01 + numero02 + numero03 + numero04
	}
	fmt.Println(calculo(43, 35))

	calculo = func(numero01 int, numero02 int) int {
		return numero01 * numero02 * numero03
	}
	fmt.Println(calculo(10, 25))
	fmt.Println(suma(10, 34))
}
