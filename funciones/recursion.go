package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 1000000 {
		return
	}
	Exponencia(valor * 2)
	fmt.Println(valor)

}
