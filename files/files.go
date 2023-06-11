package files

import (
	"bufio"
	"fmt"
	"github.com/elcorderista/godesdecero/ejercicios"
	"io/ioutil"
	"os"
)

var fileName string = "./files/txt/tabla.txt"

// GrabaTabla Crea un archivo y graba la tabla de multiplicar
func GrabaTabla() {
	var texto string = ejercicios.Multiplicar()
	//Creamos el archivo ruta completa desde main
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	//Grabar el texto en un archivo
	fmt.Fprintln(archivo, texto)

	//Cerramos el file
	archivo.Close()
}

// SumaTabla Toma un archivo que indiquemos y agrega la tabla al final
func SumaTabla() {
	var texto string = ejercicios.Multiplicar()
	if !Append(texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

// Append Creamos una funcion para concatenar el contenido de una tabla
// con el contenido de la variable
func Append(texto string) bool {
	//Abrimos el file con permisos de escritura
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}
	//Escribimos en el archivo la tabla
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}

	archivo.Close()

	return true
}

// LeerFile Funcion para leer un archivo
func LeerFile() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoFile() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hubo un error al abrir el file")
		return
	}
	scanner := bufio.NewScanner(archivo)
	//Por cada lectura de linea correcta
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
}
