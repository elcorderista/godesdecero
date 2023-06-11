package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() string {

	var texto string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			for i := 1; i <= 10; i++ {
				texto += fmt.Sprintf("%d x %d = %d \n", i, numero, i*numero)
			}
			break
		}
	}
	return texto
}
