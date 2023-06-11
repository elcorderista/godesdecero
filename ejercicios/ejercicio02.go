package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplicar() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			for i := 1; i <= 10; i++ {
				fmt.Printf("%d x %d = %d \n", i, numero, i*numero)
			}
			break
		}
	}
}
