package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

// Creacion de un midleware
func Middleware() {
	fmt.Println("Inicio")

	//Infocamos una funcion enviado una funcion como argumento y sus dos parametros
	result := operacionesMid(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMid(restar)(5, 3)
	fmt.Println(result)
	result = operacionesMid(multiplicar)(2, 3)
	fmt.Println(result)
}

// Recibe una funcion y devuelve una funcion
func operacionesMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operaion")
		return f(a, b)
	}
}
