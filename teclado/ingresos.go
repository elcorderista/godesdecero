package teclado

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numero01 int
var numero02 int
var leyenda string
var err error

// зто Метод
func IngresoNumeros() {
	//Creamos el scanner especifico del teclado
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 01: ")
	if scanner.Scan() {
		numero01, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicf("El dato ingresado es incorrecto %s ", err.Error())
		}
	}

	fmt.Println("Ingrese numero 02: ")
	if scanner.Scan() {
		numero02, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicf("El dato ingresado es incorrecto %s ", err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero01*numero02)
}
