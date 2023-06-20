package ejercicios

import (
	"fmt"
	"github.com/elcorderista/godesdecero/interfaces"
)

// Esta funcion puede recibir cualqueir objeto que implente la interfaz humano
func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un@ %s, y estoy respirando \n", hu.Genero())
}
