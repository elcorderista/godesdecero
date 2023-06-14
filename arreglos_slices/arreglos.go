package arreglos_slices

import "fmt"

// Creacion e inicializion de un vector de 10 elementos
var tabla [10]int = [10]int{10, 0, 56}
var matriz [10][20]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)

	tabla2 := [10]string{"Sabine", "Paola", "Olga"}
	fmt.Println(tabla2)

	//Recorrer un array
	for i := 0; i < len(tabla2); i++ {
		if tabla2[i] == "" {
			tabla2[i] = "F1"
		}
		println(tabla2[i])

	}

	matriz[7][19] = 15
	fmt.Println(matriz)
}
