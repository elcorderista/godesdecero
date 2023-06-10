package iteraciones

import (
	"fmt"
)

// зто метод
func Iterar() {
	for i := 0; i < 50; i += 5 {
		fmt.Println(i)
	}
}

func IterarReverso() {
	for j := 10; j >= 0; j-- {
		fmt.Println(j)
	}
}

func IterarSaltos() {
	for i := 0; i < 50; i += 5 {
		if i == 20 {
			continue
		}
		if i == 40 {
			break
		}
		fmt.Println(i)
	}
}
